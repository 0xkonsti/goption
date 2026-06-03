package goption

import "testing"

func TestSome(t *testing.T) {
	some := Some(420)
	if some.isSome != true {
		t.Errorf("Expected is_some to be true, got %v", some.isSome)
	}

	if some.value != 420 {
		t.Errorf("Expected value to be 420, got %v", some.value)
	}
}

func TestNone(t *testing.T) {
	none := None[int]()
	if none.isSome != false {
		t.Errorf("Expected is_some to be false, got %v", none.isSome)
	}
}

func TestIsSome(t *testing.T) {
	some := Some(420)
	if !some.IsSome() {
		t.Errorf("Expected IsSome to return true, got false")
	}
}

func TestIsNone(t *testing.T) {
	none := None[int]()
	if !none.IsNone() {
		t.Errorf("Expected IsNone to return true, got false")
	}
}

func TestUnwrap(t *testing.T) {
	some := Some(420)
	if some.Unwrap() != 420 {
		t.Errorf("Expected Unwrap to return 420, got %v", some.Unwrap())
	}
	none := None[int]()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Unwrap to panic on None, but it did not")
		}
	}()
	none.Unwrap()
}

func TestUnwrapOr(t *testing.T) {
	some := Some(420)
	if some.UnwrapOr(100) != 420 {
		t.Errorf("Expected UnwrapOr to return 420, got %v", some.UnwrapOr(100))
	}

	none := None[int]()
	if none.UnwrapOr(100) != 100 {
		t.Errorf("Expected UnwrapOr to return 100, got %v", none.UnwrapOr(100))
	}
}

func TestMap(t *testing.T) {
	some := Some(420)
	mapped := some.Map(func(x int) int { return x * 2 })
	if !mapped.IsSome() || mapped.Unwrap() != 840 {
		t.Errorf("Expected Map to return Some(840), got %v", mapped)
	}

	none := None[int]()
	mappedNone := none.Map(func(x int) int { return x * 2 })
	if !mappedNone.IsNone() {
		t.Errorf("Expected Map to return None, got %v", mappedNone)
	}
}

func TestUnwrapOrElse(t *testing.T) {
	some := Some(420)
	if some.UnwrapOrElse(func() int { return 100 }) != 420 {
		t.Errorf("Expected UnwrapOrElse to return 420, got %v", some.UnwrapOrElse(func() int { return 100 }))
	}

	none := None[int]()
	if none.UnwrapOrElse(func() int { return 100 }) != 100 {
		t.Errorf("Expected UnwrapOrElse to return 100, got %v", none.UnwrapOrElse(func() int { return 100 }))
	}
}

func TestUnwrapOk(t *testing.T) {
	some := Some(420)
	val, ok := some.UnwrapOk()
	if !ok || val != 420 {
		t.Errorf("Expected (420, true), got (%v, %v)", val, ok)
	}

	none := None[int]()
	val, ok = none.UnwrapOk()
	if ok {
		t.Errorf("Expected false, got true (val=%v)", val)
	}
}

func TestMapOpt(t *testing.T) {
	some := Some(420)
	mapped := MapOption(some, func(x int) string { return "val" })
	if !mapped.IsSome() || mapped.Unwrap() != "val" {
		t.Errorf("Expected MapOpt to return Some(\"val\"), got %v", mapped)
	}

	none := None[int]()
	mappedNone := MapOption(none, func(x int) string { return "val" })
	if !mappedNone.IsNone() {
		t.Errorf("Expected MapOpt to return None, got %v", mappedNone)
	}
}

func TestMatch(t *testing.T) {
	some := Some(420)
	some.Match(func(x int) {
		if x != 420 {
			t.Errorf("Expected Match to call someFunc with 420, got %v", x)
		}
	}, func() {
		t.Errorf("Expected Match to call someFunc, but it called noneFunc")
	})

	none := None[int]()
	none.Match(func(x int) {
		t.Errorf("Expected Match to call noneFunc, but it called someFunc with %v", x)
	}, func() {
		// Expected to call noneFunc, do nothing
	})
}
