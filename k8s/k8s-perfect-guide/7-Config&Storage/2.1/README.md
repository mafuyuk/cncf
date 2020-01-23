```bash
$ kubectl apply -f ./
$ kubectl exec -it sample-env env | grep MAX_CONNECTION
MAX_CONNECTION=100
$ kubectl delete pods sample-env
```