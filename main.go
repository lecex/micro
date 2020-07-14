package main

import (
	"context"
	"flag"
	"path/filepath"

	"github.com/micro/micro/v2/cmd"

	"github.com/micro/go-micro/v2/util/log"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()
	log.Log(1, kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	log.Log(2, config)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Log(3, clientset)
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	log.Log(4, deploymentsClient)
	log.Log(5, apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	log.Log(6, list)
	cmd.Init()

}
