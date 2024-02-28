package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Result struct {
		Number           string        `json:"number"`
		Hash             string        `json:"hash"`
		ParentHash       string        `json:"parentHash"`
		Nonce            string        `json:"nonce"`
		Sha3Uncles       string        `json:"sha3Uncles"`
		LogsBloom        string        `json:"logsBloom"`
		TransactionsRoot string        `json:"transactionsRoot"`
		StateRoot        string        `json:"stateRoot"`
		ReceiptsRoot     string        `json:"receiptsRoot"`
		Miner            string        `json:"miner"`
		Difficulty       string        `json:"difficulty"`
		TotalDifficulty  string        `json:"totalDifficulty"`
		ExtraData        string        `json:"extraData"`
		Size             string        `json:"size"`
		GasLimit         string        `json:"gasLimit"`
		GasUsed          string        `json:"gasUsed"`
		Timestamp        string        `json:"timestamp"`
		Transactions     []Transaction `json:"transactions"`
		Uncles           []interface{} `json:"uncles"`
	} `json:"result"`
	ID int `json:"id"`
}

func main() {
	rpcEndpoint := "https://rpc.holesky.redstone.xyz"

	// Block number to query
	blockNumber := "0x1" 

	requestBody := `{"jsonrpc":"2.0","method":"eth_getBlockByNumber","params":["` + blockNumber + `",true],"id":1}`

	resp, err := http.Post(rpcEndpoint, "application/json", strings.NewReader(requestBody))
	if err != nil {
		log.Fatal("Error sending request to Ethereum RPC endpoint:", err)
	}
	defer resp.Body.Close()

	var blockResponse BlockResponse
	if err := json.NewDecoder(resp.Body).Decode(&blockResponse); err != nil {
		log.Fatal("Error decoding JSON response:", err)
	}

	// Log block data
	log.Println("Block Number:", blockResponse.Result.Number)
	log.Println("Block Hash:", blockResponse.Result.Hash)
	log.Println("Timestamp:", blockResponse.Result.Timestamp)

	// get transaction data
	for _, tx := range blockResponse.Result.Transactions {
		log.Println("Transaction Hash:", tx.Hash)
		log.Println("From:", tx.From)
		log.Println("To:", tx.To)
		log.Println("Value:", tx.Value)
		log.Println("Gas Price:", tx.GasPrice)
		// log.Println("Input:", tx.Input)
		log.Println("Nonce:", tx.Nonce)
		log.Println("Transaction Index:", tx.TransactionIndex)
		log.Println("V:", tx.V)
		log.Println("R:", tx.R)
		log.Println("S:", tx.S)
	}

}
