package champ

type PersonInterface interface {
	Id() int
	FirstName() string
	LastName() string
	Email() string
}

type PeopleInterface interface {
	Add(firstname string, lastname string, email string)
	List() []Person
}

type Person struct {
	id                         int
	firstName, lastName, email string
}

func (p *Person) Id() int {
	return p.id
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) Email() string {
	return p.email
}
