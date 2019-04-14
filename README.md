# PKCS5 for Go

Public Key Cryptography Standards #5


```bash
go get github.com/trylife/go-PKCS5
```


```go
import "github.com/trylife/go-PKCS5"

padded := PKCS5.Padding([]byte("hi"), 8)
original := PKCS5.RemovePadding(padded)
fmt.Println(padded)
fmt.Println(original)
// Output:
// [104 105 6 6 6 6 6 6]
// [104 105]

```
