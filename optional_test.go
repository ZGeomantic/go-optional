package optional

import (
	"testing"
)

type Foo struct {
	A string
}

func TestDemo(t *testing.T) {
	var (
		nilInt *int

		nonnilInt *int
		expectInt int = 100

		defaultInt int = 1234
	)

	nonnilInt = &expectInt

	if Deref(nilInt) != 0 {
		t.Fatalf("nilInt expect 0\n")
	}

	if Deref(nonnilInt) != expectInt {
		t.Fatalf("nonnilInt expect %d, actucal: %d\n", expectInt, *nonnilInt)
	}

	if DerefMust(nilInt, defaultInt) != defaultInt {
		t.Fatalf("nonnilInt expect %d, actucal: %d\n", defaultInt, *nilInt)
	}

	if DerefMust(nonnilInt, defaultInt) != expectInt {
		t.Fatalf("nonnilInt expect %d, actucal: %d\n", expectInt, *nonnilInt)
	}

}

func TestZero(t *testing.T) {
	// var foo *Foo = new(Foo)
	var a *string
	t.Logf("Is zero: %v\n", IsZero(a))

}

func BenchmarkGetIntNil(b *testing.B) {
	var nilInt *int

	for i := 0; i < b.N; i++ {
		_ = Deref(nilInt).(int)
	}
}
func BenchmarkGetInt(b *testing.B) {
	var nilInt *int = new(int)

	for i := 0; i < b.N; i++ {
		_ = Deref(nilInt).(int)
	}
}

func BenchmarkGetIntNativeNil(b *testing.B) {
	var nilInt *int

	for i := 0; i < b.N; i++ {
		if nilInt == nil {
			_ = 0
		} else {
			_ = *nilInt
		}
	}
}

func BenchmarkGetIntNative(b *testing.B) {
	var nilInt *int = new(int)

	for i := 0; i < b.N; i++ {
		if nilInt == nil {
			_ = 0
		} else {
			_ = *nilInt
		}
	}
}

func BenchmarkGetStrNil(b *testing.B) {
	var nilStr *string

	for i := 0; i < b.N; i++ {
		_ = Deref(nilStr).(string)
	}
}
func BenchmarkGetStr(b *testing.B) {
	var nilStr *string = new(string)

	for i := 0; i < b.N; i++ {
		_ = Deref(nilStr).(string)
	}
}

func BenchmarkGetStrNativeNil(b *testing.B) {
	var nilStr *string

	for i := 0; i < b.N; i++ {
		if nilStr == nil {
			_ = 0
		} else {
			_ = *nilStr
		}
	}
}

func BenchmarkGetStrNative(b *testing.B) {
	var nilStr *string = new(string)

	for i := 0; i < b.N; i++ {
		if nilStr == nil {
			_ = 0
		} else {
			_ = *nilStr
		}
	}
}

func BenchmarkGetFooNil(b *testing.B) {
	var nilFoo *Foo
	for i := 0; i < b.N; i++ {
		_ = Deref(nilFoo).(Foo)
	}
}

func BenchmarkGetFoo(b *testing.B) {
	var nilFoo *Foo = new(Foo)

	for i := 0; i < b.N; i++ {
		_ = Deref(nilFoo).(Foo)
	}
}

func BenchmarkGetStrFooNil(b *testing.B) {
	var nilFoo *Foo

	for i := 0; i < b.N; i++ {
		if nilFoo == nil {
			_ = 0
		} else {
			_ = *nilFoo
		}
	}
}
func BenchmarkGetStrFoo(b *testing.B) {
	var nilFoo *Foo = new(Foo)

	for i := 0; i < b.N; i++ {
		if nilFoo == nil {
			_ = 0
		} else {
			_ = *nilFoo
		}
	}
}
