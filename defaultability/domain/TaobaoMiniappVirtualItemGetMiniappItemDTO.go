package domain


type TaobaoMiniappVirtualItemGetMiniappItemDTO struct {
    /*
        商品id     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        主图     */
    MainPicUrl  *string `json:"main_pic_url,omitempty" `

    /*
        商品标题     */
    Title  *string `json:"title,omitempty" `

    /*
        一口价，单位【分】     */
    ReservePrice  *int64 `json:"reserve_price,omitempty" `

    /*
        外部商品id     */
    OutItemId  *string `json:"out_item_id,omitempty" `

    /*
        0-正常，其他均为不正常     */
    AuctionStatus  *int64 `json:"auction_status,omitempty" `

}

func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetItemId(v int64) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.ItemId = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetMainPicUrl(v string) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.MainPicUrl = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetTitle(v string) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.Title = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetReservePrice(v int64) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.ReservePrice = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetOutItemId(v string) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.OutItemId = &v
    return s
}
func (s *TaobaoMiniappVirtualItemGetMiniappItemDTO) SetAuctionStatus(v int64) *TaobaoMiniappVirtualItemGetMiniappItemDTO {
    s.AuctionStatus = &v
    return s
}
