package k8kubernetes

import (
		"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//"k8s.io/api/core/v1"
	//"net/http"
		"log"
	//	"k8s.io/client-go/rest"
	//"strings"
	//"path"
	"encoding/json"
	"fmt"
	"strings"
)
type KubernetesClient struct {
	client  kubernetes.Clientset
}


func New(kubeconfig string)(*KubernetesClient,error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	clientset, err := kubernetes.NewForConfig(config)

	return &KubernetesClient{*clientset},err

	}

func (c *KubernetesClient) GetValues(endpoints []string) (map[string]string, error) {
	vars := make(map[string]string)
	for _, key := range endpoints {
		newKey:=strings.Split(key,"/")
		namespace:=newKey[1]
		//log.Println("======================",newKey[1],newKey[2])
		//log.Println(newKey[1])
		endpoint,err :=c.client.CoreV1().Endpoints(namespace).Get(newKey[2],metav1.GetOptions{})
		if err != nil {
			log.Println(err)
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
