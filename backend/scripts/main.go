package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/x24870/LesGoh/backend/contracts/accountrouter"
	"github.com/x24870/LesGoh/backend/contracts/mailbox"
	"github.com/x24870/LesGoh/backend/contracts/nft"
	"github.com/x24870/LesGoh/backend/contracts/paymaster"
)

const (
	// testnet
	endpoint = "https://ethereum-goerli.publicnode.com" // goerli
	// mainnet
	// endpoint = ""

	pk = "3f65e43afa4731c92c9570728ef74836ab2fb40ce03adb69ad3e214498feb7a0"
)

var (
	client *ethclient.Client
	ctx    = context.Background()

	chainID *big.Int

	goerliAccountRouter = common.HexToAddress("0x55486284a85d7b51a7bBfd343702414D65276fa6")
	goerliPaymaster     = common.HexToAddress("0xF90cB82a76492614D07B82a7658917f3aC811Ac1")
	goerliMailBox       = common.HexToAddress("0xCC737a94FecaeC165AbCf12dED095BB13F037685")

	accountRouterABI  abi.ABI
	mailboxABI        abi.ABI
	nftABI            abi.ABI
	dispatchIdEventID common.Hash
	gasPaymaster      *paymaster.Paymaster

	mumbaiDest     uint32 = 80001
	mumbaiReceiver        = common.HexToAddress("0x36FdA966CfffF8a9Cdc814f546db0e6378bFef35")
	mumbaiNFT             = common.HexToAddress("0xdf67E4A6B7953782987f5684635e93f24a211D74")
)

func init() {
	var err error
	client, err = ethclient.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	chainID, err = client.NetworkID(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("chain id: ", chainID)

	accountRouterABI, err = abi.JSON(bytes.NewReader([]byte(accountrouter.AccountrouterABI)))
	if err != nil {
		panic(err)
	}
	mailboxABI, err = abi.JSON(bytes.NewReader([]byte(mailbox.MailboxABI)))
	if err != nil {
		panic(err)
	}
	nftABI, err = abi.JSON(bytes.NewReader([]byte(nft.NftABI)))
	if err != nil {
		panic(err)
	}
	dispatchIdEvent, ok := mailboxABI.Events["DispatchId"]
	if !ok {
		panic(errors.New("event DispatchId not found"))
	}
	dispatchIdEventID = dispatchIdEvent.ID
	gasPaymaster, err = paymaster.NewPaymaster(
		goerliPaymaster,
		client,
	)
	if err != nil {
		panic(err)
	}
}

func getAddress(ctx context.Context) (common.Address, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return common.Address{}, err
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("failed to get public key")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}

func sendAccountAPI(ctx context.Context, user common.Address, dryrun bool) (common.Hash, error) {
	// # sendData: "fooBar(uint256,string)" 666 "DAT ASS"
	sendData := hexutil.MustDecode("0xf07c1f47000000000000000000000000000000000000000000000000000000000000029a000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000074441542041535300000000000000000000000000000000000000000000000000")
	callData, err := accountRouterABI.Pack(
		"callRemote",
		mumbaiDest,
		mumbaiReceiver,
		big.NewInt(0),
		sendData,
	)
	if err != nil {
		return common.Hash{}, err
	}
	estGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  user,
		To:    &goerliAccountRouter,
		Value: nil,
		Data:  callData,
	})
	if err != nil {
		rpcErr, ok := err.(rpc.Error)
		if !ok {
			fmt.Printf("unexpected estimate error: %+v\n", err)
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("estimate error: %+v\n", err)
		fmt.Printf("error message: %s\n", rpcErr.Error())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())

		rpcDataErr, ok := err.(rpc.DataError)
		if !ok {
			fmt.Println("not support data")
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("data error: %+v\n", rpcDataErr)
		fmt.Printf("error message: %s\n", rpcDataErr.Error())
		fmt.Printf("error data: %+v\n", rpcDataErr.ErrorData())
		return common.Hash{}, fmt.Errorf("estimate gas error: %v", err)
	}

	nonce, err := client.PendingNonceAt(ctx, user)
	if err != nil {
		return common.Hash{}, err
	}
	price, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	fmt.Println("nonce: ", nonce)
	fmt.Println("estRes: ", estGas)
	fmt.Println("price: ", price)
	fmt.Println("tip: ", tip)

	if dryrun {
		return common.Hash{}, errors.New("dryrun only")
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: price,
		GasTipCap: tip,
		Gas:       estGas,
		To:        &goerliAccountRouter,
		Value:     nil,
		Data:      callData,
	})
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err = types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		rpcErr, ok := err.(rpc.Error)
		if !ok {
			fmt.Printf("unexpected send error: %+v\n", err)
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("send error: %+v\n", err)
		fmt.Printf("error message: %s\n", rpcErr.Error())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("compare error: %v\n", err == core.ErrNonceTooLow)
		fmt.Printf("compare error: %v\n", errors.Is(err, core.ErrNonceTooLow))
		fmt.Printf("compare error: %v\n", errors.Is(rpcErr, core.ErrNonceTooLow))

		rpcDataErr, ok := err.(rpc.DataError)
		if !ok {
			fmt.Println("not support data")
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("data error: %+v\n", rpcDataErr)
		fmt.Printf("error message: %s\n", rpcDataErr.Error())
		fmt.Printf("error data: %+v\n", rpcDataErr.ErrorData())

		return common.Hash{}, err
	}

	return txHash, nil
}

