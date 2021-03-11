package exists

type Entity struct{}
type DuplicateEntity struct{}

func (e Entity) Error() string {
	return ""
}

func (d DuplicateEntity) Error() string {
	return ""
}
