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

# ConfigMapの全てのキーを環境変数に渡す Pod
```bash
$ kubectl apply \
  -f ./sample-configmap.yaml \
  -f ./sample-configmap-multi-env.yaml
$ kubectl get pods sample-configmap-multi-env

$ kubectl exec -it sample-configmap-multi-env env

$ kubectl delete pods sample-configmap-multi-env
$ kubectl delete configmaps sample-configmap
```

# ConfigMapの１つのキーをVolumeマウントするPod
```bash
$ kubectl apply \
  -f ./sample-configmap.yaml \
  -f ./sample-configmap-single-volume.yaml
$ kubectl get pods sample-configmap-single-volume

$ kubectl exec -it sample-configmap-single-volume cat /config/nginx-sample.conf
user nginx;
worker_processes auto;
error_log /var/log/nginx/error.log;
pid /run/nginx.pid;

$ kubectl delete pods sample-configmap-single-volume
$ kubectl delete configmaps sample-configmap
```

# ConfigMapの全てのキーをVolumeマウントするPod
```bash
$ kubectl apply \
  -f ./sample-configmap.yaml \
  -f ./sample-configmap-multi-volume.yaml
$ kubectl get pods sample-configmap-multi-volume

$ kubectl exec -it sample-configmap-multi-volume ls /config
connection.max  connection.min  nginx.conf  sample.properties  thread

$ kubectl delete pods sample-configmap-multi-volume
$ kubectl delete configmaps sample-configmap
```