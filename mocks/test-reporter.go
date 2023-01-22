package mocks

type Reporter struct {
}

func (r Reporter) Errorf(_ string, _ ...interface{}) {

}
func (r Reporter) Fatalf(_ string, _ ...interface{}) {

}
