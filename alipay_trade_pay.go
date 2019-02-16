package alipay

const method_alipay_trade_pay string = "alipay.trade.pay"

type AlipayTradePay struct {
	method       string
	Timestamp    string
	format       string
	charset      string
	Version      string
	Format       string
	AppAuthToken string
	Alipay
	*BizContent
}

type BizContent struct {
	OutTradeNo         string
	Scene              string
	AuthCode           string
	ProductCode        string
	Subject            string
	BuyerId            string
	SellerId           string
	TotalAmount        string
	TransCurrency      string
	SettleCurrency     string
	DiscountableAmount string
	Body               string
	GoodsDetail        *[]GoodsDetail
	OperatorId         string
	StoreId            string
	TerminalId         string
	ExtendParams       *ExtendParams
	TimeoutExpress     string
	AuthConfirmMode    string
	TerminalParams     string
	PromoParams        *PromoParams
}
type ExtendParams struct {
	SysServiceProviderId string
	IndustryRefluxInfo   string
	CardType             string
}

type GoodsDetail struct {
	GoodsId        string
	GoodsName      string
	Quantity       string
	Price          string
	GoodsCategory  string
	CategoriesTree string
	Body           string
	ShowUrl        string
}

type PromoParams struct {
	ActualOrderTime string
}

func (atp *AlipayTradePay) setBizContent(bizContent BizContent) {
	*atp.BizContent = bizContent
}

func (bc *BizContent) setPromoParams(promoParams PromoParams) {
	*bc.PromoParams = promoParams
}

func (bc *BizContent) setGoodsDetail(goodsDetail []GoodsDetail) {
	*bc.GoodsDetail = goodsDetail
}

func (bc *BizContent) setExtendParams(extendParams ExtendParams) {
	*bc.ExtendParams = extendParams
}
func (bc *BizContent) ToJson() string{
	return ""
}

func (atp *AlipayTradePay)ToUrl() interface{} {

	return nil
}