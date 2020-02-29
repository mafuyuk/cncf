# deployとservice作成
```bash
$ kubectl apply -f ./
```

# サービスの内容確認
```bash
$ kubectl get services sample-externalip
$ kubectl describe svc sample-externalip
```


# k8sノード一覧とアドレスを出力
```bash
$ kubectl get nodes -o custom-columns="NAME:{metadata.name},IP:{status.address[].address}"
```

# 削除
```bash
$ kubectl delete services sample-externalip
$ kubectl delete deployments sample-deployment
```