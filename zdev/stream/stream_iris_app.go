package stream

import (
	stdctx "context"

	"gitee.com/sienectagv/gozk/zdev/base"
	"gitee.com/sienectagv/gozk/zlogger"
	"gitee.com/sienectagv/gozk/zutils"
	"github.com/kataras/iris/v12"
)

func NewStrmIrisApp(app *iris.Application) *StrmIrisApp {
	return &StrmIrisApp{app: app}
}

type StrmIrisApp struct {
	StreamBase
	StrmEmptyWrite
	app *iris.Application
}

func (s *StrmIrisApp) ISetParams(cmd *base.Command) {
	s.StreamBase.ISetParams(cmd)
	if i, ok := cmd.BodyMap.Get(base.FieldIrisApp); ok {
		s.app = i.(*iris.Application)
	}
}

func (s *StrmIrisApp) IoOpen() error {
	if nil == s.app {
		return zutils.NewError(-1, "the iris.Application is nil")
	}
	return nil
}

func (s *StrmIrisApp) IoClose() error {
	if nil != s.app {
		s.app.Shutdown(stdctx.TODO())
	}
	return nil
}

func (s *StrmIrisApp) IoCanRead() bool {
	return nil != s.app
}

func (s *StrmIrisApp) IoRead() (bin interface{}, err error) {
	addr, ok := s.ParamsMap().GetString(base.FieldAddr)
	if !ok {
		return
	}
	zlogger.Info("virtual device:", s.ID(), "is listen on", addr)
	s.app.Listen(addr)
	return
}
