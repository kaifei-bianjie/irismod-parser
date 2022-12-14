package service

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type (
	DocMsgDefineService struct {
		Name              string   `bson:"name" yaml:"name"`
		Description       string   `bson:"description" yaml:"description"`
		Tags              []string `bson:"tags" yaml:"tags"`
		Author            string   `bson:"author" yaml:"author"`
		AuthorDescription string   `bson:"author_description" yaml:"author_description"`
		Schemas           string   `bson:"schemas"`
	}
)

func (m *DocMsgDefineService) GetType() string {
	return MsgTypeDefineService
}

func (m *DocMsgDefineService) BuildMsg(v interface{}) {
	msg := v.(*MsgDefineService)

	m.Name = msg.Name
	m.Description = msg.Description
	m.Tags = msg.Tags
	m.Author = msg.Author
	m.AuthorDescription = msg.AuthorDescription
	m.Schemas = msg.Schemas
}

func (m *DocMsgDefineService) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDefineService)
	addrs = append(addrs, msg.Author)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
