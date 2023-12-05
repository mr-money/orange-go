package Wechat

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"io/ioutil"
	"time"
	"umay-go/Library/Handler"
	"umay-go/MiddleWare"
)

//WxConf 微信支付配置
type WxConf struct {
	IsService uint8  //是否为服务商支付  0否 1是
	SpAppid   string //服务商appid
	SpMchid   string //服务商商户号
	SubAppid  string //商户appid
	SubMchid  string //商户号
	KeyPem    string //私钥
	CertPem   string //证书
	Key       string //api秘钥
}

//
// clientInit
// @Description: 微信支付创建client
// @return error
//
func clientInit(conf WxConf) (*core.Client, error) {
	var (
		mchID       string                // 商户号
		keyPem      string = conf.KeyPem  // 商户私钥
		certPem     string = conf.CertPem // 商户证书
		mchAPIv3Key string = conf.Key     // 商户APIv3密钥
	)

	//服务商判断商户号
	if conf.IsService == 1 {
		mchID = conf.SpMchid
	} else {
		mchID = conf.SubMchid
	}

	// 通过文本内容加载商户私钥
	mchPrivateKey, err := utils.LoadPrivateKey(keyPem)
	if err != nil {
		return nil, err
	}

	// 通过证书的文本内容加载证书
	certificate, err := utils.LoadCertificate(certPem)
	if err != nil {
		return nil, err
	}

	// 从证书中获取证书序列号
	mchCertificateSerialNumber := utils.GetCertificateSerialNumber(*certificate)

	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(
			mchID,
			mchCertificateSerialNumber,
			mchPrivateKey,
			mchAPIv3Key,
		),
	}

	return core.NewClient(context.Background(), opts...)
}

// PayOrder 微信支付订单数据
type PayOrder struct {
	Description string //商品说明
	OutTradeNo  string //订单号
	Attach      string //自定义数据说明
	NotifyUrl   string //回调链接
	Amount      int64  //金额 单位：分
	Openid      string //下单人openid
}

//
// WechatPayService
// @Description: 微信服务商支付
// @return error
//
func (conf WxConf) WechatPayService(payOrder PayOrder) (interface{}, error) {
	client, err := clientInit(conf)
	if err != nil {
		return nil, err
	}

	url := consts.WechatPayAPIServer + "/v3/pay/partner/transactions/jsapi"

	var prepayData = struct {
		SpAppid     string `json:"sp_appid"`
		SpMchid     string `json:"sp_mchid"`
		SubAppid    string `json:"sub_appid"`
		SubMchid    string `json:"sub_mchid"`
		Description string `json:"description"`
		OutTradeNo  string `json:"out_trade_no"`
		NotifyUrl   string `json:"notify_url"`
		Amount      struct {
			Total    int64  `json:"total"`
			Currency string `json:"currency"`
		} `json:"amount"`
		Payer struct {
			SpOpenid  string `json:"sp_openid"`
			SubOpenid string `json:"sub_openid"`
		} `json:"payer"`
	}{
		SpAppid:     conf.SpAppid,
		SpMchid:     conf.SpMchid,
		SubAppid:    conf.SubAppid,
		SubMchid:    conf.SubMchid,
		Description: payOrder.Description,
		OutTradeNo:  payOrder.OutTradeNo,
		NotifyUrl:   payOrder.NotifyUrl,
	}

	prepayData.Amount.Total = payOrder.Amount
	prepayData.Amount.Currency = "CNY"
	//prepayData.Payer.SpOpenid = payOrder.Openid
	prepayData.Payer.SubOpenid = payOrder.Openid

	apiResult, err := client.Post(
		context.Background(),
		url,
		prepayData,
	)
	if err != nil {
		return nil, err
	}

	prepayRes, _ := ioutil.ReadAll(apiResult.Response.Body)

	var prepayId map[string]string
	_ = json.Unmarshal(prepayRes, &prepayId)

	prepay := map[string]string{
		"appId":     conf.SubAppid,
		"timeStamp": cvt.String(time.Now().Unix()),
		"nonceStr":  Handler.RandString(16),
		"package":   "prepay_id=" + prepayId["prepay_id"],
	}

	message := fmt.Sprintf("%s\n%s\n%s\n%s\n",
		prepay["appId"],
		prepay["timeStamp"],
		prepay["nonceStr"],
		prepay["package"],
	)
	sign, err := client.Sign(context.Background(), message)
	if err != nil {
		return nil, err
	}

	prepay["paySign"] = sign.Signature

	return prepay, err
}

