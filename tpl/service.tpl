package {{.packageName}}

import (
	"git.zuoyebang.cc/huixuexi/classflow/layer"
)

type {{.methodName}}Service struct {
	layer.Service
}

func (entity *{{.methodName}}Service) OnCreate(param layer.IFlowParam) {
	entity.Service.OnCreate(param)
}

type {{.methodName}}Req struct {}

type {{.methodName}}Res struct {}

func (entity *{{.methodName}}Service) Handle(req *{{.methodName}}Req) (res *{{.methodName}}Res, err error) {
    return
}