package Wechat

import (
	"context"
	"encoding/json"
	"github.com/wechatpay-apiv3/wechatpay-go/core/consts"
	"io/ioutil"
)

//
// RefundOrder
// @Description: 微信退款订单数据结构
//
type RefundOrder struct {
	OutTradeNo  string //商户订单号
	OutRefundNo string //退款单号
	Reason      string //退款原因
	NotifyUrl   string //回调链接
	Total       int64  //金额 单位：分
	Refund      int64  //金额 单位：分
}

//
// ServiceRefund
// @Description: 服务商退款
// @receiver conf 配置
// @param refundOrder 退款订单数据
// @return interface{}
// @return error
//
func (conf WxConf) ServiceRefund(refundOrder *RefundOrder) (interface{}, error) {
	client, err := clientInit(conf)
	if err != nil {
		return nil, err
	}

	url := consts.WechatPayAPIServer + "/v3/refund/domestic/refunds"

	var refundData = struct {
		SubMchid string `json:"sub_mchid"`
		//TransactionId string `json:"transaction_id"`
		OutTradeNo  string `json:"out_trade_no"`
		OutRefundNo string `json:"out_refund_no"`
		Reason      string `json:"reason"`
		NotifyUrl   string `json:"notify_url"`
		Amount      struct {
			Refund   int64  `json:"refund"`
			Total    int64  `json:"total"`
			Currency string `json:"currency"`
		} `json:"amount"`
	}{
		SubMchid:    conf.SubMchid,
		OutTradeNo:  refundOrder.OutTradeNo,
		OutRefundNo: refundOrder.OutRefundNo,
		Reason:      refundOrder.Reason,
		NotifyUrl:   refundOrder.NotifyUrl,
	}

	refundData.Amount.Refund = refundOrder.Refund
	refundData.Amount.Total = refundOrder.Total
	refundData.Amount.Currency = "CNY"

	apiResult, err := client.Post(
		context.Background(),
		url,
		refundData,
	)
	if err != nil {
		return nil, err
	}

	refundRes, _ := ioutil.ReadAll(apiResult.Response.Body)

	var result map[string]interface{}
	_ = json.Unmarshal(refundRes, &result)

	return result, nil
}
