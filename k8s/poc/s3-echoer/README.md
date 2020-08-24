### テスト用のクラスタ生成
```bash
$ kind create cluster --name s3-echoer --image kindest/node:v1.18.2
$ kind get clusters
s3-echoer

$ kctx
Switched to context "s3-echoer".
```

### 構築
```bash
$ kustomize build k8s/poc/s3-echoer | kubectl apply -f -
$ k get po -w
```

### RBACを使うのでクラスタでRBAC使えるようになっているか確認
```bash
$ k describe pod -n kube-system -l component=kube-apiserver | grep authorization-mode
      --authorization-mode=Node,RBAC
```

### Cleanup
```bash
$ kind delete cluster --name s3-echoer
```

# 参考リンク
- https://katainaka0503.hatenablog.com/entry/2019/12/07/091737
- https://aws.amazon.com/jp/blogs/news/introducing-fine-grained-iam-roles-service-accounts/
