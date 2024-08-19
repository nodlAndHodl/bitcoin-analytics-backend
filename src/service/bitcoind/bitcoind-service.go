package bitcoind

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type Config struct {
	RPCUser     string
	RPCPassword string
	RPCHost     string
	RPCPort     string
}

type BitcoindService struct {
	config Config
}

func NewBitcoindService() (*BitcoindService, error) {
	var config = Config{
		RPCUser:     os.Getenv("NODE_USERNAME"),
		RPCPassword: os.Getenv("NODE_PASSWORD"),
		RPCHost:     os.Getenv("NODE_RPC_HOST"),
		RPCPort:     os.Getenv("NODE_RPC_PORT"),
	}

	if config.RPCUser == "" || config.RPCPassword == "" || config.RPCHost == "" || config.RPCPort == "" {
		return nil, fmt.Errorf("missing required environment variables")
	}
	return &BitcoindService{config: config}, nil
}

func (s *BitcoindService) GetAddressFromScriptPubKey(scriptPubKey ScriptPubKey) string {
	var address string

	if len(scriptPubKey.Address) > 0 {
		address = string(scriptPubKey.Address[0])
	} else if scriptPubKey.Asm != "" {
		parts := split(scriptPubKey.Asm, " ")
		if len(parts) >= 2 && parts[1] == "OP_CHECKSIG" {
			address = parts[0]
		} else {
			for _, part := range parts {
				if len(part) == 40 {
					address = s.convertHashToAddress(part)
					break
				}
			}
		}
	}

	return address
}

func (s *BitcoindService) GetBlock(blockHash string) (*Block, error) {
	result, err := s.makeRpcCall("getblock", []interface{}{blockHash})
	if err != nil {
		return nil, err
	}
	fmt.Printf("Result: %v\n", result)

	// Convert result to JSON and then unmarshal into Block struct
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	var block Block
	if err := json.Unmarshal(resultBytes, &block); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result to Block: %w", err)
	}

	return &block, nil
}

func (s *BitcoindService) GetBlockHash(index int) (string, error) {
	result, err := s.makeRpcCall("getblockhash", []interface{}{index})
	if err != nil {
		return "", err
	}
	hash, ok := result.(string)
	if !ok {
		return "", fmt.Errorf("failed to convert result to string")
	}
	return hash, nil
}

func (s *BitcoindService) GetBlockCount() (int, error) {
	result, err := s.makeRpcCall("getblockcount", nil)
	if err != nil {
		return 0, err
	}
	count, ok := result.(int)
	if !ok {
		return 0, fmt.Errorf("failed to convert result to int")
	}
	return count, nil
}

func (s *BitcoindService) GetBlockStats(hashOrHeight interface{}) (*BlockStats, error) {
	result, err := s.makeRpcCall("getblockstats", []interface{}{hashOrHeight})
	if err != nil {
		return nil, err
	}
	stats, ok := result.(*BlockStats)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to *BlockStats")
	}
	return stats, nil
}

func (s *BitcoindService) GetBlockchainInfo() (*BlockchainInfo, error) {
	result, err := s.makeRpcCall("getblockchaininfo", nil)
	if err != nil {
		return nil, err
	}
	info, ok := result.(*BlockchainInfo)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to *BlockchainInfo")
	}
	return info, nil
}

func (s *BitcoindService) GetDifficulty() (*Difficulty, error) {
	result, err := s.makeRpcCall("getdifficulty", nil)
	if err != nil {
		return nil, err
	}
	difficulty, ok := result.(*Difficulty)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to *Difficulty")
	}
	return difficulty, nil
}

func (s *BitcoindService) GetMempoolInfo() (*MempoolInfo, error) {
	result, err := s.makeRpcCall("getmempoolinfo", nil)
	if err != nil {
		return nil, err
	}
	info, ok := result.(*MempoolInfo)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to *MempoolInfo")
	}
	return info, nil
}

func (s *BitcoindService) GetChainTxStats(nblocks int, blockhash *string) (*ChainTxStats, error) {
	params := []interface{}{nblocks}
	if blockhash != nil {
		params = append(params, *blockhash)
	}
	result, err := s.makeRpcCall("getchaintxstats", params)
	if err != nil {
		return nil, err
	}
	stats, ok := result.(*ChainTxStats)
	if !ok {
		return nil, fmt.Errorf("failed to convert result to *ChainTxStats")
	}
	return stats, nil
}

func (s *BitcoindService) GetRawTransaction(txid string, verbose bool) (*Transaction, error) {
	result, err := s.makeRpcCall("getrawtransaction", []interface{}{txid, verbose})
	if err != nil {
		return nil, err
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	var transaction Transaction
	if err := json.Unmarshal(resultBytes, &transaction); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result to Block: %w", err)
	}

	return &transaction, nil
}

func (s *BitcoindService) makeRpcCall(method string, params []interface{}) (interface{}, error) {
	url := fmt.Sprintf("http://%s:%s@%s:%s/", s.config.RPCUser, s.config.RPCPassword, s.config.RPCHost, s.config.RPCPort)

	payload := map[string]interface{}{
		"jsonrpc": "1.0",
		"id":      method,
		"method":  method,
		"params":  params,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	defer client.CloseIdleConnections()
	for attempt := 1; attempt <= 3; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			if attempt < 3 {
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, err
		}

		if result["error"] != nil {
			if attempt < 3 {
				time.Sleep(1 * time.Second)
				continue
			}
			return nil, fmt.Errorf("RPC Error: %v", result["error"])
		}

		return result["result"], nil
	}

	return nil, fmt.Errorf("failed to make RPC call after 3 attempts")
}

func (s *BitcoindService) convertHashToAddress(hash string) string {
	// Implement the conversion logic here
	return ""
}

func split(s, sep string) []string {
	// Implement the split function here
	return []string{}
}
