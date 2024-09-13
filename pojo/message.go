package pojo

type Message struct {
	Time      string    `json:"time"`
	PlatForm  string    `json:"platform"`
	GroupName string    `json:"groupname"`
	Contests  []Context `json:"contests"`
}

type Context struct {
	Num         int    `json:"num"`
	LineName    string `json:"linename"`
	LineContext string `json:"linecontext"`
}

type Fenye struct {
	Status     string `json:"status"`
	Index      string `json:"index"`
	From       string `json:"from"`
	Size       string `json:"size"`
	SortField  string `json:"sort_field"`
	Asc        string `json:"asc"`
	Time_start string `json:"time_start"`
	Time_end   string `json:"time_end"`
}
