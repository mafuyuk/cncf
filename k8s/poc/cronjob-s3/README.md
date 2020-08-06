```bash
$ kind create cluster --name cronjob-s3 --image kindest/node:v1.18.2
$ kind get clusters
cronjpb-s3

$ kctx
Switched to context "cronjob-s3".

$ kustomize build k8s/poc/cronjob-s3 | kubectl apply -f -

$ k get po -w

$ kind delete cluster --name cronjob-s3
```
