package service

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var Service service

type service struct{}

type ServicesResp struct {
	Items []corev1.Service `json:"items"`
	Total int              `json:"total"`
}

type ServiceCreate struct {
	Name          string            `json:"name"`
	Namespace     string            `json:"namespace"`
	Type          string            `json:"type"`
	ContainerPort int32             `json:"container_port"`
	Port          int32             `json:"port"`
	NodePort      int32             `json:"node_port"`
	Label         map[string]string `json:"label"`
}

// 获取service详情
func (s *service) GetServicetDetail(serviceName, namespace string) (service *corev1.Service, err error) {
	service, err = K8s.ClientSet.CoreV1().Services(namespace).Get(context.TODO(), serviceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Service详情失败, " + err.Error()))
		return nil, errors.New("获取Service详情失败, " + err.Error())
	}

	return service, nil
}

// 创建service,,接收ServiceCreate对象
func (s *service) CreateService(data *ServiceCreate) (*corev1.Service, error) {
	//将data中的数据组装成corev1.Service对象
	service := &corev1.Service{
		//ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		//Spec中定义类型，端口，选择器
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(data.Type),
			Ports: []corev1.ServicePort{
				{
					Name:     "http",
					Port:     data.Port,
					Protocol: "TCP",
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: data.ContainerPort,
					},
				},
			},
			Selector: data.Label,
		},
	}
	service.Spec.Ports[0].Port = data.Port
	//创建Service
	svc, err := K8s.ClientSet.CoreV1().Services(data.Namespace).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建Service失败, " + err.Error()))
		return svc, errors.New("创建Service失败, " + err.Error())
	}

	return svc, nil
}

// 删除service
func (s *service) DeleteService(serviceName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Services(namespace).Delete(context.TODO(), serviceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Service失败, " + err.Error()))
		return errors.New("删除Service失败, " + err.Error())
	}

	return nil
}

// 更新service
func (s *service) UpdateService(namespace, content string) (err error) {
	var service = &corev1.Service{}

	err = json.Unmarshal([]byte(content), service)
	if err != nil {
		logger.Error(errors.New("反序列化失败, " + err.Error()))
		return errors.New("反序列化失败, " + err.Error())
	}

	_, err = K8s.ClientSet.CoreV1().Services(namespace).Update(context.TODO(), service, metav1.UpdateOptions{})
	if err != nil {
		logger.Error(errors.New("更新service失败, " + err.Error()))
		return errors.New("更新service失败, " + err.Error())
	}
	return nil
}
