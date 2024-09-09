package pojo

type Newmarkdown struct {
	Status   string `json:"status"`
	Zhiwen   string `json:"zhiwen"`
	Time     string `json:"time"`
	Markdown string `json:"markdown"`
}

func NewNewmarkdown(status string, zhiwen string, time string, markdown string) *Newmarkdown {
	return &Newmarkdown{Status: status, Zhiwen: zhiwen, Time: time, Markdown: markdown}
}
