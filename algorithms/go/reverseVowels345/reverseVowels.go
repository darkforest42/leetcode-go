package reverseVowels345

func ReverseVowels(s string) string {
	vowels := map[byte]bool{
		'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
		'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	}
	strBytes := []byte(s)
	i , j := 0, len(s)-1
	for  {
		for i < len(s) && !vowels[s[i]] {
			i++
		}
		for j >= 0 && !vowels[s[j]] {
			j--
		}
		if i > j {
			break
		}
		strBytes[i], strBytes[j] = strBytes[j], strBytes[i]
		i++
		j--
	}
	return string(strBytes)
}
