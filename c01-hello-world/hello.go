package c01

const greetingDefaultPrefix = "Hello, "
const greetingDefaultTarget = "World"

var greetingPrefixForLanguage = map[string]string{
	"English": "Hello, ",
	"Spanish": "Hola, ",
	"French":  "Bonjour, ",
}

func Hello(name string, language string) string {
	name = setName(name)
	greetingPrefix := setGreetingPrefix(language)
	return greetingPrefix + name
}

func setGreetingPrefix(language string) string {
	greetingPrefix := greetingPrefixForLanguage[language]
	if greetingPrefix == "" {
		greetingPrefix = greetingDefaultPrefix
	}
	return greetingPrefix
}

func setName(name string) string {
	if name == "" {
		name = greetingDefaultTarget
	}
	return name
}
