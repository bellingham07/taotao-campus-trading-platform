package chatApi

import (
	chatLogic "com.xpdj/go-gin/logic/chat"
	"com.xpdj/go-gin/model"
	"com.xpdj/go-gin/model/response"
	chatRepository "com.xpdj/go-gin/repository/chat"
	"com.xpdj/go-gin/router/middleware"
	"com.xpdj/go-gin/utils/jsonUtil"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)
import "github.com/gorilla/websocket"

type RoomApi struct {
}

func (RoomApi) CreateRoom(c *gin.Context) {
	cr := new(model.ChatRoom)
	err := c.ShouldBind(cr)
	fmt.Println(cr)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！🤡"))
		return
	}
	c.JSON(http.StatusOK, chatLogic.ChatRoom.CreateRoom(cr))
}

var upgrader websocket.Upgrader

var wc = make(map[int64]*myConnection)

type myConnection struct {
	*websocket.Conn
}

func (RoomApi) Chat(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	myConn := &myConnection{conn}
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, response.ErrorMsg("无法进入聊天，请重试😥"))
		return
	}
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Println("地欧弟1", err)
		c.JSON(http.StatusBadRequest, response.ErrorMsg("请求参数错误！🤡"))
		return
	}
	cr := &model.ChatRoom{Id: id}
	err = chatRepository.ChatRoom.QueryById(cr)
	if err != nil {
		log.Println("地欧弟2", err)
		c.JSON(http.StatusBadRequest, response.ErrorMsg("房间不存在！😫"))
		return
	}
	wc[id] = myConn
	userId := middleware.GetUserId(c)
	for {
		msg := &model.ChatMessage{
			RoomId: id,
			Time:   time.Now(),
			UserId: userId,
		}
		err = myConn.readJSON(msg)
		if err != nil {
			log.Println("地欧弟3", err)
			c.JSON(http.StatusBadRequest, response.ErrorMsg("消息发送失败"))
			return
		}
		if err = chatRepository.ChatMessage.Insert(msg); err != nil {
			log.Println("地欧弟4", err)
			c.JSON(http.StatusBadRequest, response.ErrorMsg("消息发送失败"))
			return
		}
		if cc, ok := wc[msg.ToUserId]; ok {
			err = cc.WriteMessage(websocket.TextMessage, []byte(msg.Content))
			if err != nil {
				log.Printf("Write Message Error:%v\n\n", err)
				return
			}
		}
	}
}

func (c *myConnection) readJSON(v interface{}) error {
	_, r, err := c.NextReader()
	if err != nil {
		return err
	}
	err = jsonUtil.Json.NewDecoder(r).Decode(v)
	if err == io.EOF {
		// One value is expected in the message.
		err = io.ErrUnexpectedEOF
	}
	return err
}
