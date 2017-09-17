package rand

const (
	RST_LOWER  = 0x1
	RST_UPPER  = 0x2
	RST_NUMBER = 0x4
	RST_SYMBOL = 0x8
)

const (
	gcLowerString  = "qwertyuiopasdfghjklzxcvbnm"
	gcUpperString  = "QWERTYUIOPASDFGHJKLZXCVBNM"
	gcNumberString = "1234567890"
	gcSymbolString = "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
)

func (r *Rand) String(length int, flag int) string {
	str := ""
	if flag == 0 {
		str = gcLowerString + gcUpperString + gcNumberString + gcSymbolString
	} else {
		if flag&RST_LOWER != 0 {
			str += gcLowerString
		}
		if flag&RST_UPPER != 0 {
			str += gcUpperString
		}
		if flag&RST_NUMBER != 0 {
			str += gcNumberString
		}
		if flag&RST_SYMBOL != 0 {
			str += gcSymbolString
		}
	}

	return r.StringLib(length, &str)
}

func (r *Rand) StringLib(length int, str *string) string {
	_s := ""
	for i := 0; i < length; i++ {
		_s += r.RangeString(str)
	}
	return _s
}

func (r *Rand) StringArray(list *[]string) string {
	return (*list)[r.rand.Intn(len(*list))]
}

func (r *Rand) RangeString(s *string) string {
	gRand.Lock()
	defer gRand.Unlock()
	str := []rune(*s)
	index := gRand.rand.Intn(len(str))
	return string(str[index])
}
