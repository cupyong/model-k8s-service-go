package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

var Deployment deployment

type deployment struct{}

// 定义DeployCreate结构体
type DeployCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Replicas      int32             `json:"replicas"`
	Image         string            `json:"image"`
	Label         map[string]string `json:"label"`
	Cpu           string            `json:"cpu"`
	Gpu           string            `json:"gpu"`
	Memory        string            `json:"memory"`
	ContainerPort int32             `json:"container_port"`
	Volumes       []Volume          `json:"volumes"`
}

type ControllerData struct {
	Name   string          `json:"name"`
	Status corev1.PodPhase `json:"status"`
}

type Volume struct {
	ClaimName string
	MountPath string
}

// 获取deployment详情
func (d *deployment) GetDeploymentDetail(deploymentName, namespace string) (deployment *appsv1.Deployment, err error) {
	deployment, err = K8s.ClientSet.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Deployment详情失败, " + err.Error()))
		return nil, errors.New("获取Deployment详情失败, " + err.Error())
	}
	return deployment, nil
}

// 创建deployment,接收DeployCreate对象
func (d *deployment) CreateDeployment(data *DeployCreate) (err error) {
	//将data中的属性组装成appsv1.Deployment对象,并将入参数据放入
	deployment := &appsv1.Deployment{
		//ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		//Spec中定义副本数、选择器、以及pod属性
		Spec: appsv1.DeploymentSpec{
			Replicas: &data.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: data.Label,
			},
			Template: corev1.PodTemplateSpec{
				//定义pod名和标签
				ObjectMeta: metav1.ObjectMeta{
					Name:   data.Name,
					Labels: data.Label,
				},
				//定义容器名、镜像和端口
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  data.Name,
							Image: data.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									Protocol:      corev1.ProtocolTCP,
									ContainerPort: data.ContainerPort,
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse(data.Cpu),
									corev1.ResourceMemory: resource.MustParse(data.Memory),
								},
								Requests: map[corev1.ResourceName]resource.Quantity{
									corev1.ResourceCPU:    resource.MustParse(data.Cpu),
									corev1.ResourceMemory: resource.MustParse(data.Memory),
								},
							},
						},
					},
				},
			},
		},
		Status: appsv1.DeploymentStatus{},
	}

	volumeMounts := make([]corev1.VolumeMount, 0)
	volumes := make([]corev1.Volume, 0)
	for i, v := range data.Volumes {

		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      data.Name + "-" + strconv.Itoa(i),
			MountPath: v.MountPath,
		})

		volumes = append(volumes, corev1.Volume{
			Name: data.Name + "-" + strconv.Itoa(i),
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: v.ClaimName,
					ReadOnly:  false,
				},
			},
		})
	}
	deployment.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
	deployment.Spec.Template.Spec.Volumes = volumes
	//调用sdk创建deployment
	_, err = K8s.ClientSet.AppsV1().Deployments(data.Namespace).Create(context.TODO(), deployment, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建Deployment失败, " + err.Error()))
		return errors.New("创建Deployment失败, " + err.Error())
	}
	return nil
}

// 删除deployment
func (d *deployment) DeleteDeployment(deploymentName, namespace string) (err error) {
	err = K8s.ClientSet.AppsV1().Deployments(namespace).Delete(context.TODO(), deploymentName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Deployment失败, " + err.Error()))
		return errors.New("删除Deployment失败, " + err.Error())
	}
	return nil
}
