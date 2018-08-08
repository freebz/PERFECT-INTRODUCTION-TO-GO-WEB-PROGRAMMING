// 기본 타입을 사용자 정의 타입으로 정의

type quantity int
type dozen []quantity

var d dozen
for i := quaintity(1); i <= 12; i++ {
	d = append(d, i)
}
fmt.Println(d)
