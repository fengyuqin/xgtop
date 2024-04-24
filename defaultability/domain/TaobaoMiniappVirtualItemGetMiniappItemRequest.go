package domain


type TaobaoMiniappVirtualItemGetMiniappItemRequest struct {
    /*
        商品id列表     */
    ItemIds  *[]int64 `json:"item_ids,omitempty" `

    /*
        小程序小游戏appId     */
    AppId  *int64 `json:"app_id,omitempty" `

}

func (s *TaobaoMiniappVirtualItemGetMiniappItemRequest) SetItemIds(v []int64) *TaobaoMiniappVirtualItemGetMiniappItemRequest {
    s.ItemIds = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemRequest) SetAppId(v int64) *TaobaoMiniappVirtualItemGetMiniappItemRequest {
    s.AppId = &v
    return s
}
