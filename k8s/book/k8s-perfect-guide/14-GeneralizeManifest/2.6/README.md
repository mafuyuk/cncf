```bash
$ helm create mychart

$ helm install mychart --dry-run --debug ./mychart
# Source: mychart/templates/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: mychart-configmap
data:
  myvalue: Hello World!

```