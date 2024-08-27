package elastics

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/olivere/elastic/v7"
	"go-track/global"
	"go-track/pojo"
	"log"
	"os"
	"strconv"
)

func GetEsClient() (*elastic.Client, error) {
	// 创建 Elasticsearch 客户端
	url := fmt.Sprintf("http://%s:%s", global.CONF.System.Elasticsearch.Eshost, global.CONF.System.Elasticsearch.Esport)
	client, err := elastic.NewClient(
		//elastics 服务地址
		elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.CONF.System.Elasticsearch.Username, global.CONF.System.Elasticsearch.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	return client, err
}

func CreateIndexESForAlert(message []pojo.Alerts, index string) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		return err, ""
	}

	for _, m := range message {
		marshaler, err := json.Marshal(m)
		if err != nil {
			log.Println(err)
		}
		log.Println(string(marshaler))
		_, err = ESclient.Index().Index(index).BodyString(string(marshaler)).Do(context.Background())
	}
	if err == nil {
		log.Print("消息成功写入索引" + ":" + index)
		return err, "消息成功写入索引" + ":" + index
	}
	return err, ""
}

func TestCreateIndexForAlert(message pojo.Alerts) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		return err, "es客户端创建失败"
	}

	marshaler, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(marshaler))
	_, err = ESclient.Index().Index("l1").BodyString(string(marshaler)).Do(context.Background())

	if err == nil {
		log.Print("消息成功写入索引" + ":" + "l1")
	}
	return err, ""
}

func CreateIndexForMarkDown(message *pojo.Desc, index string) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		return err, "es客户端创建失败"
	}
	marshaler, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(marshaler))
	num, err := JudgeMarkdownTemp(index)
	if err != nil && num != 1 {
		return err, ""
	}
	_, err = ESclient.Index().Index(index + "_t").BodyString(string(marshaler)).Do(context.Background())

	if err == nil {
		log.Print("markdown模板成功写入索引" + ":" + index + "_t")
	}
	return err, message.Markdown
}

func UpdateIndexForMarkDown(message *pojo.Markdown, index string) (error, string) {
	marshaler, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(marshaler))
	doc_id, err := SelectNewDocidByindex(index+"_t", "maketime.keyword")
	if err != nil {
		return err, ""
	}
	if doc_id == "" {
		return fmt.Errorf("模板不存在无法更新"), ""
	} else {
		err := UpdateDoc(index+"_t", doc_id, message.Desc)
		if err != nil {
			return err, ""
		}
	}
	return err, message.Desc.Markdown
}

func SelectNewMarkdownTempByIndex(index string) (error, *pojo.Markdown) {
	var markdown pojo.Markdown
	byindex, err := SelectNewDocByindex(index+"_t", "maketime.keyword", pojo.Markdown{})
	if err != nil {
		return err, nil
	}
	message := byindex.(json.RawMessage)
	err = sonic.Unmarshal(message, &markdown.Desc)
	return err, &markdown
}

func JudgeMarkdownTemp(index string) (int, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		return 0, err
	}
	do, err := ESclient.IndexExists(index + "_t").Do(context.Background())
	if !do {
		return 1, fmt.Errorf("索引%s不存在", index)
	} else {
		do, _ := ESclient.Count(index + "_t").Do(context.Background())
		//一个接收者只容许保留一个模板
		if int(do) != 0 {
			return 0, fmt.Errorf("模板已存在无法新增")
		}
		return 0, nil
	}
}

func CreateIndexES(message interface{}, index string) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		return err, ""
	}

	marshaler, err := json.Marshal(message)

	_, err = ESclient.Index().Index(index).BodyString(string(marshaler)).Do(context.Background())
	if err != nil {
		log.Print("消息成功写入索引" + ":" + index)
		return err, "消息成功写入索引" + ":" + index
	}

	return err, ""
}

func SelectEsDocByIndex2keyword(fenye *pojo.Fenye, groupname string, time string, keyword1 string) ([]pojo.Message, error) {
	var messages []pojo.Message
	ESclient, err := GetEsClient()
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
		messages = append(messages, message)
	}
	return messages, err
}

