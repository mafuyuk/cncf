# 作成
```bash
$ export AWS_PROFILE=default
$ eksctl create cluster -f eks/alb/cluster.yaml
$ aws eks update-kubeconfig --name demo
$ kctx
$ kubectl apply -k eks/alb/overlays
```

# 削除
```bash
$ export AWS_PROFILE=default
$ kctx
$ kubectl delete -k eks/alb/overlays
$ eksctl delete cluster -f eks/alb/cluster.yaml
```
# 参考リンク
- https://kubernetes-sigs.github.io/aws-alb-ingress-controller/
- https://kubernetes-sigs.github.io/aws-alb-ingress-controller/guide/controller/setup/
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/alb-ingress.html
