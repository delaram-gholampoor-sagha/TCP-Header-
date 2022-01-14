package main

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

type Packet struct {
	Header []byte

}

// SourcePort Source TCP port number (2 bytes or 16 bits):
// The source TCP port number represents the sending device.
func (p *Packet) SourcePort() uint16 {

	return binary.BigEndian.Uint16(p.Header[0:2])

}

// DestinationPort Destination TCP port number (2 bytes or 16 bits):
// The destination TCP port number is the communication endpoint for the receiving device.
func (p *Packet) DestinationPort() uint16 {

	return binary.BigEndian.Uint16(p.Header[2:4])
}

// SequenceNumber Sequence number (4 bytes or 32 bits):
// Message senders use sequence numbers to mark the ordering of a group of messages.
func (p *Packet) SequenceNumber() uint32 {

	return binary.BigEndian.Uint32(p.Header[4:8])
}

// AckNumber Acknowledgment number (4 bytes or 32 bits): Both senders and receivers
// use the acknowledgment numbers field to communicate the sequence numbers of
// messages that are either recently received or expected to be sent.
func (p *Packet) AckNumber() uint32 {

	return binary.BigEndian.Uint32(p.Header[8:12])
}

// DO TCP data offset (4 bits): The data offset field stores the total size of
// a TCP header in multiples of four bytes. A header not using the optional
// TCP field has a data offset of 5 (representing 20 bytes), while a header
// using the maximum-sized optional field has a data offset of 15 (representing 60 bytes).
func (p *Packet) DO() uint8 {

	do := fmt.Sprintf("%b", p.Header[12:14][0])
	output, _ := strconv.ParseInt(do[0:4], 2, 5)
	return uint8(output)
}

// RSV Reserved data (3 bits): Reserved data in TCP headers always has a value of zero.
// This field aligns the total header size as a multiple of four bytes,
// which is important for the efficiency of computer data processing.
func (p *Packet) RSV() uint8 {
	rs := fmt.Sprintf("%b", p.Header[12:14][0])
	output, _ := strconv.ParseInt(rs[4:7], 2, 4)

	return uint8(output)
}

// Flags Control flags (up to 9 bits): TCP uses a set of six standard and
// three extended control flags—each an individual bit representing On or Off—to manage
// data flow in specific situations.
func (p *Packet) Flags() struct {
	SYN bool
	ACK bool
	RST bool
	FIN bool
	PSH bool
	URG bool
} {
	fg1 := fmt.Sprintf("%.1b", p.Header[12:14][0])
	fg2 := fmt.Sprintf("%.8b", p.Header[12:14][1])

	data := struct {
		SYN bool
		ACK bool
		RST bool
		FIN bool
		PSH bool
		URG bool
	}{
		SYN: false, 
		ACK: false, 
		RST: false, 
		FIN: false, 
		PSH: false, 
		URG: false,
	}

	if fg1[7:8] != "0" {
		data.SYN = true
	}

	if fg2[0:1] != "0" {
		data.ACK = true
	}

	if fg2[1:2] != "0" {
		data.RST = true
	}

	if fg2[2:3] != "0" {
		data.FIN = true
	}

	if fg2[3:4] != "0" {
		data.PSH = true
	}

	if fg2[4:5] != "0"{
		data.URG = true
	}

	return data
}

// Window Window size (2 bytes or 16 bits): TCP senders use a number,
// called window size, to regulate how much data they send to a receiver before
// requiring an acknowledgment in return. If the window size is too small,
// network data transfer is unnecessarily slow. If the window size is too large,
// the network link may become saturated,
// or the receiver may not be able to process incoming data quickly enough,
// resulting in slow performance. Windowing algorithms built into the protocol
// dynamically calculate size values and use this field of TCP headers to
// coordinate changes between senders and receivers.
func (p *Packet) Window() uint16 {

	return binary.BigEndian.Uint16(p.Header[14:16])
}

// Checksum TCP checksum (2 bytes or 16 bits): The checksum value inside
// a TCP header is generated by the protocol sender as a mathematical technique
// to help the receiver detect messages that are corrupted or tampered with.
func (p *Packet) Checksum() uint16 {

	return binary.BigEndian.Uint16(p.Header[16:18])
}

// UrgentPointer Urgent pointer (2 bytes or 16 bits): The urgent pointer field
// is often set to zero and ignored, but in conjunction with one of the control
// flags, it can be used as a data offset to mark a subset of a message as
// requiring priority processing.
func (p *Packet) UrgentPointer() uint16 {

	return binary.BigEndian.Uint16(p.Header[18:20])
}

// // Options TCP optional data (0 to 40 bytes): Usages of optional TCP data
// // include support for special acknowledgment and window scaling algorithms.
// func (p *Packet) Options() {
// 	fmt.Println("Options")
// }

func main() {

	p := Packet{

		Header: []byte{
			0xb7, 0x4e,
			0x01, 0xbb,
			0xb1, 0x46,
			0xa4, 0x61,
			0x00, 0x00,
			0x00, 0x00,
			0xa0, 0x02,
			0xfa, 0xf0,
			0x9b, 0xba,
			0x00, 0x00,
		},
	}

	fmt.Println(p.SourcePort())
	fmt.Println(p.DestinationPort())
	fmt.Println(p.SequenceNumber())
	fmt.Println(p.AckNumber())
	fmt.Println(p.DO())
	fmt.Println(p.RSV())
	fmt.Println(p.Flags(), p.Flags().FIN)
	fmt.Println(p.Window())
	fmt.Println(p.Checksum())
	fmt.Println(p.UrgentPointer())

}
