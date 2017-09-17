package rand

var (
	PROVINCE_CODE        []rune = []rune("京津沪渝冀豫云辽黑湘皖鲁新苏浙赣鄂桂甘晋蒙陕吉闽贵粤青藏川宁琼")
	PROVINCE_CODE_LENGTH int    = 0
	RCPTAIL_STR          string = "1234567890ABCDEFGHJKLMNPQRSTUVWXYZ" //车牌尾号
)

const (
	RCPTYPE_NORMAL     = 0x1
	RCPTYPE_NEW_ENERGY = 0x2
)

func init() {
	PROVINCE_CODE_LENGTH = len(PROVINCE_CODE)
}

func (r *Rand) CarPlate() string {
	return string(PROVINCE_CODE[gRand.rand.Intn(PROVINCE_CODE_LENGTH)]) + r.String(1, RST_UPPER) + r.StringLib(5, &RCPTAIL_STR)
}

//新能源车牌 6位
func (r *Rand) CarPlateNewEnergy() string {
	return string(PROVINCE_CODE[gRand.rand.Intn(PROVINCE_CODE_LENGTH)]) + r.String(1, RST_UPPER) + r.StringLib(6, &RCPTAIL_STR)
}
