# service作成
```bash
$ kubectl apply -f ./
```

# サービスの内容確認
```bash
$ kubectl get services sample-lb
$ kubectl get deployments sample-deployment
```

# loadbalancerから呼び出し
```bash
$ curl -s http://localhost:8080
```

# 削除
```bash
$ kubectl delete deployments sample-deployment
$ kubectl delete services sample-lb
```