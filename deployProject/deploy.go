package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	logName      = "run.log"
	settingsPath = "deploy.yaml"
	zipName      = "deploy.zip"
)

var (
	info *log.Logger
)

type deploySettings struct {
	FileString string `yaml:"files"`
}

func initLogger(fileName string) {
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	info = log.New(logFile, "[info]", log.LstdFlags)
}

func (d *deploySettings) getConf(path string) *deploySettings {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		info.Printf("cannot read %s, errno %s.\n", path, err.Error())
	}

	err = yaml.Unmarshal(yamlFile, d)
	if err != nil {
		log.Fatal(err)
	}

	return d
}

func main() {
	initLogger(logName)

	f := deploySettings{}
	f.getConf(settingsPath)

	info.Println("items:", f.FileString)
	fileList := strings.Split(f.FileString, ",")
	info.Println("files:", fileList)
	generateZip(fileList, zipName)
}
