# Package json

http://golang.org/pkg/encoding/json/

```go
import "encoding/json"
```

Package json implements encoding and decoding of JSON objects.

- [`Marshal`](http://golang.org/pkg/encoding/json/#Marshal)
encoding to json

  ```go
  func Marshal(v interface{}) ([]byte, error)
  ```

- [`Unmarshal`](http://golang.org/pkg/encoding/json/#Unmarshal)
decoding from json

  ```go
  func Unmarshal(data []byte, v interface{} ) error
  ```
