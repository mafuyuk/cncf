https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#create-cronjob-v1beta1-batch

# CronJobの作成
$ kubectl apply -f sample-cronjob.yaml

# CronJobの一覧確認
$ kubectl get cronjobs

# Jobの一覧確認
$ kubectl get jobs

# CronJobの削除
$ kubectl delete cronjobs sample-cronjob