package utils

import (
	"bytes"
	"go-track/pojo"
	"html/template"
	"log"
	"strings"
	"unicode"
)

type ActionMessage struct {
}

type Info struct {
	Key   string
	Value string
}

func (*ActionMessage) ExtractInfo(markdownContent string, keyword bool) ([]Info, string) {
	var infoSlice []Info

	lines := strings.Split(markdownContent, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "**") {
			// 提取键值对
			parts := strings.SplitN(line, "：", 2)
			if len(parts) == 2 {
				key := strings.Trim(parts[0], "* ")
				value := strings.TrimSpace(parts[1])
				info := Info{
					Key:   key,
					Value: value,
				}
				infoSlice = append(infoSlice, info)
			}
		}
	}
	if keyword == true {
		return infoSlice, infoSlice[0].Value
	}

	return infoSlice, ""
}

func (*ActionMessage) EditFisrtCharToLower(s string) string {
	//首字母大写转小写
	runes := []rune(s)
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}

func (*ActionMessage) InsertJsonToMarkdown(desc *pojo.Desc, alert *pojo.Alerts) (string, error) {
	tmpl, err := template.New("markdown").Parse(desc.Markdown)
	if err != nil {
		log.Println(err)
		return "", err
	}
	var newmarkdown bytes.Buffer
	//该模板填入json数据源，以及实例载体
	err = tmpl.Execute(&newmarkdown, alert)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return newmarkdown.String(), err
}

var ActionMessages ActionMessage
