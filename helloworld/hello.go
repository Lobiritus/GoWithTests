package helloworld

const (
	englisHelloPrefix     = "Hello, "
	spanishHelloPrefix    = "Hola, "
	frenchHelloPrefix     = "Bonjour, "
	azerbaijanHelloPrefix = "Salam, "
	spanish               = "Spanish"
	french                = "French"
	azerbaijan            = "Azerbaijan"
)

func Hello(text, language string) string {
	if text == "" {
		text = "World"
	}
	return greetingPrefix(language) + text
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case azerbaijan:
		prefix = azerbaijanHelloPrefix
	default:
		prefix = englisHelloPrefix
	}
	return
}
