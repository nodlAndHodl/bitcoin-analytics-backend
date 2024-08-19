package bitcoind

type RpcResponse[T any] struct {
	Result T         `json:"result"`
	Error  *RpcError `json:"error"`
	ID     string    `json:"id"`
}

type RpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BlockchainInfo struct {
	Chain                string        `json:"chain"`
	Blocks               int           `json:"blocks"`
	Headers              int           `json:"headers"`
	BestBlockHash        string        `json:"bestblockhash"`
	Difficulty           float64       `json:"difficulty"`
	Mediantime           int           `json:"mediantime"`
	VerificationProgress float64       `json:"verificationprogress"`
	InitialBlockDownload bool          `json:"initialblockdownload"`
	Chainwork            string        `json:"chainwork"`
	SizeOnDisk           int64         `json:"size_on_disk"`
	Pruned               bool          `json:"pruned"`
	Softforks            []interface{} `json:"softforks"` // This can be detailed further based on your needs
	Warnings             string        `json:"warnings"`
}

type BlockStats struct {
	AvgFee     int    `json:"avgfee"`
	AvgFeeRate int    `json:"avgfeerate"`
	AvgTxSize  int    `json:"avgtxsize"`
	BlockHash  string `json:"blockhash"`
	Height     int    `json:"height"`
	Ins        int    `json:"ins"`
	MaxFee     int    `json:"maxfee"`
	MaxFeeRate int    `json:"maxfeerate"`
	MaxTxSize  int    `json:"maxtxsize"`
	MedianFee  int    `json:"medianfee"`
	Mediantime int    `json:"mediantime"`
	MinFee     int    `json:"minfee"`
	Outs       int    `json:"outs"`
	Subsidy    int    `json:"subsidy"`
	Time       int    `json:"time"`
	TotalFee   int    `json:"totalfee"`
	Txs        int    `json:"txs"`
	// Add other fields as needed
}

type ChainTxStats struct {
	Time                 int     `json:"time"`
	TxCount              int     `json:"txcount"`
	WindowBlockCount     int     `json:"window_block_count"`
	WindowFinalBlockHash string  `json:"window_final_block_hash"`
	WindowTxCount        int     `json:"window_tx_count"`
	WindowInterval       int     `json:"window_interval"`
	TxRate               float64 `json:"txrate"`
}

type MempoolInfo struct {
	Size          int     `json:"size"`
	Bytes         int     `json:"bytes"`
	Usage         int     `json:"usage"`
	MaxMempool    int     `json:"maxmempool"`
	MempoolMinFee float64 `json:"mempoolminfee"`
	MinRelayTxFee float64 `json:"minrelaytxfee"`
}

type Block struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	StrippedSize      int      `json:"strippedsize"`
	Size              int      `json:"size"`
	Weight            int      `json:"weight"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	VersionHex        string   `json:"versionHex"`
	MerkleRoot        string   `json:"merkleroot"`
	Tx                []string `json:"tx"`
	Time              int      `json:"time"`
	MedianTime        int      `json:"mediantime"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	Chainwork         string   `json:"chainwork"`
	NTx               int      `json:"nTx"`
	PreviousBlockHash string   `json:"previousblockhash"`
	NextBlockHash     string   `json:"nextblockhash"`
}

type Difficulty struct {
	Data float64 `json:"data"`
}

type BlockCount struct {
	Data int `json:"data"`
}

type TransactionInput struct {
	Txid      string `json:"txid"`
	Vout      int    `json:"vout"`
	ScriptSig struct {
		Asm string `json:"asm"`
		Hex string `json:"hex"`
	} `json:"scriptSig"`
	Sequence int    `json:"sequence"`
	Coinbase string `json:"coinbase,omitempty"`
}

type TransactionOutput struct {
	Value        float64 `json:"value"`
	N            int     `json:"n"`
	ScriptPubKey struct {
		Asm     string `json:"asm"`
		Hex     string `json:"hex"`
		Type    string `json:"type"`
		Desc    string `json:"desc,omitempty"`
		Address string `json:"address,omitempty"`
	} `json:"scriptPubKey"`
}

type TransactionDetail struct {
	InvolvesWatchonly bool    `json:"involvesWatchonly,omitempty"`
	Address           string  `json:"address"`
	Category          string  `json:"category"`
	Amount            float64 `json:"amount"`
	Label             string  `json:"label,omitempty"`
	Vout              int     `json:"vout"`
	Fee               float64 `json:"fee,omitempty"`
	Abandoned         bool    `json:"abandoned,omitempty"`
}

type Transaction struct {
	Txid              string              `json:"txid"`
	Hash              string              `json:"hash"`
	Version           int                 `json:"version"`
	Size              int                 `json:"size"`
	Vsize             int                 `json:"vsize"`
	Weight            int                 `json:"weight"`
	Locktime          int                 `json:"locktime"`
	Vin               []TransactionInput  `json:"vin"`
	Vout              []TransactionOutput `json:"vout"`
	Hex               string              `json:"hex"`
	Blockhash         string              `json:"blockhash"`
	Confirmations     int                 `json:"confirmations"`
	Time              int                 `json:"time"`
	Blocktime         int                 `json:"blocktime"`
	WalletConflicts   []string            `json:"walletconflicts"`
	Comment           string              `json:"comment,omitempty"`
	Bip125Replaceable string              `json:"bip125_replaceable"`
	Details           []TransactionDetail `json:"details"`
}

type ScriptPubKey struct {
	Asm     string `json:"asm"`
	Hex     string `json:"hex"`
	Desc    string `json:"desc,omitempty"`
	Type    string `json:"type"`
	Address string `json:"address,omitempty"`
}
