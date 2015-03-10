# rlite-go

golang bindings for rlite. For more information about rlite, go to
[rlite repository](https://github.com/seppo0010/rlite)

## Installation

First install [rlite](https://github.com/seppo0010/rlite#installation)

```bash
$ go get github.com/seppo0010/rlite-go
```

## Usage

### Using rlite-go

```go
import "rlite"

import "github.com/seppo0010/rlite-go/rlite"
import "fmt"

func main () {
    db, _ := rlite.Open(":memory:")
    rlite.Command(db, []string{"SET", "key", "value"})

    reply, err := rlite.Command(db, []string{"GET", "key"})
    if err != nil {
        // ...
    }
    if reply != "value" {
        // ...
    }
    fmt.Println(reply)
}
```
