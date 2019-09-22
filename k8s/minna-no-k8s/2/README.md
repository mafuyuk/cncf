```bash
$ kubectl kustomize . | kubectl apply -f - --prune -l hello=world

$ kubectl get pod
$ kubectl get service

$ kubectl kustomize . | kubectl delete -f -
```