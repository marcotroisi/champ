package champ

type Person interface {
	FirstName() string
	LastName() string
	Email() string
}

type People interface {
	Add(firstname string, lastname string, email string)
	List() []Person
}
