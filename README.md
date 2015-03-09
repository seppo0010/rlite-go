# rlite-go

[![Build Status](https://travis-ci.org/seppo0010/rlite-go.svg?branch=master)](https://travis-ci.org/seppo0010/rlite-go)

Java bindings for rlite. For more information about rlite, go to
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

// ...
db, _ := Open(":memory:")
Command(db, []string{"SET", "key", "value"})

reply, err := Command(db, []string{"GET", "key"})
if err != nil {
    t.Error("Got error")
}
if reply != "value" {
    t.Error("Got invalid reply")
}
```
