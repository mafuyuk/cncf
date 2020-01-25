```bash
$ openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
  -keyout ~/tls.key -out ~/tls.crt -subj "/CN=sample1.example.com"

# TLS Secretの作成
$ kubectl create secret tls --save-config tls-sample --key ~/tls.key --cert ~/tls.crt

$ kubectl get secret tls-sample
$ kubectl describe secret tls-sample

$ rm ~/tls.key ~/tls.crt
$ kubectl delete secret tls-sample
```