package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string // 이전의 hash

	// one-way function (단 방향으로만 실행할 수 있다, 결정론적이다.(입력값은 항상 같은 출력값을 얻게 된다.))
}

type blockchain struct {
	blocks []*block // block 의 slice 가 아닌 pointer 들의 slice (복사 X)
}

var b *blockchain // SingleTon
var one sync.Once // sync package

/* receiver function (Method)
Calculate Hash
*/
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash)) // sha256 Byte Code Hash 생성 (data + prevHash)
	b.Hash = fmt.Sprintf("%x", hash)                   // 바이트 코드를 16진수 String 으로 변환
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
func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// export function
func (b *blockchain) AddBlock(data string) {
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

func (b *blockchain) AllBLocks() []*block {
	//func AllBlocks() []*block {
	//	return GetBlockchain().blocks
	return b.blocks
}
