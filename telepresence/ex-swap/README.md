## 事前準備
```bash
$ k apply -f test.yaml
```

## Deploymentを置き換える
```bash
$ telepresence --swap-deployment=debugkit --docker-run --rm -it amaya382/k8s-debugkit bash

root@debugkit-33731d68433243ac907fbd571b0b9635-545c75dbd8-tkn5c:/opt/k8s-debugkit# echo ${TELEPRESENCE_ROOT}
/tmp/tel-1cqlr6tr/fs

root@debugkit-33731d68433243ac907fbd571b0b9635-545c75dbd8-tkn5c:/opt/k8s-debugkit# echo ${TELEPRESENCE_MOUNTS}
/etc/config:/var/run/secrets/kubernetes.io/serviceaccount

root@debugkit-33731d68433243ac907fbd571b0b9635-545c75dbd8-tkn5c:/opt/k8s-debugkit# cat ${TELEPRESENCE_ROOT}/etc/config/config.toml
[foo]
bar = "hello"

root@debugkit-33731d68433243ac907fbd571b0b9635-545c75dbd8-tkn5c:/opt/k8s-debugkit# echo $MY_ENV
cool!
```
