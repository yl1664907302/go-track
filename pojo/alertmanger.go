package pojo

type Alert struct {
	Receiver    string   `json:"receiver"`
	Status      string   `json:"status"`
	Alerts      []Alerts `json:"alerts"`
	GroupLabels struct {
		Alertname string `json:"alertname"`
	} `json:"groupLabels"`
	CommonLabels      interface{} `json:"commonLabels"`
	CommonAnnotations interface{} `json:"commonAnnotations"`
	ExternalURL       string      `json:"externalURL"`
	Version           string      `json:"version"`
	GroupKey          string      `json:"groupKey"`
	TruncatedAlerts   int         `json:"truncatedAlerts"`
}

type Alerts struct {
	Status       string      `json:"status"`
	Labels       interface{} `json:"labels"`
	Annotations  interface{} `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorURL string      `json:"generatorURL"`
	Fingerprint  string      `json:"fingerprint"`
}

//type Labels struct {
//	Alertname      string `json:"alertname"`
//	Attribute      string `json:"attribute"`
//	Env            string `json:"env"`
//	Idc            string `json:"idc"`
//	Instance       string `json:"instance"`
//	Job            string `json:"job"`
//	ObjectSummary  string `json:"object_summary"`
//	Quality        string `json:"quality"`
//	RuleCreateUser string `json:"ruleCreateUser"`
//	ServiceName    string `json:"service_name"`
//	Severity       string `json:"severity"`
//	Type           string `json:"type"`
//}

//type Annotations struct {
//	Description string `json:"description"`
//	Summary     string `json:"summary"`
//}

//type CommonLabels struct {
//	AdditionalData map[string]interface{} `json:"additional_data,omitempty"`
//	Alertname      string `json:"alertname"`
//	Attribute      string `json:"attribute"`
//	Env            string `json:"env"`
//	Idc            string `json:"idc"`
//	Instance       string `json:"instance"`
//	Job            string `json:"job"`
//	ObjectSummary  string `json:"object_summary"`
//	Quality        string `json:"quality"`
//	RuleCreateUser string `json:"ruleCreateUser"`
//	ServiceName    string `json:"service_name"`
//	Severity       string `json:"severity"`
//	Type           string `json:"type"`
//}
//
//type CommonAnnotations struct {
//	Description string `json:"description"`
//	Summary     string `json:"summary"`
//}
