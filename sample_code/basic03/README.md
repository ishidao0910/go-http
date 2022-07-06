# Overview
ユーザー名をjson形式でやりとりするREST APIを実装

## endpoint
- GET/PUT/DELETE /users/:code
- GET/POST /users

## Check
serverを立ち上げたあと、curlコマンドで確認可能。

```
curl -i -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine"}' http://127.0.0.1:8080/users
curl -i http://127.0.0.1:8080/users/0
curl -i -X PUT -H 'Content-Type: application/json' \
    -d '{"Name":"Antoine Imbert"}' http://127.0.0.1:8080/users/0
curl -i -X DELETE http://127.0.0.1:8080/users/0
curl -i http://127.0.0.1:8080/users
```
