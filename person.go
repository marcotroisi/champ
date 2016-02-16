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

type PersonConfig struct {
	Id                         int
	FirstName, LastName, Email string
}

func NewPerson(config PersonConfig) *Person {
	return &Person{PersonConfig: config}
}

type Person struct {
	PersonConfig
}

func (p *Person) Id() int {
	return p.PersonConfig.Id
}

func (p *Person) FirstName() string {
	return p.PersonConfig.FirstName
}

func (p *Person) LastName() string {
	return p.PersonConfig.LastName
}

func (p *Person) Email() string {
	return p.PersonConfig.Email
}
