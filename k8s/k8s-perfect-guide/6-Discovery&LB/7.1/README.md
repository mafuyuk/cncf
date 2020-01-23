# service作成
```bash
$ kubectl apply -f ./
```

# サービスの内容確認
```bash
$ kubectl get services sample-headless
$ kubectl get statefulsets sample-statefulset-headless
```

# Pod名でのサービスディスカバリ
```bash
$ kubectl run --image=centos:6 --restart=Never --rm -i testpod -- dig sample-statefulset-headless-0.sample-headless.default.svc.cluster.local
```

# 削除
```bash
$ kubectl delete statefulsets sample-statefulset-headless
$ kubectl delete services sample-headless
```