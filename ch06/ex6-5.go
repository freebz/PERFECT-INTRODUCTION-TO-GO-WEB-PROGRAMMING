// fmt.Errorf() 함수

func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