func mintNTF(ctx context.Context, user common.Address, target string, dryrun bool) (common.Hash, error) {
	mintData, err := nftABI.Pack(
		"mint",
		common.HexToAddress(target),
		// common.HexToAddress("0x9a702c347b4C0e5AFF041339f598253fab22C03C"), // me
		// common.HexToAddress("0x9FCb30E2003801C67627fB1720895cb023447d84"), // d
	)
	if err != nil {
		return common.Hash{}, err
	}

	callData, err := accountRouterABI.Pack(
		"callRemote",
		mumbaiDest,
		mumbaiNFT,
		big.NewInt(0),
		mintData,
	)
	if err != nil {
		return common.Hash{}, err
	}
	estGas, err := client.EstimateGas(ctx, ethereum.CallMsg{
		From:  user,
		To:    &goerliAccountRouter,
		Value: nil,
		Data:  callData,
	})
	if err != nil {
		rpcErr, ok := err.(rpc.Error)
		if !ok {
			fmt.Printf("unexpected estimate error: %+v\n", err)
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("estimate error: %+v\n", err)
		fmt.Printf("error message: %s\n", rpcErr.Error())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())

		rpcDataErr, ok := err.(rpc.DataError)
		if !ok {
			fmt.Println("not support data")
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("data error: %+v\n", rpcDataErr)
		fmt.Printf("error message: %s\n", rpcDataErr.Error())
		fmt.Printf("error data: %+v\n", rpcDataErr.ErrorData())
		return common.Hash{}, fmt.Errorf("estimate gas error: %v", err)
	}

	nonce, err := client.PendingNonceAt(ctx, user)
	if err != nil {
		return common.Hash{}, err
	}
	price, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	fmt.Println("nonce: ", nonce)
	fmt.Println("estRes: ", estGas)
	fmt.Println("price: ", price)
	fmt.Println("tip: ", tip)

	if dryrun {
		return common.Hash{}, errors.New("dryrun only")
	}

	tx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: price,
		GasTipCap: tip,
		Gas:       estGas,
		To:        &goerliAccountRouter,
		Value:     nil,
		Data:      callData,
	})
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return common.Hash{}, err
	}
	tx, err = types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		return common.Hash{}, err
	}
	txHash := tx.Hash()
	err = client.SendTransaction(ctx, tx)
	if err != nil {
		rpcErr, ok := err.(rpc.Error)
		if !ok {
			fmt.Printf("unexpected send error: %+v\n", err)
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("send error: %+v\n", err)
		fmt.Printf("error message: %s\n", rpcErr.Error())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("error code: %d\n", rpcErr.ErrorCode())
		fmt.Printf("compare error: %v\n", err == core.ErrNonceTooLow)
		fmt.Printf("compare error: %v\n", errors.Is(err, core.ErrNonceTooLow))
		fmt.Printf("compare error: %v\n", errors.Is(rpcErr, core.ErrNonceTooLow))

		rpcDataErr, ok := err.(rpc.DataError)
		if !ok {
			fmt.Println("not support data")
			return common.Hash{}, err
		}
		// error message: err: insufficient funds for gas * price + value: address 0x0b43BCe5Cfc9Fcc9B25E70e5d7022c2f1D5468ad have 189978999999664000 want 1000000000000000000 (supplied gas 10020264)
		// error code: -32000
		fmt.Printf("data error: %+v\n", rpcDataErr)
		fmt.Printf("error message: %s\n", rpcDataErr.Error())
		fmt.Printf("error data: %+v\n", rpcDataErr.ErrorData())

		return common.Hash{}, err
	}

	return txHash, nil
}

