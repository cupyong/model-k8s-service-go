package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var Namespace namespace

type namespace struct{}

type NamespacesResp struct {
	Items []corev1.Namespace `json:"items"`
	Total int                `json:"total"`
}

// 创建namespace
func (n *namespace) CreateNamespace(namespaceName string) (namespace *corev1.Namespace, err error) {
	return K8s.ClientSet.CoreV1().Namespaces().Create(context.TODO(), &corev1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: namespaceName,
		},
	}, metav1.CreateOptions{})
}

// 获取namespace详情
func (n *namespace) GetNamespaceDetail(namespaceName string) (namespace *corev1.Namespace, err error) {
	namespace, err = K8s.ClientSet.CoreV1().Namespaces().Get(context.TODO(), namespaceName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取Namespace详情失败, " + err.Error()))
		return nil, errors.New("获取Namespace详情失败, " + err.Error())
	}

	return namespace, nil
}

// 删除namespace
func (n *namespace) DeleteNamespace(namespaceName string) (err error) {
	err = K8s.ClientSet.CoreV1().Namespaces().Delete(context.TODO(), namespaceName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除Namespace失败, " + err.Error()))
		return errors.New("删除Namespace失败, " + err.Error())
	}
	return nil
}
