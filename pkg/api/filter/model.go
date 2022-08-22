package filter

import "fmt"

const (
	DataTypeString = "string"
	DataTypeInt    = "int"
	DataTypeDate   = "date"
	DataTypeBool   = "bool"

	OperatorEq            = "eq"
	OperatorNotEq         = "neq"
	OperatorLowerThen     = "lt"
	OperatorLowerThenEq   = "lte"
	OperatorGreaterThen   = "gt"
	OperatorGreaterThenEq = "gte"
	OperatorBetween       = "between"
	OperatorLike          = "like"
)

type options struct {
	limit  int
	fields []Field
}

func NewOptions(limit int) Options {
	return &options{limit: limit}
}

type Field struct {
	Name     string
	Value    string
	Operator string
	Type     string
}

// Options TODO move it
type Options interface {
	Limit() int
	AddField(name, operator, value, dtype string) error
	Fields() []Field
}

func (o *options) Limit() int {
	return o.limit
}

func (o *options) AddField(name, operator, value, dtype string) error {
	err := ValidateOperator(operator)
	if err != nil {
		return err
	}

	o.fields = append(o.fields, Field{
		Name:     name,
		Value:    value,
		Operator: operator,
		Type:     dtype,
	})

	return nil
}

func (o *options) Fields() []Field {
	return o.fields
}

func ValidateOperator(operator string) error {
	switch operator {
	case OperatorEq:
	case OperatorNotEq:
	case OperatorLowerThen:
	case OperatorLowerThenEq:
	case OperatorGreaterThen:
	case OperatorGreaterThenEq:
	case OperatorBetween:
	case OperatorLike:
	default:
		return fmt.Errorf("bad operator")
	}

	return nil
}
