// 문자열 내부에 접근 2 - 문자열을 []rune 타입으로 변환한 후 인덱스 사용

s1 : = "hello"
s2 := "안녕하세요"
r1 := []rune(s1)
r2 := []rune(s2)
fmt.Printf("s1: %c %c %c %c %c\n", r1[0], r1[1], r1[2], r1[3], r1[4])
fmt.Printf("s2: %c %c %c %c %c", r2[0], r2[1], r2[2], r2[3], r2[4])
