package mialipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"net/url"
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
	//编码格式
	Charset string
	//支付宝网关
	GatewayUrl string
	// 支付宝服务器主动通知商户服务器里指定的页面
	NotifyUrl string
	//异步通知地址,只有扫码支付预下单可用
	ReturnUrl string
}

type RespMsg struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code"`
	SubMsg  string `json:"sub_msg"`
	Sign    string `json:"sign"`
}

func NewAlipay(appid string, alipayPublicKey string, merchantPrivateKey string, signType string, charset string, gatewayUrl string, notifyUrl string, returnUrl string) *Alipay {
	alipay := &Alipay{
		appid,
		alipayPublicKey,
		merchantPrivateKey,
		signType,
		charset,
		gatewayUrl,
		notifyUrl,
		returnUrl,
	}
	return alipay
}

/**
alipay.trade.pay
*/
func (ali *Alipay) AlipayTradePay() *AlipayTradePay {
	alipayTradePay := &AlipayTradePay{}
	//alipayTradePay := new(AlipayTradePay)
	alipayTradePay.Alipay = ali
	alipayTradePay.Format = "JSON"
	alipayTradePay.method = method_alipay_trade_pay
	alipayTradePay.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	alipayTradePay.Version = version
	return alipayTradePay
}

/**
请求执行
*/

func GetRequest(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body), err
}

/**
build Sign
*/
func BuildSign(strSign string, priveKey string) string {
	result := Filter(strSign)
	key, _ := base64.StdEncoding.DecodeString(result)
	privateKey, _ := x509.ParsePKCS1PrivateKey([]byte(priveKey))
	hash := sha256.New()
	hash.Write([]byte(key))
	signature, _ := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash.Sum(nil))
	signatureHex := base64.StdEncoding.EncodeToString(signature)
	return signatureHex
}

/**
filter parameter sort
*/
func Filter(quer string) string {
	pars, _ := url.ParseQuery(quer)
	newPare := url.Values{}
	for k, v := range pars {
		if v[0] != "" {
			newPare.Set(k, v[0])
		}
	}
	str := newPare.Encode()
	return str
}

