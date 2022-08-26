package main

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ParsedBlock struct {
	Version            uint32 `json:"version"`
	PreviousHeaderhash string `json:"previousHeaderHash"`
	MerkleRootHash     string `json:"merkleRootHash"`
	Timestamp          uint32 `json:"timestamp"`
	Nbits              uint32 `json:"nbits"`
	Nonce              uint32 `json:"nonce"`
}

func ParseBlock(binaryBlock []byte) ParsedBlock {
	binHeader := binaryBlock[0:80]
	version := binary.LittleEndian.Uint32(binHeader[0:4])
	previousHeaderHash := hex.EncodeToString(binHeader[4:36])
	merkleRootHash := hex.EncodeToString(binHeader[36:68])
	timestamp := binary.LittleEndian.Uint32(binHeader[68:72])
	nbits := binary.LittleEndian.Uint32(binHeader[72:76])
	nonce := binary.LittleEndian.Uint32(binHeader[76:80])

	return ParsedBlock{version, previousHeaderHash, merkleRootHash, timestamp, nbits, nonce}
}

var parsedBlocks = []ParsedBlock{}

func BlockHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		blockNumber, err := strconv.Atoi(r.URL.Query().Get("blockNumber"))
		if err != nil || blockNumber >= len(parsedBlocks) {
			http.Error(w, "Invalid Block Number", http.StatusBadRequest)
			return
		}
		block := parsedBlocks[blockNumber]
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(block)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {

	bin, _ := os.ReadFile("blk00000.dat")

	for len(bin) > 0 {
		blockSize := binary.LittleEndian.Uint32(bin[4:8])

		parsedBlocks = append(parsedBlocks, ParseBlock(bin[8:8+blockSize]))

		bin = bin[8+blockSize:]
	}
	fmt.Println("Blocks loaded into memory, starting http server")

	mux := http.NewServeMux()

	mux.HandleFunc("/block", BlockHandler)

	http.ListenAndServe(":8080", mux)
}
