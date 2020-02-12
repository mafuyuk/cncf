# 事前準備
```bash
$ export AWS_PROFILE=bellface-dev
$ eksctl create cluster -f cluster.yaml
$ aws eks update-kubeconfig --name demo
$ kctx
```

## 作成
```bash
$ kubectl apply -k eks/alb/overlays
```

# 削除
```bash
$ kubectl delete -k eks/alb/overlays
```
# 参考リンク
- https://kubernetes-sigs.github.io/aws-alb-ingress-controller/
- https://kubernetes-sigs.github.io/aws-alb-ingress-controller/guide/controller/setup/
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alb-ingress.html
