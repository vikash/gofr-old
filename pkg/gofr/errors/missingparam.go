package errors

import "fmt"

type MissingParam struct {
	Param string
}

func (m MissingParam) Error() string {
	return fmt.Sprintf("%v is missing", m.Param)
}
