package utils

import "strings"

var base62List = []string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
	"l", "m", "n", "o", "p", "q", "r", "s", "t", "u",
	"v", "w", "x", "y", "z", "A",
	"B", "C", "E", "F", "D", "G", "H", "I", "J", "K", "L", "M",
	"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
}

const bas62Len int64 = 62

func Int2Base62(num int64) string {
	sb := strings.Builder{}
	for num > 0 {
		index := num % bas62Len
		sb.WriteString(base62List[index])
		num /= bas62Len
	}

	return sb.String()
}
