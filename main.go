/*
-------------------------------------------------
   Author :       Zhang Fan
   date：         2020/12/4
   Description :
-------------------------------------------------
*/

package main

import (
	"fmt"
	"sort"

	jsoniter "github.com/json-iterator/go"
	"github.com/zly-app/service/api"
	"github.com/zly-app/zapp"
	"github.com/zly-app/zapp/core"
)

func valuesToText(values map[string][]string) string {
	var temp []string
	for k, vs := range values {
		for _, v := range vs {
			temp = append(temp, k+": "+v)
		}
	}
	sort.Strings(temp)
	text, _ := jsoniter.MarshalIndent(temp, "", "    ")
	return string(text)
}

func Process(ctx *api.Context) interface{} {
	ctx.Request().Header.Add("host", ctx.Request().Host)
	headerText := valuesToText(ctx.Request().Header)
	paramsText := valuesToText(ctx.Request().URL.Query())

	body, _ := ctx.GetBody()

	return fmt.Sprintf("ip: %s\npath:%s\nparams:%s\n\nheaders:\n%s\n\nbody:\n%s\n",
		ctx.RemoteAddr(),
		ctx.Path(),
		paramsText,
		headerText,
		string(body),
	)
}

func Router(c core.IComponent, route api.Party) {
	route.Any("/", api.Wrap(Process))
	route.Any("/{a:path}", api.Wrap(Process))
}

func main() {
	app := zapp.NewApp("http-print", api.WithService())
	api.RegistryRouter(Router)
	api.SetWriteResponseFunc(func(ctx *api.Context, code int, message string, data interface{}) {
		switch v := data.(type) {
		case string:
			_, _ = ctx.WriteString(v)
		default:
			_, _ = ctx.JSON(api.Response{
				ErrCode: code,
				ErrMsg:  message,
				Data:    data,
			})
		}
	})
	app.Run()
}
