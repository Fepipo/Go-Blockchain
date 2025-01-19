package miner

import (
  "blockchain/models"
  "crypto/sha256"
  "encoding/hex"
  "strconv"
  "strings"
  "time"
)

func Miner(block models.Block, dificuldade int) (int, string, time.Duration) {

  var hash [32]byte
  var string_hash string
  var duration time.Duration

  info_block := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PreviousHash

  quant_zeros	:= strings.Repeat("0", dificuldade)
	
  loop := true
  nonce := 0
	start := time.Now()

	for loop {
		info_block = info_block + strconv.Itoa(nonce)
		hash = sha256.Sum256([]byte(info_block))

		string_hash = hex.EncodeToString(hash[:])
		
		if strings.HasPrefix(string_hash, quant_zeros) {
			duration = time.Since(start)
			break
			} else {
				nonce += 1
			}
	}

	return nonce, string_hash, duration
}
