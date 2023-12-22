package config

type ConfigStruct struct {
	AppPort      string `yaml:"app_port"`
	Name         string `yaml:"name"`
	Env          string `yaml:"env"`
	Namespace    string `yaml:"namespace"`
	Master       string `yaml:"master"`
	CsiSecret    string `yaml:"csi_secret"`
	CsiNamespace string `yaml:"csi_namespace"`
	CsiDriver    string `yaml:"csi_driver"`
	Capacity     string `yaml:"capacity"`
	K8sUrl       string `yaml:"k8s_url"`
}

var Config ConfigStruct
