// Function to create Statefulset.
package state

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

func State(client *kubernetes.Clientset, deployName string, imageName string, namespaceName string, replicaCount int, podPort int) {

	// Changing the DataType of the podPort
	podport := int32(podPort)

	// Changing the DataType of the replicaCount
	replicacount := int32(replicaCount)

	// Changing the storage class Data Type
	storageClassName := string("standard")

	deployState := client.AppsV1().StatefulSets(namespaceName)

	// Statefull Manifest
	stateSpec := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployName,
			Namespace: namespaceName,
		},
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicacount,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": deployName,
				},
			},
			ServiceName: deployName,
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": deployName,
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:  deployName,
							Image: imageName,
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: podport,
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []apiv1.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name: "data",
					},
					Spec: apiv1.PersistentVolumeClaimSpec{
						AccessModes: []apiv1.PersistentVolumeAccessMode{
							apiv1.ReadWriteOnce,
						},
						StorageClassName: &storageClassName,
					},
				},
			},
		},
	}
	// Message to let user know that Statefulset is getting created
	fmt.Println("Deploying your Statefulset", deployName)

	// Deploying the above Statefulset manifest using deployState reference and it's associated methods we have created in the first step in this file.
	if _, err := deployState.Create(context.Background(), stateSpec, metav1.CreateOptions{}); err != nil {
		log.Fatalf("Failed to deploy the Statefulset manifest %v \n", err)
		os.Exit(1)
	}
	// Message to let user know that Statefulset is deployed
	fmt.Printf("Statefulset is deployed successfully %v in the %v Namespace. \n \t", deployName, namespaceName)
}
