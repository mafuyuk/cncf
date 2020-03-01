# 使い方
## 事前準備
```bash
$ k create deploy --image=amaya382/k8s-debugkit debugkit
$ k expose deploy debugkit --port=80 --type=NodePort
$ k edit svc debugkit # spec.ports[0].nodePortを31380に更新
```

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
$ curl example-server:8080
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
