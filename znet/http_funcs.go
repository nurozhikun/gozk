package znet

import (
	"io"
	"net/http"

	"github.com/nurozhikun/gozk/zlogger"
)

func HttpGetFile(url string, fn func(body io.Reader)) error {
	zlogger.Info(url)
	res, err := http.Get(url)
	if nil != err {
		return err
	}
	if nil != res.Body {
		fn(res.Body)
		res.Body.Close()
	}
	return nil
}
