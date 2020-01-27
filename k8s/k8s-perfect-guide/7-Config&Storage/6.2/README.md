```bash
$ kubectl apply -f ./
$ kubectl describe pods sample-hostpath
Volumes:
  hostpath-sample:
    Type:          HostPath (bare host directory volume)
    Path:          /etc
    HostPathType:  DirectoryOrCreate

# ホストOSのイメージを確認
$ kubectl exec -it sample-hostpath cat /srv/os-release | grep PRETTY_NAME
PRETTY_NAME="Docker Desktop"

# コンテナOSのイメージを確認
$ kubectl exec -it sample-hostpath cat /etc/os-release | grep PRETTY_NAME
PRETTY_NAME="Debian GNU/Linux 9 (stretch)"

$ kubectl delete pods sample-hostpath
```
