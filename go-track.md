go-track



es

```bash
docker run -d --restart=always --name es --network es-net -p 9200:9200 -p 9300:9300 --privileged -v /data/es/data:/usr/share/elasticsearch/data -v /data/es/plugins:/usr/share/elasticsearch/plugins -e "discovery.type=single-node" -e "ES_JAVA_OPTS=-Xms512m -Xmx512m" elasticsearch:8.6.0

```

go-es

```
package elastic

import (
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"log"
	"os"
)

func GetEsClient(eshost string, esport string, username string, password string) (*elastic.Client, error) {
	// 创建 Elasticsearch 客户端
	url := fmt.Sprintf("http://%s:%s", eshost, esport)
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(username, password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		return nil, err
	}
	return client, err
}

```



企业微信机器人

```
promethues配置rule时annotations配置一个标签填写，组群名称，gomessage中消息模板配置为第一条数据引用这个”组名标签"
```

飞书

```
gomessage中客户端机器人配置“标题名称”为组名
```

钉钉

```
gomessage中客户端机器人配置“放行关键字”为组名
```



功能设计

```
1.完成接收aleatmanger告警消息
2.自定义告警信息的转换样式
2.兼容GoMessage的告警信息接入
```

细节要点

```
1.新增告警消息模块，只需要aleatmanger中新增接收器，自定义url中的路径即可
2.web端自动新增面包屑，可以进行告警消息的格式转换。
```









## v2

1. 原始消息接收，过滤指定json后存入es（完成）
2. 前端获取指定索引的json数据
3. 前端预显示一个markdown编辑器页面，以及该索引最新一个json原始消息页面
4. 前端点击保存后发送markdown数据到后端，后端使用sdk将json数据取出填入markdown
5. 后端api接收到数据后，保存markdown格式，全局变量key为真
6. 编写函数完成json转markdown格式
7. postapi保存告警消息的同时判断全局变量key为真，执行6的函数，并发送到钉钉



1. 是否存在模板
2. 存在执行转换函数
3. 发送钉钉
