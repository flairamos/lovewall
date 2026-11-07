package wechat

import (
	"time"
)

const PREPAY = "https://api.mch.weixin.qq.com/v3/pay/transactions/jsapi"

type PrepayParam struct {
	Appid         string    `json:"appid"`
	Mchid         string    `json:"mchid"`
	Description   string    `json:"description"`
	OutTradeNo    string    `json:"out_trade_no"`
	TimeExpire    time.Time `json:"time_expire"`
	Attach        string    `json:"attach"`
	NotifyURL     string    `json:"notify_url"`
	GoodsTag      string    `json:"goods_tag"`
	SupportFapiao bool      `json:"support_fapiao"`
	Amount        struct {
		Total    int    `json:"total"`
		Currency string `json:"currency"`
	} `json:"amount"`
	Payer struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Detail struct {
		CostPrice   int    `json:"cost_price"`
		InvoiceID   string `json:"invoice_id"`
		GoodsDetail []struct {
			MerchantGoodsID  string `json:"merchant_goods_id"`
			WechatpayGoodsID string `json:"wechatpay_goods_id"`
			GoodsName        string `json:"goods_name"`
			Quantity         int    `json:"quantity"`
			UnitPrice        int    `json:"unit_price"`
		} `json:"goods_detail"`
	} `json:"detail"`
	SceneInfo struct {
		PayerClientIP string `json:"payer_client_ip"`
		DeviceID      string `json:"device_id"`
		StoreInfo     struct {
			ID       string `json:"id"`
			Name     string `json:"name"`
			AreaCode string `json:"area_code"`
			Address  string `json:"address"`
		} `json:"store_info"`
	} `json:"scene_info"`
	SettleInfo struct {
		ProfitSharing bool `json:"profit_sharing"`
	} `json:"settle_info"`
}

type PrepayResp struct {
	PrepayId string `json:"prepay_id"`
}

//func MakeSign(param PrepayParam) (string, error) {
//	var (
//		mchID                      string = config.Config.Miniprogram.Mchid                      // 商户号
//		mchCertificateSerialNumber string = config.Config.Miniprogram.MchCertificateSerialNumber // 商户证书序列号
//		mchAPIv3Key                string = config.Config.Miniprogram.MchAPIv3Key                // 商户APIv3密钥
//	)
//	// 使用 utils 提供的函数从本地文件中加载商户私钥，商户私钥会用来生成请求的签名
//	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(config.Config.Path + "/apiclient_key.pem")
//	if err != nil {
//		log.Fatal("load merchant private key error")
//		return "", err
//	}
//}
