package errors

import "fmt"

type InvalidParam struct {
	Param string
}

func (i InvalidParam) Error() string {
	return fmt.Sprintf("%v is invalid", i.Param)
}
