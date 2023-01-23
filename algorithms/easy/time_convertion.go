package easy

import (
	"fmt"
	"strings"
)

func TimeConversion(time string) string {
	letter := fmt.Sprintf("%c", time[len(time)-2])
	hour := fmt.Sprintf("%c%c", time[0], time[1])
	newHour := time
	militaryTimes := map[string]string{
		"01": "13",
		"02": "14",
		"03": "15",
		"04": "16",
		"05": "17",
		"06": "18",
		"07": "19",
		"08": "20",
		"09": "21",
		"10": "22",
		"11": "23",
	}
	if letter == "P" && hour != "12" {
		mHour := militaryTimes[hour]
		newHour = strings.Replace(time, hour, mHour, 1)
	}
	if letter == "A" && hour == "12" {
		newHour = strings.Replace(time, "12", "00", 1)
	}
	return newHour[:len(time)-2]
}
