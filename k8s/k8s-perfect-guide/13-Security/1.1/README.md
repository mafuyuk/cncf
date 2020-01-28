```bash
$ kubectl apply -f ./
$ kubectl get serviceaccounts sample-serviceaccount -o yaml
secrets:
- name: sample-serviceaccount-token-vn22f


# ServiceAccountに紐づくSecret
$ kubectl get secrets sample-serviceaccount-token-vn22f -o yaml

# トークンがVolumeとしてマウントされていることを確認
$ kubectl get pods sample-serviceaccount-pod -o yaml
  volumes:
  - name: sample-serviceaccount-token-vn22f
    secret:
      defaultMode: 420
      secretName: sample-serviceaccount-token-vn22f

$ kubectl delete pods sample-serviceaccount-pod
$ kubectl delete serviceaccounts sample-serviceaccount
```