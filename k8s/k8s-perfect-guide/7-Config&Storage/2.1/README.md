```bash
$ kubectl apply -f ./
$ kubectl exec -it sample-env env | grep MAX_CONNECTION
MAX_CONNECTION=100

$ kubectl delete pods sample-env
```


# Podの情報から値を環境変数に乗せる
対象の記述
```bash
        - name: K8S_NODE
          valueFrom:
            filedRef:
              filedPath: spec.nodeName
```
コマンド
```bash
$ kubectl apply -f ./

# podの情報を確認
$ kubectl get pods sample-env -o yaml
※ 抜粋
  nodeName: docker-desktop
$ kubectl exec -it sample-env env | grep K8S_NODE
K8S_NODE=docker-desktop
$ kubectl delete pods sample-env
```

# コンテナの情報から値を環境変数に乗せる
対象の記述
```bash
        - name: K8S_NODE
          valueFrom:
            filedRef:
              filedPath: spec.nodeName
```
コマンド
```bash
$ kubectl apply -f ./

$ kubectl exec -it sample-env env | grep CPU_REQUESTS
CPU_REQUESTS=0
$ kubectl exec -it sample-env env | grep CPU_LIMITS
CPU_LIMITS=4

$ kubectl delete pods sample-env
```