package selection

import (
	"testing"
)

func TestNewSelector(t *testing.T) {
	a := NewSelector(NewRequirement("name", Like, 1))
	for _, v := range a.Query {
		query, value := v.Patch()
		t.Log(query, value)
	}
}

func TestSelector_ToBytes(t *testing.T) {
	a := NewSelector(NewRequirement("name", Like, 1),
		NewRequirement("test", Equals, "content"))

	b := a.ToBytes()
	t.Log(string(b))

	res, err := ParseFromBytes(b)
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range res.Query {
		query, value := v.Patch()
		t.Log(query, value)
	}
}
