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
	_, err = ESclient.Index().Index(index + "_t").BodyString(string(marshaler)).Do(context.Background())

	if err == nil {
		log.Print("markdown模板成功写入索引" + ":" + index + "_t")
	}
	return err, message.Markdown
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

func CreateIndexForDingTalkRobot(robot *pojo.Robot, index string) (error, string) {
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

func DelIndexForDingTalkRobot(index string) error {
	ESclient, err := GetEsClient()
	ctx := context.Background()
	if err != nil {
		log.Println(err)
	}
	deleteIndex, err := ESclient.DeleteIndex(index + "_r").Do(ctx)
	if deleteIndex.Acknowledged {
		log.Printf("已删除索引: %s", index)
	}
	return err
}
