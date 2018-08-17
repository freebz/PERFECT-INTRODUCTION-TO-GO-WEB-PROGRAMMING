// switch 문으로 확인

switch e := err.(type) {
case SqrtError:
	fmt.Println("Sqrt Error", e)
default:
	fmt.Println("Default Error", e)
}
