package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"model-k8s-service-go/config"
)

var Pvc pvc

type pvc struct{}

type PvcsResp struct {
	Items []corev1.PersistentVolumeClaim `json:"items"`
	Total int                            `json:"total"`
}

type PvcCreate struct {
	Name       string            `json:"name"`
	Namespace  string            `json:"namespace"`
	Label      map[string]string `json:"label"`
	VolumeName string            `json:"volume_name"`
}

// 创建pvc
func (p *pvc) CreatePvc(data *PvcCreate) (*corev1.PersistentVolumeClaim, error) {
	storageClassName := ""
	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes:      []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
			StorageClassName: &storageClassName,
			VolumeName:       data.VolumeName,
		},
	}

	pvc.Spec.Resources.Requests = map[corev1.ResourceName]resource.Quantity{
		corev1.ResourceStorage: resource.MustParse(config.Config.Capacity),
	}

	return K8s.ClientSet.CoreV1().PersistentVolumeClaims(data.Namespace).Create(context.TODO(), pvc, metav1.CreateOptions{})
}

// 获取pvc详情
func (p *pvc) GetPvcDetail(pvcName, namespace string) (pvc *corev1.PersistentVolumeClaim, err error) {
	pvc, err = K8s.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Get(context.TODO(), pvcName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Pvc详情失败, " + err.Error()))
		return nil, errors.New("获取Pvc详情失败, " + err.Error())
	}
	return pvc, nil
}

// 删除pvc
func (p *pvc) DeletePvc(pvcName, namespace string) (err error) {
	err = K8s.ClientSet.CoreV1().PersistentVolumeClaims(namespace).Delete(context.TODO(), pvcName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Pvc失败, " + err.Error()))
		return errors.New("删除Pvc失败, " + err.Error())
	}
	return nil
}
