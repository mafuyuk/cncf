# service作成
```bash
$ kubectl apply -f ./
```

# 内容確認
```bash
$ kubectl get services
NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes               ClusterIP   10.96.0.1       <none>        443/TCP          8d
sample-ingress-default   NodePort    10.99.249.242   <none>        8888:30169/TCP   39s
sample-ingress-svc-1     NodePort    10.103.122.55   <none>        8888:30997/TCP   39s
sample-ingress-svc-2     NodePort    10.100.181.86   <none>        8888:32122/TCP   39s

$ kubectl get pods                                                                [/Users/mafuyuk/repo/src/github.com/mafuyuk/cncf/k8s/k8s-perfect-guide/6-Discovery&LB/10.4]
NAME                     READY   STATUS    RESTARTS   AGE
sample-ingress-apps-1    1/1     Running   0          106s
sample-ingress-apps-2    1/1     Running   0          106s
sample-ingress-default   1/1     Running   0          106s

$ kubectl get ingresses
NAME             HOSTS                ADDRESS   PORTS     AGE
sample-ingress   sample.example.com             80, 443   18s

```

# index.htmlの追加
```bash
$ kubectl exec -it sample-ingress-apps-1 -- mkdir /usr/share/nginx/html/path1/
$ kubectl exec -it sample-ingress-apps-1 -- cp /etc/hostname /usr/share/nginx/html/path1/index.html

$ kubectl exec -it sample-ingress-apps-2 -- mkdir /usr/share/nginx/html/path2/
$ kubectl exec -it sample-ingress-apps-2 -- cp /etc/hostname /usr/share/nginx/html/path2/index.html

$ kubectl exec -it sample-ingress-default -- cp /etc/hostname /usr/share/nginx/html/index.html
```

# シークレットの作成
```bash
# 自己証明書の作成
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ~/tls.key -out ~/tls.crt -subj "/CN=sample.example.com"

# Secretの作成(証明書ファイルを指定した場合)
$ kubectl create secret tls --save-config tls-sample --key ~/tls.key --cert ~/tls.crt
```



# 動作確認
```bash
# L7ロードバランサの仮想IPを環境変数に保存
$ INGRESS_IP=`kubectl get ingresses sample-ingress \
  -o jsonpath='{.status.loadBalancer.ingress[0].ip}'`

# Ingressリソース経由のHTTPリクエスト
$ curl http://${INGRESS_IP}/path1/index.html -H  "Host: sample.example.com"

# Ingressリソース経由のHTTPリクエスト
$ curl http://${INGRESS_IP}/path2/index.html -H  "Host: sample.example.com"

# Ingressリソース経由のHTTPリクエスト
$ curl http://${INGRESS_IP}/index.html -H  "Host: sample.example.com"

# Ingressリソース経由のHTTPリクエスト insecure
$ curl https://${INGRESS_IP}/path1/index.html -H  "Host: sample.example.com" --insecure
```
# 削除
```bash
$ kubectl delete services sample-ingress-svc-1
$ kubectl delete services sample-ingress-svc-2
$ kubectl delete services sample-ingress-default

$ kubectl delete pods sample-ingress-apps-1
$ kubectl delete pods sample-ingress-apps-2
$ kubectl delete pods sample-ingress-default

$ kubectl delete secrets tls-sample
$ rm ~/tls.crt
```