package day16

import (
	"advent-of-code-2021/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolvePart1(t *testing.T) {
	assert.Equal(t, 14, solvePart1("C200B40A82"))
	assert.Equal(t, 16, solvePart1("8A004A801A8002F478"))
	assert.Equal(t, 12, solvePart1("620080001611562C8802118E34"))
	assert.Equal(t, 23, solvePart1("C0015000016115A2E0802F182340"))
	assert.Equal(t, 31, solvePart1("A0016C880162017C3686B18A3D4780"))
	assert.Equal(t, 963, solvePart1(utils.ReadLines("day16", "day-16-input.txt")[0]))
}

func TestSolvePart2(t *testing.T) {
	assert.Equal(t, 3, solvePart2("C200B40A82"))
	assert.Equal(t, 54, solvePart2("04005AC33890"))
	assert.Equal(t, 7, solvePart2("880086C3E88112"))
	assert.Equal(t, 9, solvePart2("CE00C43D881120"))
	assert.Equal(t, 1, solvePart2("D8005AC2A8F0"))
	assert.Equal(t, 0, solvePart2("F600BC2D8F"))
	assert.Equal(t, 0, solvePart2("9C005AC2F8F0"))
	assert.Equal(t, 1, solvePart2("9C0141080250320F1802104A08"))
	assert.Equal(t, 1549026292886, solvePart2(utils.ReadLines("day16", "day-16-input.txt")[0]))
}

func TestHexToBits(t *testing.T) {
	assert.Equal(t, "110100101111111000101000", hexToBits("D2FE28"))
	assert.Equal(t, "00111000000000000110111101000101001010010001001000000000", hexToBits("38006F45291200"))
	assert.Equal(t, "11101110000000001101010000001100100000100011000001100000", hexToBits("EE00D40C823060"))
}

func TestBitsToDecimal(t *testing.T) {
	assert.Equal(t, 1, bitsToDecimal("001"))
	assert.Equal(t, 6, bitsToDecimal("110"))
	assert.Equal(t, 3, bitsToDecimal("00000000011"))
	assert.Equal(t, 2021, bitsToDecimal("011111100101"))
	assert.Equal(t, 27, bitsToDecimal("000000000011011"))
}

func TestReadPacket(t *testing.T) {
	p, o := readPacket("110100101111111000101000", 0)
	assert.Equal(t, 21, o)
	assert.Equal(t, 6, p.version)
	assert.Equal(t, 4, p.typeID)
	assert.Equal(t, 2021, p.literalVal)

	p, o = readPacket("00110100101111111000101000", 2)
	assert.Equal(t, 23, o)
	assert.Equal(t, 6, p.version)
	assert.Equal(t, 4, p.typeID)
	assert.Equal(t, 2021, p.literalVal)

	p, o = readPacket("00111000000000000110111101000101001010010001001000000000", 0)
	assert.Equal(t, 49, o)
	assert.Equal(t, 1, p.version)
	assert.Equal(t, 6, p.typeID)
	assert.Equal(t, 0, p.lengthType)
	assert.Equal(t, 27, p.length)
	assert.Equal(t, 2, len(p.subPackets))
	assert.Equal(t, 10, p.subPackets[0].literalVal)
	assert.Equal(t, 20, p.subPackets[1].literalVal)

	p, o = readPacket("11101110000000001101010000001100100000100011000001100000", 0)
	assert.Equal(t, 51, o)
	assert.Equal(t, 7, p.version)
	assert.Equal(t, 3, p.typeID)
	assert.Equal(t, 1, p.lengthType)
	assert.Equal(t, 3, p.length)
	assert.Equal(t, 3, len(p.subPackets))
	assert.Equal(t, 1, p.subPackets[0].literalVal)
	assert.Equal(t, 2, p.subPackets[1].literalVal)
	assert.Equal(t, 3, p.subPackets[2].literalVal)
}
