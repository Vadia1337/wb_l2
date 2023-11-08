package otherTests

import "testing"

func BenchmarkDFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = DFunc()
	}
}

func BenchmarkSomeStruct_StructFunc(b *testing.B) {
	ss := SomeStruct{}
	for i := 0; i < b.N; i++ {
		_ = ss.StructFunc()
	}
}

func BenchmarkSomeStruct_StructFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ss := SomeStruct{}
		_ = ss.StructFunc()
	}
}

func BenchmarkAnonFunc(b *testing.B) {
	f := func() int { return 1 + 2 + 3 }
	for i := 0; i < b.N; i++ {
		_ = f()
	}
}

func BenchmarkAnonFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := func() int { return 1 + 2 + 3 }
		_ = f()
	}
}

func BenchmarkAsyncFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go AsyncFunc()
	}
}

func BenchmarkAsyncAnonFunc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go func() {
			_ = 1 + 2 + 3
		}()
	}
}
