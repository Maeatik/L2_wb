package Animals

type Duck struct {
	Name string
}

func (d *Duck) GetName() string {
	return d.Name
}
