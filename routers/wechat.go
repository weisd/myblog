package routers

import(
    wechat "github.com/chanxuehong/wechat/server"
    "github.com/chanxuehong/wechat/message/request"
    "net/http"
	"github.com/astaxie/beego"
)

func NewWechatServer() *wechat.Server{
    setting := wechat.ServerSetting{
            Token                          : "vshida",
            // unknown request handler
            UnknownRequestHandler          : UnknownRequestHandler,
        
            // common request handler
            TextRequestHandler             : TextRequestHandler,
            ImageRequestHandler            : ImageRequestHandler,
            VoiceRequestHandler            : VoiceRequestHandler,
            VoiceRecognitionRequestHandler : VoiceRecognitionRequestHandler,
            VideoRequestHandler            : VideoRequestHandler,
            LocationRequestHandler         : LocationRequestHandler,
            LinkRequestHandler             : LinkRequestHandler,
            // event handler,
            SubscribeEventHandler          : SubscribeEventHandler,
            UnsubscribeEventHandler        : UnsubscribeEventHandler,
            SubscribeByScanEventHandler    : SubscribeByScanEventHandler,
            ScanEventHandler               : ScanEventHandler,
            LocationEventHandler           : LocationEventHandler,
            MenuClickEventHandler          : MenuClickEventHandler,
            MenuViewEventHandler           : MenuViewEventHandler,
            MassSendJobFinishEventHandler  : MassSendJobFinishEventHandler,
            MerchantOrderEventHandler      : MerchantOrderEventHandler,
    }
    return wechat.NewServer(&setting)
}

func InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error){
    beego.Info(err)
}

// 非法请求
func UnknownRequestHandler(w http.ResponseWriter, r *http.Request, msg []byte){
    beego.Info(msg)
}

func TextRequestHandler(http.ResponseWriter, *http.Request, *request.Text){
}

func ImageRequestHandler (http.ResponseWriter, *http.Request, *request.Image){
}
func VoiceRequestHandler (http.ResponseWriter, *http.Request, *request.Voice){
}
func VoiceRecognitionRequestHandler (http.ResponseWriter, *http.Request, *request.VoiceRecognition){
}
func VideoRequestHandler (http.ResponseWriter, *http.Request, *request.Video){
}
func LocationRequestHandler (http.ResponseWriter, *http.Request, *request.Location){
}
func LinkRequestHandler (http.ResponseWriter, *http.Request, *request.Link){
}
func SubscribeEventHandler (http.ResponseWriter, *http.Request, *request.SubscribeEvent){
}
func UnsubscribeEventHandler (http.ResponseWriter, *http.Request, *request.UnsubscribeEvent){
}
func SubscribeByScanEventHandler (http.ResponseWriter, *http.Request, *request.SubscribeByScanEvent){
}
func ScanEventHandler (http.ResponseWriter, *http.Request, *request.ScanEvent){
}
func LocationEventHandler (http.ResponseWriter, *http.Request, *request.LocationEvent){
}
func MenuClickEventHandler (http.ResponseWriter, *http.Request, *request.MenuClickEvent){
}
func MenuViewEventHandler (http.ResponseWriter, *http.Request, *request.MenuViewEvent){
}
func MassSendJobFinishEventHandler (http.ResponseWriter, *http.Request, *request.MassSendJobFinishEvent){
}
func MerchantOrderEventHandler (http.ResponseWriter, *http.Request, *request.MerchantOrderEvent){
}

