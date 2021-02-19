package model

// 风向标志转换
func ConvertWind1(name string) string {
	switch name {
	case "无持续风向":
		return "0"
	case "东北风":
		return "1"
	case "东风":
		return "2"
	case "东南风":
		return "3"
	case "南风":
		return "4"
	case "西南风":
		return "5"
	case "西风":
		return "6"
	case "西北风":
		return "7"
	case "北风":
		return "8"
	default:
		return "9"
	}
}

// 风向标志转换
func ConvertWind(name string) string {

	if name == "无持续风向" {
		return "0"
	} else if name == "东北风" {
		return "1"
	} else if name == "东风" {
		return "2"
	} else if name == "东南风" {
		return "3"
	} else if name == "南风" {
		return "4"
	} else if name == "西南风" {
		return "5"
	} else if name == "西风" {
		return "6"
	} else if name == "西北风" {
		return "7"
	} else if name == "北风" {
		return "8"
	} else {
		return "9"
	}
}

// 天气状况代码转换
func ConvertWeather(name string) string {
	if name == "晴" {
		return "00"
	} else if name == "多云" {
		return "01"
	} else if name == "阴" {
		return "02"
	} else if name == "阵雨" {
		return "03"
	} else if name == "雷阵雨" {
		return "04"
	} else if name == "雷阵雨伴有冰雹" {
		return "05"
	} else if name == "雨夹雪" {
		return "06"
	} else if name == "小雨" {
		return "07"
	} else if name == "中雨" {
		return "08"
	} else if name == "大雨" {
		return "09"
	} else if name == "暴雨" {
		return "10"
	} else if name == "大暴雨" {
		return "11"
	} else if name == "特大暴雨" {
		return "12"
	} else if name == "阵雪" {
		return "13"
	} else if name == "小雪" {
		return "14"
	} else if name == "中雪" {
		return "15"
	} else if name == "大雪" {
		return "16"
	} else if name == "暴雪" {
		return "17"
	} else if name == "雾" {
		return "18"
	} else if name == "冻雨" {
		return "19"
	} else if name == "沙尘暴" {
		return "20"
	} else if name == "小到中雨" {
		return "21"
	} else if name == "中到大雨" {
		return "22"
	} else if name == "大到暴雨" {
		return "23"
	} else if name == "暴雨到大暴雨" {
		return "24"
	} else if name == "大暴雨到特大暴雨" {
		return "25"
	} else if name == "小到中雪" {
		return "26"
	} else if name == "中到大雪" {
		return "27"
	} else if name == "大到暴雪" {
		return "28"
	} else if name == "浮尘" {
		return "29"
	} else if name == "扬沙" {
		return "30"
	} else if name == "强沙尘暴" {
		return "31"
	} else if name == "霾" {
		return "53"
	} else {
		return "99"
	}
}
