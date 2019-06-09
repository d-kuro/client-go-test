package main

import (
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	client, err := newClient()
	if err != nil {
		log.Fatal(err)
	}

	watchInterface, err := client.CoreV1().Pods("").Watch(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	watchChan := watchInterface.ResultChan()

	for {
		select {
		case event := <-watchChan:
			po, ok := event.Object.(*v1.Pod)
			if !ok {
				continue
			}
			fmt.Printf("EventType: %s\n", event.Type)
			fmt.Printf("%#v\n", po)
		}
	}
}

func newClient() (kubernetes.Interface, error) {
	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(kubeConfig)
}
