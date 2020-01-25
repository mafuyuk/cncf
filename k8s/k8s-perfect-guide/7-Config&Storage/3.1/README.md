
```bash
$ kubectl create secret generic --save-config sample-db-auth \
  --from-env-file ./env-secret.txt

$ kubectl get secrets

$ kubectl delete secrets sample-db-auth
```