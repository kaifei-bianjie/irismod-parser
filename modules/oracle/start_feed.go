package oracle

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocMsgStartFeed struct {
	FeedName string `bson:"feed_name" yaml:"feed_name"`
	Creator  string `bson:"creator"`
}

func (m *DocMsgStartFeed) GetType() string {
	return TxTypeStartFeed
}

func (m *DocMsgStartFeed) BuildMsg(v interface{}) {
	msg := v.(*MsgStartFeed)

	m.FeedName = msg.FeedName
	m.Creator = msg.Creator
}

func (m *DocMsgStartFeed) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgStartFeed)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
