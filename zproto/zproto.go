package zproto

import (
	"gitee.com/sienectagv/gozk/znet"
	"github.com/kataras/iris/v12"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

var (
	defMarshalJson   *protojson.MarshalOptions
	defUnmarshalJson *protojson.UnmarshalOptions
)

func init() {
	defMarshalJson = &protojson.MarshalOptions{UseProtoNames: false,
		EmitUnpopulated: true,
		AllowPartial:    true}
	defUnmarshalJson = &protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
}

// type Message interface {
// 	proto.Message
// 	// Marshal() (data []byte, err error)
// }
type Message = proto.Message

func UnmarshalString(bs []byte, msg proto.Message) error {
	return defUnmarshalJson.Unmarshal(bs, msg)
}

func MarshalString(msg Message) string {
	return defMarshalJson.Format(msg)
}

func ReadContext(ctx iris.Context, msg Message) error {
	contentType := ctx.GetHeader(znet.ContentType)
	bs, _ := ctx.GetBody()
	if contentType == znet.ContentValueJson {
		o := &protojson.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		}
		return o.Unmarshal(bs, msg)
	} else {
		o := &proto.UnmarshalOptions{
			AllowPartial:   true,
			DiscardUnknown: true,
		}
		return o.Unmarshal(bs, msg)
	}
}

func WriteContext(ctx iris.Context, msg Message, contentType string) (int, error) {
	ctx.Header(znet.ContentType, contentType)
	if contentType == znet.ContentValueJson || contentType == znet.ContentValuePlain {
		o := protojson.MarshalOptions{UseProtoNames: false,
			EmitUnpopulated: true,
			AllowPartial:    true}
		return ctx.Text(o.Format(msg))
	} else {
		o := proto.MarshalOptions{AllowPartial: true}
		bs, _ := o.Marshal(msg)
		return ctx.Binary(bs)
	}
}
