package pojo

type DingtalkMarkdownMessage struct {
	MsgType  string `json:"msgtype"`
	Markdown struct {
		Title string `json:"title"`
		Text  string `json:"text"`
	} `json:"markdown"`
	At struct {
		AtMobiles []interface{} `json:"atMobiles"`
		AtAll     bool          `json:"atAll"`
	} `json:"at"`
}
