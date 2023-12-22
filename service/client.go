package service

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"model-k8s-service-go/config"
)

type Client struct {
	ClientSet *kubernetes.Clientset
}

var K8s Client

func InitClient() Client {
	config, err := clientcmd.BuildConfigFromFlags(config.Config.K8sUrl, "config/config")
	if err != nil {
		panic(err)
	}
	c, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	client := Client{}
	client.ClientSet = c
	K8s = client
	return client
}

/*
监听c
*/
func (client Client) Watch() {
	sharedInformerFactory := informers.NewSharedInformerFactory(client.ClientSet, 0)
	podInformer := sharedInformerFactory.Core().V1().Pods()

	informer := podInformer.Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("[Add Event] pod name:%s\n", pod.Name)
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			pod := newObj.(*v1.Pod)
			fmt.Println("[Update Event] new pod name", pod.Name, pod.Namespace, pod.Status.Phase)
		},
		DeleteFunc: func(obj interface{}) {
			pod := obj.(*v1.Pod)
			fmt.Printf("[Delete Event] pod name:%s\n", pod.Name)
		},
	})
	stopCh := make(chan struct{})
	sharedInformerFactory.Start(stopCh)
	<-stopCh
}
