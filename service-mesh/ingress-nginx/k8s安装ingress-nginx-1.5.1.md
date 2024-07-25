#k8s安装ingress-nginx-1.5.1

## 下载配置文件

```
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.5.1/deploy/static/provider/cloud/deploy.yaml
```

## 修改deploy.yaml

vim deploy.yaml
```
image: registry.k8s.io/ingress-nginx/controller:v1.5.1@sha256:4ba73c697770664c1e00e9f968de14e08f606ff961c76e5d7033a4a9c593c629
```
改为
```
image: registry.cn-hangzhou.aliyuncs.com/google_containers/nginx-ingress-controller:v1.5.1
```

```
image: registry.k8s.io/ingress-nginx/kube-webhook-certgen:v20220916-gd32f8c343@sha256:39c5b2e3310dc4264d638ad28d9d1d96c4cbb2b2dcfb52368fe4e3c63f61e10f
```
改为
```
image: registry.cn-hangzhou.aliyuncs.com/google_containers/kube-webhook-certgen:v20220916-gd32f8c343
```

## 运行

```
kubectl apply -f deploy.yaml
#等待几分钟部署完成
[root@k8s-master ingress-nginx]# kubectl get pods --namespace=ingress-nginx
NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-r66nq        0/1     Completed   0          11h
ingress-nginx-admission-patch-2wcrj         0/1     Completed   2          11h
ingress-nginx-controller-547d59555f-vcwgb   1/1     Running     0          11h
```
