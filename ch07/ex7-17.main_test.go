// main_test.go

func TestLenForChan(t *testing.T) {
	v := make(chan int)
	actual, expected := Len(v), 1
	if actual != expected {
		t.Errorf("%d != %d", actual, expected)
	}
}
