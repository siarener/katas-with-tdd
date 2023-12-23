package kata

const lenAlphabet = 26
const lenRotation = 13

func Rot13(msg string) string {
	var result []rune
	for _, oneChar := range msg {
		switch {
		case 'A' <= oneChar && oneChar <= 'Z':
			result = append(result, 'A'+(oneChar-'A'+lenRotation)%lenAlphabet)
		case 'a' <= oneChar && oneChar <= 'z':
			result = append(result, 'a'+(oneChar-'a'+lenRotation)%lenAlphabet)
		default:
			result = append(result, oneChar)
		}

	}
	return string(result)
}
