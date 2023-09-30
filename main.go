package main

import (
	"flag"
	"fmt"
	"k8sUsingJson/daemonset"
	"k8sUsingJson/deploy"
	headlessvc "k8sUsingJson/headlesSvc"
	"k8sUsingJson/state"
	"k8sUsingJson/svc"
	"log"
	"os"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// Setting Default Values for Deployment, as we are using flag package values can we override from the CLI while executing the main.go
	// for eg:- go run main.go -deploymentName sagar -image test
	deployName := flag.String("deploymentName", "sagar", "Name of the deployment/Service")
	replicaCount := flag.Int("replica", 2, "Number of repplicas you want to create")
	imageName := flag.String("image", "sagar27/petclinic-demo", "The name of the container Image")
	namespaceName := flag.String("namespace", "default", "Name of the namespace where you want to deploy the app")
	podPort := flag.Int("targetPort", 0, "The target port on the pods that this Service will forward traffic to (Must Require to create service and deployment)")
	serviceType := flag.String("servicetype", "ClusterIP", "Service Type that you want to deploy")
	servicePort := flag.Int("servicePort", 0, "The port on the service itself (Must Required to create service)")
	nodePort := flag.Int("nodePort", 0, "The static port on each node (can be omitted to let Kubernetes choose)")
	protocol := flag.String("protocol", "TCP", "The IP protocol for this port. Supports 'TCP', 'UDP' and 'SCTP'")
	ss := flag.String("ss", "", "Pass the Value as 'Yes' to deploy the Statefulset")
	ds := flag.String("ds", "", "Pass the Value as 'Yes' to deploy the Daemonset")

	flag.Parse()
	// Verifying the required values are passed or not
	// Checking the value of Target Port and Service Port
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
	if *serviceType == "NodePort" && *nodePort == 0 {
		fmt.Println("Since you have selected the service type NodePort and you haven't provided any port, Now a random nodeport will get assigned to the service \n \t")
	}

	// Getting Home Directory
	configPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("error in configPath %v \n \t", err)
		os.Exit(1)
	}

	// Setting complete path of homeDIr
	kubeconfigPath := filepath.Join(configPath, ".kube", "config")

	// Storing content of ~/.kube/config file in a variable
	kubeconfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Printf("error in kubeconfig %v\n \t", err)
		os.Exit(1)
	}

	//  Client is kubernetes which will use the configuration provided in the kubeconfig and help us to interact with the Kubernetes API.
	client, err := kubernetes.NewForConfig(kubeconfig)
	if err != nil {
		fmt.Printf("error in setting client value %v\n \t", err)
		os.Exit(1)
	}

	if *ss == "Yes" && *ds == "Yes" || *ss == "yes" && *ds == "yes" {
		log.Fatal("You cannot deploy both Statefulset and Daemonset in one Go")
		os.Exit(1)

	} else if *ss == "Yes" || *ss == "yes" {
		// Printing Kubeconfig file details
		fmt.Printf("Hey!! we are using kubeconfig present on this path to deploy the application %v \n \t", kubeconfigPath)

		// Printing the passed argument
		fmt.Printf("Your app %s will use Image %s and will get deployed in the %s namespace \n \t", *deployName, *imageName, *namespaceName)

		// Calling State function to create the StateFulset with the passed variables.
		state.State(client, *deployName, *imageName, *namespaceName, *replicaCount, *podPort)

		// Calling Headless Service function to create the Service with the passed variables.
		headlessvc.HdlsSvc(client, *deployName, *namespaceName, *podPort, *servicePort)

		// Terminating loop
		return

	} else if *ds == "Yes" || *ds == "yes" {
		// Printing Kubeconfig file details
		fmt.Printf("Hey!! we are using kubeconfig present on this path to deploy the application %v \n \t", kubeconfigPath)

		// Printing the passed argument
		fmt.Printf("Your app %s will use Image %s and will get deployed in the %s namespace \n \t", *deployName, *imageName, *namespaceName)

		// Calling DaemonApply function to create the DaemonSet with the passed variables.
		daemonset.DaemonApply(client, *deployName, *imageName, *namespaceName, *podPort)

		// Calling SvcApply function to create the Service with the passed variables.
		svc.SvcApply(client, *deployName, *serviceType, *namespaceName, *nodePort, *podPort, *servicePort, *protocol)

		// Terminating loop
		return
	} else {
		// Printing Kubeconfig file details
		fmt.Printf("Hey!! we are using kubeconfig present on this path to deploy the application %v \n \t", kubeconfigPath)

		// Printing the passed argument
		fmt.Printf("Your app %s will use Image %s and will get deployed in the %s namespace \n \t", *deployName, *imageName, *namespaceName)

		// Calling DeployAPP function to deploy the app with the passed variables.
		deploy.DeployApp(client, *deployName, *imageName, *namespaceName, *replicaCount, *podPort)

		// Calling SvcApply function to create the Service with the passed variables.
		svc.SvcApply(client, *deployName, *serviceType, *namespaceName, *nodePort, *podPort, *servicePort, *protocol)

		// Terminating loop
		return
	}
}
