```bash
$ kind create cluster --name 9.5-cluster --image kindest/node:v1.18.2
$ kind get clusters
9.5-cluster
$ kctx
Switched to context "kind-9.5-cluster".

$ k apply -f k8s/book/k8s-native-development\&deployment/9/9.5/demo-cron.yaml
$ k get cronjob
NAME        SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
demo-cron   */1 * * * *   False     0        33s             60s

$ k get po -w
demo-cron-1596746580-6lg67   0/1     Pending     0          0s
demo-cron-1596746580-6lg67   0/1     ContainerCreating   0          0s
demo-cron-1596746580-6lg67   1/1     Running             0          2s
demo-cron-1596746580-6lg67   0/1     Completed           0          7s
demo-cron-1596746400-2dv7b   0/1     Terminating         0          3m11s

$ kind delete cluster --name 9.5-cluster
```
