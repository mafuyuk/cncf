```bash
$ kubectl apply -f ./
$ kubectl describe pods sample-projected
Volumes:
  projected-volume:
    Type:                Projected (a volume that contains injected data from multiple sources)
    SecretName:          sample-db-auth
    SecretOptionalName:  <nil>
    ConfigMapName:       sample-configmap
    ConfigMapOptional:   <nil>

# ファイルの確認
$ kubectl exec -it sample-projected ls /srv
configmap  podname  secret

$ kubectl delete pods sample-projected
$ kubectl delete configmaps sample-configmap
$ kubectl delete secret sample-db-auth
```
