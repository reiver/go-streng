# go-streng

Package **streng** provides a **string** _option type_, _result type_, and _nullable type_, for the Go programming language.

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-streng

[![GoDoc](https://godoc.org/github.com/reiver/go-streng?status.svg)](https://godoc.org/github.com/reiver/go-streng)

## Option Types

Here is how you would declare a **string** _option type_:
```go
var stringOption streng.Option
```

## Result Types

Here is how you would declare a **string** _result type_:
```go
var stringResult streng.Result
```

## Nullable Types

Here is how you would declare a **string** _nullable type_:
```go
var nullableString streng.Nullable
```

## JSON Data Format

If you want to use one of these types with the JSON data format, and in particular using the built-in `"encoding/json"` package,
then you should use `streng.Nullable`.
