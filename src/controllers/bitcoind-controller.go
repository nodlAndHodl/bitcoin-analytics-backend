package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dbconfig "github.com/nodlandhodl/bitcoin-analytics-backend/src/db-config"
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
var dbConnect *gorm.DB = dbconfig.ConnectDB()

// @Summary      Get Blockhash by block height
// @Description  get blockhash by block height
// @Tags         bitcoind
// @Accept       json
// @Produce      json
// @Param        blockHeight path int true "Block Height"
// @Success      200 {string} hash
// @Router       /blockhash/{blockHeight} [get]
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

// @Summary      Get Block by block hash
// @Description  get block by block hash
// @Tags         bitcoind
// @Accept       json
// @Produce      json
// @Param        blockHash path string true "Block Hash"
// @Success      200 {object} models.Block
// @Router       /block/{blockHash} [get]
func (bc *BitcoindController) GetBlock(
	context *gin.Context) {
	// Extract the block hash parameter from the request
	blockHash := context.Param("blockHash")

	// Call the GetBlock function with the extracted block hash
	block, err := bc.bitcoindService.GetBlock(blockHash)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Creating http response
	context.JSON(http.StatusOK, block)
}
