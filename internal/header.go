package internal

import (
	"encoding/binary"
	"fmt"
)


type DNSHeader struct{
	ID  uint16
	Flags uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

func ParseDNSHeader(msg []byte)(*DNSHeader, error){
	if len(msg)<12{
		return nil, fmt.Errorf("packet too short for response")
	}
	h := &DNSHeader{
		ID: binary.BigEndian.Uint16(msg[0:2]),
		Flags: binary.BigEndian.Uint16(msg[2:4]),
		QDCOUNT: binary.BigEndian.Uint16(msg[4:6]),
		ANCOUNT: binary.BigEndian.Uint16(msg[6:8]),
		NSCOUNT: binary.BigEndian.Uint16(msg[8:10]),
		ARCOUNT: binary.BigEndian.Uint16(msg[10:12]),
	}
	return h, nil
}

func BuildDNSServer(h *DNSHeader) []byte{
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:2], h.ID)
	binary.BigEndian.PutUint16(buf[0:2], h.Flags)
	binary.BigEndian.PutUint16(buf[0:2], h.QDCOUNT)
	binary.BigEndian.PutUint16(buf[0:2], h.ANCOUNT)
	binary.BigEndian.PutUint16(buf[0:2], h.NSCOUNT)
	binary.BigEndian.PutUint16(buf[0:2], h.ARCOUNT)

	return buf
}