package environment

type Environment = int

const (
	Development int = iota
	Test
	Production
)

var env = Development

func Set(e Environment) {
	env = e
}

func Get() Environment {
	return env
}

func IsDevelopment() bool {
	return env == Development
}

func IsTest() bool {
	return env == Test
}

func IsProduction() bool {
	return env == Production
}
