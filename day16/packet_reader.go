package day16

import (
	"fmt"
	. "github.com/ahmetalpbalkan/go-linq"
	"log"
	"strconv"
)

type PacketHeader struct {
	version, packetType int
}

type Packet struct {
	header       PacketHeader
	literalValue int64
	subPackets   []Packet
}

func (h PacketHeader) String() string {
	return fmt.Sprintf("Header{v=%d,t=%d}", h.version, h.packetType)
}

func (p Packet) PrintPacket(level int) {
	if p.header.packetType == 4 {
		fmt.Println("Literal(", p.header, "):", p.literalValue)
	} else {
		fmt.Println("Operator(", p.header, "):")
		for _, subPacket := range p.subPackets {
			for i := 0; i < level; i++ {
				fmt.Print("\t")
			}
			subPacket.PrintPacket(level + 1)
		}
		fmt.Println()
	}
}

func (p Packet) versionSum() int {
	sum := p.header.version
	for _, subPacket := range p.subPackets {
		sum += subPacket.versionSum()
	}
	return sum
}

func (p Packet) compute() int64 {
	var result int64 = 0

	switch p.header.packetType {
	case 4:
		return p.literalValue
	case 0:
		return From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).SumInts()
	case 1:
		return From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).AggregateT(func(p1, p2 int64) int64 { return p1 * p2 }).(int64)
	case 2:
		return From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).Min().(int64)
	case 3:
		return From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).Max().(int64)
	case 5:
		subValues := make([]int64, 0)
		From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).ToSlice(&subValues)
		if subValues[0] > subValues[1] {
			return 1
		}
		return 0
	case 6:
		subValues := make([]int64, 0)
		From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).ToSlice(&subValues)
		if subValues[0] < subValues[1] {
			return 1
		}
		return 0
	case 7:
		subValues := make([]int64, 0)
		From(p.subPackets).SelectT(func(p Packet) int64 { return p.compute() }).ToSlice(&subValues)
		if subValues[0] == subValues[1] {
			return 1
		}
		return 0
	}

	return result
}

func parseBinary(binaryStr string, offset int) (Packet, int) {
	packetHeader, newOffset := parseHeader(binaryStr, offset)
	offset = newOffset

	fmt.Println("reading packet with header = ", packetHeader)

	if packetHeader.packetType == 4 {
		literal, newOffset := parseLiteral(binaryStr, offset)
		//fmt.Println(", literal = ", literal, ", offset = ", offset, "]")

		return Packet{
			header:       packetHeader,
			literalValue: literal,
		}, newOffset

	} else {
		subPackets, newOffset := parseOperator(binaryStr, offset)

		return Packet{
			header:       packetHeader,
			literalValue: -1,
			subPackets:   subPackets,
		}, newOffset
	}

	//return Packet{literalValue: -999}, offset
}

func parseHeader(binaryStr string, offset int) (PacketHeader, int) {
	// new entry
	versionBin := binaryStr[offset : offset+3]
	version, err := strconv.ParseInt(versionBin, 2, 32)
	if err != nil {
		log.Fatalln("version parsing error")
	}
	offset += 3
	typeBin := binaryStr[offset : offset+3]
	packetType, err := strconv.ParseInt(typeBin, 2, 32)
	if err != nil {
		log.Fatalln("packet type parsing error")
	}
	offset += 3

	//fmt.Print("[version =", version)
	//fmt.Print(", type =", packetType)
	return PacketHeader{version: int(version), packetType: int(packetType)}, offset
}

func parseLiteral(binaryStr string, offset int) (int64, int) {
	//fmt.Println("Literal value !")

	literalBin := ""
	isLast := false
	for !isLast {
		isLast = binaryStr[offset:offset+1] == "0"
		offset += 1
		literalBin += binaryStr[offset : offset+4]
		offset += 4
	}

	literal, err := strconv.ParseInt(literalBin, 2, 64)
	if err != nil {
		//literal = -777
		log.Fatalln("error while parsing literal")
	}

	// skip padding
	//for offset % 4 != 0 {
	//	offset++
	//}

	return literal, offset
}

func parseOperator(binaryStr string, offset int) ([]Packet, int) {
	subPackets := make([]Packet, 0)

	// 0 = 15 bit // 1 = 11 bit
	lengthTypeId := binaryStr[offset : offset+1]
	offset += 1
	switch lengthTypeId {
	case "0":
		subPacketLength := binToLiteral(binaryStr[offset : offset+15])
		offset += 15

		var innerOffset int
		for innerOffset = offset; innerOffset < offset+subPacketLength; {
			var packet Packet
			packet, innerOffset = parseBinary(binaryStr, innerOffset)
			subPackets = append(subPackets, packet)
		}
		offset = innerOffset

	case "1":
		subPacketCount := binToLiteral(binaryStr[offset : offset+11])
		offset += 11

		for i := 0; i < subPacketCount; i++ {
			var packet Packet
			packet, offset = parseBinary(binaryStr, offset)
			subPackets = append(subPackets, packet)
		}
	default:
		log.Fatalln("Error while reading operator")
	}

	// skip padding
	//for offset % 4 != 0 {
	//	offset++
	//}

	return subPackets, offset
}

func hexaToBinary(hexa string) string {

	n, err := strconv.ParseInt(hexa, 16, 32)
	if err != nil {
		log.Fatalln("Hexadecimal parsing error")
	}

	// https://stackoverflow.com/questions/43003700/convert-a-hexadecimal-number-to-binary-in-go-and-be-able-to-access-each-bit
	// padding 4 ==> %04b !!
	return fmt.Sprintf("%04b", n)
}

func binToHexa(hexa string) string {

	n, err := strconv.ParseInt(hexa, 2, 32)
	if err != nil {
		log.Fatalln("Binary parsing error")
	}

	return strconv.FormatInt(n, 16)
}

func binToLiteral(binaryStr string) int {
	n, err := strconv.ParseInt(binaryStr, 2, 32)
	if err != nil {
		log.Fatalln("Binary parsing error")
	}

	return int(n)
}