func fundGas(ctx context.Context, user common.Address, txHash common.Hash) (common.Hash, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return common.Hash{}, err
	}

	// get $GAS_PAYMENT_QUOTE
	quote, err := gasPaymaster.QuoteGasPayment(nil, mumbaiDest, big.NewInt(10000))
	if err != nil {
		return common.Hash{}, err
	}
	fmt.Println("quote: ", quote)

	// get message id
	messageID, err := getMessageID(ctx, txHash)
	if err != nil {
		return common.Hash{}, err
	}
	fmt.Println("messageID: ", messageID)

	// call payForGas
	nonce, err := client.PendingNonceAt(ctx, user)
	if err != nil {
		return common.Hash{}, err
	}
	price, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	tip, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		return common.Hash{}, err
	}
	fmt.Println("nonce: ", nonce)
	fmt.Println("price: ", price)
	fmt.Println("tip: ", tip)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return common.Hash{}, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasFeeCap = price
	auth.GasTipCap = tip
	auth.GasLimit = uint64(300000) // in units
	var messageID32Bytes [32]byte
	copy(messageID32Bytes[:], messageID.Bytes())
	tx, err := gasPaymaster.PayForGas(auth, messageID32Bytes, mumbaiDest, quote, user)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

const (
	// the maximum retry count of getMessageID
	maxRetryCount = 100
)

func getMessageID(ctx context.Context, txHash common.Hash) (common.Hash, error) {
	var isPending bool = true
	var count int
	var logs []*types.Log
	for isPending && count < maxRetryCount {
		receipt, err := client.TransactionReceipt(ctx, txHash)
		isPending = (err != nil || receipt == nil || receipt.Status != types.ReceiptStatusSuccessful)
		if isPending {
			fmt.Println("keep polling...")
			time.Sleep(5 * time.Second)
			count += 1
			continue
		}
		fmt.Printf("receipt: %+v\n", receipt)
		logs = receipt.Logs
	}
	if isPending {
		return common.Hash{}, fmt.Errorf("transaction %s timed out", txHash)
	}

	for _, log := range logs {
		if len(log.Topics) == 0 {
			continue
		}
		if bytes.Equal(log.Topics[0].Bytes(), dispatchIdEventID.Bytes()) {
			return log.Topics[1], nil
		}
	}
	return common.Hash{}, fmt.Errorf("transaction %s not found", txHash)
}

func getBurnedEvents(ctx context.Context) error {
	return nil
}

func main() {
	ctx := context.Background()
	address, err := getAddress(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("address: ", address.Hex())

	// txHash, err := sendAccountAPI(ctx, address, false)
	// if err != nil {
	// 	panic(err)
	// }
	// target := "0xF423e1cE4d81d1389E1b0C0925C912644FB0fcd7"
	target := "0x3C127Be7c30e4f0c863047ca566684d66a264893"
	txHash, err := mintNTF(ctx, address, target, false)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash: ", txHash.Hex())

	msgID, err := getMessageID(ctx, txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println("msgID: ", msgID.Hex()) // 0x189810303918bd1c139d6267e596b072a84dfa67fe9071ff914924cd23b773ae

	// txHash := common.HexToHash("0xb88b1bb88c998f2958d0d13738074a62bccc162d1f76bbbc0dd22b2652289d8a")
	txHash2, err := fundGas(ctx, address, txHash)
	if err != nil {
		panic(err)
	}
	fmt.Println("txHash2: ", txHash2.Hex())
}
