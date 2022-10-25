package go_template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Masterminds/sprig"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"
)

type MultipleFileData struct {
	Files []struct {
		FileName string
		Data     interface{}
	}
}

func GoTemplate(dataFileName, templateFileName, outputDirectory, outputFile, multipleFiles *string, byteValue []byte) {

	hasMultipleFiles := (multipleFiles != nil) && (*multipleFiles != "") && (*multipleFiles == "multi")

	template, err := template.New(path.Base(*templateFileName)).Funcs(template.FuncMap(MyFuncMap)).Funcs(sprig.TxtFuncMap()).ParseFiles(*templateFileName)
	if err != nil {
		fmt.Println(err)
	}

	if byteValue == nil || len(byteValue) == 0 {
		var dataFile *os.File
		dataFile, err = os.Open(*dataFileName)
		if err != nil {
			fmt.Println(err)
		}

		defer dataFile.Close()

		byteValue, err = ioutil.ReadAll(dataFile)
		if err != nil {
			fmt.Println(err)
		}
		byteValue, err = ToJSON(byteValue)
	}

	if hasMultipleFiles {
		var multidata MultipleFileData

		err = json.Unmarshal(byteValue, &multidata)
		if err != nil {
			fmt.Println(err)
		}

		for _, mfile := range multidata.Files {
			outputFileName := path.Join(path.Dir(*dataFileName), mfile.FileName)
			mdata := mfile.Data
			generateFile(template, *outputDirectory, outputFileName, mdata)
		}

	} else {
		var data interface{}
		err = json.Unmarshal(byteValue, &data)
		if err != nil {
			fmt.Println(err)
		}
		//判断传入文件名
		if *outputFile == "" {
			tmp := strings.TrimSuffix(*dataFileName, filepath.Ext(*dataFileName))
			outputFile = &tmp
		}
		outputFileName := *outputFile + ".go"
		generateFile(template, *outputDirectory, outputFileName, data)
	}
}

func generateFile(template *template.Template, outputDirectory string, outputFileName string, data interface{}) {
	absOutputFileName := path.Join(outputDirectory, outputFileName)
	os.MkdirAll(path.Dir(absOutputFileName), os.ModePerm)
	outputFile, err := os.Create(absOutputFileName)
	fmt.Println("Generating file : " + absOutputFileName)
	defer outputFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = template.Execute(outputFile, data)
	if err != nil {
		fmt.Println(err)
	}
}

func ToJSON(data []byte) ([]byte, error) {
	if hasJSONPrefix(data) {
		return data, nil
	}
	return yaml.YAMLToJSON(data)
}

var jsonPrefix = []byte("{")

func hasJSONPrefix(buf []byte) bool {
	trim := bytes.TrimLeftFunc(buf, unicode.IsSpace)
	return bytes.HasPrefix(trim, jsonPrefix)
}
