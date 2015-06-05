## go-text

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ferhatelmas/go-text)
[![Build Status](https://travis-ci.org/ferhatelmas/go-text.png?branch=master)](https://travis-ci.org/ferhatelmas/go-text)

> Check if a filepath is a text file.

### Install

```
go get github.com/ferhatelmas/go-text
```

### Usage

```go
import "github.com/ferhatelmas/go-text"

text.Is("src/unicorn.go")
//=> true

text.Is("src/unicorn.exe")
//=> false
```

### License

MIT © [ferhat elmas](http://ferhatelmas.com)
