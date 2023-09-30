# k8sUsingJson
Creating resources in kubernetes without using Yaml file

### Prerequisite:- 
This script will use the kubeconfig file stored in home directory `~/.kube/config` path. So, before excuting this script make sure you have kubeconfig file in the right directory and right context has been set.

### Info To run the code
To run this code clone the repo on to your local machine and go inside the `k8sUsingJson` directory and execute this command `go run main.go -h` or `go run main.go`.  It will show you all the flags which you can pass to run the code. Flags are also listed below.

```
  -deploymentName string
        Name of the deployment / Service (default "sagar")
  -ds string
        Pass the Value as 'Yes' to deploy the Daemonset
  -image string
        The name of the container Image (default "sagar27/petclinic-demo")
  -namespace string
        Name of the namespace where you want to deploy the app (default "default")
  -nodePort int
        The static port on each node (can be omitted to let Kubernetes choose)
  -protocol string
        The IP protocol for this port. Supports 'TCP', 'UDP' and 'SCTP' (default "TCP")
  -replica int
        Number of repplica you want to create (default 2)
  -servicePort int
        The port on the service itself (Must Required to create service)
  -servicetype string
        Service Type that you want to deploy (default "ClusterIP")
  -ss string
        Pass the Value as 'Yes' to deploy the Statefulset
  -targetPort int
        The target port on the pods that this Service will forward traffic to (Must Require to create service and deployment)
```

If you want to deploy the daemonset pass the flag -ds with argument yes and If you want to deploy the Statefulset pass the flag -ss with argument yes. The script will not deploy the Daemonset and Statefulset together if you pass the value of both the flag -ds and -ss as yes. For deployment no Flag is required.

You can only deploy Satetfulset, Daemonset or Deployment in one go.