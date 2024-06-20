package tpwd

import (
	"coupon-go-api/pkg/taobao"
	"coupon-go-api/pkg/taobao/tpwd/request"
	"coupon-go-api/pkg/taobao/tpwd/response"
	"coupon-go-api/pkg/taobao/util"
	"errors"
	"log"
)

type AliTpwd struct {
	Client *taobao.TopClient
}

func NewAliTpwd(client *taobao.TopClient) *AliTpwd {
	return &AliTpwd{client}
}

// 淘口令查询优惠券
func (ali *AliTpwd) TpwdQuery(req *request.TpwdQueryRequest) (*response.TpwdQueryResponse, error) {
	if ali.Client == nil {
		return nil, errors.New("topClient is nil")
	}

	var jsonStr, err = ali.Client.Execute("taobao.wireless.share.tpwd.query", req.ToMap(), req.ToFileMap())
	var resStruct = response.TpwdQueryResponse{}
	if err != nil {
		log.Println("taobaoWirelessShareTpwdQuery error", err)
		return &resStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &resStruct)
	if resStruct.Body == "" || len(resStruct.Body) == 0 {
		resStruct.Body = jsonStr
	}
	return &resStruct, err
}
