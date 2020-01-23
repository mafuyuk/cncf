```bash
$ kubectl apply -f ./
$ kubectl exec -it sample-env env | grep MAX_CONNECTION
MAX_CONNECTION=100

$ kubectl delete pods sample-env
```


# 追加
```bash
$ kubectl apply -f ./

# podの情報を確認
$ kubectl get pods sample-env -o yaml
※ 抜粋
  nodeName: docker-desktop

$ kubectl delete pods sample-env


# specの値を環境変数に乗せる
$ vim sample-env.yaml
        - name: K8S_NODE
          valueFrom:
            filedRef:
              filedPath: spec.nodeName
を追加
$ kubectl apply -f ./
$ kubectl exec -it sample-env env | grep K8S_NODE
K8S_NODE=docker-desktop
$ kubectl delete pods sample-env
```