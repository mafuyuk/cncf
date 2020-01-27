# ConfigMapの１つのキーを環境変数に渡す Pod
```bash
$ kubectl apply \
  -f ./sample-configmap.yaml \
  -f ./sample-configmap-single-env.yaml
$ kubectl get pods sample-configmap-single-env

$ kubectl exec -it sample-configmap-single-env env | grep CONNECTION_MAX
CONNECTION_MAX=100

$ kubectl delete pods sample-configmap-single-env
$ kubectl delete configmaps sample-configmap
```