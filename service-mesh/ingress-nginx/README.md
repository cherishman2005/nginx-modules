# ingress-nginx

## whoami部署

* 部署
```
kubectl apply -f whoami.yaml
```

* 将 Deployment 暴露出来
```
kubectl expose deployment whoami --type=NodePort --port=80
```

# FAQ

## Error obtaining Endpoints for Service "default/whoami": no object matching key "default/whoami" in local store

```
W0731 09:56:33.489512      30 controller.go:1021] Error obtaining Endpoints for Service "default/whoami": no object matching key "default/whoami" in local store
```

* 将 Deployment 暴露出来
```
kubectl expose deployment whoami --type=NodePort --port=80
```
