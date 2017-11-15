package wechat

import "reflect"

//https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=8_1
//app pay method
type AppPay struct {
	TradeType      string            `json:"trade_type"`
	Body           string            `json:"body"`
	NonceStr       string            `json:"nonce_str" json:"nonce_str"`
	NotifyUrl      string            `json:"notify_url" json:"-"`
	SpbillCreateIp string            `json:"spbill_create_ip" json:"-"`
	TotalFee       string            `json:"total_fee" json:"-"`
	OutTradeNo     string            `json:"out_trade_no" json:"-"`
	Sign           string            `json:"sign" json:"sign"`
	Attach         string            `json:"attach" json:"attach"`
	Data           map[string]string `json:"-"`
}

func (*AppPay) ChooseHost() string {
	return WechatApiUnifyOrder
}

func (a *AppPay) SetExtraParam(key, value string) {
	a.Data[key] = value
}

func (a *AppPay) SetActionParam() {
	e := reflect.ValueOf(a)
	t := e.Type()
	for i := 0; i < t.NumField(); i++ {
		if v := e.Field(i).Interface(); v != nil && t.Field(i).Tag.Get("json") != "-" {
			a.Data[t.Field(i).Tag.Get("json")] = v.(string)
		}
	}
}

func (*AppPay) DoRequest(body []byte) ([]byte, error) {
	panic("implement me")
}

func (a *AppPay)GetMapData()map[string]string{
	return a.Data
}