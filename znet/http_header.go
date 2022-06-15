package znet

const (
	AcCtrlOrigin       = "Access-Control-Allow-Origin"  //value maybe "*", "*.sohu.com"
	AcCtrlHeaders      = "Access-Control-Allow-Headers" //"DNT"
	AcCtrlMethods      = "Access-Control-Allow-Methods" //"GET, POST, OPTIONS"
	ContentType        = "content-type"                 //below is value
	ContentValueJson   = "json"
	ContentValueBytes  = "bytes"
	ContentValuePlain  = "plain"
	ContentValueStream = "stream"
	//sele defined keys in header of http
	ZkCmd       = "zk-cmd"       //int64 ReqCmd in proto
	ZkTimestamp = "zk-timestamp" //int64 timestamp to UTC
	ZkJwt       = "zk-jwt"       //string
	ZkCode      = "zk-code"      //int64
	ZkError     = "zk-error"     //string
	ZkHeader    = "zk-header"    //string Base64 Header in proto
	ZkTailer    = "zk-tailer"    //string Base64 Tailer in proto
)

type ReqBodyType = interface{}
type ResBodyType = interface{}

type Command struct {
	Cmd        int
	Path       string
	MethodName string
	// 要用函数的原因是每次都要生成一个对象，不能用共享的一个对象
	FnRequestBody func() ReqBodyType
}
