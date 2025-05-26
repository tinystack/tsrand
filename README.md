# tsrand
Go rand Package

[![Go Report Card](https://goreportcard.com/badge/github.com/tinystack/tsrand)](https://goreportcard.com/report/github.com/tinystack/tsrand)
![Go Version](https://img.shields.io/badge/go%20version-%3E=1.18-61CFDD.svg?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/mod/github.com/tinystack/tsrand)](https://pkg.go.dev/mod/github.com/tinystack/tsrand)

## 安装

go get -u github.com/tinystack/tsrand

## API
### 生成数字类型随机整数
- rand.Int32()
- rand.Int64()
- rand.Uint32()
- rand.Uint64()
- rand.Int()
- rand.Uint()

### 在范围内生成随机整数
- rand.RangeInt(min, max int) int
- rand.RangeInt64(min, max int64) int64

### 生成随机串
- rand.String(len int) string
- rand.VisibleString(len int) string
- rand.UUID() string

## 示例
```go
import "github.com/tinystack/tsrand"

func main() {
	fmt.Printf("rand int32 is %d\n", rand.Int32())
    fmt.Printf("rand int64 is %d\n", rand.Int64())
    fmt.Printf("rand int is %d\n", rand.Int())
    
    fmt.Printf("rand normal strings is %s\n", rand.String(40))
    fmt.Printf("rand visible strings is %s\n", rand.VisibleString(40))
    fmt.Printf("uuid is %s\n", rand.UUID())
}
```

```shell
output:

rand int32 is 449921072
rand int64 is 276967040279418390
rand int is 3327377255262805722
rand normal strings is QRzNfQwhRbGd4V27y76gV5acfVHmK2gOmeusVhCc
rand visible strings is jVezpsYggbxqHHRswbBkM7TYZAPQE4R2xpYqZh8V
uuid is 12da435c-f61f-4f3e-aade-f43296caa9fe

```