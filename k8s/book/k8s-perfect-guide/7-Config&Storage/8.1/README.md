```bash
$ gcloud compute disks create --size=10GB sample-gce-pv --zone asis-northeast1-a

$ kubectl apply -f ./
$ kubectl get persistentvolumes sample-pv
$ kubectl get persistentvolumeclaims sample-pvc

$ kubectl delete persistentvolumeclaims sample-pvc

# 一旦どうなっているのか確認
$ kubectl get persistentvolumes sample-pv
$ kubectl get persistentvolumeclaims sample-pvc

$ kubectl delete persistentvolumes sample-pv
```
