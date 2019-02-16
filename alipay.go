package alipay

import (
	"time"
)

/**
alipay
 */

const (
	version            = "1.0"
	format_json string = "JSON"
)

type Alipay struct {
	//应用ID
	AppId string
	//支付宝公钥
	AlipayPublicKey string
	//商户私钥
	MerchantPrivateKey string
	//签名方式
	SignType string
	//支付宝网关
	GatewayUrl string
	//编码格式
	//charset string
	// 支付宝服务器主动通知商户服务器里指定的页面
	NotifyUrl string
	//异步通知地址,只有扫码支付预下单可用
	ReturnUrl string
}
type Request interface {
	Execute() interface{}
	ToUrl()   interface{}
}

func NewAlipay(appid string, alipayPublicKey string, merchantPrivateKey string, signType string, gatewayUrl string, notifyUrl string, returnUrl string) *Alipay {
	alipay := &Alipay{
		appid,
		alipayPublicKey,
		merchantPrivateKey,
		signType,
		gatewayUrl,
		notifyUrl,
		returnUrl,
	}
	return alipay
}

/**
alipay.trade.pay
 */
func (a *Alipay) AlipayTradePay() *AlipayTradePay {
	alipayTradePay := &AlipayTradePay{}
	//alipayTradePay := new(AlipayTradePay)
	alipayTradePay.Alipay = *a
	alipayTradePay.method = method_alipay_trade_pay
	alipayTradePay.format = format_json
	alipayTradePay.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	alipayTradePay.Version = version
	return alipayTradePay
}



//func (a *Alipay) Request() *AlipayTradePay {
//	fmt.Println(1111)
//	return  nil
//}
