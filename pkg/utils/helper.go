package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	Timezone = "Asia/Bangkok"
)

func Debug(obj any) {
	raw, _ := json.MarshalIndent(obj, "", "\t")
	fmt.Println(raw)
}

func LocalTime() time.Time {
	loc, _ := time.LoadLocation(Timezone)
	return time.Now().In(loc)
}

func ConvertStringTimeToTime(t string) (time.Time, error) {
	layout := "2006-01-02 15:04:05.999 -0700 MST"
	result, err := time.Parse(layout, t)
	if err != nil {
		return time.Time{}, err
	}

	return result, nil
}
