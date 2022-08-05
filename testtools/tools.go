package testtools

type MockCall struct {
	Called bool
	Input  []interface{}
	Output []interface{}
}
