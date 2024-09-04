package pojo

type Robot struct {
	Robot_name  string `json:"robot_name"`
	Receiver    string `json:"receiver"`
	Robot_class string `json:"robot_class"`
	Robot_id    int    `json:"robot_id"`
	Switch      bool   `json:"switch"`
	Accesstoken string `json:"accesstoken"`
	Secret      string `json:"secret"`
}
