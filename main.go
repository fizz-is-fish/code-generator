package main

import (
	"encoding/json"
	"go-code-generator/go_template"
	"go-code-generator/utils"
)

// 配置项，classflow只需修改此项即可
var (
	projectPath = "/Users/missfizz/go/src/bawu"
	projectName = "bawu"
	packageName = "school"
	methodName  = "list"
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
	serviceTemplateFileName = "tpl/service.tpl"
)

func main() {
	genController()
	genService()
}

func genController() {
	data := struct {
		PackageName    string `json:"packageName"`
		MethodName     string `json:"methodName"`
		ServicePackage string `json:"servicePackage"`
		ProjectName    string `json:"projectName"`
	}{
		PackageName:    controllerDirectory,
		MethodName:     utils.FirstUpper(methodName),
		ServicePackage: serviceDirectory,
		ProjectName:    projectName,
	}
	byteValue, _ := json.Marshal(data)
	go_template.GoTemplate(&controllerDataFileName, &controllerTemplateFileName, &controllerPath, &methodName, nil, byteValue)
}

func genService() {
	data := struct {
		PackageName string `json:"packageName"`
		MethodName  string `json:"methodName"`
	}{
		PackageName: serviceDirectory,
		MethodName:  utils.FirstUpper(methodName),
	}
	byteValue, _ := json.Marshal(data)
	go_template.GoTemplate(&serviceDataFileName, &serviceTemplateFileName, &servicePath, &methodName, nil, byteValue)
}
