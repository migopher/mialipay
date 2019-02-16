package alipay

import "time"

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
	GoodsDetail        string
	OperatorId         string
	StoreId            string
	TerminalId         string
	ExtendParams       string
	TimeoutExpress     string
	AuthConfirmMode    string
	TerminalParams     string
	PromoParams        string
}

func (atp *AlipayTradePay) init() {
	atp.method = method_alipay_trade_pay
	atp.format = format_json
	atp.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	atp.Version = version
}

//func (atp *AlipayTradePay) setBizContent() *BizContent {
//	return atp.BizContent
//}
