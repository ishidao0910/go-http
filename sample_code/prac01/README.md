# Overview
httpサーバーがちゃんと動いているかcurlで確認できます。

# Terminal
1. Run server

```go run main.go```

2. curl command 

```curl -i http://127.0.0.1:8080/lookup/google.com
curl -i http://127.0.0.1:8080/lookup/tenseiblog.com
curl -i http://127.0.0.1:8080/lookup/notadomain```

