package nft

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
	"strings"
)

type DocMsgTransferDenom struct {
	Id        string `bson:"id"`
	Sender    string `bson:"sender"`
	Recipient string `bson:"recipient"`
}

func (m *DocMsgTransferDenom) GetType() string {
	return MsgTypeTransferDenom
}

func (m *DocMsgTransferDenom) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferDenom)

	m.Sender = msg.Sender
	m.Recipient = msg.Recipient
	m.Id = strings.ToLower(msg.Id)
}

func (m *DocMsgTransferDenom) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransferDenom)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
