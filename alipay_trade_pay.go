package alipay

import (
	"encoding/json"
	"fmt"
)

const method_alipay_trade_pay string = "alipay.trade.pay"
const alipay_trade_pay_url  = "?app_id=%v&method=%v&format=%v&charset=%v&sign_type=%v&sign=%v&timestamp=%v&version=%v&notify_url=%v&app_auth_token=%v&biz_content=%v"

type AlipayTradePay struct {
	method       string
	Timestamp    string
	format       string
	charset      string
	Version      string
	Format       string
	AppAuthToken string
	*Alipay
	*BizContent
}

type BizContent struct {
	OutTradeNo         string `json:"out_trade_no"`
	Scene              string `json:"scene"`
	AuthCode           string `json:"auth_code"`
	ProductCode        string `json:"product_code"`
	Subject            string `json:"subject"`
	BuyerId            string `json:"buyer_id"`
	SellerId           string `json:"seller_id"`
	TotalAmount        string `json:"total_amount"`
	TransCurrency      string `json:"trans_currency"`
	SettleCurrency     string `json:"settle_currency"`
	DiscountableAmount string `json:"discountable_amount"`
	Body               string `json:"body"`
	GoodsDetail        *[]GoodsDetail `json:"goods_detail"`
	OperatorId         string `json:"operator_id"`
	StoreId            string `json:"store_id"`
	TerminalId         string `json:"terminal_id"`
	ExtendParams       *ExtendParams `json:"extend_params"`
	TimeoutExpress     string `json:"timeout_express"`
	AuthConfirmMode    string `json:"auth_confirm_mode"`
	TerminalParams     string `json:"terminal_params"`
	PromoParams        *PromoParams `json:"promo_params"`
}
type ExtendParams struct {
	SysServiceProviderId string	 `json:"sys_service_provider_id"`
	IndustryRefluxInfo   string	 `json:"industry_reflux_info"`
	CardType             string	 `json:"card_type"`
}

type GoodsDetail struct {
	GoodsId        string `json:"goods_id"`
	GoodsName      string `json:"goods_name"`
	Quantity       string `json:"quantity"`
	Price          string `json:"price"`
	GoodsCategory  string `json:"goods_category"`
	CategoriesTree string `json:"categories_tree"`
	Body           string `json:"body"`
	ShowUrl        string `json:"show_url"`
}

type PromoParams struct {
	ActualOrderTime string
}

func (atp *AlipayTradePay) SetBizContent(bizContent BizContent) {
	atp.BizContent = &bizContent
}

//func (bc *BizContent) setPromoParams(promoParams PromoParams) {
//	*bc.PromoParams = promoParams
//}

func (bc *BizContent) SetGoodsDetail(goodsDetail []GoodsDetail) {
	bc.GoodsDetail = &goodsDetail
}

func (bc *BizContent) SetExtendParams(extendParams ExtendParams) {
	*bc.ExtendParams = extendParams
}
func (bc *BizContent) ToJson()(string,error) {
	byteJson, err := json.Marshal(bc)
	if err != nil {
		return "",err
	}
	return string(byteJson),nil
}

func (atp *AlipayTradePay)ToUrl() interface{} {
	request:=atp.method+alipay_trade_pay_url
	//url:=fmt.Printf(request,atp.AppId,atp.method,atp.format,atp.charset,atp.SignType,atp.)
	return nil
}