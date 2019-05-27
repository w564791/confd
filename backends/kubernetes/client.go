package k8kubernetes

import (
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"encoding/json"
	"fmt"
	"strings"
	"k8s.io/client-go/rest"
	"github.com/kelseyhightower/confd/log"
)
type KubernetesClient struct {
	client  kubernetes.Clientset
}


func New(kubeconfig string,InCluster bool)(*KubernetesClient,error) {
	var config *rest.Config
	var err error
	if InCluster{
		config, err = rest.InClusterConfig()
		if err != nil {
			log.Info(err.Error())
		}
		log.Info("Use InCluster config")
	}else{
		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
		log.Info("Use OutOfCluster config")
	}

	clientset, err := kubernetes.NewForConfig(config)

	return &KubernetesClient{*clientset},err

	}

func (c *KubernetesClient) GetValues(endpoints []string) (map[string]string, error) {
	vars := make(map[string]string)
	for _, key := range endpoints {
		newKey:=strings.Split(key,"/")
		namespace:=newKey[1]

		endpoint,err :=c.client.CoreV1().Endpoints(namespace).Get(newKey[2],metav1.GetOptions{})
		if err != nil {
			log.Info(err.Error())
			return vars,err
		}
		jsendpoint,_:=json.Marshal(endpoint)
		vars[key]=string(jsendpoint)
	}
	return vars, nil
}
func (c *KubernetesClient) WatchPrefix(prefix string, keys []string, waitIndex uint64, stopChan chan bool) (uint64, error) {
	return 1.0,fmt.Errorf("ddd")
}
