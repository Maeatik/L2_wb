package Animals

type Rhino struct {
	Name string
}

func (r *Rhino) GetName() string {
	return r.Name
}
