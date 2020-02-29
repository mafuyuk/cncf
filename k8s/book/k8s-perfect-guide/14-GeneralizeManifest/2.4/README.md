```bash
$ helm pull stable/datadog --untar
$ helm template datadog/
$ helm template datadog/ -x templates/serviceaccount.yaml

$ helm install sample-mychart mychart
$ kubectl get configmaps mychart-configmap -o yaml
$ helm uninstall sample-mychart
```

