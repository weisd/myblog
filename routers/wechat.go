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
    beego.Info("xx")
    beego.Info(err)
}

// 非法请求
func UnknownRequestHandler(w http.ResponseWriter, r *http.Request, msg []byte){
    beego.Info("xx2")
    beego.Info(msg)
}

func TextRequestHandler(w http.ResponseWriter, r *http.Request, text *request.Text){
    beego.Info(text)
}

func ImageRequestHandler (w http.ResponseWriter, r *http.Request, img *request.Image){
}
func VoiceRequestHandler (w http.ResponseWriter, r *http.Request, voice *request.Voice){
}
func VoiceRecognitionRequestHandler (w http.ResponseWriter, r *http.Request, voiceRecognition *request.VoiceRecognition){
}
func VideoRequestHandler (w http.ResponseWriter, r *http.Request, video *request.Video){
}
func LocationRequestHandler (w http.ResponseWriter, r *http.Request, location *request.Location){
}
func LinkRequestHandler (w http.ResponseWriter, r *http.Request, link *request.Link){
}
func SubscribeEventHandler (w http.ResponseWriter, r *http.Request, event *request.SubscribeEvent){
}
func UnsubscribeEventHandler (w http.ResponseWriter, r *http.Request, event *request.UnsubscribeEvent){
}
func SubscribeByScanEventHandler (w http.ResponseWriter, r *http.Request, event *request.SubscribeByScanEvent){
}
func ScanEventHandler (w http.ResponseWriter, r *http.Request, event *request.ScanEvent){
}
func LocationEventHandler (w http.ResponseWriter, r *http.Request, event *request.LocationEvent){
}
func MenuClickEventHandler (w http.ResponseWriter, r *http.Request, event *request.MenuClickEvent){
}
func MenuViewEventHandler (w http.ResponseWriter, r *http.Request, event *request.MenuViewEvent){
}
func MassSendJobFinishEventHandler (w http.ResponseWriter, r *http.Request, event *request.MassSendJobFinishEvent){
}
func MerchantOrderEventHandler (w http.ResponseWriter, r *http.Request, event *request.MerchantOrderEvent){
}

