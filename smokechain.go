package smokechain

import (
	"time"
)

// BlockData defines the structure of data in each block in the block-chain
// Multiple Info are added to each block.
// All the transactions that take place in one minute is stored in a single block.
type BlockData struct {
	PreviousHash string        `json:"previous"`
	Nonce        int64         `json:"nonce"`
	Info         []Information `json:"info"`
}

// Information interface type for the data to be stored in each block.
// EachInformation stored in a block has an id
// using this id user can get the info from the block
type Information interface{}

// Block defines the structure for the block.
// It contains the metadata for each block
type Block struct {
	Hash    string
	Data    BlockData
	Created time.Time
}

func WriteBlockFile(){
	// TODO: Create the function to write the block data on a file so we can extract the data later
}

func GenerateHash(){
	// TODO: Create the function to generate the hash of each function
}

func BlockValid(){
	// TODO: Create the function tot check the validity of the block
}

func CreateBlock(){
	// TODO: Create the new block
}

func ReadBlock(){
	// TODO: Create the function to Read the info from a single block
}

func ReadInfo(){
	// TODO: Create the function to read an info from the single block
}