package logic

import (
	"model-k8s-service-go/config"
	"model-k8s-service-go/service"
	"model-k8s-service-go/util"
	"strconv"
)

type ServiceDto struct {
	Port     int32     `json:"port"`
	Image    string    `json:"image"`
	CPU      int       `json:"cpu"`
	GPU      int       `json:"gpu"`
	Memory   int       `json:"memory"`
	Storages []Storage `json:"storages"`
}

type Storage struct {
	MountPath string         `json:"mountPath"`
	Nfs       service.NFS    `json:"nfs"`
	MiniIo    service.MiniIo `json:"miniio"`
}

type ServiceResult struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}

type logic struct{}

var Logic logic

/*
 *创建服务
 */
func (l logic) CreateService(dto ServiceDto) (interface{}, error) {
	name := util.SetName("service" + "-" + util.GetDate())
	labels := make(map[string]string, 0)
	labels["name"] = name

	volumes := make([]service.Volume, 0)
	// step1 pv、pvc创建
	for i, v := range dto.Storages {
		var pvCreate service.PvCreate
		pvCreate.Name = name + "-" + strconv.Itoa(i)
		pvCreate.Nfs = v.Nfs
		pvCreate.MiniIo = v.MiniIo
		_, err := service.Pv.CreatePv(&pvCreate)
		if err != nil {
			return "", err
		}
		//step2 pvc创建
		var pvcCreate service.PvcCreate
		pvcCreate.Name = name + "-" + strconv.Itoa(i)
		pvcCreate.Namespace = config.Config.Namespace
		pvcCreate.VolumeName = name + "-" + strconv.Itoa(i)
		_, err = service.Pvc.CreatePvc(&pvcCreate)
		if err != nil {
			return "", err
		}
		volumes = append(volumes, service.Volume{
			ClaimName: name + "-" + strconv.Itoa(i),
			MountPath: v.MountPath,
		})
	}

	//step2 deloy创建
	var deployCreate service.DeployCreate
	deployCreate.Namespace = config.Config.Namespace
	deployCreate.Cpu = strconv.Itoa(dto.CPU)
	deployCreate.Image = dto.Image
	deployCreate.Memory = strconv.Itoa(dto.Memory) + "G"
	deployCreate.Name = name
	deployCreate.Replicas = 1
	deployCreate.ContainerPort = dto.Port
	deployCreate.Label = labels
	deployCreate.Gpu = strconv.Itoa(dto.GPU)
	deployCreate.Volumes = volumes

	//step3 svc创建
	err := service.Deployment.CreateDeployment(&deployCreate)
	if err != nil {
		return "", err
	}
	var serviceCreate service.ServiceCreate
	serviceCreate.Namespace = config.Config.Namespace
	serviceCreate.Label = labels
	serviceCreate.Name = name
	serviceCreate.Port = dto.Port
	serviceCreate.Type = "NodePort"
	serviceCreate.ContainerPort = dto.Port

	svc, err := service.Service.CreateService(&serviceCreate)
	if err != nil {
		return "", err
	}
	return ServiceResult{
		"http://" + config.Config.Master + ":" + util.Int32ToString(svc.Spec.Ports[0].NodePort),
		name,
	}, nil
}

func (l logic) DeleteService(name string) (interface{}, error) {
	//step1 删除关联pv、pvc
	deploy, err := service.Deployment.GetDeploymentDetail(name, config.Config.Namespace)
	for _, v := range deploy.Spec.Template.Spec.Volumes {
		service.Pvc.DeletePvc(v.PersistentVolumeClaim.ClaimName, config.Config.Namespace)
		service.Pv.DeletePv(v.PersistentVolumeClaim.ClaimName)
	}
	if err != nil {
		return "", err
	}
	//step2 删除关联deploy
	err = service.Deployment.DeleteDeployment(name, config.Config.Namespace)
	if err != nil {
		return "", err
	}
	//step3 删除关联svc
	err = service.Service.DeleteService(name, config.Config.Namespace)
	if err != nil {
		return "", err
	}
	return true, nil
}
