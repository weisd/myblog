package routers

import(
    wechat "github.com/chanxuehong/wechat/server"
    "github.com/chanxuehong/wechat/message/request"
    "net/http"
	"github.com/astaxie/beego"
)

func NewWechatServer() *wechat.Server{
    setting := wechat.ServerSetting{
        Token:"xx",
        TextRequestHandler:wechatTextRequest,
    }
    return wechat.NewServer(&setting)
}

// 处理用户发过来的文本信息
func wechatTextRequest(w http.ResponseWriter, r *http.Request, msg *request.Text){

    beego.Info(msg)

}
