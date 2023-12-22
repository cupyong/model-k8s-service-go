package service

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/wonderivan/logger"
	"io"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

var Pod pod

type pod struct{}

type PodsNp struct {
	Namespace string
	PodNum    int
}

type PodData struct {
	Name   string          `json:"name"`
	Status corev1.PodPhase `json:"status"`
}

// 获取pod详情
func (p *pod) GetPodDetail(podName, namespace string) (pod *corev1.Pod, err error) {
	pod, err = K8s.ClientSet.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
	if err != nil {
		logger.Error("获取Pod详情失败，" + err.Error())
		return nil, errors.New("获取Pod详情失败，" + err.Error())
	}
	return pod, nil
}

// 根据控制器获取pod名称
func (p *pod) GetPodList(controller, namespace string) ([]PodData, error) {
	list := make([]PodData, 0)

	podList1, err := K8s.ClientSet.CoreV1().Events(namespace).List(context.TODO(), metav1.ListOptions{})
	b, _ := json.Marshal(podList1)
	fmt.Println(string(b))

	podList, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取Pod列表失败，" + err.Error())
		return list, errors.New("获取Pod列表失败，" + err.Error())
	}
	for _, v := range podList.Items {
		if strings.HasPrefix(v.Name, controller) {
			list = append(list, PodData{
				Name:   v.Name,
				Status: v.Status.Phase,
			})
		}
	}
	return list, nil
}

// 根据控制器获取pod名称
func (p *pod) GetControllerList(controllers []string, namespace string) ([]ControllerData, error) {
	list := make([]ControllerData, 0)
	podList, err := K8s.ClientSet.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		logger.Info("获取Pod列表失败，" + err.Error())
		return list, errors.New("获取Pod列表失败，" + err.Error())
	}
	for _, v := range podList.Items {
		for _, c := range controllers {
			if strings.HasPrefix(v.Name, c) {
				list = append(list, ControllerData{
					Name:   c,
					Status: v.Status.Phase,
				})
			}
		}
	}
	return list, nil
}

// 删除Pod
func (p *pod) DeletePod(podName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().Pods(namespace).Delete(context.TODO(), podName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error("删除Pod详情失败，" + err.Error())
		return errors.New("删除Pod详情失败，" + err.Error())
	}
	return nil
}

// 获取Pod中的容器名
func (p *pod) GetPodContainer(podName string, namespace string) (containers []string, err error) {
	pod, err := p.GetPodDetail(podName, namespace)
	if err != nil {
		return nil, err
	}
	for _, container := range pod.Spec.Containers {
		containers = append(containers, container.Name)
	}
	return containers, nil
}

// 获取Pod内容器日志
func (p *pod) GetPodLog(num int, podName, namespace string) (log string, err error) {
	// 设置日志配置，容器名，获取内容的配置
	lineLimit := int64(num)
	option := &corev1.PodLogOptions{
		//Container: containerName,
		TailLines: &lineLimit,
	}
	// 获取一个request实例
	req := K8s.ClientSet.CoreV1().Pods(namespace).GetLogs(podName, option)
	// 发起stream连接，获取到Response.body
	podLogs, err := req.Stream(context.TODO())
	if err != nil {
		return "", errors.New("连接失败," + err.Error())
	}
	defer podLogs.Close()
	// 将Response.body写入到缓存区，目的为了转换成string类型
	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, podLogs)
	if err != nil {
		logger.Error("复制podLog失败，" + err.Error())
		return "", errors.New("复制podLog失败，" + err.Error())
	}
	return buf.String(), nil
}
