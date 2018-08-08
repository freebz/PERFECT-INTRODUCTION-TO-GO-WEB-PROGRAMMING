// 타입 어설션으로 타입 변환

if w, ok := v.(io.Writer); ok {
	fmt.Fprintln(w, "v has Write() method")
}
