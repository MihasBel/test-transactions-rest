package mocks

// Reporter mock
type Reporter struct {
}

// Errorf mock
func (r Reporter) Errorf(_ string, _ ...interface{}) {

}

// Fatalf mock
func (r Reporter) Fatalf(_ string, _ ...interface{}) {

}
