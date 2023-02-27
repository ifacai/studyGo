func msg(content, color string) string {
	var startColorString string
	switch color {
	case "greenBg", "gbg":
		startColorString = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
		break
	case "whiteBg", "wbg":
		startColorString = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
		break
	case "yellowBg", "ybg":
		startColorString = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
		break
	case "redBg", "rbg":
		startColorString = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
		break
	case "blueBg", "bbg":
		startColorString = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
		break
	case "purpleBg", "pbg":
		startColorString = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
		break
	case "cyanBg", "cbg":
		startColorString = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
		break
	case "green", "g":
		startColorString = string([]byte{27, 91, 51, 50, 109})
		break
	case "white", "w":
		startColorString = string([]byte{27, 91, 51, 55, 109})
		break
	case "yellow", "y":
		startColorString = string([]byte{27, 91, 51, 51, 109})
		break
	case "red", "r":
		startColorString = string([]byte{27, 91, 51, 49, 109})
		break
	case "blue", "b":
		startColorString = string([]byte{27, 91, 51, 52, 109})
		break
	case "purple", "p":
		startColorString = string([]byte{27, 91, 51, 53, 109})
		break
	case "cyan", "c":
		startColorString = string([]byte{27, 91, 51, 54, 109})
		break
	default:
		startColorString = ""
	}
	return startColorString + content + string([]byte{27, 91, 48, 109})
}
