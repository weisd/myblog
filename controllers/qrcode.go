package controllers

import(
	"github.com/astaxie/beego"
)

type QrcodeController struct{
    beego.Controller
}

func (this *QrcodeController) Get(){
    this.TplNames = "qrcode/index.html"
}
