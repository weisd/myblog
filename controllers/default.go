package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	CommonControllers
}

func (this *MainController) Get() {
    beego.Info("main")
    /*
    this.Data["name"] = "weisd"

    this.TplNames = "main/index.html"
    */
    this.Redirect("/qrcode", 302)
}
