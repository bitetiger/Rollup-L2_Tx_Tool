package tokentransfer

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func sendToken() {
	// 이더리움 클라이언트 생성
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/YOUR_PROJECT_ID")
	if err != nil {
		log.Fatal(err)
	}

	// 개인 키 생성
	privateKey, err := crypto.HexToECDSA("YOUR_PRIVATE_KEY")
	if err != nil {
		log.Fatal(err)
	}

	// ERC20 토큰 컨트랙트 주소 설정
	tokenAddress := common.HexToAddress("TOKEN_CONTRACT_ADDRESS")

	// 토큰 인스턴스 생성
	tokenInstance, err := NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	// 전송할 수량과 대상 주소 설정
	amount := big.NewInt(1000000000000000000) // 1 ETH
	toAddress := common.HexToAddress("TO_ADDRESS")

	// 토큰 전송 트랜잭션 생성
	data, err := tokenInstance.Transfer(nil, toAddress, amount)
	if err != nil {
		log.Fatal(err)
	}

	// 트랜잭션 필드 설정
	gasLimit := uint64(21000) // 이더 전송과 같은 가스 리밋
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := client.PendingNonceAt(context.Background(), common.HexToAddress("YOUR_ADDRESS"))
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, tokenAddress, big.NewInt(0), gasLimit, gasPrice, data)

	// 트랜잭션 서명
	signedTx, err := types.SignTx(tx, types.HomesteadSigner{}, privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// 트랜잭션 전송
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Transaction sent: %s\n", signedTx.Hash().Hex())
}
