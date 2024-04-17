package pojo

type FeishuMarkdownMessage struct {
	MsgType string `json:"msg_type"`
	Card    struct {
		Header struct {
			Title    Title  `json:"title"`
			Template string `json:"template"`
		} `json:"header"`
		Elements []Element `json:"elements"`
	} `json:"card"`
}

type Title struct {
	Tag     string `json:"tag"`
	Content string `json:"content"`
}

type Element struct {
	Lines   int    `json:"lines"`
	Tag     string `json:"tag"`
	Content string `json:"Content"`
}
