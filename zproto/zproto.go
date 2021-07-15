package zproto

import (
	"gitee.com/sienectagv/gozk/znet"
	"github.com/kataras/iris/v12"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Message interface {
	proto.Message
	// Marshal() (data []byte, err error)
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
	if contentType == znet.ContentValueJson {
		o := protojson.MarshalOptions{UseProtoNames: true,
			EmitUnpopulated: true,
			AllowPartial:    true}
		// return o.Format(msg), nil
		return ctx.Text(o.Format(msg))
	} else {
		o := proto.MarshalOptions{AllowPartial: true}
		bs, _ := o.Marshal(msg)
		return ctx.Binary(bs)
	}
}
