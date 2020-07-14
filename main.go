package main

import (
	"flag"
	"path/filepath"

	"github.com/micro/micro/v2/cmd"

	"github.com/micro/go-micro/v2/util/log"
	apiv1 "k8s.io/api/core/v1"
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
	log.Log(kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	log.Log(config)
	if err != nil {
		log.Fatal(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	log.Log(clientset)
	deploymentsClient := clientset.AppsV1().Deployments(apiv1.NamespaceDefault)
	log.Log(deploymentsClient)

	cmd.Init()

}
