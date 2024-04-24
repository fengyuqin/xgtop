package defaultability

import (
    "log"
    "errors"
    "github.com/fengyuqin/xgtop"
    "github.com/fengyuqin/xgtop/defaultability/request"
    "github.com/fengyuqin/xgtop/defaultability/response"
    "github.com/fengyuqin/xgtop/util"
)

type Defaultability struct {
    Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability{
    return &Defaultability{client}
}

/*
    小程序关联虚拟商品查询
*/
func (ability *Defaultability) TaobaoMiniappVirtualItemGet(req *request.TaobaoMiniappVirtualItemGetRequest,session string) (*response.TaobaoMiniappVirtualItemGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.miniapp.virtual.item.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoMiniappVirtualItemGetResponse{}
    if(err != nil){
        log.Println("taobaoMiniappVirtualItemGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest,session string) (*response.TaobaoKfcKeywordSearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Defaultability topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoKfcKeywordSearchResponse{}
    if(err != nil){
        log.Println("taobaoKfcKeywordSearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
