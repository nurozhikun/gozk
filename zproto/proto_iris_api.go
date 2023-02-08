package zproto

import "github.com/kataras/iris/v12"

type ProtoMethodHandler interface {
	ReqBodyOfCmd(cmd int) Message
}

type ProtoApiCmd struct {
	SubUrl     string //Url 子路径
	MethodName string //对应的处理函数的名称
}

type ProtoApiParty struct {
	PartyUrl string //
	Handler  ProtoMethodHandler
	Cmds     map[int]*ProtoApiCmd
}

func (pap *ProtoApiParty) InstallToApp(app *iris.Application) error {
	return nil
}
