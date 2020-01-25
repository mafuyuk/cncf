# 環境変数として渡す
## Secretの1つのキーを環境変数に渡す
```bash
$ kubectl apply \
  -f ./sample-db-auth.yaml \
  -f ./sample-secret-single-env.yaml

$ kubectl get secrets sample-db-auth
$ kubectl get pods sample-secret-single-env
$ kubectl exec -it sample-secret-single-env env | grep DB_USERNAME
DB_USERNAME=root
$ kubectl delete pods sample-secret-single-env
$ kubectl delete secrets sample-db-auth
```

# Volumeとしてマウントする
## Secretの１つのキーをマウントするPod
```bash
$ kubectl apply \
  -f ./sample-db-auth.yaml \
  -f ./sample-secret-single-volume.yaml

$ kubectl get secrets sample-db-auth
$ kubectl get pods sample-secret-single-volume
$ kubectl exec -it sample-secret-single-volume cat /config/username.txt

$ kubectl delete pods sample-secret-single-volume
$ kubectl delete secrets sample-db-auth
```