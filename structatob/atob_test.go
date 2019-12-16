package structatob

import (
	"testing"
)

type TestA struct {
	Name    string `json:"name"`
	Account string `json:"account"`
}

type TestB struct {
	Name    string `json:"name,omitempty"`
	Account int    `json:"account"`
}

func TestAtoB(t *testing.T) {
	a := &TestA{
		Name:    "123",
		Account: "asd",
	}

	b := &TestB{}

	t.Log(AtoB(a, b))

	t.Logf("%+v", b)
}
