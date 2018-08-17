func BenchmarkLenforString(b *testing.B) {
	b.StopTimer()
	v := make([]string, 1000000)
	for i := 0; i < 1000000; i++ {
		v[i] = strconv.Itoa(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Len(v[i%1000000])
	}
}
