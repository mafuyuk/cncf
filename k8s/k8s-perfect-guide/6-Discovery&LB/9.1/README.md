# service作成
```bash
$ kubectl apply -f ./
```

# サービスの内容確認
```bash
$ kubectl describe svc sample-none-selector
Name:              sample-none-selector
Namespace:         default
Labels:            <none>
Annotations:       kubectl.kubernetes.io/last-applied-configuration:
                     {"apiVersion":"v1","kind":"Service","metadata":{"annotations":{},"name":"sample-none-selector","namespace":"default"},"spec":{"ports":[{"p...
Selector:          <none>
Type:              ClusterIP
IP:                10.105.149.228
Port:              <unset>  8080/TCP
TargetPort:        80/TCP
Endpoints:         192.168.1.1:80,192.168.1.2:80
Session Affinity:  None
Events:            <none>
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