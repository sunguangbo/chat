package service

import (
	"chat/dal"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"net/http"
	"strings"
	"time"
)

func GetAccessToken(c context.Context) (string, error) {
	result, err := dal.Redis.Get(c, AccessTokenKey).Result()
	if err == nil {
		return result, nil
	}
	if err == redis.Nil {
		resp, err := Request(GET, AccessTokenUrl, "")
		if err != nil {
			return "", err
		}
		var tokenModel AccessTokenModel

		err = json.Unmarshal(resp, &tokenModel)
		if err != nil {
			return "", err
		}
		dal.Redis.Set(c, AccessTokenKey, tokenModel.AccessToken, AccessTokenExpire)
		return tokenModel.AccessToken, nil
	}
	return "", err
}

func GetSendMsgHandler(c *gin.Context) {
	c.String(http.StatusOK, "abc")
}

func PostSendMsgHandler(c *gin.Context) {
	var text TextReqMSG
	err := c.ShouldBind(&text)
	if err != nil {
		fmt.Println("参数错误")
		c.String(http.StatusOK, "")
		return
	}
	fmt.Println("请求体：", text)

	_, err = dal.Redis.Get(c, text.MsgId).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("消息已经处理")
		c.String(http.StatusOK, "")
		return
	}

	dal.Redis.Set(c, text.MsgId, "1", 60*time.Second)

	lastID, err := dal.Redis.Get(c, text.FromUserName).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("获取最后一个id失败")
		c.String(http.StatusOK, "")
		return
	}
	fmt.Println("最后一个id：", lastID)
	ChatResult := RequestChat(text.Content, lastID)

	array := strings.Split(ChatResult, "\n")
	lastElem := array[len(array)-1]
	var chat ChatMsg
	err = json.Unmarshal([]byte(lastElem), &chat)
	if err != nil {
		fmt.Println("gpt返回数据序列化失败")
		c.String(http.StatusOK, "")
		return
	}
	fmt.Println("本次id", chat.ID)
	fmt.Println("父id", chat.ParentMessageId)

	dal.Redis.Set(c, text.FromUserName, chat.ID, UserMsgExpire)
	var respMsg TextResMSG
	respMsg.MsgType = CDATA{"text"}
	respMsg.Content = CDATA{chat.Text}
	respMsg.FromUserName = CDATA{text.ToUserName}
	respMsg.ToUserName = CDATA{text.FromUserName}
	respMsg.CreateTime = time.Now().Unix()

	fmt.Println("返回数据", respMsg)
	c.XML(http.StatusOK, respMsg)
}
