package rand

var (
	g_province_code_list []string = []string{"京",
		"津",
		"沪",
		"渝",
		"冀",
		"豫",
		"云",
		"辽",
		"黑",
		"湘",
		"皖",
		"鲁",
		"新",
		"苏",
		"浙",
		"赣",
		"鄂",
		"桂",
		"甘",
		"晋",
		"蒙",
		"陕",
		"吉",
		"闽",
		"贵",
		"粤",
		"青",
		"藏",
		"川",
		"宁",
		"琼"}
)

func CarPlate() string {
	return StringArray(&g_province_code_list) + String(1, RST_UPPER) + String(5, RST_UPPER|RST_NUMBER)
}
