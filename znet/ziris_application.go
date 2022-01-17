package znet

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/core/router"
)

type IrisApp = iris.Application
type IrisCtx = context.Context
type IrisParty = router.Party

func IrisCopyHeaderKeys(ctx IrisCtx, keys ...string) {
	for _, k := range keys {
		ctx.Header(k, ctx.GetHeader(k))
	}
}
