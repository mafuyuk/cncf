### Istioインストール
```bash
$ curl -L https://git.io/getLatestIstio | sh -
$ export PATH=$PATH:$PWD/istio-1.3.0/bin
$  istioctl version --remote=false
1.3.0
$ kubectl apply -f ./istio-1.3.0/install/kubernetes/istio-demo.yaml
$ kubectl get services -n istio-system
$ kubectl get pods -n istio-system
```

### 立ち上げ
```bash
$ kubectl apply -f helloworld-v1.yaml
$ kubectl apply -f helloworld-v2.yaml
$ kubectl get pod

$ kubectl apply -f service.yaml
$ kubectl get service

$ kubectl apply -f virtualservice.yaml
$ kubectl apply -f app-gateway.yaml
```