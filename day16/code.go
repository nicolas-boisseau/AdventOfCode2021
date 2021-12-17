package day16

import (
	"AdventOfCode2021/common"
	"bytes"
	"fmt"
)

func Process(fileName string) (int, int64) {
	lines := common.ReadLinesFromFile(fileName)

	strBuff := bytes.NewBufferString("")
	for i := 0; i < len(lines[0]); i += 1 {
		fmt.Fprintf(strBuff, hexaToBinary(lines[0][i:i+1]))
	}

	binaryStr := strBuff.String()
	//fmt.Println("bin str=", binaryStr)

	packet, _ := parseBinary(binaryStr, 0)

	//packet.PrintPacket(1)

	return packet.versionSum(), packet.compute()
}
