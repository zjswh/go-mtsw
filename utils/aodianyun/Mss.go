package aodianyun

import "encoding/json"

type ClipResult struct {
	Flag       int     `json:"flag"`
	FlagString string  `json:"flagString"`
	Duration   float32 `json:"duration"`
	Thumbnail  string  `json:"thumbnail"`
	Location   string  `json:"location"`
}

func ClipDvrByEndTime(uin int, appId string, stream string, startTime int, endTime int, title string, stype string) (ClipResult, error) {
	diyParam, _ := json.Marshal(map[string]interface{}{
		"type":      stype,
		"programme": true,
	})
	param := map[string]interface{}{
		"uin":       uin,
		"app":       appId,
		"stream":    stream,
		"startTime": startTime,
		"endTime":   endTime,
		"title":     title,
		"diyParam": diyParam,
	}
	url := "DVR.ClipDvrByEndTime"

	res, err := OpenApi(url, uin, param)

	var clipResult ClipResult
	json.Unmarshal(res, &clipResult)
	return clipResult, err
}
