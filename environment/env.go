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

func IsNotDevelopment() bool {
	return env != Development
}

func IsDevelopment() bool {
	return env == Development
}

func IsNotTest() bool {
	return env != Test
}

func IsTest() bool {
	return env == Test
}

func IsNotProduction() bool {
	return env != Production
}

func IsProduction() bool {
	return env == Production
}