// 分页查询函数
func PaginateSearchEsDoc(fenye *pojo.Fenye) ([]pojo.Message, error) {
	var messages []pojo.Message
	ESclient, err := GetEsClient()
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

	for _, hit := range searchResult.Hits.Hits {
		var message pojo.Message
		err := json.Unmarshal(hit.Source, &message)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}
	return messages, err
}

func SearchBySortAndUnique(fenye *pojo.Fenye, key bool) ([]pojo.Alerts, []pojo.Newmarkdown, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}

	size, _ := strconv.Atoi(fenye.Size)
	var zhiwen string
	var shijian string
	if key == false {
		zhiwen = "fingerprint.keyword"
		shijian = "startsAt"
	} else {
		zhiwen = "zhiwen.keyword"
		shijian = "time"
	}

	// 配置聚合语句
	aggs := elastic.NewTermsAggregation().Field(zhiwen).Size(1000).SubAggregation("latest_alert",
		elastic.NewTopHitsAggregation().Sort(shijian, false).Size(1).FetchSourceContext(elastic.NewFetchSourceContext(true).Include("*")))
	// 执行查询
	searchResult, err := ESclient.Search().
		Index(fenye.Index).
		Size(size).
		Aggregation("unique_alerts", aggs).
		Do(context.Background())
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil, nil, err
	}
	if key == false {
		var alerts []pojo.Alerts
		agg, found := searchResult.Aggregations.Terms("unique_alerts")
		if found {
			for _, bucket := range agg.Buckets {
				tophits, _ := bucket.TopHits("latest_alert")
				for _, hit := range tophits.Hits.Hits {
					var alert pojo.Alerts
					err := sonic.UnmarshalString(string(hit.Source), &alert)
					if err != nil {
						log.Println(err)
					}
					alerts = append(alerts, alert)
				}
			}
		}
		return alerts, nil, err
	}

	var markdowns []pojo.Newmarkdown
	agg, found := searchResult.Aggregations.Terms("unique_alerts")
	if found {
		for _, bucket := range agg.Buckets {
			tophits, _ := bucket.TopHits("latest_alert")
			for _, hit := range tophits.Hits.Hits {
				var newmarkdown pojo.Newmarkdown
				err := sonic.UnmarshalString(string(hit.Source), &newmarkdown)
				if err != nil {
					log.Println(err)
				}
				markdowns = append(markdowns, newmarkdown)
			}
		}
	}
	return nil, markdowns, err
}

func SearchMarkDown(index string) (bool, *pojo.Desc, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	// 执行查询
	searchResult, err := ESclient.Search().
		Index(index).
		Size(1).
		Sort("maketime.keyword", false).
		Do(context.Background())
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return false, nil, err
	}

	if searchResult.TotalHits() == 0 {
		log.Println("No documents found")
		return false, nil, err
	}

	//目前只获取最新的模板
	var desc pojo.Desc
	for _, hit := range searchResult.Hits.Hits {
		err := sonic.Unmarshal(hit.Source, &desc)
		if err != nil {
			log.Printf("Error unmarshalling document: %s", err)
		}
	}

	return true, &desc, err
}

func SearchRobot(index string) ([]pojo.Robot, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	// 执行查询
	searchResult, err := ESclient.Search().
		Index(index + "_r").
		Do(context.Background())
	if err != nil {
		log.Printf("Error getting response: %s", err)
		return nil, err
	}

	if searchResult.TotalHits() == 0 {
		log.Println("No documents found")
		return nil, err
	}

	var robots []pojo.Robot
	for _, hit := range searchResult.Hits.Hits {
		var robot pojo.Robot
		err := sonic.Unmarshal(hit.Source, &robot)
		if err != nil {
			log.Printf("Error unmarshalling document: %s", err)
		}
		robots = append(robots, robot)
	}

	return robots, err
}

func CreateIndexForNewMarkDown(newmarkdown *pojo.Newmarkdown, index string) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	marshaler, err := json.Marshal(&newmarkdown)
	if err != nil {
		log.Println(err)
	}
	_, err = ESclient.Index().Index(index + "_n").BodyString(string(marshaler)).Do(context.Background())
	return err, ""
}

