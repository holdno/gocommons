package selection

import (
	"bytes"
	"fmt"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json jsoniter.API

func init() {
	json = jsoniter.ConfigCompatibleWithStandardLibrary
}

type Operator string

const (
	DoesNotExist    Operator = "!"
	Equals          Operator = "="
	EqualString     Operator = "'='"
	DoubleEquals    Operator = "=="
	In              Operator = "IN"
	NotEquals       Operator = "!="
	NotEqualsString Operator = "'!='"
	NotIn           Operator = "NOT IN"
	GreaterThan     Operator = ">"
	LessThan        Operator = "<"
	Like            Operator = "LIKE"
	FindIn          Operator = "FIND_IN_SET"
)

// Requirement contains values, a key, and an operator that relates the key and values.
type Requirement struct {
	Key      string   `json:"k"`
	Operator Operator `json:"o"`
	// In huge majority of cases we have at most one value here.
	// It is generally faster to operate on a single-element slice
	// than on a single-element map, so we have a slice here.
	Values interface{} `json:"v"`
}

func NewRequirement(key string, operator Operator, values interface{}) Requirement {
	return Requirement{key, operator, values}
}

type Requirements []Requirement

func (r Requirements) Append(requirements ...Requirement) Requirements {
	var res []Requirement
	for index := range r {
		res = append(res, r[index])
	}

	for index := range requirements {
		res = append(res, requirements[index])
	}

	return res
}

func (r Requirements) String() string {
	var res []string
	for _, v := range r {
		con, value := v.Patch()
		res = append(res, strings.Replace(con, "?", fmt.Sprintf("%v", value), -1))
	}
	return strings.Join(res, ";")
}

// String returns a human-readable string that represents this
// Requirement. If called on an invalid Requirement, an error is
// returned. See NewRequirement for creating a valid Requirement.
func (r *Requirement) Patch() (string, interface{}) {
	var buffer bytes.Buffer
	if r.Operator == DoesNotExist {
		buffer.WriteString("!")
	}

	if r.Operator != FindIn {
		buffer.WriteString(r.Key)
	}

	switch r.Operator {
	case Equals, EqualString:
		buffer.WriteString("=")
	case DoubleEquals:
		buffer.WriteString("==")
	case NotEquals, NotEqualsString:
		buffer.WriteString("!=")
	case In:
		buffer.WriteString(" IN (")
	case NotIn:
		buffer.WriteString(" NOT IN (")
	case GreaterThan:
		buffer.WriteString(">")
	case LessThan:
		buffer.WriteString("<")
	case Like:
		buffer.WriteString(" LIKE ")
		//case Exists, DoesNotExist:
		//	return buffer.String()
	case FindIn:
		buffer.WriteString(" FIND_IN_SET(")
	}

	if r.Operator != FindIn {
		buffer.WriteString("?")
	} else {
		buffer.WriteString(r.Key + ",?")
	}

	switch r.Operator {
	case In, NotIn, FindIn:
		buffer.WriteString(")")
	case Like:
		return buffer.String(), fmt.Sprintf("%%%v%%", r.Values)
	}

	return buffer.String(), r.Values
}

type Selector struct {
	queryIndex int
	query      Requirements `json:"q,omitempty"`
	Page       int          `json:"p,omitempty"`
	Pagesize   int          `json:"ps,omitempty"`
	OrderBy    string       `json:"ob,omitempty"`
	Select     string       `json:"s,omitempty"`
}

func (s *Selector) ToBytes() []byte {
	res, _ := json.Marshal(s)
	return res
}

func ParseFromBytes(b []byte) (Selector, error) {
	s := NewSelector()
	if err := json.Unmarshal(b, &s); err != nil {
		return Selector{}, err
	}
	return s, nil
}

func NewSelector(r ...Requirement) Selector {
	s := Selector{}
	if len(r) == 0 {
		return s
	}
	s.AddQuery(r...)
	return s
}

func (s *Selector) Len() int {
	return len(s.query)
}

func (s *Selector) Patch() (string, interface{}) {
	i := s.queryIndex - 1
	return s.query[i].Patch()
}

func (s *Selector) GetCurrentQuery() Requirement {
	i := s.queryIndex - 1
	return s.query[i]
}

func (s *Selector) NextQuery() bool {
	s.queryIndex++
	if len(s.query) < s.queryIndex {
		return false
	}
	return true
}

func (s *Selector) AddQuery(r ...Requirement) {
	s.query = append(s.query, r...)
}

func (s *Selector) AddOrder(str string) {
	if s.OrderBy != "" {
		s.OrderBy += "," + str
	} else {
		s.OrderBy = str
	}
}

func (s Selector) AddSelect(str string) Selector {
	if s.OrderBy != "" {
		s.OrderBy += "," + str
	} else {
		s.OrderBy = str
	}
	return s
}
