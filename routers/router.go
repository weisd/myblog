package routers

import (
	"github.com/weisd/myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/qrcode", &controllers.QrcodeController{})

    beego.Handler("/wechat", NewWechatServer())
}
