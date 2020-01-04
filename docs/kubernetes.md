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
   
expose service `kubectl expose deployment/piggy-api --type ClusterIP --port 80 --target-port 8080`

pod ip `kubectl get pods -l run=my-app -o yaml | grep podIP`

create secrete 
```
# Create files needed for rest of example.
echo -n 'admin' > ./username.txt
echo -n '1f2d1e2e67df' > ./password.txt
kubectl create secret generic jwaggs27-chase-user-pass --from-file=./username.txt --from-file=./password.txt --from-file=./clientuid.txt
```

or

`kubectl create secret generic dev-db-secret --from-literal=username=devuser --from-literal=password='S!B\*d$zDsb'`

### Cron
format: minute hour dayOfMonth month dayOfWeek

every second: `* * * * *`
every 5th second: `*/5 * * * *`
one minute past every 5th hour: `1 */5 * * *`

### Labels

```
kubectl run nginx1 --image=nginx --restart=Never --labels=app=v1
kubectl run nginx2 --image=nginx --restart=Never --labels=app=v1
kubectl run nginx3 --image=nginx --restart=Never --labels=app=v1
```