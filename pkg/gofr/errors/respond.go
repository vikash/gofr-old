package errors

import "fmt"

type Response struct {
	Code    int
	Message string
}

func (r Response) Error() string {
	return fmt.Sprint(r.Message)
}
