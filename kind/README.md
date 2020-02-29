
## クラスタ作成方法
### オプション使用
```bash
$ kind create cluster \
  --name dev-cluster \
  --loglevel debug \
  --image kindest/node:v1.14.3
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
```

## クラスタ削除
```bash
kind delete cluster \
  --name dev-cluster
```
## ノードコンテナにアクセス
```bash
docker exec -it kind-control-plane bash
```

# 参考リンク
- https://kind.sigs.k8s.io/
- https://kind.sigs.k8s.io/docs/user/quick-start/#setting-kubernetes-version
- https://hub.docker.com/r/kindest/node/tags
- https://blog.cybozu.io/entry/2019/07/03/170000
