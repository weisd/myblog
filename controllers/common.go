package controllers

import(
	"github.com/astaxie/beego"
)

type CommonControllers struct{
    beego.Controller
}

func (this *CommonControllers) Prepare(){
    this.Layout = "layout.html"
}
