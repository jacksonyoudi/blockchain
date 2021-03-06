package v1

/*
type Block struct {
	Version int64
	PrevBlockHash []byte
	Hash []byte
	TimeStamp int64
	TargetBits int64
	Nonce int64
	MerkeRoot []byte
	Data []byte
}

func NewBlock(data string, prevBlockHash []byte)  *Block{
	block :=  &Block{
		Version:1,
		PrevBlockHash:prevBlockHash,
		TimeStamp:time.Now().Unix(),
		TargetBits:10,
		Nonce:5,
		MerkeRoot: []byte{},
		Data: []byte(data)}
	block.SetHash()

	return block

}

// int -> byte

func (block *Block)SetHash()  {
	tmp := [][]byte{
		Int2Byte(block.Version),
		block.PrevBlockHash,
		Int2Byte(block.TimeStamp),
		block.MerkeRoot,
		Int2Byte(block.Nonce),
		block.Data}
	data := bytes.Join(tmp, []byte{})
	hash := sha256.Sum256(data)
	block.Hash = hash[:]
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block!", []byte{})
}*/