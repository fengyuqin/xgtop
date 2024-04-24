package response

import (
    "github.com/fengyuqin/xgtop/defaultability/domain"
)

type TaobaoMiniappVirtualItemGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否成功
    */
    Suc  bool `json:"suc,omitempty" `
    /*
        业务错误信息
    */
    ECode  int64 `json:"e_code,omitempty" `
    /*
        返回列表
    */
    Model  []domain.TaobaoMiniappVirtualItemGetMiniappItemDTO `json:"model,omitempty" `
    /*
        错误信息
    */
    ErrMessage  string `json:"err_message,omitempty" `
}
