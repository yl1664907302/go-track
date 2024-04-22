package utils

import (
	"strings"
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
