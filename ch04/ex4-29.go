// 기본 라이브러리의 인터페이스 활용 - io.Writer

type Writer interface {
	Write(p []byte) (n int, err os.Error)
}
	
func Fprintln(w io.Writer, a ...interface{}) (n int, err error)


func handle(w io.Writer, msg string) {
	fmt.Frpintln(w, msg)
}


func main() {
	msg := []string{"This", "is", "an", "example", "of", "io.Writer"}

	for _, s := range msg {
		time.Sleep(100 * time.Millisecond)
		handle(os.Stdout, s)
	}
}



func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handle(w, r.URL.Path[1:])
	})

	fmt.Println("start listening on port 4000")
	http.ListenAndServe(":4000", nil)
}
