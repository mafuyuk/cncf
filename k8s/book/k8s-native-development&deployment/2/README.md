```bash
$ kubectl create deploy --image=amaya382/k8s-debugkit debugkit \
  --dry-run -o=yaml > deployment-by-dry-run.yaml
```
