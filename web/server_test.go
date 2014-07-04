package web

import (
	"testing"
)

func Test_AddRoute(t *testing.T) {
	sv := &Server{}
	beforeCount := len(sv.routes)
	sv.AddRoute("/",HandleFunc(func(ctx *Context){}),"Get")
	afterCount := len(sv.routes)
	if afterCount > beforeCount {
		t.Log("添加路由测试通过")
	} else {
		t.Error("添加路由测试未通过")
	}
}

