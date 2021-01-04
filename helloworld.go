package main

const (
	englishGreeting = "Hello"
	arabicGreeting = "Asalaam alaikum"
	frenchGreeting = "Bonjour"
)

func Greeting(name, lang string) string {

	if name == "" {
		return greetPrefix(lang) + " World"
	}

	return greetPrefix(lang) + " " + name
}

func greetPrefix(lang string) (prefix string) {
	switch lang {
	case "Arabic":
		prefix = arabicGreeting
	case "French": 
		prefix = frenchGreeting
	default:
		prefix = englishGreeting
	}

	return
}
func main() {
	Greeting("Chris", "")
}