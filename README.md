# k8sUsingJson
Creating resources in kubernetes without using Yaml file

To run this code clone the repo on to your local machine and go inside the `k8sUsingJson` directory and execute this command.
`go run main.go`
It will show you all the flags which you can pass to run the code. Flags are also listed below.
```
  -deploymentName string
        Name of the deployment / Service (default "sagar")
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
  -targetPort int
        The target port on the pods that this Service will forward traffic to (Must Require to create service and deployment)
```