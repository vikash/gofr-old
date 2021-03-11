package errors

type NewEntity struct{}

func (e NewEntity) Error() string {
	return ""
}
