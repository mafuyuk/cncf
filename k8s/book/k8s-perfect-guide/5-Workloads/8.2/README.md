# CronJobの一時停止
メンテナンスなどでJobが作成されてほしくない場合にsuspend(一時停止)する
```bash
# クライアントからのCronJobの一時停止
$ kubectl patch cronjob sample-cronjob -p '{"spec":{"suspend": true}}'
```