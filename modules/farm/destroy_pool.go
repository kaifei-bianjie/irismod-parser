package farm

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type DocTxMsgDestroyPool struct {
	PoolId  string `bson:"pool_id"`
	Creator string `bson:"creator"`
}

func (m *DocTxMsgDestroyPool) GetType() string {
	return MsgTypeDestroyPool
}

func (m *DocTxMsgDestroyPool) BuildMsg(v interface{}) {
	msg := v.(*MsgDestroyPool)
	m.PoolId = msg.PoolId
	m.Creator = msg.Creator
}

func (m *DocTxMsgDestroyPool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDestroyPool)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
