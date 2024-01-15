package Wechat

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/services/transferbatch"
)

//
// ServiceTransfer
// @Description: 商家转账
// @receiver conf
// @return interface{}
// @return error
//
func (conf WxConf) ServiceTransfer() (interface{}, error) {
	wechatClient, err := clientInit(conf)
	if err != nil {
		return nil, err
	}

	svc := transferbatch.TransferBatchApiService{
		Client: wechatClient,
	}

	resp, result, err := svc.InitiateBatchTransfer(
		context.Background(),
		transferbatch.InitiateBatchTransferRequest{
			Appid:       core.String("wxf636efh567hg4356"),
			OutBatchNo:  core.String("plfk2020042013"),
			BatchName:   core.String("2019年1月深圳分部报销单"),
			BatchRemark: core.String("2019年1月深圳分部报销单"),
			TotalAmount: core.Int64(4000000),
			TotalNum:    core.Int64(200),
			TransferDetailList: []transferbatch.TransferDetailInput{transferbatch.TransferDetailInput{
				OutDetailNo:    core.String("x23zy545Bd5436"),
				TransferAmount: core.Int64(200000),
				TransferRemark: core.String("2020年4月报销"),
				Openid:         core.String("o-MYE42l80oelYMDE34nYD456Xoy"),
				UserName:       core.String("757b340b45ebef5467rter35gf464344v3542sdf4t6re4tb4f54ty45t4yyry45"),
			}},
			TransferSceneId: core.String("1000"),
		},
	)

	fmt.Println("==========resp==========")
	fmt.Println(resp)
	fmt.Println("==========result==========")
	fmt.Println(result)

	return resp, err
}
