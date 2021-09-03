package timex

import "time"

var DefaultTimeFormat = "2006-01-02 15:04:05"

func TimeStampToDate(date int64, sType string) string {
	tm := time.Unix(date, 0)
	layout := "2006-01-02"
	switch sType {
	case "h":
		layout = "2006-01-02 15"
	case "m":
		layout = "2006-01-02 15:04"
	case "ymd":
		layout = "20060102"
	case "y-m-dTH:i:sZ":
		layout = "2006-01-02T15:04:05Z"
	default:
		layout = "2006-01-02 15:04:05"
	}
	return tm.Format(layout)
}

func DateToTimeStamp(date string, sType string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	layout := "2006-01-02"
	switch sType {
	case "h":
		layout = "2006-01-02 15"
	case "m":
		layout = "2006-01-02 15:04"
	case "y/m/d":
		layout = "2006/01/02 15:04:05"
	default:
		layout = "2006-01-02 15:04:05"
	}
	tt, _ := time.ParseInLocation(layout, date, loc)
	return tt.Unix()
}

func DateToTime(date string, sType string) time.Time {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	layout := "2006-01-02"
	switch sType {
	case "h":
		layout = "2006-01-02 15"
	case "m":
		layout = "2006-01-02 15:04"
	case "ymd":
		layout = "20060102"
	case "y/m/d":
		layout = "2006/01/02 15:04:05"
	default:
		layout = "2006-01-02 15:04:05"
	}
	tt, _ := time.ParseInLocation(layout, date, loc)
	return tt
}

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func DateStringToTimeStamp(date string) int64 {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	tt, _ := time.ParseInLocation("2006-01-02T15:04:05+08:00", date, loc)
	return tt.Unix()
}
