package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

var (
	client  *elastic.Client
	logChan chan *LogData
)

// LogData ...
type LogData struct {
	Topic string
	Data  map[string]interface{}
}

// Init 初始化es一个client连接
func Init(address string, maxSize int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		return err
	}
	fmt.Println("connect to es success")
	// 初始化channel
	logChan = make(chan *LogData, maxSize)
	// 开启一个后台goroutine循环往ES中发送日志
	go sendToEs()
	return
}

// SendToChan 向外提供一个往内部通道中发送数据的函数
func SendToChan(topic string, data map[string]interface{}) {
	ld := &LogData{
		Topic: topic,
		Data:  data,
	}
	logChan <- ld
}

// sendToEs 向外提供一个往es发送数据的函数
func sendToEs() {
	for {
		select {
		case ld := <-logChan:
			put1, err := client.Index().
				Index(ld.Topic).
				Type("xxx").
				BodyJson(ld.Data).
				Do(context.Background())
			if err != nil {
				// Handle error
				panic(err)
			}
			fmt.Printf("Indexed user %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}

}
