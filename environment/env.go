package environment

import "os"

type Environment = int

const (
	Development int = iota
	Test
	Production
)

var env = Development

func Init() {
	systemEnv, ok := os.LookupEnv("ICEBERG_ENV")
	if ok {
		switch systemEnv {
		case "production":
			Set(Production)
		case "development":
			Set(Development)
		case "test":
			Set(Test)
		}
	}
}

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

func Name() string {
	switch env {
	case Production:
		return "production"
	case Test:
		return "test"
	case Development:
		return "development"

	}
	return ""
}
