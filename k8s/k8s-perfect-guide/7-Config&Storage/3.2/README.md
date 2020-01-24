# 3.1.1
```bash
$ kubectl create secret generic --save-config sample-db-auth \
  --from-env-file ./env-secret.txt

$ kubectl get secrets

$ kubectl delete secrets sample-db-auth
```

# 3.1.2
```bash
$ kubectl apply -f ./

$ kubectl get secrets sample-db-auth
$ kubectl describe secrets sample-db-auth

$ kubectl delete secrets sample-db-auth
```