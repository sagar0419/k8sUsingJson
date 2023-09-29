package deploy

import (
	"context"
	"fmt"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// DeployApp(client, deployName, imageName, namespaceName)
func DeployApp(client *kubernetes.Clientset, deployName string, imageName string, namespaceName string, replica int, podPort int) {

	// Changing Replica count variable Data Type
	replicaCount := int32(replica)

	// Changing podPort variable Data Type
	port := int32(podPort)

	// Setting Reference to the Kubernetes Deployment resource within a given namespace.
	// Now can create, update, delete, or list Deployments in that namespace using the deploy variable and its associated methods from the Kubernetes client library.
	deploy := client.AppsV1().Deployments(namespaceName)

	deploymentSpec := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      string(deployName),
			Namespace: string(namespaceName),
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicaCount,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deployName,
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deployName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  "sagar",
							Image: imageName,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: port,
									Protocol:      apiv1.ProtocolTCP,
								},
							},
						},
					},
				},
			},
		},
	}
	// Message to let user know that Deployment is getting created
	fmt.Println("Deploying your deployment", deployName)

	// Deploying the above deployment manifest using deploy reference and it's associated methods we have created in the first step in this file.
	if _, err := deploy.Create(context.Background(), deploymentSpec, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to deploy the deployment manifest %v \n", err)
		os.Exit(1)
	}
	// Message to let user know that Deployment is deployed
	fmt.Printf("Deployment is deployed successfully %v in the %v Namespace. \n \t", deployName, namespaceName)
}
