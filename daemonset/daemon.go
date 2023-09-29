package daemonset

import (
	"context"
	"fmt"
	"log"
	"os"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func DaemonApply(client *kubernetes.Clientset, deployName string, imageName string, namespaceName string, podPort int) {

	// Changing podPort variable Data Type
	podport := int32(podPort)

	// Converting termination grace period to int64
	gracePeriodSeconds := int64(30)

	daemon := client.AppsV1().DaemonSets(namespaceName)

	daemonsetSpec := &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName,
			Namespace: namespaceName,
			Labels: map[string]string{
				"app": deployName,
			},
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"name": deployName,
				},
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"name": deployName,
					},
				},
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						{
							Name:            deployName,
							Image:           imageName,
							ImagePullPolicy: v1.PullAlways,
							Ports: []v1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: podport,
								},
							},
						},
					},
					TerminationGracePeriodSeconds: &gracePeriodSeconds,
				},
			},
		},
	}

	fmt.Printf("Deploying your Daemonset %v \n\t", deployName)
	if _, err := daemon.Create(context.Background(), daemonsetSpec, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to deploy the Daemonset manifest %v \n", err)
		os.Exit(1)
	}
	fmt.Printf("Your Daemonset %v is deployed successfully in the %v namespace \n\t", deployName, namespaceName)
}
