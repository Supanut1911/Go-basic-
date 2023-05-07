package person

type Person struct {
	id int
	name string
}

func (p Person) GetName() string {
	return p.name
}

func (p Person) GetId() int { 
	return p.id
}

func (p *Person) SetName(name string) {
	p.name = name
}