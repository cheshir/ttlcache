[![Build Status](https://travis-ci.com/cheshir/ttlcache.svg?branch=main)](https://travis-ci.com/cheshir/ttlcache)
[![codecov](https://codecov.io/gh/cheshir/ttlcache/branch/main/graph/badge.svg?token=WsaH2t9dGh)](https://codecov.io/gh/cheshir/ttlcache)
[![Go Report Card](https://goreportcard.com/badge/cheshir/ttlcache)](https://goreportcard.com/report/github.com/cheshir/ttlcache)
[![GoDoc](https://godoc.org/github.com/cheshir/ttlcache?status.svg)](https://godoc.org/github.com/cheshir/ttlcache)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/avelino/awesome-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cheshir/go-mq/blob/master/LICENSE)

# ttlcache

## About

**ttlcache** – is a simple and efficient in-memory key value storage with TTL for each record.

The key is of uint64 type. Library provides wrappers for common types:

* IntKey
* ByteKey
* Int8Key
* Uint8Key
* Int16Key
* Uint16Key
* Int32Key
* Uint32Key
* Int64Key
* Uint64Key
* BytesKey ([]byte)
* StringKey
* AnyKey (interface{})

`AnyKey` is not suitable for _very_ intensive usage. Consider writing your own hash function if your keys are complex types, 
and you faced performance degradation.

## Installation

`go get -u github.com/cheshir/ttlcache`

## Usage

```go
package main

import (
    "github.com/cheshir/ttlcache"
)

const (
    minute = 60 * time.Second
    hour = 60 * minute
)

func main() {
    // How often we need to stop the world and remove outdated records.
    resolution := minute

    cache := ttlcache.New(resolution)
    cache.Set(ttlcache.StringKey("some key"), "value", hour)
    value, ok := cache.Get(ttlcache.StringKey("some key"))
    if !ok {
        // there is no record with key "some key" in the cache. Probably it has been expired.
    }

    fmt.Println(value.(string)) // This is necessary type assertion, because returned value is of interface{} type.
}
```

## Performance

If you're interested in benchmarks you can check them in repository.
Just play with numbers and types and check that library is suitable for your purposes.

`go test -bench=. -benchmem`

For those of us who wants to get some numbers without downloading unknown stuff (MacBook Pro 16"):

```go
BenchmarkCache_Set_100-16           8959221               125 ns/op
BenchmarkCache_Set_1000-16          9177854               123 ns/op
BenchmarkCache_Set_10000-16         9247304               131 ns/op
BenchmarkCache_Get_100-16           50562800              23.9 ns/op
BenchmarkCache_Get_1000-16          47270793              26.9 ns/op
BenchmarkCache_Get_10000-16         42578484              27.7 ns/op

```
