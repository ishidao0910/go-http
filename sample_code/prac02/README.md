# Overview
国の名前をjson形式でやりとりするREST APIを実装

## endpoint
- GET /countries/:code
- GET/POST /countries
- DELETE /countries/:code

## Check
serverを立ち上げたあと、curlコマンドで確認可能。

```
curl -i -H 'Content-Type: application/json' \
    -d '{"Code":"FR","Name":"France"}' http://127.0.0.1:8080/countries
curl -i -H 'Content-Type: application/json' \
    -d '{"Code":"US","Name":"United States"}' http://127.0.0.1:8080/countries
curl -i http://127.0.0.1:8080/countries/FR
curl -i http://127.0.0.1:8080/countries/US
curl -i http://127.0.0.1:8080/countries
curl -i -X DELETE http://127.0.0.1:8080/countries/FR
curl -i http://127.0.0.1:8080/countries
curl -i -X DELETE http://127.0.0.1:8080/countries/US
curl -i http://127.0.0.1:8080/countries
```
