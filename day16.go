package main

import (
	"bufio"
	"fmt"
	"os"
)

func advent161(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	versionSum := 0
	type packet struct {
		Version    int
		Type       int
		Value      int
		LengthType int
		Size       int
		SubPackets []packet
	}
	packets := []packet{}

	var decodePacket func(bits string, index int) (int, int, packet)
	decodePacket = func(bits string, index int) (int, int, packet) {
		sum := 0
		newIndex := index
		newPacket := packet{}

		newPacket.Version = 0
		for j := 0; j < 3; j++ {
			newPacket.Version = newPacket.Version * 2
			newPacket.Version = newPacket.Version + int(bits[newIndex+j]-'0')
		}
		newIndex = newIndex + 3
		sum = sum + newPacket.Version

		newPacket.Type = 0
		for j := 0; j < 3; j++ {
			newPacket.Type = newPacket.Type * 2
			newPacket.Type = newPacket.Type + int(bits[newIndex+j]-'0')
		}
		newIndex = newIndex + 3

		if newPacket.Type == 4 {
			bitValue := ""
			for newIndex < len(bits) {
				bitValue = bitValue + bits[newIndex+1:newIndex+5]
				newIndex = newIndex + 5
				if bits[newIndex-5] == '0' {
					break
				}
			}
			newPacket.Value = 0
			for i := 0; i < len(bitValue); i++ {
				newPacket.Value = newPacket.Value * 2
				newPacket.Value = newPacket.Value + int(bitValue[i]-'0')
			}
		} else {
			newPacket.LengthType = int(bits[newIndex] - '0')
			newIndex = newIndex + 1

			if newPacket.LengthType == 0 {
				newPacket.Size = 0
				for j := 0; j < 15; j++ {
					newPacket.Size = newPacket.Size * 2
					newPacket.Size = newPacket.Size + int(bits[newIndex+j]-'0')
				}
				newIndex = newIndex + 15

				endPacket := newIndex + newPacket.Size
				for newIndex < endPacket {
					returnSum, returnIndex, returnSubPacket := decodePacket(bits, newIndex)
					sum = sum + returnSum
					newIndex = returnIndex
					newPacket.SubPackets = append(newPacket.SubPackets, returnSubPacket)
				}
			} else {
				newPacket.Size = 0
				for j := 0; j < 11; j++ {
					newPacket.Size = newPacket.Size * 2
					newPacket.Size = newPacket.Size + int(bits[newIndex+j]-'0')
				}
				newIndex = newIndex + 11

				subPacketCount := 0
				for subPacketCount < newPacket.Size {
					returnSum, returnIndex, returnSubPacket := decodePacket(bits, newIndex)
					sum = sum + returnSum
					newIndex = returnIndex
					newPacket.SubPackets = append(newPacket.SubPackets, returnSubPacket)
					subPacketCount++
				}
			}
		}

		fmt.Println("New Packet:", sum, newIndex, newPacket)
		return sum, newIndex, newPacket
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	hexStream := scanner.Text()
	fd.Close()

	bitStream := ""
	for _, hex := range hexStream {
		switch hex {
		case '0':
			bitStream = bitStream + "0000"
		case '1':
			bitStream = bitStream + "0001"
		case '2':
			bitStream = bitStream + "0010"
		case '3':
			bitStream = bitStream + "0011"
		case '4':
			bitStream = bitStream + "0100"
		case '5':
			bitStream = bitStream + "0101"
		case '6':
			bitStream = bitStream + "0110"
		case '7':
			bitStream = bitStream + "0111"
		case '8':
			bitStream = bitStream + "1000"
		case '9':
			bitStream = bitStream + "1001"
		case 'A':
			bitStream = bitStream + "1010"
		case 'B':
			bitStream = bitStream + "1011"
		case 'C':
			bitStream = bitStream + "1100"
		case 'D':
			bitStream = bitStream + "1101"
		case 'E':
			bitStream = bitStream + "1110"
		case 'F':
			bitStream = bitStream + "1111"
		}
	}
	fmt.Println("BitStream is", bitStream)

	i := 0
	for i < len(bitStream)-4 {
		returnSum, returnIndex, newPacket := decodePacket(bitStream, i)
		versionSum = versionSum + returnSum
		i = returnIndex
		packets = append(packets, newPacket)
	}

	fmt.Println("Version Sum is", versionSum)
}

func advent162(filename string) {
	var fd *os.File
	var err error
	var scanner *bufio.Scanner
	operations := map[int]string{0: "Sum", 1: "Product", 2: "Min", 3: "Max", 4: "Value", 5: ">", 6: "<", 7: "=="}
	type packet struct {
		Version    int
		Type       int
		Value      int
		LengthType int
		Size       int
		SubPackets []packet
	}

	var decodePacket func(bits string, index int) (int, packet)
	decodePacket = func(bits string, index int) (int, packet) {
		newIndex := index
		newPacket := packet{}

		newPacket.Version = 0
		for i := 0; i < 3; i++ {
			newPacket.Version = newPacket.Version * 2
			newPacket.Version = newPacket.Version + int(bits[newIndex+i]-'0')
		}
		newIndex = newIndex + 3

		newPacket.Type = 0
		for i := 0; i < 3; i++ {
			newPacket.Type = newPacket.Type * 2
			newPacket.Type = newPacket.Type + int(bits[newIndex+i]-'0')
		}
		newIndex = newIndex + 3

		if newPacket.Type == 4 {
			bitValue := ""
			for newIndex < len(bits) {
				bitValue = bitValue + bits[newIndex+1:newIndex+5]
				newIndex = newIndex + 5
				if bits[newIndex-5] == '0' {
					break
				}
			}
			newPacket.Value = 0
			for i := 0; i < len(bitValue); i++ {
				newPacket.Value = newPacket.Value * 2
				newPacket.Value = newPacket.Value + int(bitValue[i]-'0')
			}
		} else {
			newPacket.LengthType = int(bits[newIndex] - '0')
			newIndex = newIndex + 1

			if newPacket.LengthType == 0 {
				newPacket.Size = 0
				for i := 0; i < 15; i++ {
					newPacket.Size = newPacket.Size * 2
					newPacket.Size = newPacket.Size + int(bits[newIndex+i]-'0')
				}
				newIndex = newIndex + 15

				endPacket := newIndex + newPacket.Size
				for newIndex < endPacket {
					returnIndex, returnSubPacket := decodePacket(bits, newIndex)
					newIndex = returnIndex
					newPacket.SubPackets = append(newPacket.SubPackets, returnSubPacket)
				}
			} else {
				newPacket.Size = 0
				for i := 0; i < 11; i++ {
					newPacket.Size = newPacket.Size * 2
					newPacket.Size = newPacket.Size + int(bits[newIndex+i]-'0')
				}
				newIndex = newIndex + 11

				subPacketCount := 0
				for subPacketCount < newPacket.Size {
					returnIndex, returnSubPacket := decodePacket(bits, newIndex)
					newIndex = returnIndex
					newPacket.SubPackets = append(newPacket.SubPackets, returnSubPacket)
					subPacketCount++
				}
			}

			switch newPacket.Type {
			case 0:
				newPacket.Value = 0
				for _, sp := range newPacket.SubPackets {
					newPacket.Value = newPacket.Value + sp.Value
				}
			case 1:
				newPacket.Value = 1
				for _, sp := range newPacket.SubPackets {
					newPacket.Value = newPacket.Value * sp.Value
				}
			case 2:
				newPacket.Value = -1
				for _, sp := range newPacket.SubPackets {
					if sp.Value < newPacket.Value || newPacket.Value == -1 {
						newPacket.Value = sp.Value
					}
				}
			case 3:
				newPacket.Value = -1
				for _, sp := range newPacket.SubPackets {
					if sp.Value > newPacket.Value || newPacket.Value == -1 {
						newPacket.Value = sp.Value
					}
				}
			case 5:
				if newPacket.SubPackets[0].Value > newPacket.SubPackets[1].Value {
					newPacket.Value = 1
				} else {
					newPacket.Value = 0
				}
			case 6:
				if newPacket.SubPackets[0].Value < newPacket.SubPackets[1].Value {
					newPacket.Value = 1
				} else {
					newPacket.Value = 0
				}
			case 7:
				if newPacket.SubPackets[0].Value == newPacket.SubPackets[1].Value {
					newPacket.Value = 1
				} else {
					newPacket.Value = 0
				}
			}
		}

		fmt.Println("New Packet: ", operations[newPacket.Type], newPacket.Value)
		return newIndex, newPacket
	}

	fd, err = os.Open(filename)
	check(err)
	scanner = bufio.NewScanner(fd)
	scanner.Scan()
	hexStream := scanner.Text()
	fd.Close()

	bitStream := ""
	for _, hex := range hexStream {
		switch hex {
		case '0':
			bitStream = bitStream + "0000"
		case '1':
			bitStream = bitStream + "0001"
		case '2':
			bitStream = bitStream + "0010"
		case '3':
			bitStream = bitStream + "0011"
		case '4':
			bitStream = bitStream + "0100"
		case '5':
			bitStream = bitStream + "0101"
		case '6':
			bitStream = bitStream + "0110"
		case '7':
			bitStream = bitStream + "0111"
		case '8':
			bitStream = bitStream + "1000"
		case '9':
			bitStream = bitStream + "1001"
		case 'A':
			bitStream = bitStream + "1010"
		case 'B':
			bitStream = bitStream + "1011"
		case 'C':
			bitStream = bitStream + "1100"
		case 'D':
			bitStream = bitStream + "1101"
		case 'E':
			bitStream = bitStream + "1110"
		case 'F':
			bitStream = bitStream + "1111"
		}
	}
	fmt.Println("BitStream is", bitStream)

	_, newPacket := decodePacket(bitStream, 0)

	fmt.Println("Result is", newPacket.Value)
}
