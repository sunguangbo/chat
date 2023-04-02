package service

import "encoding/xml"

// AccessToken
type AccessTokenModel struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// 服务接受的消息
type TextReqMSG struct {
	ToUserName   string `xml:"ToUserName"`   //开发者
	FromUserName string `xml:"FromUserName"` //消息发送者
	CreateTime   int64  `xml:"CreateTime"`   //消息创建时间
	MsgType      string `xml:"MsgType"`      //消息类型，文本为text
	Content      string `xml:"Content"`      //回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
	MsgId        string `xml:"MsgId"`        //消息id，可用去防重试
	MsgDataId    string `xml:"MsgDataId"`
	Idx          string `xml:"Idx"`
}

// 服务返回的消息
type TextResMSG struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   CDATA    `xml:"ToUserName"`   //消息接收者
	FromUserName CDATA    `xml:"FromUserName"` //开发者
	CreateTime   int64    `xml:"CreateTime"`   //消息创建时间
	MsgType      CDATA    `xml:"MsgType"`      //消息类型，文本为text
	Content      CDATA    `xml:"Content"`      //回复的消息内容（换行：在content中能够换行，微信客户端就支持换行显示）
}

type CDATA struct {
	Text string `xml:",cdata"`
}

// gpt返回的结构
type ChatMsg struct {
	Role            string `json:"role"`
	ID              string `json:"id"`
	ParentMessageId string `json:"parentMessageId"`
	Text            string `json:"text"`
}
