package zproto

import (
	"github.com/kataras/iris/v12"
	"github.com/nurozhikun/gozk/znet"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	defMarshalJson   *protojson.MarshalOptions
	defUnmarshalJson *protojson.UnmarshalOptions
)

func init() {
	defMarshalJson = &protojson.MarshalOptions{
		UseProtoNames:   false,
		EmitUnpopulated: true,
		AllowPartial:    true,
	}
	defUnmarshalJson = &protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}
}

// type Message interface {
// 	proto.Message
// 	// Marshal() (data []byte, err error)
// }

// TODO 这样自定义的EmptyMessage使用时会panic 已在protbee中重新定义
type Message = proto.Message

type EmptyMessage struct{}

func (*EmptyMessage) ProtoMessage() {}

func (*EmptyMessage) ProtoReflect() protoreflect.Message {
	return nil
}

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
		o := protojson.MarshalOptions{
			UseProtoNames:   false,
			EmitUnpopulated: true,
			AllowPartial:    true,
		}
		return ctx.Text(o.Format(msg))
	} else {
		o := proto.MarshalOptions{AllowPartial: true}
		bs, _ := o.Marshal(msg)
		return ctx.Binary(bs)
	}
}
