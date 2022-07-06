# Overview
httpサーバーがちゃんと動いているかcurlで確認できます。

# Terminal
1. Run server

```go run main.go```

2. curl command 

```
curl -i http://127.0.0.1:8080/lookup/google.com
curl -i http://127.0.0.1:8080/lookup/tenseiblog.com
curl -i http://127.0.0.1:8080/lookup/notadomain
```

![スクリーンショット 2022-07-05 21 49 20](https://user-images.githubusercontent.com/73809994/177331441-b81f21ec-950e-463c-8f06-e5df355fc4e6.png)
