package dem0chain

import (
	"bytes"
	"encoding/gob"

	"github.com/boltdb/bolt"
)

// storage related
func (block Block) Serialize() []byte {
	var result bytes.Buffer            // set buff as result
	encoder := gob.NewEncoder(&result) // new encoder using result as buff
	encoder.Encode(block)              // encode and store the encoded data in buff result
	return result.Bytes()              // return encoded data
}

func Unserialze(encoded_block []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(encoded_block))
	decoder.Decode(&block)
	return &block
}

type FileBlockchain struct {
	LastHash []byte
	DB       *bolt.DB
}

// func OpenBlockchainFile(file_path string) *FileBlockchain {
// 	db, _ := bolt.Open(file_path, 0600, nil)
// 	db.Update(func(tx *bolt.Tx) error {
// 		blocks_bucket := tx.Bucket([]byte("BlocksBucket"))
// 		if blocks_bucket == nil {
// 			// unable to fetch bucket 'BlocksBucket'
// 			// which means chain doesn't exist
// 		}
// 	})
// }
