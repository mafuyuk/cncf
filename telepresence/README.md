# 使い方
## ローカルでサービスを起動しつつpodとして認識させる
```bash
# ctxの向け先を変える
$ kctx

# python2系がある前提
$ telepresence -n example-server --expose 8080:8080 --run python -m SimpleHTTPServer 8080

# logが不要な場合は --logfile=/dev/null をつければよい

$ k get deploy,svc
NAME                                   READY   UP-TO-DATE   AVAILABLE   AGE
deployment.extensions/debugkit         1/1     1            1           78m
deployment.extensions/example-server   1/1     1            1           57s

NAME                     TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)        AGE
service/debugkit         NodePort    10.103.32.219    <none>        80:31380/TCP   78m
service/example-server   ClusterIP   10.106.116.143   <none>        8080/TCP       56s
service/kubernetes       ClusterIP   10.96.0.1        <none>        443/TCP        78m

$ k get po -o wide
NAME                              READY   STATUS    RESTARTS   AGE     IP           NODE                 NOMINATED NODE   READINESS GATES
debugkit-798cb977b-xp9pj          1/1     Running   0          79m     10.244.1.2   dev-cluster-worker   <none>           <none>
example-server-569dd46b97-9mtzw   1/1     Running   0          2m18s   10.244.1.5   dev-cluster-worker   <none>           <none>
$ curl example-server:8080                                                                                                          [/Users/mafuyuk/repo/src/github.com/mafuyuk/cncf/kind]
<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN"><html>
<title>Directory listing for /</title>
<body>
<h2>Directory listing for /</h2>
<hr>
<ul>
<li><a href="my-k8s.yaml">my-k8s.yaml</a>
<li><a href="README.md">README.md</a>
<li><a href="telepresence.log">telepresence.log</a>
</ul>
<hr>
</body>
</html>

```

## ローカルをコンテナを起動しつつpodとして認識させる
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

# 参考リンク
- https://www.telepresence.io/
