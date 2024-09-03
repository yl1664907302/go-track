package pojo

type Stepform struct {
	Niname      string `json:"niname"`
	Robot_ok    bool   `json:"robot_ok"`
	Receiver    string `json:"receiver"`
	Robot_class string `json:"robot_class"`
	Robot_id    int    `json:"robot_id"`
	Switch      string `json:"switch"`
	Accesstoken string `json:"accesstoken"`
	Secret      string `json:"secret"`
	Markdown_ok bool   `json:"markdown_ok"`
	Desc        Desc   `json:"desc"`
}
