package errors

import "fmt"

type Respond struct {
	Code    int
	Message string
}

func (r Respond) Error() string {
	return fmt.Sprint(r.Message)
}
