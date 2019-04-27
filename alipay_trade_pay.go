package mialipay

import (
	"encoding/json"
	"fmt"
	"net/url"
)

const method_alipay_trade_pay string = "alipay.trade.pay"
const alipay_common_parame = "app_id=%v&method=%v&format=%v&charset=%v&sign_type=%v&timestamp=%v&version=%v&notify_url=%v&app_auth_token=%v&biz_content=%v"

type AlipayTradePay struct {
	method       string
	Timestamp    string
	Version      string
	Sign         string
	Format       string
	AppAuthToken string
	*Alipay
	BizContent
}

type BizContent struct {
	OutTradeNo         string         `json:"out_trade_no,omitempty"`
	Scene              string         `json:"scene,omitempty"`
	AuthCode           string         `json:"auth_code,omitempty"`
	ProductCode        string         `json:"product_code,omitempty"`
	Subject            string         `json:"subject,omitempty"`
	BuyerId            string         `json:"buyer_id,omitempty"`
	SellerId           string         `json:"seller_id,omitempty"`
	TotalAmount        float64         `json:"total_amount,omitempty"`
	TransCurrency      string         `json:"trans_currency,omitempty"`
	SettleCurrency     string         `json:"settle_currency,omitempty"`
	DiscountableAmount string         `json:"discountable_amount,omitempty"`
	Body               string         `json:"body,omitempty"`
	GoodsDetail        *[]GoodsDetail `json:"goods_detail,omitempty"`
	OperatorId         string         `json:"operator_id,omitempty"`
	StoreId            string         `json:"store_id,omitempty"`
	TerminalId         string         `json:"terminal_id,omitempty"`
	ExtendParams       *ExtendParams  `json:"extend_params,omitempty"`
	TimeoutExpress     string         `json:"timeout_express,omitempty"`
	AuthConfirmMode    string         `json:"auth_confirm_mode,omitempty"`
	TerminalParams     string         `json:"terminal_params,omitempty"`
	PromoParams        *PromoParams   `json:"promo_params,omitempty"`
}
type ExtendParams struct {
	SysServiceProviderId string `json:"sys_service_provider_id,omitempty"`
	IndustryRefluxInfo   string `json:"industry_reflux_info,omitempty"`
	CardType             string `json:"card_type,omitempty"`
}

type GoodsDetail struct {
	GoodsId        string `json:"goods_id,omitempty"`
	GoodsName      string `json:"goods_name,omitempty"`
	Quantity       string `json:"quantity,omitempty"`
	Price          string `json:"price,omitempty"`
	GoodsCategory  string `json:"goods_category,omitempty"`
	CategoriesTree string `json:"categories_tree,omitempty"`
	Body           string `json:"body,omitempty"`
	ShowUrl        string `json:"show_url,omitempty"`
}

type PromoParams struct {
	ActualOrderTime string `json:"actual_order_time,omitempty"`
}

func (atp *AlipayTradePay) BuildUrl() string {
	biz, _ := atp.BizContent.ToJson()
	strUrl := fmt.Sprintf(alipay_common_parame, atp.AppId, method_alipay_trade_pay, atp.Format, atp.Charset, atp.SignType, atp.Timestamp, atp.Version, atp.NotifyUrl, atp.AppAuthToken, biz)
	//signUrl := fmt.Sprintf(alipay_common_parame, atp.AppId, method_alipay_trade_pay, atp.Format, atp.Charset, "", atp.Timestamp, atp.Version, atp.NotifyUrl, atp.AppAuthToken, biz)
	strUrl = FilterUrl(strUrl)
	fmt.Println("排序结果", strUrl)
	sign := BuildSign(strUrl, atp.MerchantPrivateKey)
	v, _ := url.ParseQuery(strUrl)
	fmt.Println(v.Encode())
	return atp.GatewayUrl + "?" + v.Encode() + "&sign=" + sign
}

func (atp *AlipayTradePay) SetBizContent(bizContent BizContent) {
	atp.BizContent = bizContent
}

func (bc *BizContent) setPromoParams(promoParams PromoParams) {
	*bc.PromoParams = promoParams
}

func (bc *BizContent) SetGoodsDetail(goodsDetail []GoodsDetail) {
	bc.GoodsDetail = &goodsDetail
}

func (bc *BizContent) SetExtendParams(extendParams ExtendParams) {
	*bc.ExtendParams = extendParams
}
func (bc *BizContent) ToJson() (string, error) {
	byteJson, err := json.Marshal(bc)
	if err != nil {
		return "", err
	}
	return string(byteJson), nil
}

//func ExecRequest()  {
//
//}
