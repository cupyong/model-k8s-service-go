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

var Pv pv

type pv struct{}

type PvsResp struct {
	Items []corev1.PersistentVolume `json:"items"`
	Total int                       `json:"total"`
}

type PvCreate struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Label     map[string]string `json:"label"`
	Nfs       NFS               `json:"nfs"`
	MiniIo    MiniIo            `json:"miniio"`
}

type NFS struct {
	Path   string `json:"path"`
	Server string `json:"server"`
}

type MiniIo struct {
	Path string `json:"path"`
}

// 创建pv
func (p *pv) CreatePv(data *PvCreate) (*corev1.PersistentVolume, error) {
	pv := &corev1.PersistentVolume{
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: corev1.PersistentVolumeSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{corev1.ReadWriteMany},
			Capacity: map[corev1.ResourceName]resource.Quantity{
				corev1.ResourceStorage: resource.MustParse(config.Config.Capacity),
			},
		},
	}
	if len(data.Nfs.Path) > 0 {
		pv.Spec.NFS = &corev1.NFSVolumeSource{
			Path:   data.Nfs.Path,
			Server: data.Nfs.Server,
		}
	}
	if len(data.MiniIo.Path) > 0 {
		pv.Spec.CSI = &corev1.CSIPersistentVolumeSource{
			Driver: config.Config.CsiDriver,
			ControllerPublishSecretRef: &corev1.SecretReference{
				Name:      config.Config.CsiSecret,
				Namespace: config.Config.CsiNamespace,
			},
			NodePublishSecretRef: &corev1.SecretReference{
				Name:      config.Config.CsiSecret,
				Namespace: config.Config.CsiNamespace,
			},
			NodeStageSecretRef: &corev1.SecretReference{
				Name:      config.Config.CsiSecret,
				Namespace: config.Config.CsiNamespace,
			},
			VolumeAttributes: map[string]string{
				"capacity": config.Config.Capacity,
				"mounter":  "geesefs",
				"options":  "--memory-limit 1000 --dir-mode 0777 --file-mode 0666",
			},
			VolumeHandle: data.MiniIo.Path,
		}
	}
	return K8s.ClientSet.CoreV1().PersistentVolumes().Create(context.TODO(), pv, metav1.CreateOptions{})
}

// 获取pv详情
func (p *pv) GetPvDetail(pvName string) (pv *corev1.PersistentVolume, err error) {
	pv, err = K8s.ClientSet.CoreV1().PersistentVolumes().Get(context.TODO(), pvName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Pv详情失败, " + err.Error()))
		return nil, errors.New("获取Pv详情失败, " + err.Error())
	}

	return pv, nil
}

// 删除pv
func (p *pv) DeletePv(pvName string) (err error) {
	err = K8s.ClientSet.CoreV1().PersistentVolumes().Delete(context.TODO(), pvName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Pv失败, " + err.Error()))
		return errors.New("删除Pv失败, " + err.Error())
	}
	return nil
}
