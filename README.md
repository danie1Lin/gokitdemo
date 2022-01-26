# Go kit developing flow by using truss

## 評估
1. 決定 project structure
2. 定義標準 error response 格式，http and Grpc

## Some Issues 
1. In latest Go version 1.17, truss will throw error like: `Imported file not found`.
Because the go package location is changed, no longer storing in `$GOPATH/src`. 
but truss is still using that path.

- [ ] try using buf to acquire deps.

2. need to modify the api response case json:"Event,omitempty"`

3. need to modify cmd/liveevent/main.go or svc/server.go
