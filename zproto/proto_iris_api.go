package zproto

import (
	"encoding/base64"
	"fmt"
	"reflect"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/gozk/zlogger"
	"github.com/nurozhikun/gozk/znet"
	"github.com/nurozhikun/gozk/zproto/zpbf"
	"github.com/nurozhikun/gozk/zutils"
)

type ProtoMethodHandler interface {
	// ReqBodyOfCmd(cmd int) Message
}

// type ReqBodyType = Message
// type ResBodyType = Message

type Command struct {
	Cmd        int
	Path       string
	MethodName string
	// 要用函数的原因是每次都要生成一个对象，不能用共享的一个对象
	CreateRequestBody func() Message
}

// 保证Handler对象在整个过程中有效就行了
type ProtoApiParty struct {
	PartyUrl string //
	Handler  ProtoMethodHandler
	Cmds     map[int]*Command
}

type ProtoApiFunc = func(h *zpbf.Header, req Message) (Message, error)

var (
	ApiUseBase64 = false
)

func (pap *ProtoApiParty) InstallToApp(app *iris.Application) error {
	if nil == app {
		return zutils.ErrNullParam
	}
	party := app.Party(pap.PartyUrl)
	for _, cmd := range pap.Cmds {
		if len(cmd.Path) == 0 || len(cmd.MethodName) == 0 {
			continue
		}
		fn := HandleByFuncName(pap.Handler, cmd.MethodName)
		if nil == fn {
			//LOGO
			continue
		}
		reqBody := cmd.CreateRequestBody()
		if reqBody == nil {
			reqBody = &EmptyMessage{}
		}
		ctxFn := func(ctx znet.IrisCtx) {
			var resMsg Message
			h, err := ParserHeader(ctx) //get header
			// zlogger.Info("Parse header:", h, err)
			defer func() {
				SetHeader(ctx, h, err)
				if nil != resMsg && nil == err {
					ctx.Text(MarshalString(resMsg))
				} else {
					ctx.Text(err.Error())
				}
				// zlogger.Info("request ack with error:", err)
			}()
			if nil != err {
				zlogger.Error("Iris api error in Parser Header:", err)
				return
			}
			//has header
			if err = ParseBody(ctx, reqBody); err != nil {
				zlogger.Error("Iris api error in ParseBody:", err)
				return
			}
			//has body
			resMsg, err = fn(h, reqBody)
			if nil != err {
				zlogger.Error("Iris api error in fn:", err)
				return
			}
		}
		party.Post(cmd.Path, ctxFn)
		party.Get(cmd.Path, ctxFn)
		zlogger.Info("install api", pap.PartyUrl, cmd.Path)
	}
	return nil
}

func HandleByFuncName(api interface{}, method string) ProtoApiFunc {
	if nil == api {
		return nil
	}
	value := reflect.ValueOf(api)
	if !value.IsValid() {
		return nil
	}
	vf := value.MethodByName(method)
	var apiFunc ProtoApiFunc
	if vf.Type() != reflect.TypeOf(apiFunc) {
		return nil
	}
	return func(h *zpbf.Header, req Message) (Message, error) {
		ins := make([]reflect.Value, 2)
		ins[0] = reflect.ValueOf(h)
		ins[1] = reflect.ValueOf(req)
		out := vf.Call(ins)
		if len(out) < 2 {
			return nil, zutils.NewError(-1, fmt.Sprintf("the method(%s)'s response return is wrrong", method))
		}
		msg, _ := out[0].Interface().(Message)
		err, _ := out[1].Interface().(error)
		return msg, err
	}
}

func ParserHeader(ctx znet.IrisCtx) (header *zpbf.Header, err error) {
	header = &zpbf.Header{}
	s := ctx.GetHeader(znet.ZkHeader)
	if len(s) > 0 {
		UnmarshalString([]byte(s), header)
	}
	header.Cmd, _ = zutils.InterfaceToInt(ctx.GetHeader(znet.ZkCmd))
	header.Timestamp, _ = zutils.InterfaceToInt(ctx.GetHeader(znet.ZkTimestamp))
	header.Jwt = ctx.GetHeader(znet.ZkJwt)
	header.Code, _ = zutils.InterfaceToInt(ctx.GetHeader(znet.ZkCode))
	header.Error = ctx.GetHeader(znet.ZkHeader)
	return
}

func SetHeader(ctx znet.IrisCtx, h *zpbf.Header, err error) {
	if nil != h {
		if nil != err {
			h.Code = int64(zutils.ErrorCode(err))
			h.Error = err.Error()
		}
		h.Timestamp = time.Now().UTC().UnixMilli()
		if ApiUseBase64 {
			js := MarshalString(h)
			s := base64.StdEncoding.EncodeToString([]byte(js))
			ctx.Header(znet.ZkHeader, s)
		}
		ctx.Header(znet.ZkCmd, zutils.StringFromInterface(h.Cmd))
		ctx.Header(znet.ZkTimestamp, zutils.StringFromInterface(h.Timestamp))
		ctx.Header(znet.ZkJwt, h.Jwt)
		ctx.Header(znet.ZkCode, zutils.StringFromInterface(h.Code))
		ctx.Header(znet.ZkError, h.Error)
	} else {
		CopyHeader(ctx)
		if nil != err {
			ctx.Header(znet.ZkCode, zutils.StringFromInterface(zutils.ErrorCode(err)))
			ctx.Header(znet.ZkError, err.Error())
		}
	}
}

func CopyHeader(ctx znet.IrisCtx) {
	znet.IrisCopyHeaderKeys(ctx, znet.ZkCmd, znet.ZkJwt)
	t := time.Now().UTC().UnixMilli()
	ctx.Header(znet.ZkTimestamp, zutils.StringFromInterface(t))
}

func ParseBody(ctx znet.IrisCtx, msg Message) error {
	bs, err := ctx.GetBody()
	// zlogger.Info("GetBody", bs, err)
	if nil != err {
		return err
	}
	if len(bs) == 0 {
		return nil
	}
	return UnmarshalString(bs, msg)
}
