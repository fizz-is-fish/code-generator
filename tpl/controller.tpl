package {{.packageName}}

import (
	"{{.projectName}}/service/{{.servicePackage}}"
	"git.zuoyebang.cc/huixuexi/classflow/layer"
	"github.com/gin-gonic/gin/binding"
)

type {{.methodName}}Controller struct {
	layer.Controller
	dto *{{.servicePackage}}.{{.methodName}}Req
}

func (entity *{{.methodName}}Controller) GetDtoRequest() interface{} {
	entity.dto = new({{.servicePackage}}.{{.methodName}}Req)
	return entity.dto
}

func (entity *{{.methodName}}Controller) GetBindingType() binding.Binding {
	return binding.Query
}

func (entity *{{.methodName}}Controller) Action() (interface{}, error) {
	svc := entity.Create(new({{.servicePackage}}.{{.methodName}}Service)).(*{{.servicePackage}}.{{.methodName}}Service)
	return svc.Handle(entity.dto)
}
