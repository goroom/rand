package rand

const (
	RST_LOWER  = 0x1
	RST_UPPER  = 0x2
	RST_NUMBER = 0x4
	RST_SYMBOL = 0x8
)

const (
	RST_LOWER_STR  = "qwertyuiopasdfghjklzxcvbnm"
	RST_UPPER_STR  = "QWERTYUIOPASDFGHJKLZXCVBNM"
	RST_NUMBER_STR = "1234567890"
	RST_SYMBOL_STR = "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
)

func String(length int, flag int) string {
	str := ""
	if flag == 0 {
		str = RST_LOWER_STR + RST_UPPER_STR + RST_NUMBER_STR + RST_SYMBOL_STR
	} else {
		if flag&RST_LOWER != 0 {
			str += RST_LOWER_STR
		}
		if flag&RST_UPPER != 0 {
			str += RST_UPPER_STR
		}
		if flag&RST_NUMBER != 0 {
			str += RST_NUMBER_STR
		}
		if flag&RST_SYMBOL != 0 {
			str += RST_SYMBOL_STR
		}
	}

	return StringLib(length, &str)
}

func StringLib(length int, str *string) string {
	_s := ""
	for i := 0; i < length; i++ {
		_s += RangeString(str)
	}
	return _s
}

func StringArray(list *[]string) string {
	return (*list)[Intn(len(*list))]
}

func RangeString(s *string) string {
	str := []rune(*s)
	index := g_rand.Intn(len(str))
	return string(str[index])
}
