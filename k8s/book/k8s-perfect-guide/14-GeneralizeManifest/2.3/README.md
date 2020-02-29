```bash
$ helm install datadog-chart-sample stable/datadog \
  --values values.yaml
NAME: datadog-chart-sample
LAST DEPLOYED: Wed Jan 29 11:46:40 2020
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None

$ helm uninstall datadog-chart-sample
```