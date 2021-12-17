package day16

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_HexaToBinary(t *testing.T) {

	str := hexaToBinary("A")
	if str != "1010" {
		t.Error("hexa to binary error")
	}

	str = hexaToBinary("D2FE28")
	if str != "110100101111111000101000" {
		t.Error("hexa parsing KO")
	}

}

func Test_examples(t *testing.T) {

	printVersionSum("8A004A801A8002F478")
	printVersionSum("620080001611562C8802118E34")
	printVersionSum("C0015000016115A2E0802F182340")
	printVersionSum("A0016C880162017C3686B18A3D4780")
}

func printVersionSum(str string) {
	strBuff := bytes.NewBufferString("")
	for i := 0; i < len(str); i += 1 {
		fmt.Fprintf(strBuff, hexaToBinary(str[i:i+1]))
	}
	packet, _ := parseBinary(strBuff.String(), 0)
	fmt.Println(packet.versionSum())
}

func Test_nico(t *testing.T) {
	fmt.Println(binToLiteral("010001000000000"))
}
