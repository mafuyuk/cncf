
## クラスタ作成方法
### config
```bash
$ kind create cluster \
  --name crd-cluster \
  --config my-k8s.yaml

$ kind get clusters
crd-cluster
$ kctx
kind-crd-cluster
```

## クラスタ削除
```bash
$ kind delete cluster \
  --name crd-cluster
```
## ノードコンテナにアクセス
```bash
$ docker exec -it crd-cluster-control-plane bash
```

##
```bash
$ k create deploy --image=amaya382/k8s-debugkit debugkit
$ k expose deploy debugkit --port=80 --type=NodePort
$ k edit svc debugkit # spec.ports[0].nodePortを31380に更新
```
