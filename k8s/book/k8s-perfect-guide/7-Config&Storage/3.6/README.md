# 暗号化して上書き
```bash
$ kubesec encrypt \
  --key=arn:aws:kms:us-west-2:**********:key/00008853-64e6-43e8-a303-125f7394e902 -i sample-db-auth.yaml
```

# 復号して上書き
```bash
$ kubesec decrypt -i sample-db-auth.yaml 
```
# Secretの作成
```bash
$ kubesec decrypt sample-db-auth.yaml | kubectl apply -f -
```

# Secretの削除
```bash
$ kubesec decrypt sample-db-auth.yaml | kubectl delete -f -
```