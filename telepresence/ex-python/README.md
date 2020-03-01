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
$ docker build -t app-a .
$ telepresence -n app-a --expose 8080:8080 \
  --docker-run --rm -v $(pwd):/opt/app-a app-a python main.py

$ k get deploy,svc                                                                                                                  [/Users/mafuyuk/repo/src/github.com/mafuyuk/cncf/kind]
NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/app-a      1/1     1            1           76s
deployment.extensions/debugkit   1/1     1            1           18h

NAME                 TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
service/app-a        ClusterIP   10.111.175.141   <none>        8080/TCP       76s
service/debugkit     NodePort    10.103.32.219    <none>        80:31380/TCP   18h
service/kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP       18h

$ k get po -o wide                                                                                                                  [/Users/mafuyuk/repo/src/github.com/mafuyuk/cncf/kind]
NAME                       READY   STATUS    RESTARTS   AGE    IP           NODE                 NOMINATED NODE   READINESS GATES
app-a-5b9fbb5b4-vkccq      1/1     Running   0          113s   10.244.1.7   dev-cluster-worker   <none>           <none>
debugkit-798cb977b-xp9pj   1/1     Running   0          18h    10.244.1.2   dev-cluster-worker   <none>           <none>

# アクセス確認用telepresence起動
$ telepresence

# アクセス確認
@kind-dev-cluster|bash-3.2$ curl app-a:8080
app-a@kind-dev-cluster|bash-3.2$ 
```
