package mialipay

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
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
generate md5
return string
*/
func Md5(str string) string {
	acti := md5.New()
	acti.Write([]byte(str))
	resu := acti.Sum(nil)
	return hex.EncodeToString(resu)
}

/**
build Sign
*/
func BuildSign(quer string, priveKey string) string {
	result := Filter(quer)
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

//func (atp *Alipay) BuildUrl() string {
//	//url:=fmt.Formatter(alipay_trade_pay_url,)
//	//biz, _ := atp.ToJson()
//	strUrl := fmt.Sprintf(alipay_common_parame, atp.AppId, method_alipay_trade_pay, atp.Format, atp.Charset, atp.SignType, atp.Timestamp, atp.Version, atp.NotifyUrl, atp.AppAuthToken, biz)
//	//parse, _ := url.ParseQuery(strUrl)
//	//u := parse.Encode()
//	//uu,_:=url.QueryUnescape(u)
//	//h := sha256.New()
//	//h.Write([]byte(u))
//
//	//parse.Set("sign",strings.ToUpper(string(h.Sum(nil))))
//	//pq,_:=url.Parse(strUrl)
//	//sign := strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
//	//bas := base64.StdEncoding.EncodeToString([]byte(sign))
//	//pg, _ := url.QueryUnescape(parse.Encode())
//	ccc:=BuildSign(strUrl)
//	return atp.GatewayUrl + "?" + strUrl + "&sign=" +ccc
//}

//func FFF(quer string,)  {
//	//pre:="MIIEpAIBAAKCAQEA6I1hmSInXVVklh1Km0UMIWf0k+ffx9vQesQ66wpVlpnLDSuIoxZWYCLD98S5BId+S7RPDdR1vIrPsWD6wMdqzsCdVgIbjYPnNtsYBguns9B6Q9XDMcshuCitlqcydiFSGWTxB16aD6eZhVa8khKcTFvD4+cve6iXC6ie+KCa3SPuiTYLBYJK0w5rA/Q2+vK/Zo+g+gM9+mTShdiFT3fzdS8G+KOylqBpujEPBA/vWcs/fdeu7MRl7/on3SjpXrJ8LAOZcCz6Km0ebR+zrmDoxyQ1B45/1N+piT1H9jaLNv0J9F1ADqeRypy8Gy7kytGx7iUuUXhazi+QwDMSOQ/mLwIDAQABAoIBAD370RO74roh1JmeXRBdqwoMZ0g0+ZSPplXSj9GuA3oMq+5quBSgE420Nn+H07i9VJBhEuEgy/DhHtKsgql3faR1+nm+PhHOIoaghxq1aJGo8625ADn5ZpdrYDlTf0O1Rei22veX44Bgr662m7RjeX1wyBmquSFAomHdI+IpDJ38P6tPxdelRqIFkDIjJl0acuFqDLaPS8YYGlxUCOHQtDaiV9LgIGhxUAldfJ2Pi0C0rSMg7pjlMWfC6FQs/NvxyU88DSIqUxF6W7mjL/FYzb4+w5dzI2OAxh+fzA42lDpKT8IheqQo0HWWkvL/mXvUW8BeqMNbYlmbflgw+00x5qECgYEA+3eIMG3fsdP8sBTv2alzFVvrjgDdoGTrdeiDT1Px1sZmDnTA3pKEiPptZA36J0qGLSFZmfQsMEfVuLF+Z+louZQnG9gTH59S/ZhuptjhJdvlxLHLE+Ujc1yPbq+RqP5VY/r3U5fB/5n3v+8lkuewh10TcJmK+TeyARS0Fb01rV0CgYEA7L6P4aKrPX+fj8rqck0XFXAo4RPRz69T9GEX7wX+bqUxOOwdotMEex5qarng5JrKO5M28uKdVaZwDf1V9VbTJociXKzNZKiRWMF6zVbOnuFyWbu5iBg9gbtV69RaKgUW4cL6sZWfVV1FA4kIBewSV1fCwuf+pkcXCG9NOVwP3PsCgYEAkyFv/K38ySY4XaoSX+8mF6QtoWteQP9rSRSe56Y9BKYWhnTHeDKP+zlTEcyfxadz9CnkLBDSXflZ0fN/+zp0/MfzTfZR0hm9TAWJEHQq36TWvgwsvto9sLzCa0esWQ1eVt47IZzUdEHY1GMPP/HxtnG98w7TYle0/a7oYyPOqT0CgYEA31a+4bvWFARMARkKp532QvE2f291JZpfd1IJhCKBbtxEXVDSfILZ3VRKTDji4obsddthoS0MBHsphukRqIuqUtR1JDyY33bu924/bWuRLO2+NM0WMD+99p9GZV5FWRLIDV8lpVZjo+KDctxZR0z32JIgWEMIRtEV940Yqx4gkPUCgYABdIiqsUPxIgEir7jkCiVzXKz0Fepp736HGaBWyOk5Yp4MmV5jnH6JUQosRySKg4QfuzV9d27Xa+HubzNTrBqWkmgyK7R3F9+itWSE023cSxJf+ZodN4U4yFO1gjHoxsfOaKvW5L3gflE+67uPBP6KDGjlcp7cwYQlwdmF1TwUzg=="
//	//pare:="app_id=2018100261579217&biz_content={\"out_trade_no\":\"20190406010101015\",\"scene\":\"bar_code\",\"auth_code\":\"284978823302537804\",\"subject\":\"测试\",\"total_amount\":\"0.1\"}&charset=utf-8&format=JSON&method=alipay.trade.pay&notify_url=http://pay.ixuzl.com/site/not&sign_type=RSA2&timestamp=2019-04-06 18:47:21&version=1.0"
//
//	key, _ := base64.StdEncoding.DecodeString(pre)
//	privateKey, _ := x509.ParsePKCS1PrivateKey([]byte(key))
//	rng := rand.Reader
//	//message := []byte(pare)
//	hash := sha256.New()
//	hash.Write([]byte(pare))
//	//wenzi:=strings.ToUpper(string(hash.Sum(nil)))
//	signature, _ := rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hash.Sum(nil))
//	signatureHex := base64.StdEncoding.EncodeToString(signature)
//
//	fmt.Println(signatureHex)
//}
