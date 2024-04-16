package borrows

type Borrowed struct {
	Id       string
	BookId   string
	UserId   string
	Returned bool
}
