package appConst

import "fmt"

const (
	//ErrorMessage
	NotFind         = "记录不存在或已被删除"
	ReplyMe         = "不能回复自己"
	NotLiving       = "直播间未开始直播"
	Empty           = "内容不得为空"
	PicLimit        = "上传图片不得多于九张"
	TimeError       = "自定义时间不得大于当前时间"
	PleaseCloseLive = "请先关闭直播后再删除"
	GuestNotAllow   = "游客不允许评论"

	//Redis
	VirtualViewCountKey = "pic_text_virtual_viewcount"
	ViewCountKey        = "pic_text_viewcount"
	WatchAnalysisKey    = "pic_text_watch_analysis"
	VirtualStarKey      = "pic_text_virtual_star"
	StarCountKey        = "pic_text_star_count"
	StarAnalysisKey     = "pic_text_star_analysis"
	RoomInfoKey         = "pic_text_room_info"
	ReportKey           = "pic_text_report"
	ContentKey          = "pic_text_content"
	CommentKey          = "pic_text_comment"
)

func GetMailInfoKey(uin, aid int) string {
	return fmt.Sprintf("gdy_mtsw_mail_info:%d_%d", uin, aid)
}
