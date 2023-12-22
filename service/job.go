package service

import (
	"context"
	"errors"
	"github.com/wonderivan/logger"
	batchsv1 "k8s.io/api/batch/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strconv"
)

var Job job

type job struct{}

type JobCreate struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Image     string            `json:"image"`
	Label     map[string]string `json:"label"`
	Cpu       string            `json:"cpu"`
	Memory    string            `json:"memory"`
	Command   []string          `json:"command"`
	Volumes   []Volume          `json:"volumes"`
}

func (j *job) GetJobDetail(jobName, namespace string) (job *batchsv1.Job, err error) {
	job, err = K8s.ClientSet.BatchV1().Jobs(namespace).Get(context.TODO(), jobName, metav1.GetOptions{})
	if err != nil {
		logger.Error(errors.New("获取job详情失败, " + err.Error()))
		return nil, errors.New("获取job详情失败, " + err.Error())
	}
	return job, nil
}

func (j *job) CreateJob(data *JobCreate) (err error) {
	job := &batchsv1.Job{
		//ObjectMeta中定义资源名、命名空间以及标签
		ObjectMeta: metav1.ObjectMeta{
			Name:      data.Name,
			Namespace: data.Namespace,
			Labels:    data.Label,
		},
		Spec: batchsv1.JobSpec{
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
							Name:    data.Name,
							Image:   data.Image,
							Command: data.Command,
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
					RestartPolicy: corev1.RestartPolicyNever,
				},
			},
		},
		Status: batchsv1.JobStatus{},
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
	job.Spec.Template.Spec.Containers[0].VolumeMounts = volumeMounts
	job.Spec.Template.Spec.Volumes = volumes

	_, err = K8s.ClientSet.BatchV1().Jobs(data.Namespace).Create(context.TODO(), job, metav1.CreateOptions{})
	if err != nil {
		logger.Error(errors.New("创建job失败, " + err.Error()))
		return errors.New("创建job失败, " + err.Error())
	}
	return nil
}

func (d *job) DeleteJob(jobName, namespace string) (err error) {
	err = K8s.ClientSet.BatchV1().Jobs(namespace).Delete(context.TODO(), jobName, metav1.DeleteOptions{})
	if err != nil {
		logger.Error(errors.New("删除job失败, " + err.Error()))
		return errors.New("删除job失败, " + err.Error())
	}
	return nil
}
