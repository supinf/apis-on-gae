# ローカル起動

GAE に合わせて全て Docker コンテナ内で動作すべきところだが  
サンプルであるため簡易的にローカルに go 環境があるものとして

## go-swagger を使い、swagger.yaml からのソース生成

```
$ swagger generate server -A demo-apis -P models.Principal -t api
```

## golang/dep などを使い、依存性を解決

```
$ pushd api
$ dep init
$ popd
```

## 起動

```
$ go run api/cmd/demo-apis-server/main.go --host 0.0.0.0 --port 8080
```

## e2e テスト

```
$ newman run --environment test/postman_environment.json test/postman_collection.json
```

# GAE での起動

以下を全て CI で実行する

## Build, Test, Ship

実際にはバージョンは git のコミットハッシュ、ブランチ名やタグなどを利用

```
$ version=$( date +%Y%m%d%H%M ) && echo ${version}
$ docker build -t asia.gcr.io/${PROJECT_ID}/apis-on-gae:${version} .
```

テストが通ったら

```
$ # do tests
```

コンテナレジストリに push

```
$ token=$( gcloud auth application-default print-access-token )
$ docker login -u oauth2accesstoken -p "${token}" https://gcr.io
$ docker push asia.gcr.io/${PROJECT_ID}/apis-on-gae:${version}
```

## Cloud Endpoints の設定

swagger.yaml のバージョンなどを更新して

```
$ sed -i -e "s/^  version: \'.*\'/  version: \'${version}\'/" swagger.yaml
$ sed -i -e "s/^host: your-project-id/host: ${PROJECT_ID}/" swagger.yaml
```

デプロイ

```
$ gcloud service-management deploy swagger.yaml
```

## App Engine の更新

app.yaml の Cloud Endpoints バージョンなどを更新

```
$ config_id=$( gcloud service-management configs list \
    --service ${PROJECT_ID}.appspot.com \
    --format=json \
    | jq "map(select(.name | index(\"${PROJECT_ID}.appspot.com\")).id)" \
    | jq -r "sort | .[-1]" \
  ) && echo ${config_id}
$ sed -i -e "s/^  config_id: \'.*\'/  config_id: \'${config_id}\'/" app.yaml
$ sed -i -e "s/^  name: \'your-project-id/  name: \'${PROJECT_ID}/" app.yaml
```

デプロイ

```
$ gcloud app deploy app.yaml --no-promote --no-stop-previous-version \
    --image-url asia.gcr.io/${PROJECT_ID}/apis-on-gae:${version} \
    --version ${version} --quiet
```

## e2e テスト

https://${PROJECT_ID}-dot-tokyo-devops.appspot.com/ に対して e2s テストを実行

## トラフィックの切り替え

```
$ gcloud app services set-traffic default --splits ${version}=1 --quiet
```

# 状況確認

## ログの確認

```
$ gcloud app logs tail --service=default --version=${version}
```

## Cloud Endpoints で状況を確認

```
$ open "https://console.cloud.google.com/endpoints/api/${PROJECT_ID}.appspot.com/overview"
```

## GAE, Cloud Endpoints 過去バージョンの削除
