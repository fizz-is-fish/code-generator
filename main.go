package main

import (
	"encoding/json"
	"go-code-generator/go_template"
	"go-code-generator/utils"
)

// 配置项，classflow只需修改此项即可
var (
	projectPath = "/Users/missfizz/go/src/hxx-aiclass-course"
	packageName = "wrongQuestion"
	fileName    = "list"
)

// classflow默认路径，可根据项目实际情况修改
var (
	controllerDirectory        = "ctl_" + packageName
	controllerPath             = projectPath + "/controllers/http/" + controllerDirectory
	controllerDataFileName     = ""
	controllerTemplateFileName = "tpl/controller.tpl"

	serviceDirectory        = "svc_" + packageName
	servicePath             = projectPath + "/service/" + serviceDirectory
	serviceDataFileName     = ""
	serviceTemplateFileName = "tpl/sercice.tpl"
)

func main() {
	genController()
}

func genController() {
	data := struct {
		PackageName string `json:"packageName"`
		MethodName  string `json:"methodName"`
	}{
		PackageName: controllerDirectory,
		MethodName:  utils.FirstUpper(fileName),
	}
	byteValue, _ := json.Marshal(data)
	go_template.GoTemplate(&controllerDataFileName, &controllerTemplateFileName, &controllerPath, &fileName, nil, byteValue)
}
