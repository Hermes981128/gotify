package gotify

import "time"

type Application struct {
	DefaultPriority int       `json:"defaultPriority,omitempty" comment:"此应用程序发送的消息的默认优先级。默认为0。"`
	Description     string    `json:"description" comment:"应用程序的描述。"`
	Id              int       `json:"id"`
	Image           string    `json:"image"`
	Internal        bool      `json:"internal" comment:"应用程序是否为内部应用程序。不应删除内部应用程序。"`
	LastUsed        time.Time `json:"lastUsed" comment:"最后一次使用应用程序令牌的时间。"`
	Name            string    `json:"name" comment:"应用程序名称。这就是应用程序应该如何显示给用户的方式。"`
	Token           string    `json:"token" comment:"应用程序令牌。"`
}

type Message struct {
	Extras   map[string]map[string]any `json:"extras,omitempty" comment:"保存有关由应用程序发送的消息的信息。"`
	Message  string                    `json:"message" comment:"消息内容。"`
	Priority int                       `json:"priority,omitempty" comment:"消息优先级。默认为0。"`
	Title    string                    `json:"title,omitempty" comment:"消息标题。"`
}

type MessageExternal struct {
	Appid    int                       `json:"appid" comment:"应用程序ID。"`
	Date     time.Time                 `json:"date" comment:"消息发送的时间。"`
	Extras   map[string]map[string]any `json:"extras" comment:"保存有关由应用程序发送的消息的信息。"`
	Id       int                       `json:"id" comment:"消息ID。"`
	Message  string                    `json:"message" comment:"消息内容。"`
	Priority int                       `json:"priority,omitempty" comment:"消息优先级。默认为0。"`
	Title    string                    `json:"title" comment:"消息标题。"`
}
