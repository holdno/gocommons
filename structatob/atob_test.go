package structatob

import (
	"testing"
)

type TestA struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Tel      int64  `json:"tel"`
	OpenID   string `json:"openid"`
	UnionID  string `json:"unionid"`
	Password string `json:"password"`
}

type TestB struct {
	Name     string `json:"name"`
	Account  string `json:"account"`
	Tel      int64  `json:"tel"`
	OpenID   string `json:"openid"`
	UnionID  string `json:"unionid"`
	Password string `json:"password"`
}

func TestAtoB(t *testing.T) {
	a := &TestA{
		Name:    "123",
		Account: "asd",
	}

	b := &TestB{}

	err := AtoB(a, b)
	if err != nil {
		t.Error(err)
	}

	t.Log(AtoB(a, b))

	t.Logf("%+v", b)
}

func BenchmarkAtoB(tb *testing.B) {
	a := &TestA{
		Name:     "123",
		Account:  "asd",
		Password: "cccccccc",
		Tel:      18888888888,
		OpenID:   "djklsjfklsjdklfs",
		UnionID:  "129nmcsdf912n321kjj213b123",
	}

	b := &TestB{}

	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		AtoB(a, b)
	}

}

func BenchmarkAtoB2(tb *testing.B) {
	a := &TestA{
		Name:     "123",
		Account:  "asd",
		Password: "cccccccc",
		Tel:      18888888888,
		OpenID:   "djklsjfklsjdklfs",
		UnionID:  "129nmcsdf912n321kjj213b123",
	}

	b := &TestB{}

	tb.ResetTimer()
	for i := 0; i < tb.N; i++ {
		b.Name = a.Name
		b.Account = a.Account
		b.Password = a.Password
		b.Tel = a.Tel
		b.OpenID = a.OpenID
		b.UnionID = a.UnionID
	}
}
