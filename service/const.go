package service

import "time"

var ( //微信相关
	AppID     = "wxb7482365869f9d2b"
	AppSecret = "de3d0e8ac0ae4a3f1b1ffdbf4a488b75"

	//获取accesstoken
	AccessTokenUrl = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=wxb7482365869f9d2b&secret=de3d0e8ac0ae4a3f1b1ffdbf4a488b75"

	//
)

var ( //redis
	AccessTokenKey = "accesstokenkey"
	//token过期时间
	AccessTokenExpire = 7000 * time.Second

	//记录用户上下文超时时间 12x小时
	UserMsgExpire = 43200 * time.Second
)

var (
	GET  = "GET"
	POST = "POST"
)
