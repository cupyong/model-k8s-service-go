package initialize

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"model-k8s-service-go/config"
	"model-k8s-service-go/service"
	"os"
)

func Init(filename string) {
	config.Config = config.ConfigStruct{}
	service.InitClient()
	initConfig(filename)
}

func initConfig(filename string) {
	local := os.Getenv("local")
	if local != "false" {
		conf := new(config.ConfigStruct)
		yamlFile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Printf("yamlFile.Get err #%v ", err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}
		config.Config.AppPort = conf.AppPort
		config.Config.Name = conf.Name
		config.Config.Env = conf.Env
		config.Config.Namespace = conf.Namespace
		config.Config.Master = conf.Master
		config.Config.CsiSecret = conf.CsiSecret
		config.Config.CsiNamespace = conf.CsiNamespace
		config.Config.CsiDriver = conf.CsiDriver
		config.Config.Capacity = conf.Capacity
		config.Config.K8sUrl = conf.K8sUrl
	} else {
		config.Config.AppPort = os.Getenv("app_port")
		config.Config.Name = os.Getenv("name")
		config.Config.Env = os.Getenv("env")
		config.Config.Namespace = os.Getenv("namespace")
		config.Config.Master = os.Getenv("master")
		config.Config.CsiSecret = os.Getenv("csi_secret")
		config.Config.CsiNamespace = os.Getenv("csi_namespace")
		config.Config.CsiDriver = os.Getenv("csi_driver")
		config.Config.Capacity = os.Getenv("capacity")
		config.Config.K8sUrl = os.Getenv("k8s_url")
	}
}
