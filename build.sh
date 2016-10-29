echo "$(date '+%H:%M:%S') 开始编译"
gofmt -w .
go build -o incr main.go
echo "$(date '+%H:%M:%S') 结束编译"
