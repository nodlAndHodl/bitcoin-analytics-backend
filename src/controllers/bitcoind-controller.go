package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/config"
	"github.com/nodlandhodl/bitcoin-analytics-backend/src/service/bitcoind"
	"gorm.io/gorm"
)

type BitcoindController struct {
	bitcoindService *bitcoind.BitcoindService
}

func NewBitcoindController(bitcoindService *bitcoind.BitcoindService) *BitcoindController {
	return &BitcoindController{
		bitcoindService: bitcoindService,
	}
}

// Define database client
var dbConnect *gorm.DB = config.ConnectDB()

// Define database client

// GetAllTodos lists all existing todos
//
//	@Summary      Get Blockhash by block height
//	@Description  get all todo
//	@Tags         bitcoind
//	@Accept       json
//	@Produce      json
//	@Param        blockHeight path int true "Block Height"
//	@Success      200 {string} hash
//	@Router       /blockhash/{blockHeight} [get]
func (bc *BitcoindController) GetBlockHash(
	context *gin.Context) {
	// Extract the block height parameter from the request
	blockHeightStr := context.Param("blockHeight")

	// Convert the block height parameter to an integer
	blockHeight, err := strconv.Atoi(blockHeightStr)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid block height parameter"})
		return
	}

	// Call the GetBlockHash function with the extracted block height
	hash, err := bc.bitcoindService.GetBlockHash(blockHeight)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, hash)
}
