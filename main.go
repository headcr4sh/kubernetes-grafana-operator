package main

import (
	"flag"
	"fmt"
	"os"

	//"k8s.io/api/core/v1"
	//apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	//"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
)

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to KUBECONFIG for running out of cluster. (Default: null)")
}

func main() {
	flag.Parse()
	fmt.Printf("Getting Kubernetes context (KUBECONFIG=\"%s\")\n", kubeconfig)
	config, err := getClientConfig(kubeconfig)
	if err != nil {
		fmt.Printf("Failed to create context: %+v\n", err)
		os.Exit(1)
	}
	fmt.Printf("%v\n", config)
}

func getClientConfig(kubeconfig string) (*rest.Config, error) {
	if kubeconfig != "" {
		return clientcmd.BuildConfigFromFlags("", kubeconfig)
	}
	return rest.InClusterConfig()
}
