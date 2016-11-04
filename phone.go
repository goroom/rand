package rand

var (
	g_phone_prefix_list []string = []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
)

func RandPhone() string {
	return StringArray(&g_phone_prefix_list) + String(8, RST_NUMBER)
}
