# Overview
基本的な認証がgoで実装できているかを確認する。
Demonstrate how to setup AuthBasicMiddleware as a pre-routing middleware.

## endpoint
- GET /

## Check
serverを立ち上げたあと、curlコマンドで確認可能。

```
//　こちらのコマンドでは、auth情報が含まれていないためエラーが出る
curl -i http://127.0.0.1:8080/

//--------------------------------
HTTP/1.1 401 Unauthorized
Content-Type: application/json; charset=utf-8
Www-Authenticate: Basic realm=test zone
X-Powered-By: go-json-rest
Date: Wed, 06 Jul 2022 05:03:08 GMT
Content-Length: 31

{
  "Error": "Not Authorized"
}
```

```
// これにはauthが含まれているので200responseが返る
curl -i -u admin:admin http://127.0.0.1:8080/

//--------------------------------
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Wed, 06 Jul 2022 05:03:17 GMT
Content-Length: 27

{
  "Body": "Hello World"
}
```
