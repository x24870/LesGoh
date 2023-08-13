package chain

import (
	"context"
	"net/http"

	"github.com/coming-chat/go-sui/account"
	"github.com/coming-chat/go-sui/client"
	"github.com/coming-chat/go-sui/types"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	rpcUrl = "https://fullnode.devnet.sui.io:443"
	// mnemonic = "tag couch antenna aerobic dignity boost equal unfold peanut humble seek ten"
	// pkgId    = "0x37a7b60b7e8c4ae9799b0b1e748f98aa03a818bb"
	// gasId    = "0x19d4f97e22da92811478b8c9a69fc920a190aa8b"
	// maxGas   = 30000
	// seed     = "AH+Jiv+xw/G5t/8wmBsAv+bbTIez6qeqqA/0xETqQsGz55Uow52GZi58ExLkhgYJnt7vZKZ45Qf5brj7HmRlqjQ="
)

func RegisterAPI(r *gin.Engine, apt *Adapter) {
	hdl := &hander{
		apt: apt,
	}

	r.POST("/notif", hdl.notif)
}

type notifRequest struct {
	TransactionHash string `json:"transactionHash"`
	Address         string `json:"address"`
}

type notifResponse struct {
	Message string `json:"message"`
}

type hander struct {
	apt *Adapter
}

func (h *hander) notif(c *gin.Context) {
	req := notifRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.WithField("err", err).Error("c.ShouldBindJSON error")
		c.JSON(http.StatusBadRequest, notifResponse{Message: "invalid request"})
		return
	}

	if req.TransactionHash == "" || req.Address == "" {
		logrus.WithField("req", req).Error("empty request")
		c.JSON(http.StatusBadRequest, notifResponse{Message: "empty request"})
		return
	}

	if err := h.apt.Save(req.TransactionHash, req.Address); err != nil {
		logrus.WithFields(logrus.Fields{
			"hash": req.TransactionHash,
			"err":  err,
		}).Error("adapter.Save error")
		c.JSON(http.StatusBadRequest, notifResponse{Message: err.Error()})
	}

	c.JSON(http.StatusOK, notifResponse{Message: "ok"})
}

func Exec() {
	acc, err := account.NewAccountWithMnemonic(mnemonic)
	if err != nil {
		logrus.Error("account.NewAccountWithMnemonic err: ", err)
		return
	}
	// acc := account.NewAccount([]byte(seed))

	signer, _ := types.NewAddressFromHex(acc.Address)

	ctx := context.Background()
	cli, err := client.Dial(rpcUrl)
	if err != nil {
		logrus.Error("client.Dial err: ", err)
		return
	}

	pkgId, err := types.NewHexData(pkgId)
	if err != nil {
		logrus.Error("types.NewHexData err: ", err)
		return
	}
	gasObjId, err := types.NewHexData(gasId)
	if err != nil {
		logrus.Error("types.NewHexData err: ", err)
		return
	}

	args := []interface{}{
		"0x9817db7e507a682b29354f1df62b3d1289ee0da6", 42, 7, "0xb0152567ea827539013c0c08a223125de938021f",
	}
	txnBytes, err := cli.MoveCall(ctx, *signer, *pkgId, "my_module", "sword_create", []string{}, args, gasObjId, uint64(maxGas))
	if err != nil {
		logrus.Error("MoveCall err: ", err)
		return
	}
	logrus.Info("txnBytes = ", *txnBytes)

	// recipient, err := types.NewAddressFromHex("0xa80da231c6388f7e6ab7641c3adb1fd3987887bd")
	// if err != nil {
	// 	logrus.Error("types.NewAddressFromHex err: ", err)
	// 	return
	// }
	// suiObjectId, err := types.NewHexData("0x36d3176a796e167ffcbd823c94718e7db56b955f")
	// if err != nil {
	// 	logrus.Error("types.NewHexData err: ", err)
	// 	return
	// }
	// transferAmount := uint64(10000)

	// txnBytes, err := cli.TransferSui(ctx, *signer, *recipient, *suiObjectId, transferAmount, uint64(maxGas))
	// if err != nil {
	// 	logrus.Error("client.Dial err: ", err)
	// 	return
	// }
	// logrus.Info("txnBytes = ", *txnBytes)

	// Sign
	signedTxn := txnBytes.SignWith(acc.PrivateKey)
	txnResponse, err := cli.ExecuteTransaction(ctx, *signedTxn, types.TxnRequestTypeWaitForLocalExecution)
	if err != nil {
		logrus.Error("ExecuteTransaction err: ", err)
		return
	}
	logrus.Info("transaction resp = ", *txnResponse)
	logrus.Info("transaction resp digest = ", txnResponse.TransactionDigest())

}
func Do() {
	cli, err := client.Dial(rpcUrl)
	if err != nil {
		logrus.Error(err)
		return
	}
	ctx := context.Background()

	// call get transaction
	digest, err := types.NewBase64Data("KevFVGiLaLbSQh2cq0SvBYJvAH0OUl4/+847biyspkg=")
	if err != nil {
		logrus.Error("types.NewBase64Data err: ", err)
		return
	}
	resp := types.TransactionResponse{}
	err = cli.CallContext(ctx, &resp, "sui_getTransaction", digest)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info("transaction resp = ", resp)

	// // And you can call some predefined methods
	// digest, err := types.NewBase64Data("/KXvTwNRHKKzAB+/Dz1O64LjVbISgIW4VUCmuuPyEfU=")
	// resp, err := cli.GetTransaction(ctx, digest)
	// print("transaction status = ", resp.Effects.Status)
	// print("transaction timestamp = ", resp.TimestampMs)
}
