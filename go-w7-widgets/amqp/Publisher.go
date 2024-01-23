package amqp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"w7_widgets/amqp/encrypt"
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	// 队列名称
	QueueName string
	// 交换机
	Exchange string
	// routing Key
	RoutingKey string
	//MQ链接字符串
	Mqurl string
}

func NewRabbitMQ(queueName, exchange, routingKey, ak, sk, instanceId, host, port, vHost string) *RabbitMQ {

	rabbitMQ := RabbitMQ{
		QueueName:  queueName,
		Exchange:   exchange,
		RoutingKey: routingKey,
	}

	var buf bytes.Buffer

	// amqp://username:password@host:port/vHost
	userName := encrypt.GetUserName(ak, instanceId)
	password := encrypt.GetPassword(sk)
	buf.WriteString("amqp://")
	buf.WriteString(userName)
	buf.WriteString(":")
	buf.WriteString(password)
	buf.WriteString("@")
	buf.WriteString(host)
	buf.WriteString(":")
	buf.WriteString(port)
	buf.WriteString("/")
	buf.WriteString(vHost)
	rabbitMQ.Mqurl = buf.String()

	var err error
	//创建rabbitmq连接
	rabbitMQ.Conn, err = amqp.Dial(rabbitMQ.Mqurl)
	checkErr(err, "创建连接失败")

	//创建Channel
	rabbitMQ.Channel, err = rabbitMQ.Conn.Channel()
	checkErr(err, "创建channel失败")

	return &rabbitMQ

}

// ReleaseRes 释放资源,建议NewRabbitMQ获取实例后 配合defer使用
func (mq *RabbitMQ) ReleaseRes() {
	connErr := mq.Conn.Close()
	if connErr != nil {
		return
	}
	channelErr := mq.Channel.Close()
	if channelErr != nil {
		return
	}
}

func checkErr(err error, meg string) {
	if err != nil {
		log.Fatalf("%s:%s\n", meg, err)
	}
}

func DoTest2() {

	// 初始化mq
	mq := NewRabbitMQ("hd-live-callback-common-entrance-queue", "hd-live-callback-common-entrance-exchange",
		"hd-live-callback-common-entrance-routing-key",
		"", "", "",
		"", "5672", "hd-live-task-callback-jd")

	defer mq.ReleaseRes() // 完成任务释放资源

	err := mq.Channel.ExchangeDeclare(mq.Exchange, "direct",
		true, false, false, false, nil)
	failOnError(err, "Failed to Declare a exchange")

	q, err := mq.Channel.QueueDeclare(
		mq.QueueName,
		true,
		false,
		false,
		false,
		map[string]interface{}{"x-retry-interval": 1800},
	)
	failOnError(err, "Failed to declare a queue"+q.Name)

	err = mq.Channel.QueueBind(mq.QueueName, mq.RoutingKey, mq.Exchange, false, nil)
	failOnError(err, "Failed to bind a queue")

	/*
	 * rabbitmq client 向Server发起connection,新建channel大约需要进行15+个TCP报文的传输，会消耗大量网络资源和Server端的资源，甚至引起Server端SYN flooding 攻击保护。
	 * 因此我们建议消息的发送和消费尽量采用长链接的模式。
	 */
	body := map[string]interface{}{
		"accountId":       38,
		"entranceTypeStr": "SENSITIVE_WORDS_MANAGE",
		"creditData": map[string]interface{}{
			"action": "add",
			"words":  "kkkk,ddd,哈哈哈",
		},
	}
	byteBody, _ := json.Marshal(body)
	err = mq.Channel.Publish(
		mq.Exchange,
		mq.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteBody,
		})

	failOnError(err, "Failed to publish a message")

	fmt.Println(123123)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
