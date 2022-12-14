package htlc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type HtlcClient struct {
}

func NewClient() HtlcClient {
	return HtlcClient{}
}

func (htlc HtlcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgClaimHTLC:
		docMsg := DocTxMsgClaimHTLC{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgCreateHTLC:
		docMsg := DocTxMsgCreateHTLC{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}

}
