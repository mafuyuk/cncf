# Dockerレジストリの認証情報のSecret作成
```bash
$ kubectl create secret docker-registry \
  --save-config sample-registry-auth \
  --docker-server=REGISTRY_SERVER \
  --docker-username=REGISTRY_USER \
  --docker-password=REGISTRY_USER_PASSWORD \
  --docker-email=REGISTRY_USER_EMAIL
```

# base64でエンコードされたdockercfg形式のデータ
```bash
$ kubectl get secrets -o json sample-registry-auth | jq .data
{
  ".dockerconfigjson": "eyJhdXRocyI6eyJSRUdJU1RSWV9TRVJWRVIiOnsidXNlcm5hbWUiOiJSRUdJU1RSWV9VU0VSIiwicGFzc3dvcmQiOiJSRUdJU1RSWV9VU0VSX1BBU1NXT1JEIiwiZW1haWwiOiJSRUdJU1RSWV9VU0VSX0VNQUlMIiwiYXV0aCI6IlVrVkhTVk5VVWxsZlZWTkZVanBTUlVkSlUxUlNXVjlWVTBWU1gxQkJVMU5YVDFKRSJ9fX0="
}
```

# イメージをSecret情報を使用して取得
```bash
$ kubectl apply -f ./
```

# 確認と削除
```bash
$ kubectl get secrets sample-registry-auth
NAME                   TYPE                             DATA   AGE
sample-registry-auth   kubernetes.io/dockerconfigjson   1      3m39s

$ kubectl delete secrets sample-registry-auth
$ kubectl delete pods sample-pull-secret
```