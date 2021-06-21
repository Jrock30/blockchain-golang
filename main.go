package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	data     string
	hash     string
	prevHash string // 이전의 hash

	// one-way function (단 방향으로만 실행할 수 있다, 결정론적이다.(입력값은 항상 같은 출력값을 얻게 된다.))
}

type blockchain struct {
	blocks []block
}

// Get Last Block Hash
func (b *blockchain) getLastHash() string {
	// 블록이 기존에 하나라도 존재하면 마지막 블록 Hash 값 가져 온다.
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

// receiver function
func (b *blockchain) addBlock(data string) {
	// block struct 생성
	newBlock := block{data, "", b.getLastHash()}
	// sha256 Byte Code Hash 생성
	hash := sha256.Sum256([]byte(newBlock.data + newBlock.prevHash))
	// 바이트 코드를 16진수 String 으로 변환
	newBlock.hash = fmt.Sprintf("%x", hash)
	// 새로운 블록을 기존 블록에 추가
	b.blocks = append(b.blocks, newBlock)
}

// receiver function
func (b *blockchain) listBlocks() {
	// index 무시
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Prev Hash: %s\n\n", block.prevHash)
	}
}

/*
  Main Package (Entry Point)
*/
func main() {
	chain := blockchain{}
	chain.addBlock("Genesis Block")
	chain.addBlock("Second Block")
	chain.addBlock("Third Block")
	chain.listBlocks()
}
