// json 패키지의 SyntaxError 타입

type SyntaxError struct {
	msg    string  // 에러 설명
	Offset int64   // 에러가 발생한 지점의 오프셋(byte 단위)
}

func (e *SyntaxError) Error() string { return e.msg }
