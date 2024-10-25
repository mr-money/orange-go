package Wechat

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/partnertransferbatch"
	"io/ioutil"
	"log"
)

//
// TransferBatch
// @Description: 批量转账数据
//
type TransferBatch struct {
	OutBatchNo  string //商家批次单号 商户系统内部唯一
	BatchName   string //批次名称
	BatchRemark string //批次备注
	TotalAmount int64  //转账总金额 分
	TotalNum    int64  //转账总笔数
	DetailList  []struct {
		Openid         string
		OutDetailNo    string //商家明细单号
		TransferAmount int64  //	转账金额 分
		TransferRemark string // 转账备注 最多允许32个字符
		UserName       string //收款用户姓名
		//UserIdCard     string //收款用户身份证号
	}
}

//
// TransferBatch
// @Description: 发起商家转账 商户可以通过该接口同时向多个用户微信零钱进行转账操作
// @receiver conf 支付配置
// @param TransferBatch 提现数据
// @return interface{} 微信返回
// @return error
//
func (conf WxConf) TransferBatch(TransferBatch *TransferBatch) (interface{}, error) {
	wechatClient, err := clientInit(conf)
	if err != nil {
		return nil, err
	}

	//服务商转账请求数据
	transferBatchRequest := partnertransferbatch.InitiateTransferBatchRequest{
		SubMchid:          core.String(conf.SubMchid),
		SubAppid:          core.String(conf.SubAppid),
		AuthorizationType: partnertransferbatch.AUTHTYPE_INFORMATION_AUTHORIZATION_TYPE.Ptr(),
		OutBatchNo:        core.String(TransferBatch.OutBatchNo),
		BatchName:         core.String(TransferBatch.BatchName),
		BatchRemark:       core.String(TransferBatch.BatchRemark),
		TotalAmount:       core.Int64(TransferBatch.TotalAmount),
		TotalNum:          core.Int64(TransferBatch.TotalNum),
		SpAppid:           core.String(conf.SpAppid),
		TransferPurpose:   partnertransferbatch.TRANSFERUSETYPE_GOODSPAYMENT.Ptr(),
		TransferScene:     partnertransferbatch.TRANSFERSCENE_ORDINARY_TRANSFER.Ptr(),
	}

	var TransferDetailList []partnertransferbatch.TransferDetailInput
	for _, detail := range TransferBatch.DetailList {
		TransferDetailList = append(TransferDetailList, partnertransferbatch.TransferDetailInput{
			Openid:         core.String(detail.Openid),
			OutDetailNo:    core.String(detail.OutDetailNo),
			TransferAmount: core.Int64(detail.TransferAmount),
			TransferRemark: core.String(detail.TransferRemark),
			UserName:       core.String(detail.UserName),
			//UserIdCard:     core.String(detail.UserIdCard),
		})
	}

	transferBatchRequest.TransferDetailList = TransferDetailList

	fmt.Printf("transferBatchRequest---------%v\n", transferBatchRequest)

	svc := partnertransferbatch.TransferBatchApiService{Client: wechatClient}

	resp, apiResult, err := svc.InitiateTransferBatch(
		context.Background(),
		transferBatchRequest,
	)

	//商家转账请求数据
	/*transferBatchRequest := transferbatch.InitiateBatchTransferRequest{
		Appid:           core.String(conf.SubAppid),
		OutBatchNo:      core.String(TransferBatch.OutBatchNo),
		BatchName:       core.String(TransferBatch.BatchName),
		BatchRemark:     core.String(TransferBatch.BatchRemark),
		TotalAmount:     core.Int64(TransferBatch.TotalAmount),
		TotalNum:        core.Int64(TransferBatch.TotalNum),
		TransferSceneId: core.String("1001"),
	}

	var TransferDetailList []transferbatch.TransferDetailInput
	for _, detail := range TransferBatch.DetailList {
		TransferDetailList = append(TransferDetailList, transferbatch.TransferDetailInput{
			Openid:         core.String(detail.Openid),
			OutDetailNo:    core.String(detail.OutDetailNo),
			TransferAmount: core.Int64(detail.TransferAmount),
			TransferRemark: core.String(detail.TransferRemark),
			UserName:       core.String(detail.UserName),
		})
	}
	transferBatchRequest.TransferDetailList = TransferDetailList

	fmt.Printf("transferBatchRequest---------%v\n", transferBatchRequest)

	svc := transferbatch.TransferBatchApiService{Client: wechatClient}

	resp, apiResult, err := svc.InitiateBatchTransfer(
		context.Background(),
		transferBatchRequest,
	)*/

	if err != nil {
		// 处理错误
		log.Printf("call InitiateBatchTransfer err:%s", err)
	} else {
		// 处理返回结果
		log.Printf("status=%d resp=%s", apiResult.Response.StatusCode, resp)
	}

	transferRes, _ := ioutil.ReadAll(apiResult.Response.Body)

	var result map[string]interface{}
	_ = json.Unmarshal(transferRes, &result)

	return transferRes, nil
}
