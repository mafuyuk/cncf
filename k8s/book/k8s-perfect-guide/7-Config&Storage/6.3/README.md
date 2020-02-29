```bash
$ kubectl apply -f ./
$ kubectl describe pods sample-downward-api
Volumes:
  downward-api-volume:
    Type:  DownwardAPI (a volume populated by information about the pod)
    Items:
      metadata.name -> podname
      requests.cpu -> cpu-request


# ファイルの確認
$ kubectl exec -it sample-downward-api ls /srv

$ kubectl delete pods sample-downward-api
```
