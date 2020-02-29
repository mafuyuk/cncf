# service作成
```bash
$ kubectl apply -f ./
```

# サービスの内容確認
```bash
$ kubectl get services sample-externalname
NAME                  TYPE           CLUSTER-IP   EXTERNAL-IP            PORT(S)   AGE
sample-externalname   ExternalName   <none>       external.example.com   <none>    23s
```

# CNAMEの名前解決が行われることの確認
```bash
$ kubectl run --image=centos:6 --restart=Never --rm -i testpod -- dig sample-externalname.default.svc.cluster.local CNAME
```

# 削除
```bash
$ kubectl delete services sample-externalname
```