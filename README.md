# KeepCase

This module provides you with a `map[string]T`-like struct with some pretty specific behaviour.

You shouldn't need to use this. My very particular use case is to override HTTP headers while not affecting the existing casing.

## Behaviour

It fulfills the following requirements:

- Set key and value, maintaining the casing of the key.
- If a key is already set, setting it again maintains the existing casing of the key.
- Retrieve the value for a key in a case-insensitive way.

The implementation is thead-safe.

## Sample
Its behaviour is probably best demonstrated with an example. Note how:

- The `Content-Type` key casing is kept intact.
- The value for the `Content-Type` key is overwritten, even though the input key is `content-type`.
- The value can be retrieved with both `Content-Type` and `content-type`.

```go
import "github.com/aertje/keepcase"

m := keepcase.NewMap[string]()

v, ok := m.Get("content-type") // => "", false
l := m.Len() // => 0
mm := m.AsMap() // => {}

m.Set("Content-Type", "application/json")
v, ok = m.Get("content-type") // => "application/json", true
l = m.Len() // => 1
mm = m.AsMap() // => {"Content-Type": "application/json"}

m.Set("content-type", "text/html")

v, ok = m.Get("content-type") // => "text/html", true
v, ok = m.Get("Content-Type") // => "text/html", true
mm = m.AsMap() // => {"Content-Type": "text/html"}
l = m.Len() // => 1
```

## Use it
Since it uses generics, Go 1.18 is required.

```sh
go get github.com/aertje/keepcase
```