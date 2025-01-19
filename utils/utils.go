package utils

import (
	"fmt"
  "time"
  "blockchain/models"
)

func Time() string {
	time := time.Now()
	timeF := time.Format("2006-01-02 15:04:05")
	
	return timeF
}

func NewBlock(index int, timeBlock string, previous_hash string) models.Block {
	info_new_block := models.Block{
		Index: index,
		Timestamp: timeBlock,
		Data: fmt.Sprintf("Block %d", index),
		PreviousHash: previous_hash,
	}

	return info_new_block
}
