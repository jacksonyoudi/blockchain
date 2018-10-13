package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func Int2Byte(num int64) []byte {
	var buffer bytes.Buffer
	err := binary.Write(&buffer, binary.BigEndian, num)
	CheckErr(err)
	return buffer.Bytes()
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println("err:", err)
		os.Exit(1)
	}
}