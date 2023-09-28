package svc

import (
	"context"
	"fmt"
	"log"
	"os"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
)

func SvcApply(client *kubernetes.Clientset, deployName string, serviceType string, namespaceName string, nodePort int, podPort int, servicePort int, protocol string) {

	nodeport := int32(nodePort)
	podport := intstr.FromInt(podPort)
	serviceport := int32(servicePort)

	serviceDeploy := client.CoreV1().Services(namespaceName)

	serviceSpec := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name: string(deployName),
			Labels: map[string]string{
				"app": deployName,
			},
		},
		Spec: v1.ServiceSpec{
			Selector: map[string]string{
				"app": deployName,
			},
			Type: v1.ServiceType(serviceType),
			Ports: []v1.ServicePort{
				{
					Protocol:   v1.Protocol(protocol),
					TargetPort: podport,
					Port:       serviceport,
					NodePort:   nodeport,
				},
			},
		},
	}
	fmt.Printf("Creating service of name %v of type %v", deployName, serviceType)

	// Deploying the above service Manifest
	// Using serviceDeploy reference and it's associated methods we have created in the first step in this file.

	if _, err := serviceDeploy.Create(context.Background(), serviceSpec, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to create the service %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Service is deployed successfully %v in the %v Namespace.\n", deployName, namespaceName)
}
