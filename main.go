package main

import (
	"flag"
	"fmt"
	"k8sUsingJson/deploy"
	"k8sUsingJson/svc"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Setting Default Values for Deployment, as we are using flag package values can we override from the CLI while executing the main.go
	//  for eg:- go run main.go -deploymentName sagar -image test
	deployName := flag.String("deploymentName", "sagar", "Name of the deployment / Service")
	replicaCount := flag.Int("replica", 2, "Number of repplica you want to create")
	imageName := flag.String("image", "sagar27/petclinic-demo", "The name of the container Image")
	namespaceName := flag.String("namespace", "default", "Name of the namespace where you want to deploy the app")
	podPort := flag.Int("targetPort", 0, "The target port on the pods that this Service will forward traffic to")
	serviceType := flag.String("servicetype", "ClusterIP", "Service Type that you want to deploy")
	servicePort := flag.Int("servicePort", 0, "The port on the service itself (inside the cluster)")
	nodePort := flag.Int("nodePort", 0, "The static port on each node (can be omitted to let Kubernetes choose)")
	protocol := flag.String("protocol", "TCP", "The IP protocol for this port. Supports 'TCP', 'UDP' and 'SCTP'")

	flag.Parse()

	// Verifying the required values are passed or not

	// Checking the value of Target Port
	if *podPort == 0 {
		fmt.Println("Error: The 'targetPort' flag is required.\n \t")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// Checking the value of Service Port
	if *servicePort == 0 {
		fmt.Println("Error: The 'servicePort' flag is required. \n \t")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Checking if service type nordport selected and port of the service has been passed or not.
	// if *serviceType == "NodePort" && *nodePort == 0 {
	// 	fmt.Println("Error: While passing the Service Type NodePort you have to pass the value of nodeport also. \n \t")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

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

	fmt.Printf("Hey!! we are using kubeconfig present on this path to deploy the application %v", kubeconfigPath)

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

	deploy.DeployApp(client, *deployName, *imageName, *namespaceName, *replicaCount, *podPort)
	svc.SvcApply(client, *deployName, *serviceType, *namespaceName, *nodePort, *podPort, *servicePort, *protocol)
}
