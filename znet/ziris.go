package znet

import "github.com/kataras/iris/v12/context"

func Cors(ctx context.Context) {
	// zlogger.Info("in cores ....")
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Headers", "DNT,X-Mx-ReqToken,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type,Authorization")
	ctx.Header("Access-Control-Allow-Methods", "GET,POST,OPTIONS,HEAD,DELETE")
	if ctx.Method() == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return

	}
	ctx.Next()
}

func IrisResHeader(ctx context.Context) {
	s := ctx.GetHeader(ZkCmd)
	if len(s) > 0 {
		ctx.Header(ZkCmd, s)
	}
	s = ctx.GetHeader(ZkTimestamp)
	if len(s) > 0 {
		ctx.Header(ZkTimestamp, s)
	}
	s = ctx.GetHeader(ZkHeader)
	if len(s) > 0 {
		ctx.Header(ZkHeader, s)
	}
}
