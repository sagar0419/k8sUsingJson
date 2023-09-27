package main

import (
	"flag"
	"fmt"
	"k8sUsingJson/deploy"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Setting Default Values for Deployment, as we are using flag package values can we override from the CLI while executing the main.go
	//  for eg:- go run main.go -deploymentName sagar -image test
	deployName := flag.String("deploymentName", "sagar", "the name of the deployment")
	imageName := flag.String("image", "sagar27/petclinic-demo", "The name of the container Image")
	namespaceName := flag.String("namespace", "default", "Name of the namespace where you want to deploy the app")

	flag.Parse()

	// Printing the passed argument
	fmt.Printf("Your app %s will use Image %s and will get deployed in the %s namespace \n", *deployName, *imageName, *namespaceName)

	// Getting Home Directory
	configPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error in configPath %v \n", err)
		os.Exit(1)
	}

	// Setting complete path of homeDIr
	kubeconfigPath := filepath.Join(configPath, ".kube", "config")

	fmt.Printf("Hey!! we are using kubeconfig present on this path %v", kubeconfigPath)

	// Storing content of ~/.kube/config file in a variable
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf("error in kubeconfig %v", err)
		os.Exit(1)
	}

	//  Client is kubernetes which will use the configuration provided in the kubeconfig and help us to interact with the Kubernetes API.
	client, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Printf("error in setting client value %v", err)
		os.Exit(1)
	}

	deploy.DeployApp(client, *deployName, *imageName, *namespaceName)
}
