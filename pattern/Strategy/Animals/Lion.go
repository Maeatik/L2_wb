package Animals

type Lion struct {
	Name string
}

func (l *Lion) GetName() string {
	return l.Name
}
