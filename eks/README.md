# EKS作成
```bash
# EKS Clusterの作成
$ eksctl create cluster \
  -f cluster.yaml \
   --profile default
$ aws eks update-kubeconfig \
  --name demo-cluster \
  --profile default
```

# eksctlでのFlux導入
```bash
$ EKSCTL_EXPERIMENTAL=true eksctl enable repo \
  -f cluster.yaml \
  --git-url git@github.com:mafuyuk/demo-k8s-manifest.git \
  --git-email mafuyuk.m@gmail.com \
  --with-helm true \
  --namespace flux \
  --profile default
```

# EKS削除
```bash
$ eksctl delete cluster -f cluster.yaml --profile=default
```

# 参考リンク
- https://eksctl.io/usage/experimental/gitops-flux/

# 参考リンク
- https://eksctl.io/usage/experimental/gitops-flux/
- https://eksctl.io/
- https://github.com/weaveworks/eksctl/tree/master/examples
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-cluster.html
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-eksctl.html
