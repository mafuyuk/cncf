```bash
$ kctx docker-desktop

# Fluxのチャートの追加
$ helm repo add fluxcd https://charts.fluxcd.io
$ helm search repo fluxcd
NAME                    CHART VERSION   APP VERSION     DESCRIPTION                                       
fluxcd/flux             1.1.0           1.17.1          Flux is a tool that automatically ensures that ...
fluxcd/helm-operator    0.6.0           1.0.0-rc8       Flux Helm Operator is a CRD controller for decl...
 
$ kubectl create ns flux
$ helm upgrade -i flux \
  --set helmOperator.create=true \
  --set helmOperator.createCRD=false \
  --set git.url=git@github.com:mafuyuk/demo-k8s-manifest \
  --git-branch=master
  --namespace flux \
  fluxcd/flux

$ kubectl get pods -n flux

# Deploy Keysに公開鍵を追加
$ fluxctl --k8s-fwd-ns flux identity
```

## 掃除
```bash
$ eksctl delete cluster -f cluster.yaml
```
# 参考リンク
- https://eksctl.io/
- https://github.com/weaveworks/eksctl/tree/master/examples
- https://github.com/fluxcd/helm-operator-get-started
- https://eksworkshop.com/weave_flux/
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/create-cluster.html
- https://docs.aws.amazon.com/ja_jp/eks/latest/userguide/getting-started-eksctl.html
- https://www.weave.works/technologies/gitops/
- https://dev.classmethod.jp/cloud/aws/eks-weave-flux-gitops/