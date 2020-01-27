```bash
$ gcloud compute disks create --size=10GB sample-gce-pv --zone asis-northeast1-a

$ kubectl apply -f ./
$ kubectl get persistentvolumes

$ kubectl delete persistentvolumes sample-pv
```
