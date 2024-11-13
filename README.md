# gocollect

```
go build -o target/simple -cover -covermode=atomic examples/simple/main.go
```

```
GOCOVERDIR=$(pwd) ./simple
```

```
curl localhost:8082/collect
```

```
go tool covdata textfmt -i=$(pwd) -o cover.out
go tool cover -html=cover.out -o cover.html
```
