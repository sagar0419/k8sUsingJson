// Function to create Headless service.
package headlessvc

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

func HdlsSvc(client *kubernetes.Clientset, deployName string, namespaceName string, podPort int, servicePort int) {

	// Changing DataType of podPort to int32
	podport := int32(podPort)

	// Changing DataType of servicePort to int32
	serviceport := intstr.FromInt(servicePort)

	deployHeadless := client.CoreV1().Services(namespaceName)

	// Headless Manifest
	headless := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName,
			Namespace: namespaceName,
			Labels: map[string]string{
				"app": deployName,
			},
		},
		Spec: v1.ServiceSpec{
			Ports: []v1.ServicePort{
				{
					Name:       "http",
					Port:       podport,
					TargetPort: serviceport,
				},
			},
			ClusterIP: "None", //Cluste IP is none because we are creating headless service
			Selector: map[string]string{
				"app": deployName,
			},
		},
	}
	// Message to let user know the Headless service is getting created
	fmt.Printf("Creating Headless Service of name %v of type Headless Service \n \t", deployName)

	// Deploying the above service Manifest
	// Using deployHeadless reference and it's associated methods we have created in the first step in this file.
	if _, err := deployHeadless.Create(context.Background(), headless, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to create the service %v \n", err)
		os.Exit(1)
	}
	// Message to let user know that service is deployed
	fmt.Printf("Headless Service is deployed successfully %v in the %v Namespace.\n ", deployName, namespaceName)
}