func CreateIndexForRobot(robot *pojo.Robot, index string) (error, string) {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	marshaler, err := json.Marshal(&robot)
	if err != nil {
		log.Println(err)
	}
	_, err = ESclient.Index().Index(index + "_r").BodyString(string(marshaler)).Do(context.Background())
	return err, ""
}

func DelIndex(index string) error {
	ESclient, err := GetEsClient()
	ctx := context.Background()
	if err != nil {
		log.Println(err)
	}
	deleteIndex, err := ESclient.DeleteIndex(index).Do(ctx)
	if deleteIndex.Acknowledged {
		log.Printf("已删除索引: %s", index)
	}
	return err
}

func DelDocByKey(index string, key string, value string) error {
	//查询是否存在该文档
	doc_id, err := SelectDocidBySome(index, key, value)
	if doc_id == "" {
		log.Printf("索引%s中不存在%s值为%s的doc", index, key, value)
		return fmt.Errorf("索引%s中不存在%s值为%s的doc", index, key, value)
	} else {
		err := DelDoc(index, doc_id)
		if err != nil {
			return err
		}
	}
	return err
}

func DelDoc(index string, doc_id string) error {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	do, err := ESclient.Delete().Index(index).Id(doc_id).Do(context.Background())
	if do != nil {
		log.Printf("doc%s删除成功", do.Id)
	}
	return err
}

func UpdateDoc(index string, doc_id string, any interface{}) error {
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	_, err = ESclient.Update().Index(index).Id(doc_id).Doc(any).Do(context.Background())
	return err
}

func UpdateDocForRobot(index string, robot pojo.Robot) error {
	doc_id, err := SelectDocidBySome(index, "robot_id", robot.Robot_id)
	if doc_id == "" {
		log.Printf("索引%s中不存在%s值为%s的doc", index, "robot_id", robot.Robot_id)
		return fmt.Errorf("索引%s中不存在%s值为%s的doc", index, "robot_id", robot.Robot_id)
	} else {
		err = UpdateDoc(index, doc_id, robot)
		if err != nil {
			return err
		}
	}
	return err
}

func SelectDocidBySome(index string, key string, value string) (string, error) {
	var doc_id string
	ESclient, err := GetEsClient()
	if err != nil {
		log.Println(err)
	}
	//创建复合查询
	boolquery := elastic.NewBoolQuery()
	boolquery.Must(elastic.NewTermsQuery(key, value))

	do, err := ESclient.Search(index).Query(boolquery).Do(context.Background())
	fmt.Println(do.Hits.TotalHits)
	for _, hit := range do.Hits.Hits {
		if hit == nil {
			return "", err
		}
		doc_id = hit.Id
	}
	return doc_id, err
}

func SelectNewDocidByindex(index string, key string) (string, error) {
	var doc_id string
	ESclient, err := GetEsClient()
	if err != nil {
		return "", err
	}
	searchResult, err := ESclient.Search().
		Index(index).            // 设置索引名称
		Sort(key, false).        // 根据时间戳字段排序，false表示降序
		Size(1).                 // 只获取一条记录
		Do(context.Background()) // 执行查询
	if err != nil {
		return "", err
	}

	fmt.Println(searchResult.Hits.TotalHits)
	for _, hit := range searchResult.Hits.Hits {
		if hit == nil {
			return "", err
		}
		doc_id = hit.Id
	}
	return doc_id, err
}
func SelectNewDocByindex(index string, key string, any interface{}) (interface{}, error) {
	ESclient, err := GetEsClient()
	if err != nil {
		return nil, err
	}
	searchResult, err := ESclient.Search().
		Index(index).            // 设置索引名称
		Sort(key, false).        // 根据时间戳字段排序，false表示降序
		Size(1).                 // 只获取一条记录
		Do(context.Background()) // 执行查询
	if err != nil {
		return nil, err
	}

	for _, hit := range searchResult.Hits.Hits {
		if hit == nil {
			return nil, err
		}
		any = hit.Source
	}
	return any, err
}
