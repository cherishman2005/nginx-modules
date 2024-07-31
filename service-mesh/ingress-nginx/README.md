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

### 解决后运行日志

```
W0731 10:06:41.313678      30 controller.go:1123] serviceEndpoints Service "default/web" upstreams: [{10.244.0.13 8080 &ObjectReference{Kind:Pod,Namespace:default,Name:web-68487bc957-w578d,UID:080c4b70-f876-4fa0-819b-a8521858200c,APIVersion:,ResourceVersion:,FieldPath:,}}]
W0731 10:06:41.313762      30 controller.go:1123] serviceEndpoints Service "default/web2" upstreams: [{10.244.0.14 8080 &ObjectReference{Kind:Pod,Namespace:default,Name:web2-6459878f46-wtzls,UID:e923a502-0850-4656-ba12-d7626daaf3f3,APIVersion:,ResourceVersion:,FieldPath:,}}]
W0731 10:06:41.313784      30 controller.go:1123] serviceEndpoints Service "default/whoami" upstreams: [{10.244.0.21 80 &ObjectReference{Kind:Pod,Namespace:default,Name:whoami-74db784c96-gq8s8,UID:6fba3643-f24c-4ab0-be46-7b149912dbf9,APIVersion:,ResourceVersion:,FieldPath:,}} {10.244.0.20 80 &ObjectReference{Kind:Pod,Namespace:default,Name:whoami-74db784c96-rpt4k,UID:448e1f32-1a37-4481-a51e-1c76c66f6a21,APIVersion:,ResourceVersion:,FieldPath:,}}]
I0731 10:06:41.313846      30 controller.go:168] "Configuration changes detected, backend reload required"
I0731 10:06:41.414129      30 controller.go:185] "Backend successfully reloaded"
I0731 10:06:41.414618      30 event.go:285] Event(v1.ObjectReference{Kind:"Pod", Namespace:"ingress-nginx", Name:"ingress-nginx-controller-59fcf8978c-g8hqm", UID:"b0d2ca90-5f13-4071-a358-7c486999a6cc", APIVersion:"v1", ResourceVersion:"62940", FieldPath:""}): type: 'Normal' reason: 'RELOAD' NGINX reload triggered due to a change in configuration
```
