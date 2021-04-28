package main

import (
	"boot/db"
	"boot/protos"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
	"time"
)

func notice(w http.ResponseWriter, r *http.Request) {

	fmt.Println("connect")
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	defer c.Close()
	_, message, err := c.ReadMessage()
	noticeRequest := new(protos.NoticeRequest)
	if err = proto.Unmarshal(message, noticeRequest); err != nil {
		log.Println(err.Error())
	}
	// 获取该用户未读信息
	userId := noticeRequest.UserId
	ctx := context.Background()
	notices, err := db.FindNoticeByUserIdAndType(ctx, int(userId), int(noticeRequest.Type))
	if err != nil {
		log.Println(err.Error())
	}
	// 发送信息
	var noticeResponse protos.NoticeResponse
	for i := 0; i < len(notices); i++ {
		n := &protos.Notice{
			Id:      int32(notices[i].ID),
			Date:    notices[i].CreateTime.Format("06-01-02 15:04:05"),
			Content: notices[i].Content,
			UserId:  int32(notices[i].UserId),
		}
		noticeResponse.Notices = append(noticeResponse.Notices, n)
	}
	errorNotify(c, &noticeResponse)

}

// 警告信息⚠
func errorNotify(c *websocket.Conn, notices *protos.NoticeResponse) {
	//notice := new(protos.Notice)
	//notice.Content = fmt.Sprintf("刀具%d已经发生了损坏", id)
	//notice.Date = time.Now().Format("15:04:05")

	marshal, err := proto.Marshal(notices)

	if err != nil {
		log.Println("消息序列化失败")
		return
	}

	err = c.WriteMessage(websocket.BinaryMessage, marshal)
	if err != nil {
		log.Println("消息发送失败")
		return
	}
}

// 错误信息❌
func waringNotify(c *websocket.Conn, id int) {
	notice := new(protos.Notice)
	notice.Content = fmt.Sprintf("刀具%d已经发生了损坏!!", id)
	notice.Date = time.Now().Format("15:04:05")

	marshal, err := proto.Marshal(notice)

	if err != nil {
		log.Println("消息序列化失败", notice.Content)
		return
	}

	err = c.WriteMessage(websocket.BinaryMessage, marshal)
	if err != nil {
		log.Println("消息发送失败", notice.Content)
		return
	}
}
