# 使い方
## 事前準備
```bash
$ k create deploy --image=amaya382/k8s-debugkit debugkit
$ k expose deploy debugkit --port=80 --type=NodePort
$ k edit svc debugkit # spec.ports[0].nodePortを31380に更新
```

## ローカルでコンテナを起動しつつpodとして認識させる
```bash
# サーバー用telepresence起動
$ docker build -t app-b-builder --target builder .
$ telepresence --expose=8080:8080 -n app-b \
  --docker-run --rm -it -v $(pwd):/opt/app-b app-b-builder sh
/opt/app-b # APPB_PORT=8080 go run main.go

$ k get deploy,svc                                                                                                              
NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/app-b      1/1     1            1           3m50s
deployment.extensions/debugkit   1/1     1            1           19h

NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
service/app-b        ClusterIP   10.102.202.175   <none>        8080/TCP       3m50s
service/debugkit     NodePort    10.103.32.219    <none>        80:31380/TCP   19h
service/kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP        19h


$ k get po -o wide                                                                                                                  
NAME                       READY   STATUS    RESTARTS   AGE     IP           NODE                 NOMINATED NODE   READINESS GATES
app-b-5cbf587784-br9bd     1/1     Running   0          4m21s   10.244.1.9   dev-cluster-worker   <none>           <none>
debugkit-798cb977b-xp9pj   1/1     Running   0          19h     10.244.1.2   dev-cluster-worker   <none>           <none>

# アクセス確認用telepresence起動
$ telepresence

# アクセス確認
@kind-dev-cluster|bash-3.2$ curl app-b:8080
app-b@kind-dev-cluster
```

