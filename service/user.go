package service

import (
	"encoding/json"
	"fmt"
	"github.com/Chain-Zhang/pinyin"
	"github.com/gin-gonic/gin"
	"github.com/zjswh/go-tool/ipSearch"
	"math"
	"mtsw/config"
	"mtsw/middleware"
	response2 "mtsw/types/response"
	"mtsw/utils"
	"strconv"
	"strings"
)

type UserInfo struct {
	AccountId int    `json:"accountId"`
	Aid       int    `json:"aid"`
	Uin       int    `json:"uin"`
	Name      string `json:"name"`
}

type CUserInfo struct {
	Uin         int    `json:"uin"`
	Id          int    `json:"id"`
	IsSafe      int    `json:"isSafe"`
	Phone       string `json:"phone"`
	UserNick    string `json:"userNick"`
	LoginType   string `json:"loginType"`
	UserHeadImg string `json:"userHeadImg"`
	UserIp      string `json:"userIp"`
	OpenUid     string `json:"openUid"`
}

func GetUserInfo(c *gin.Context) config.UserInfo {
	parse := c.GetString("userInfo")
	userInfo := config.UserInfo{}
	json.Unmarshal([]byte(parse), &userInfo)
	return userInfo
}

func CheckLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		if token == "" {
			token = c.Query("token")
		}

		if token == "" {
			response2.Result(1, "", "未登录", c)
			c.Abort()
			return
		}
		realToken := middleware.JwtAuth(token)
		if realToken == "" {
			response2.Result(1, "", "登录过期", c)
			c.Abort()
			return
		}
		url := config.GVA_CONFIG.Param.BGatewayHost + "/v1/Passport/Index/getLoginInfo"
		url = fmt.Sprintf("%s?token=%s&path=/%s&method=%s", url, token, c.FullPath(), c.Request.Method)
		res, _ := utils.RequestGet(url)
		var result response2.Response
		json.Unmarshal(res, &result)
		if result.Code != 200 || result.ErrorCode != 0 {
			response2.Result(result.ErrorCode, "", result.ErrorMessage, c)
			c.Abort()
			return
		}

		userInfo, _ := json.Marshal(result.Data)
		c.Set("userInfo", string(userInfo))
		c.Next()
	}
}


func GetCUserInfo(c *gin.Context) CUserInfo {
	userInfo := CUserInfo{}
	token := c.GetHeader("token")
	if token == "" {
		return defaultCUserInfo(c)
	}

	realToken := middleware.JwtAuth(token)
	if realToken == "" {
		return defaultCUserInfo(c)
	}

	info, _ := config.GVA_REDIS.Get(realToken).Result()
	json.Unmarshal([]byte(info), &userInfo)
	return userInfo
}

func defaultCUserInfo(c *gin.Context) CUserInfo {
	ip := c.ClientIP()
	ipHandle, _ := ipSearch.New()
	res := ipHandle.Get(ip)
	ipAddressArr := strings.Split(res, "|")
	userNick := fmt.Sprintf("%s%s", ipAddressArr[2], ipAddressArr[3])
	str, _ := pinyin.New(userNick).Split("").Mode(pinyin.WithoutTone).Convert()
	//四位数生成规则
	ipArr := strings.Split(ip, ".")
	ip1, _ := strconv.Atoi(ipArr[1])
	ip2, _ := strconv.Atoi(ipArr[2])
	ip3, _ := strconv.Atoi(ipArr[3])
	randomStr := int64(math.Abs(float64(ip2 * (ip3 - ip1))))
	openUid := fmt.Sprintf("ady_%s_%d", str, randomStr)
	userNick = fmt.Sprintf("%s网友%d", userNick, randomStr)

	defaultInfo := CUserInfo{
		UserNick: userNick,
		OpenUid:  openUid,
		UserIp:   ip,
	}
	return defaultInfo
}
