package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type Transaction struct {
	TXID      []byte
	TXInputs  []Input
	TXOutputs []Output
}

type Input struct {
	Txid             []byte
	ReferOutputIndex int64
	UnlockScript     string
	//ScriptSigu=
}

type Output struct {
	Value      float64
	lockScript string
	//ScriptPublicKey
}

func (tx *Transaction)SetTXID()  {
	var buffer bytes.Buffer

	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(tx)
	//tmp := bytes.Join([][]byte{tx.TXInputs}, []byte{})
	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]

}


func (input *Input)ConUnlockedByAddress(unlockdata string) bool {

	return input.UnlockScript == unlockdata
}

func (output *Output)ConBeUnlockedByAddress(unlocked string) bool {
	return  output.lockScript == unlocked
}


func (tx * Transaction)IsCoinbase() bool {
	if len(tx.TXInputs) == 1 {
		if tx.TXInputs[0].Txid == nil && tx.TXInputs[0].ReferOutputIndex == -1 {
			return true
		}
	}
	return false
}