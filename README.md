# ローカル起動

GAE に合わせて全て Docker コンテナ内で動作すべきところだが  
サンプルであるため簡易的にローカルに go 環境があるものとして

## go-swagger を使い、swagger.yaml からのソース生成

```
$ swagger generate server -A demo-apis -P models.Principal -t api
```

## golang/dep を使い、依存性を解決

```
$ pushd api
$ dep ensure
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
