# goption

A generic `Option[T]` type for Go, inspired by Rust's `Option`.

## Installation

```bash
go get github.com/0xkonsti/goption
```

## Usage

```go
import "github.com/0xkonsti/goption"
```

### Basic

```go
some := goption.Some(42)
none := goption.None[int]()

fmt.Println(some.IsSome()) // true
fmt.Println(none.IsNone()) // true
```

### Unwrapping

```go
val := some.Unwrap()        // 42
val, ok := some.UnwrapOk()  // 42, true
val, ok := none.UnwrapOk()  // 0, false
val := none.UnwrapOr(0)     // 0
val := none.UnwrapOrElse(func() int { return 99 }) // 99
```

### Mapping

```go
doubled := some.Map(func(x int) int { return x * 2 })
//        → Some(84)

str := goption.MapOption(some, func(x int) string { return fmt.Sprint(x) })
// → Some("42")
```

### Pattern matching

```go
some.Match(
    func(x int) { fmt.Println("got", x) },
    func()      { fmt.Println("got nothing") },
)
```

## API

### Constructors

- `Some[T](value T) Option[T]` — wraps a value
- `None[T]() Option[T]` — represents absence

### Methods

- `IsSome() bool` / `IsNone() bool`
- `Unwrap() T` — panics on None
- `UnwrapOk() (T, bool)` — safe unwrap, comma-ok style
- `UnwrapOr(T) T` — fallback value
- `UnwrapOrElse(func() T) T` — lazy fallback
- `Map(func(T) T) Option[T]` — transform, same type

### Functions

- `MapOption[T, U](Option[T], func(T) U) Option[U]` — cross-type transform

## Contributing

Contributions are welcome! Please open an issue or submit a PR.

- Run `go test ./...` before submitting
- Keep the API surface minimal and consistent
- Follow standard Go formatting (`go fmt`)

## License

MIT
