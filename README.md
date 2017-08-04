# go-updater

[![Build Status](https://travis-ci.org/swhite24/go-updater.svg?branch=master)](https://travis-ci.org/swhite24/go-updater)

Provides a simple mechanism to conditionally update a struct's values from a map.

## Installation

```
go get github.com/swhite24/go-updater
```

## Usage

```go
import (
  "github.com/swhite24/go-updater/updater"
)

type Foo struct {
  Name          string `json:"name" update:"true"`
  Age           int    `json:"age" update:"true"`
  FavoriteColor string `json:"favorite_color"`
}

func main() {
  f := &Foo{Name: "jim", Age: 50, FavoriteColor: "red"}
  update := map[string]interface{}{
    "name": "jimbob",
    "age": 51,
    "favorite_color": "blue",
  }

  updater.Struct(f, update)
  fmt.Println(f)
  // &{jimbob 51 red}
}
```
