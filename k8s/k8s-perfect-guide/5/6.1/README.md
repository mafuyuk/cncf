# StatefulSetの作成
$ kubectl apply -f sample-statefulset.yaml

# StatefulSetの確認
$ kubectl get statefulsets

# Podの一覧
$ kubectl get pods -o wide

# StatefulSetに紐づくPersistentVolumeClaimの確認
$ kubectl get persistentvolumeclaims

# StatefulSetに紐づくPersistentVolumeの確認
$ kubectl get persistentvolumes

# StatefulSetの削除
$ kubectl delete statefulsets samaple-statefulset
