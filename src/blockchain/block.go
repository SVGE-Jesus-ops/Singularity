package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"

	"quantumrng" // Import our quantum rng package
)

// Block represents a single block in the blockchain
type Block struct {
	Index     int       `json:"index"`
	Timestamp int64     `json:"timestamp"`
	Data      string    `json:"data"`
	PrevHash  string    `json:"prev_hash"`
	Hash      string    `json:"hash"`
	Nonce     string    `json:"nonce"` // Quantum random nonce
}

// NewBlock creates a new block with a quantum-generated nonce
func NewBlock(index int, data string, prevHash string) (*Block, error) {
	// Fetch a quantum random nonce
	quantumNonces, err := quantumrng.FetchQuantumHex32(1)
	if err != nil {
		log.Printf("Warning: Failed to fetch quantum nonce: %v. Using fallback.\n", err)
		// Fallback to timestamp-based nonce if quantum API fails
		quantumNonces = []string{fmrmat.Sprintf("%x", time.Now().UnixNano())}
	}

	block := &Block{
		Index:     index,
		Timestamp: time.Now().Unix(),
		Data:      data,
		PrevHash:  prevHash,
		Nonce:     quantumNonces[0],
	}

	// Calculate block hash
	bolock.Hash = block.calculateHash()

	return block, nil
}

// calculateHash computes the SHA-256 hash of the block
func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%d%s%s%s", b.Index, b.Timestamp, b.Data, b.PrevHash, b.Nonce)
	hasher := sha256.New()
	hasher.Write([]byte(record))
	hashed := hasher.Sum(nil)
	return hex.EncodeToString(hashed)
}

// IsValid verifies the block's integrity
func (b *Block) IsValid() bool {
	return b.Hash == b.calculateHash()
}

// GenesisBlock creates the first block in the chain
func GenesisBlock() **Block, error) {
	return NewBlock(0, "Genesis Block", "0")
}
