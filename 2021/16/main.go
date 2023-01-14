package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

type Binary string
type Hexidecimal string
type Packets []Packet

type Packet struct {
	version     int
	type_id     int
	literal     int
	sub_packets Packets
}

func (packet Packet) SumVersions() int {
	total := packet.version
	for _, sub_packet := range packet.sub_packets {
		total += sub_packet.SumVersions()
	}
	return total
}

func (hex Hexidecimal) ToBinary() Binary {
	binary_map := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}
	binary := ""
	for _, char := range hex {
		binary += binary_map[string(char)]
	}

	return Binary(binary)
}

func (binary Binary) ParseLiteralValue(i int) (int, Binary) {
	bin_str := ""
	for i = i; string(binary[i]) == "1"; i += 5 {
		bin_str += string(binary[i+1 : i+5])
	}
	bin_str += string(binary[i+1 : i+5])

	literal, _ := strconv.ParseInt(bin_str, 2, 64)
	return int(literal), binary[i+5:]
}

func (binary Binary) Version() int {
	version, _ := strconv.ParseInt(string(binary[:3]), 2, 64)
	return int(version)
}

func (binary Binary) TypeId() int {
	type_id, _ := strconv.ParseInt(string(binary[3:6]), 2, 64)
	return int(type_id)
}

func (binary Binary) LengthId() int {
	length_id, _ := strconv.ParseInt(string(binary[6]), 2, 64)
	return int(length_id)
}

func (binary Binary) SubPacketsLength() int {
	length, _ := strconv.ParseInt(string(binary[7:22]), 2, 64)
	return int(length)
}

func (binary Binary) SubPacketsCount() int {
	count, _ := strconv.ParseInt(string(binary[7:18]), 2, 64)
	return int(count)
}

func (packets Packets) Sum() int {
	sum := 0
	for _, packet := range packets {
		sum += packet.literal
	}
	return sum
}

func (packets Packets) Product() int {
	product := 1
	for _, packet := range packets {
		product *= packet.literal
	}
	return product
}

func (packets Packets) Min() int {
	min := math.MaxInt
	for _, packet := range packets {
		if packet.literal < min {
			min = packet.literal
		}
	}
	return min
}

func (packets Packets) Max() int {
	max := math.MinInt
	for _, packet := range packets {
		if packet.literal > max {
			max = packet.literal
		}
	}
	return max
}

func (packets Packets) GreaterThan() int {
	if packets[0].literal > packets[1].literal {
		return 1
	} else {
		return 0
	}
}

func (packets Packets) LessThan() int {
	if packets[0].literal < packets[1].literal {
		return 1
	} else {
		return 0
	}
}

func (packets Packets) Equal() int {
	if packets[0].literal == packets[1].literal {
		return 1
	} else {
		return 0
	}
}

func (packet *Packet) Operate() {
	switch packet.type_id {
	case 0:
		packet.literal = packet.sub_packets.Sum()
	case 1:
		packet.literal = packet.sub_packets.Product()
	case 2:
		packet.literal = packet.sub_packets.Min()
	case 3:
		packet.literal = packet.sub_packets.Max()
	case 5:
		packet.literal = packet.sub_packets.GreaterThan()
	case 6:
		packet.literal = packet.sub_packets.LessThan()
	case 7:
		packet.literal = packet.sub_packets.Equal()
	}
}

func (binary Binary) ParseSubPackets() ([]Packet, Binary) {
	sub_packets, sub_packet, binary_remains := []Packet{}, Packet{}, Binary("")

	if binary.LengthId() == 0 {
		sub_binary := binary[22:(22 + binary.SubPacketsLength())]
		binary_remains = binary[(22 + binary.SubPacketsLength()):]

		for len(sub_binary) > 0 {
			sub_packet, sub_binary = sub_binary.Parse()
			sub_packets = append(sub_packets, sub_packet)
		}
	} else {
		num_sub_packets := binary.SubPacketsCount()
		sub_binary := binary[18:]

		for num_sub_packets > 0 {
			sub_packet, sub_binary = sub_binary.Parse()
			sub_packets = append(sub_packets, sub_packet)
			num_sub_packets -= 1
		}

		binary_remains = sub_binary
	}

	return sub_packets, binary_remains
}

func (binary Binary) Parse() (Packet, Binary) {
	packet := Packet{version: binary.Version(), type_id: binary.TypeId()}
	binary_remains := Binary("")

	if packet.type_id == 4 {
		packet.literal, binary_remains = binary.ParseLiteralValue(6)
		return packet, binary_remains
	}

	packet.sub_packets, binary_remains = binary.ParseSubPackets()
	packet.Operate()

	return packet, binary_remains
}

func main() {
	start := time.Now()
	f, _ := os.Open("input.txt")
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Scan()

	hex := Hexidecimal(scanner.Text())
	bin := hex.ToBinary()

	packet, _ := bin.Parse()
	packet_versions_sum := packet.SumVersions()

	time_elapsed := time.Since(start)

	log.Printf(`
The sum of the packet versions is %d.
The result of evaluating the BITS transmission is %d.
Solution generated in %s.`,
		packet_versions_sum,
		packet.literal,
		time_elapsed,
	)
}
