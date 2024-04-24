package request

import (
        "github.com/fengyuqin/xgtop/defaultability/domain"
        "github.com/fengyuqin/xgtop/util"
    )

type TaobaoMiniappVirtualItemGetRequest struct {
    /*
        请求     */
    ItemRequest  *domain.TaobaoMiniappVirtualItemGetMiniappItemRequest `json:"item_request" required:"true" `
}

func (s *TaobaoMiniappVirtualItemGetRequest) SetItemRequest(v domain.TaobaoMiniappVirtualItemGetMiniappItemRequest) *TaobaoMiniappVirtualItemGetRequest {
    s.ItemRequest = &v
    return s
}

func (req *TaobaoMiniappVirtualItemGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ItemRequest != nil) {
        paramMap["item_request"] = util.ConvertStruct(*req.ItemRequest)
    }
    return paramMap
}

func (req *TaobaoMiniappVirtualItemGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}