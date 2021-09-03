package elasticSearch

import (
	"encoding/json"
	"mtsw/global"
	"mtsw/utils"
)

const (
	INDEX = "aggregate"
	TYPE = "content"
)

type EsResult struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`
	Hits struct {
		Total    int     `json:"total"`
		MaxScore float64 `json:"max_score"`
		Hits     []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Score  float64 `json:"_score"`
			Source struct {
				Createtime    int    `json:"createTime"`
				ID            int    `json:"id"`
				Livenowstatus int    `json:"liveNowStatus"`
				Logo          string `json:"logo"`
				Name          string `json:"name"`
				Sourceid      int    `json:"sourceId"`
				Type          string `json:"type"`
				Watchnum      int    `json:"watchNum"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

func Add(typeId string, data map[string]interface{}, eid string) string {
	url := global.GVA_CONFIG.Param.EsHost + "/" + INDEX + "_" + typeId + "/" +TYPE
	if eid != "" {
		url = url + "/" + eid
	}
	res, _ := utils.Request(url, data, map[string]interface{}{}, "POST", "json")
	return string(res)
}

func Search(typeId string, condition map[string]interface{}, from int, num int) (EsResult, error) {
	url := global.GVA_CONFIG.Param.EsHost + "/" + INDEX + "_" + typeId + "/" +TYPE + "/_search"
	param := map[string]interface{}{
		"query" : map[string]interface{} {
			"match_phrase" : condition,
		},
		"size" : num,
		"from" : from,
	}
	res, err := utils.Request(url, param, map[string]interface{}{}, "POST", "json")
	var esResult EsResult
	json.Unmarshal(res, &esResult)
	return esResult, err
}

func DeleteIndex(typeId string) string {
	url := global.GVA_CONFIG.Param.EsHost + "/" + INDEX + "_" + typeId
	res, _ := utils.Request(url, map[string]interface{}{}, map[string]interface{}{}, "DELETE", "json")
	return string(res)
}
