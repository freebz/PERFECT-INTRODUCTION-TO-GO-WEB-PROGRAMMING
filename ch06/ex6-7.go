// 타입 어설션으로 확인

if e, ok := err.(SqrtError); ok {
	fmt.Println("Sqrt Error", e)
}