//
// WechatPay
// @Description: 微信商户支付
// @return error
//
func (conf WxConf) WechatPay(payOrder PayOrder) (interface{}, error) {
	wechatClient, err := clientInit(conf)
	if err != nil {
		return nil, err
	}

	svc := jsapi.JsapiApiService{Client: wechatClient}
	// 得到prepay_id，以及调起支付所需的参数和签名
	resp, result, err := svc.PrepayWithRequestPayment(
		context.Background(),
		jsapi.PrepayRequest{
			Appid:       core.String(conf.SpAppid),
			Mchid:       core.String(conf.SpMchid),
			Description: core.String(payOrder.Description),
			OutTradeNo:  core.String(payOrder.OutTradeNo),
			Attach:      core.String(""),
			NotifyUrl:   core.String(payOrder.NotifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(1),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(MiddleWare.UserInfo.Username),
			},
		},
	)

	fmt.Println("==========resp==========")
	fmt.Println(resp)
	fmt.Println("==========result==========")
	fmt.Println(result)

	return resp, err
}

// DecryptedData 回调解密资源对象
type DecryptedData struct {
	SpAppID        string `json:"sp_appid"`
	SpMchID        string `json:"sp_mchid"`
	SubAppID       string `json:"sub_appid"`
	SubMchID       string `json:"sub_mchid"`
	OutTradeNo     string `json:"out_trade_no"`
	TradeStateDesc string `json:"trade_state_desc"`
	TradeType      string `json:"trade_type"`
	Attach         string `json:"attach"`
	TransactionID  string `json:"transaction_id"`
	TradeState     string `json:"trade_state"`
	BankType       string `json:"bank_type"`
	SuccessTime    string `json:"success_time"`
	Amount         struct {
		PayerTotal    int    `json:"payer_total"`
		Total         int    `json:"total"`
		Currency      string `json:"currency"`
		PayerCurrency string `json:"payer_currency"`
	} `json:"amount"`
	PromotionDetail []struct {
		Amount              int    `json:"amount"`
		WeChatPayContribute int    `json:"wechatpay_contribute"`
		CouponID            string `json:"coupon_id"`
		Scope               string `json:"scope"`
		MerchantContribute  int    `json:"merchant_contribute"`
		Name                string `json:"name"`
		OtherContribute     int    `json:"other_contribute"`
		Currency            string `json:"currency"`
		StockID             string `json:"stock_id"`
		GoodsDetail         []struct {
			GoodsRemark    string `json:"goods_remark"`
			Quantity       int    `json:"quantity"`
			DiscountAmount int    `json:"discount_amount"`
			GoodsID        string `json:"goods_id"`
			UnitPrice      int    `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"promotion_detail"`
	Payer struct {
		OpenID string `json:"openid"`
	} `json:"payer"`
	SceneInfo struct {
		DeviceID string `json:"device_id"`
	} `json:"scene_info"`
}

//
// DecryptWechatData
// @Description: 微信支付回调解密
// @param key
// @param nonce 随机串
// @param associatedData 附加数据
// @param ciphertext 数据密文
// @return DecryptedData
// @return error
//
func DecryptWechatData(key, nonce, associatedData, ciphertext string) (DecryptedData, error) {
	keyBytes := []byte(key)
	nonceBytes := []byte(nonce)
	associatedDataBytes := []byte(associatedData)
	//ciphertextBytes := []byte(ciphertext)

	// 将Base64编码的密文解码为字节数组
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return DecryptedData{}, errors.New("Error decoding Base64 ciphertext:" + err.Error())
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return DecryptedData{}, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return DecryptedData{}, err
	}

	// 解密数据
	decryptedByte, err := gcm.Open(nil, nonceBytes, ciphertextBytes, associatedDataBytes)
	if err != nil {
		fmt.Println(key, nonce, associatedData, ciphertext)

		return DecryptedData{}, err
	}

	var decryptedData DecryptedData

	// 解码JSON
	err = json.Unmarshal(decryptedByte, &decryptedData)
	if err != nil {
		return DecryptedData{}, err
	}

	return decryptedData, nil
}
