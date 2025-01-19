package main

import (
  "blockchain/miner"
  "blockchain/models"
  "blockchain/utils"
  "fmt"
  "os"
  "strconv"
)

func main() {
  var previous_hash string
  var new_block models.Block

  quant_blocos, err := strconv.Atoi(os.Args[1])
  
  if err != nil {
    fmt.Println("Numero de quantidades de blocos invalido")
    os.Exit(1)
  }


  quant_zeros, err := strconv.Atoi(os.Args[2])

  if err != nil {
    fmt.Println("Numero de dificuldade invalido")
    os.Exit(1)
  } 

  fmt.Println("---BLOCKCHAIN---")

  for i := 0; i < quant_blocos; i ++ {
    timestamp := utils.Time()

    new_block = utils.NewBlock(i, timestamp, previous_hash)

    nonce, hash, duration := miner.Miner(new_block, quant_zeros)

    new_block.Nonce = nonce
    new_block.Hash = hash

    fmt.Printf("\nIndex: %d\nTimestamp: %s\nData: %s\nPreviousHash: %s\nNonce: %d\nHash: %s\nMinered In: %s\n", new_block.Index, new_block.Timestamp, new_block.Data, new_block.PreviousHash, new_block.Nonce, new_block.Hash, duration)

    previous_hash = hash
  }
}
