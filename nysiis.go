package jellyfish

import "strings"

func Nysiis(s string) string {
	var key []rune
	runes := []rune(strings.ToUpper(s))
	rlen := len(runes)

	// step 1 - prefixes
	switch {
	case runes[0] == 'M' && runes[1] == 'A' && runes[2] == 'C':
		runes[1] = 'C'
	case runes[0] == 'K' && runes[1] == 'N':
		runes = runes[1:]
		rlen--
	case runes[0] == 'K':
		runes[0] = 'C'
	case runes[0] == 'P' && (runes[1] == 'H' || runes[1] == 'F'):
		runes[0] = 'F'
		runes[1] = 'F'
	case runes[0] == 'S' && runes[1] == 'C' && runes[2] == 'H':
		runes[1] = 'S'
		runes[2] = 'S'
	}

	// step 2 - suffixes
	switch {
	case (runes[rlen-2] == 'I' || runes[rlen-2] == 'E') && runes[rlen-1] == 'E':
		runes = append(runes[:rlen-2], 'Y')
		rlen--
	case runes[rlen-2] == 'D' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'R' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'R' && runes[rlen-1] == 'D',
		runes[rlen-2] == 'N' && runes[rlen-1] == 'T',
		runes[rlen-2] == 'N' && runes[rlen-1] == 'D':
		runes = append(runes[:rlen-2], 'D')
		rlen--
	}

	// step 3 - first character from name
	key = append(key, runes[0])

	// step 4 - translate remaining
	var keypiece []rune
	for i := 1; i < rlen; i++ {
		ch := runes[i]
		switch {
		case ch == 'E' && i+1 < rlen && runes[i+1] == 'V':
			keypiece = []rune{'A', 'F'}
			i++
		case isVowel(ch):
			keypiece = []rune{'A'}
		case ch == 'Q':
			keypiece = []rune{'G'}
		case ch == 'Z':
			keypiece = []rune{'S'}
		case ch == 'M':
			keypiece = []rune{'N'}
		case ch == 'K':
			if i+1 < rlen && runes[i+1] == 'N' {
				keypiece = []rune{'N'}
			} else {
				keypiece = []rune{'C'}
			}
		case ch == 'S' && i+2 < rlen && runes[i+1] == 'C' && runes[i+2] == 'H':
			keypiece = []rune{'S', 'S'}
			i += 2
		case ch == 'P' && i+1 < rlen && runes[i+1] == 'H':
			keypiece = []rune{'F'}
			i++
		case ch == 'H' && (!isVowel(runes[i-1]) || (i+1 < rlen && !isVowel(runes[i+1]))):
			if isVowel(runes[i-1]) {
				keypiece = []rune{'A'}
			} else {
				keypiece = []rune{runes[i-1]}
			}
		case ch == 'W' && isVowel(runes[i-1]):
			keypiece = []rune{runes[i-1]}
		default:
			keypiece = []rune{ch}
		}

		// step 8 done here - avoid double chars
		if keypiece[len(keypiece)-1] != key[len(key)-1] {
			key = append(key, keypiece...)
		}
	}

	// step 5 - remove trailing S
	if key[len(key)-1] == 'S' {
		key = key[:len(key)-1]
	}

	// step 6 - AY=>Y
	if key[len(key)-2] == 'A' && key[len(key)-1] == 'Y' {
		key = key[:len(key)-1]
		key[len(key)-1] = 'Y'
	}

	// step 7 - remove trailing A
	if key[len(key)-1] == 'A' {
		key = key[:len(key)-1]
	}

	return string(key)
}