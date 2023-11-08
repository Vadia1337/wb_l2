package otherTests

func DFunc() int {
	return 1 + 2 + 3
}

type SomeStruct struct{}

func (ss *SomeStruct) StructFunc() int {
	return 1 + 2 + 3
}

func AsyncFunc() {
	_ = 1 + 2 + 3
}
