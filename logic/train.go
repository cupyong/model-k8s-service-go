package logic

import (
	"github.com/pkg/errors"
	"model-k8s-service-go/config"
	"model-k8s-service-go/service"
	"model-k8s-service-go/util"
	"strconv"
)

type TrainDto struct {
	Image    string    `json:"image"`
	CPU      int       `json:"cpu"`
	GPU      int       `json:"gpu"`
	Memory   int       `json:"memory"`
	Storages []Storage `json:"storages"`
	Command  []string  `json:"command"`
}

type TrainResult struct {
	Name string `json:"name"`
}

func (l logic) CreateTrain(dto TrainDto) (interface{}, error) {
	name := util.SetName("job" + "-" + util.GetDate())
	labels := make(map[string]string, 0)
	labels["name"] = name
	volumes := make([]service.Volume, 0)
	//step1 pv、pvc创建
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

	//step2 job创建
	var jobCreate service.JobCreate
	jobCreate.Namespace = config.Config.Namespace
	jobCreate.Name = name
	jobCreate.Label = labels
	jobCreate.Memory = strconv.Itoa(dto.Memory) + "G"
	jobCreate.Cpu = strconv.Itoa(dto.CPU)
	jobCreate.Image = dto.Image
	jobCreate.Command = dto.Command
	jobCreate.Volumes = volumes
	err := service.Job.CreateJob(&jobCreate)

	if err != nil {
		return "", err
	}
	return TrainResult{
		name,
	}, nil
}

func (l logic) DeleteTrain(name string) (interface{}, error) {
	//step1 删除关联pv、pvc
	job, err := service.Job.GetJobDetail(name, config.Config.Namespace)
	for _, v := range job.Spec.Template.Spec.Volumes {
		service.Pvc.DeletePvc(v.PersistentVolumeClaim.ClaimName, config.Config.Namespace)
		service.Pv.DeletePv(v.PersistentVolumeClaim.ClaimName)
	}
	if err != nil {
		return "", err
	}

	//step2 删除关联pod
	podList, err := service.Pod.GetPodList(name, config.Config.Namespace)
	for _, v := range podList {
		service.Pod.DeletePod(v.Name, config.Config.Namespace)
	}

	//step2 删除关联job
	err = service.Job.DeleteJob(name, config.Config.Namespace)
	if err != nil {
		return "", err
	}
	return true, nil
}

func (l logic) GetSingleStatus(name string) (interface{}, error) {
	podList, err := service.Pod.GetPodList(name, config.Config.Namespace)
	if err != nil {
		return "", err
	}
	if len(podList) < 1 {
		return "", errors.New("不存在该训练任务")
	}
	return service.ControllerData{
		Name:   name,
		Status: podList[0].Status,
	}, nil
}

func (l logic) GetLogs(name, line string) (interface{}, error) {
	podList, err := service.Pod.GetPodList(name, config.Config.Namespace)
	if err != nil {
		return "", err
	}
	for _, v := range podList {
		service.Pod.DeletePod(v.Name, config.Config.Namespace)
	}
	if len(podList) < 1 {
		return "", errors.New("该服务不存在")
	}
	i, err := strconv.Atoi(line)
	if err != nil {
		return "", errors.New("line必须是数字")
	}
	return service.Pod.GetPodLog(i, podList[0].Name, config.Config.Namespace)
}

func (l logic) GetListStatus(names []string) (interface{}, error) {
	return service.Pod.GetControllerList(names, config.Config.Namespace)
}
