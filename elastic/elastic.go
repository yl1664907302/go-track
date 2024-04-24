package elastic

import (
	"context"
	"encoding/json"
	"fmt"
	elastic "github.com/olivere/elastic/v7"
	"kube-auto/global"
	"kube-auto/pojo"
	"log"
	"os"
	"strconv"
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

func CreateIndexES(message *pojo.Message, index string) error {
	ESclient, err := GetEsClient(
		global.CONF.System.Elasticsearch.Eshost,
		global.CONF.System.Elasticsearch.Esport,
		global.CONF.System.Elasticsearch.Username,
		global.CONF.System.Elasticsearch.Password)
	if err != nil {
		return err
	}
	marshaler, err := json.Marshal(message)
	if index == "dingtalk" || index == "wechat_robot" || index == "feishu" {
		_, err = ESclient.Index().Index(index).BodyString(string(marshaler)).Do(context.Background())
		log.Print("消息成功写入索引" + ":" + index)
	} else {
		return fmt.Errorf("不支持的消息类型！")
	}
	return err
}

func SelectEsDocByIndex2keyword(fenye *pojo.Fenye, groupname string, time string, keyword1 string) ([]pojo.Message, error) {
	var messages []pojo.Message
	ESclient, err := GetEsClient(
		global.CONF.System.Elasticsearch.Eshost,
		global.CONF.System.Elasticsearch.Esport,
		global.CONF.System.Elasticsearch.Username,
		global.CONF.System.Elasticsearch.Password)
	if err != nil {
		return messages, err
	}

	// 构建模糊查询

	query := elastic.NewBoolQuery().
		Must(elastic.NewWildcardQuery("groupname.keyword", "*"+groupname+"*")).
		Must(elastic.NewWildcardQuery("time.keyword", "*"+time+"*")).
		Must(elastic.NewWildcardQuery("contests.linecontext.keyword", "*"+keyword1+"*"))
	//转换参数类型
	asc, _ := strconv.ParseBool(fenye.Asc)
	from, _ := strconv.Atoi(fenye.From)
	size, _ := strconv.Atoi(fenye.Size)
	// 执行查询
	searchResult, err := ESclient.Search().
		Index(fenye.Index). // 指定索引名称
		Query(query).       // 设置查询
		From(from).
		Size(size).
		Sort(fenye.SortField, asc). // 排序字段和顺序
		Do(context.Background())
	if err != nil {
		return messages, err
	}

	for _, hit := range searchResult.Hits.Hits {
		var message pojo.Message
		err := json.Unmarshal(hit.Source, &message)
		if err != nil {
			return messages, err
		}
		//fmt.Printf("Document ID: %s\n", hit.Id)
		//fmt.Printf("%s\n", hit.Source)
		//fmt.Println()
		messages = append(messages, message)
	}
	return messages, err
}

// 分页查询函数
func PaginateSearchEsDoc(fenye *pojo.Fenye) ([]pojo.Message, error) {
	var messages []pojo.Message
	ESclient, err := GetEsClient(
		global.CONF.System.Elasticsearch.Eshost,
		global.CONF.System.Elasticsearch.Esport,
		global.CONF.System.Elasticsearch.Username,
		global.CONF.System.Elasticsearch.Password)
	if err != nil {
		return messages, err
	}

	// 指定查询条件
	query := elastic.NewMatchAllQuery() // 例如：匹配所有文档
	//转换参数类型
	asc, _ := strconv.ParseBool(fenye.Asc)
	from, _ := strconv.Atoi(fenye.From)
	size, _ := strconv.Atoi(fenye.Size)
	// 执行查询
	searchResult, err := ESclient.Search().
		Index(fenye.Index).
		Query(query).
		From(from).
		Size(size).
		Sort(fenye.SortField, asc). // 排序字段和顺序
		Do(context.Background())
	if err != nil {
		log.Println(err)
	}
	//// 打印查询结果
	//fmt.Printf("Total hits: %d\n", searchResult.TotalHits())
	for _, hit := range searchResult.Hits.Hits {
		var message pojo.Message
		err := json.Unmarshal(hit.Source, &message)
		if err != nil {
			return messages, err
		}
		//fmt.Printf("Document ID: %s\n", hit.Id)
		//fmt.Printf("%s\n", hit.Source)
		//fmt.Println()
		messages = append(messages, message)
	}
	return messages, err
}
