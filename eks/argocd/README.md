# 構築
## 作成
```bash
$ eksctl create cluster -f eks/argocd/cluster.yaml
$ kubectl apply -k eks/argocd/manifest/overlays/dev/ap-northeast-1/default
$ kubectl apply -k eks/argocd/manifest/overlays/dev/ap-northeast-1/argocd
```

## 削除
```bash
$ kubectl delete ns argocd
$ kubectl delete ns default
$ eksctl delete cluster -f eks/argocd/cluster.yaml
```

# 調査
## Podに直接接続
```sh
$ kubectl port-forward nginx-deployment-554d546bdd-4hjzd 8080:80 &
$ curl --silent 127.0.0.1:8080 | head -n 10
```

## クラスタ外からClusterIPに接続
```sh
$ kubectl proxy &
[2] 65740
Starting to serve on 127.0.0.1:8001
$ curl 127.0.0.1:8001/api/v1/namespaces/default/services/nginx-service/proxy/

$ kill -9 65740

$ curl 127.0.0.1:8001/api/v1/namespaces/argocd/services/argocd-server/proxy/
```

## Argocd接続
```sh
$ kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
$ kubectl port-forward svc/argocd-server -n argocd 8081:443

$ open 127.0.0.1:8081
ID: admin
Password: ↓の結果

$ kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2
```
