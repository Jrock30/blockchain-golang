package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type Block struct {
	Data     string
	Hash     string
	PrevHash string // 이전의 hash
}

type blockchain struct {
	blocks []*Block // block 의 slice 가 아닌 pointer 들의 slice (복사 X)
}

var b *blockchain   // SingleTon
var one sync.Once	// sync package

/* receiver function (Method)
	Calculate Hash
 */
func (b *Block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))  // sha256 Byte Code Hash 생성 (data + prevHash)
	b.Hash = fmt.Sprintf("%x", hash) 		    // 바이트 코드를 16진수 String 으로 변환
}

// Get Last Block Hash
func getLastHash() string {
	totalBlocks := len(GetBlockchain().blocks)
	if totalBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totalBlocks-1].Hash
}

// Create New Block
func createBlock(data string) *Block {
	newBlock := Block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// export function
func (b *blockchain) AddBlock(data string)  {
	b.blocks = append(b.blocks, createBlock(data))
}

// GetBlockchain SingleTon Pattern (소문자를 사용해서 밖에서 접근할 수 없도록 한다)
func GetBlockchain() *blockchain {
	if b == nil {
		// go routine, thread 몇 개든 병렬적으로 실행해도 한번만 실행된다.
		one.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis Block")
		})
	}
	return b
}

func (b *blockchain) AllBLocks() []*Block {
//func AllBlocks() []*block {
//	return GetBlockchain().blocks
	return b.blocks
}