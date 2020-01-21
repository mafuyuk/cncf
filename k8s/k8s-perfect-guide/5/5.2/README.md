# DeamonSetの作成
$ kubectl apply -f sample-ds-ondelete.yaml

# Podの一覧
$ kubectl get pods -o wide
# daemonSetsの一覧
$ kubectl get daemonSets

# daemonSetsの削除
$ kubectl delete daemonSets sample-ds-ondelete
