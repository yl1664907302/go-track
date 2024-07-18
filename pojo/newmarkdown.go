package pojo

type Newmarkdown struct {
	Zhiwen   string `json:"zhiwen"`
	Time     string `json:"time"`
	Markdown string `json:"markdown"`
}

func NewNewmarkdown(zhiwen string, time string, markdown string) *Newmarkdown {
	return &Newmarkdown{Zhiwen: zhiwen, Time: time, Markdown: markdown}
}
