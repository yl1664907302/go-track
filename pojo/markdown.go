package pojo

type Markdown struct {
	Receiver string `json:"receiver"`
	Desc     Desc   `json:"desc"`
}

type Desc struct {
	Markdown string `json:"markdown"`
	Maketime string `json:"maketime"`
}
