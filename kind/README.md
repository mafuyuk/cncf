
## クラスタ作成方法
### オプション使用
```bash
$ kind create cluster \
  --name dev-cluster \
  --loglevel debug \
  --image kindest/node:v1.18.2
$ kind get clusters
dev-cluster
$ kctx 
```

### config
```bash
$ kind create cluster \
  --name dev-cluster \
  --config my-k8s.yaml

$ kind get clusters
dev-cluster
$ kctx
kind-dev-cluster
```

## クラスタ削除
```bash
$ kind delete cluster \
  --name dev-cluster
```
## ノードコンテナにアクセス
```bash
$ docker exec -it dev-cluster-control-plane bash
```

##
```bash
$ k create deploy --image=amaya382/k8s-debugkit debugkit
$ k expose deploy debugkit --port=80 --type=NodePort
$ k edit svc debugkit # spec.ports[0].nodePortを31380に更新
```

# 参考リンク
- https://kind.sigs.k8s.io/
- https://kind.sigs.k8s.io/docs/user/quick-start/#setting-kubernetes-version
- https://hub.docker.com/r/kindest/node/tags
- https://blog.cybozu.io/entry/2019/07/03/170000
