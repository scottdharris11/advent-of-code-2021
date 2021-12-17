package day16

import (
	"advent-of-code-2021/utils"
	"log"
	"strings"
	"time"
)

type Puzzle struct{}

func (Puzzle) Solve() {
	lines := utils.ReadLines("day16", "day-16-input.txt")
	solvePart1(lines[0])
	solvePart2(lines[0])
}

func solvePart1(line string) int {
	start := time.Now().UnixMilli()
	bits := hexToBits(line)
	packet, _ := readPacket(bits, 0)
	ans := sumPacketVersions(packet)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 1 (%dms): Version Sum = %d", end-start, ans)
	return ans
}

func solvePart2(line string) int {
	start := time.Now().UnixMilli()
	bits := hexToBits(line)
	packet, _ := readPacket(bits, 0)
	ans := evaluatePacket(packet)
	end := time.Now().UnixMilli()
	log.Printf("Day 16, Part 2 (%dms): Packet = %d", end-start, ans)
	return ans
}

type Packet struct {
	version    int
	typeID     int
	literalVal int
	lengthType int
	length     int
	subPackets []*Packet
}

func hexToBits(hex string) string {
	sb := strings.Builder{}
	for _, c := range hex {
		switch c {
		case '0':
			sb.WriteString("0000")
		case '1':
			sb.WriteString("0001")
		case '2':
			sb.WriteString("0010")
		case '3':
			sb.WriteString("0011")
		case '4':
			sb.WriteString("0100")
		case '5':
			sb.WriteString("0101")
		case '6':
			sb.WriteString("0110")
		case '7':
			sb.WriteString("0111")
		case '8':
			sb.WriteString("1000")
		case '9':
			sb.WriteString("1001")
		case 'A':
			sb.WriteString("1010")
		case 'B':
			sb.WriteString("1011")
		case 'C':
			sb.WriteString("1100")
		case 'D':
			sb.WriteString("1101")
		case 'E':
			sb.WriteString("1110")
		case 'F':
			sb.WriteString("1111")
		}
	}
	return sb.String()
}

func bitsToDecimal(bits string) int {
	out := 0
	posCnt := len(bits)
	for i := 0; i < posCnt; i++ {
		bitIdx := posCnt - i - 1
		if bits[i] == '1' {
			out |= 1 << bitIdx
		}
	}
	return out
}

func readPacket(s string, offset int) (*Packet, int) {
	// parse packet header
	p := Packet{}
	o := offset
	p.version = bitsToDecimal(s[o : o+3])
	o += 3
	p.typeID = bitsToDecimal(s[o : o+3])
	o += 3

	if p.typeID == 4 {
		// literal packet
		data := strings.Builder{}
		for {
			last := s[o] == '0'
			data.WriteString(s[o+1 : o+5])
			o += 5
			if last {
				break
			}
		}
		p.literalVal = bitsToDecimal(data.String())
	} else {
		// operator packet
		if s[o] == '1' {
			p.lengthType = 1
			p.length = bitsToDecimal(s[o+1 : o+12])
			o += 12
			for i := 0; i < p.length; i++ {
				sp, nOffset := readPacket(s, o)
				o = nOffset
				p.subPackets = append(p.subPackets, sp)
			}
		} else {
			p.lengthType = 0
			p.length = bitsToDecimal(s[o+1 : o+16])
			o += 16
			bitsRead := 0
			for bitsRead < p.length {
				sp, nOffset := readPacket(s, o)
				bitsRead += nOffset - o
				o = nOffset
				p.subPackets = append(p.subPackets, sp)
			}
		}
	}

	return &p, o
}

func sumPacketVersions(p *Packet) int {
	sum := p.version
	if p.subPackets != nil {
		for _, sp := range p.subPackets {
			sum += sumPacketVersions(sp)
		}
	}
	return sum
}

func evaluatePacket(p *Packet) int {
	if p.typeID == 4 {
		return p.literalVal
	}

	var spValues []int
	for _, sp := range p.subPackets {
		spValues = append(spValues, evaluatePacket(sp))
	}

	switch p.typeID {
	case 0:
		val := 0
		for _, spVal := range spValues {
			val += spVal
		}
		return val
	case 1:
		val := 1
		for _, spVal := range spValues {
			val *= spVal
		}
		return val
	case 2:
		val := spValues[0]
		for _, spVal := range spValues {
			if spVal < val {
				val = spVal
			}
		}
		return val
	case 3:
		val := spValues[0]
		for _, spVal := range spValues {
			if spVal > val {
				val = spVal
			}
		}
		return val
	case 5:
		val := 0
		if spValues[0] > spValues[1] {
			val = 1
		}
		return val
	case 6:
		val := 0
		if spValues[0] < spValues[1] {
			val = 1
		}
		return val
	case 7:
		val := 0
		if spValues[0] == spValues[1] {
			val = 1
		}
		return val
	}

	return 0
}
