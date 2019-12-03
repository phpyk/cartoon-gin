package configs

const (
	HipopayMerchantNo = "WC5d6f2fab507f8"
	HipopayCurrency = "USD"
	HipopayCurrencySymbol = "$"
	HipopayReturnUrl = ""//TODO
	HipopayAlipayNotifyUrl = "/api/notify-v2/alipay"
	HipopayWechatNotifyUrl = "/api/notify-v2/wxpay"
)

const (
	//付款渠道
	PayChannelAlipay = 1
	PayChannelWechat = 2
	PayChannelApple = 3
)