# Kubernetes liveness/readiness probes playground

```
eval $(minikube docker-env)
make VERSION=v4

kubectl apply -f deplyment.yml
kubectl apply -f service.yml

curl -s -D- $(minikube service readiness --url)/ok
```

## Recover after timeout

### 2 replicas
```
# scale
kubectl scale --replicas=2 deployment/readiness

# timeout for 1 min
curl -D- -XPOST $(minikube service readiness --url)/set/timeout

# watch how it switches traffic to a single pod. Second pod will be back after 1 min
watch -n 1 curl -s -D- $(minikube service readiness --url)/ok
```

### 1 replica
```
# scale
kubectl scale --replicas=1 deployment/readiness

# timeout for 1 min
curl -D- -XPOST $(minikube service readiness --url)/set/timeout

# not available. recover after 1 min
watch -n 1 curl -s -D- $(minikube service readiness --url)/ok
```

## Not healthy

```
# scale
kubectl scale --replicas=1 deployment/readiness

# set unhealthy
curl -D- -XPOST $(minikube service readiness --url)/set/not/healthy

# not available
watch -n 1 curl -s -D- $(minikube service readiness --url)/ok
```
