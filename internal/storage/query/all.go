package query

type Limit struct {
	Size int
	Page int
}

func (l Limit) Validate() {
	if l.Size < 1 {
		l.Size = 1
	}
	if l.Page < 0 {
		l.Page = 0
	}
}
