### テスト用のクラスタ生成
```bash
$ kind create cluster --name cronjob-s3 --image kindest/node:v1.18.2
$ kind get clusters
cronjpb-s3

$ kctx
Switched to context "cronjob-s3".
```

### 構築
```bash
$ kustomize build k8s/poc/cronjob-s3 | kubectl apply -f -
$ k get po -w
```

### RBACを使うのでクラスタでRBAC使えるようになっているか確認
```bash
$ k describe pod -n kube-system -l component=kube-apiserver | grep authorization-mode
      --authorization-mode=Node,RBAC
```

### Cleanup
```bash
$ kind delete cluster --name cronjob-s3
```
