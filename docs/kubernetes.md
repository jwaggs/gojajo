### Notes

ServiceSpec types:

    ClusterIP (default) - Exposes the Service on an internal IP in the cluster. This type makes the Service only reachable from within the cluster.
    NodePort - Exposes the Service on the same port of each selected Node in the cluster using NAT. Makes a Service accessible from outside the cluster using <NodeIP>:<NodePort>. Superset of ClusterIP.
    LoadBalancer - Creates an external load balancer in the current cloud (if supported) and assigns a fixed, external IP to the Service. Superset of NodePort.
    ExternalName - Exposes the Service using an arbitrary name (specified by externalName in the spec) by returning a CNAME record with the name. No proxy is used. This type requires v1.7 or higher of kube-dns.
    
### Basics & Helpers

capture pod name `export POD_NAME=$(kubectl get pods -o go-template --template '{{range .items}}{{.metadata.name}}{{"\n"}}{{end}}')`

logs `kubectl logs $POD_NAME`

interactive bash shell `kubectl exec -ti $POD_NAME bash`

create deployment `kubectl create deployment piggy-api --image=gcr.io/piggy-police/pig-api:latest`
   
expose service `kubectl expose deployment/piggy-api --type NodePort --port 80 --target-port 8080`
