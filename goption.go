// Package goption provides a generic Option type for Go, similar to Rust's Option.
// It represents an optional value: either Some(value) or None.
package goption

// Option represents an optional value of type T.
// Use Some to create a value, None to represent absence.
type Option[T any] struct {
	value  T
	isSome bool
}

// Some creates an Option containing the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{value: value, isSome: true}
}

// None returns an Option representing the absence of a value.
func None[T any]() Option[T] {
	return Option[T]{isSome: false}
}

// IsSome returns true if the Option contains a value.
func (opt Option[T]) IsSome() bool {
	return opt.isSome
}

// IsNone returns true if the Option is None.
func (opt Option[T]) IsNone() bool {
	return !opt.isSome
}

// Unwrap returns the contained value. Panics if the Option is None.
func (opt Option[T]) Unwrap() T {
	if !opt.isSome {
		panic("called Unwrap on a None value")
	}
	return opt.value
}

// UnwrapOk returns the contained value and true, or zero value and false if None.
func (opt Option[T]) UnwrapOk() (T, bool) {
	if opt.isSome {
		return opt.value, true
	}
	var zero T
	return zero, false
}

// UnwrapOr returns the contained value, or defaultValue if None.
func (opt Option[T]) UnwrapOr(defaultValue T) T {
	if opt.isSome {
		return opt.value
	}
	return defaultValue
}

// Map applies f to the contained value and returns the result wrapped in Some.
// Returns None if the Option is None.
func (opt Option[T]) Map(f func(T) T) Option[T] {
	if opt.isSome {
		return Some(f(opt.value))
	}
	return None[T]()
}

// MapOption applies f to the contained value, returning an Option of a different type.
// Returns None if the Option is None.
func MapOption[T, U any](opt Option[T], f func(T) U) Option[U] {
	if opt.isSome {
		return Some(f(opt.value))
	}
	return None[U]()
}

// UnwrapOrElse returns the contained value, or calls f to produce a fallback.
func (opt Option[T]) UnwrapOrElse(f func() T) T {
	if opt.isSome {
		return opt.value
	}
	return f()
}

// Match calls someFunc with the contained value, or noneFunc if None.
func (opt Option[T]) Match(someFunc func(T), noneFunc func()) {
	if opt.isSome {
		someFunc(opt.value)
	} else {
		noneFunc()
	}
}
