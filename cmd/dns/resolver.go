package dns

import (
	"encoding/binary"
	"fmt"
	"net"
)


type DNSHeader struct{
	ID  uint16
	Flags uint16
	QDCOUNT uint16
	ANCOUNT uint16
	NSCOUNT uint16
	ARCOUNT uint16
}

type Question struct{
	Name string
	QType uint16
	QClass uint16
	Raw []byte
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

func BuildDNSHeader(h *DNSHeader) []byte{
	buf := make([]byte, 12)
	binary.BigEndian.PutUint16(buf[0:2], h.ID)
	binary.BigEndian.PutUint16(buf[2:4], h.Flags)
	binary.BigEndian.PutUint16(buf[4:6], h.QDCOUNT)
	binary.BigEndian.PutUint16(buf[6:8], h.ANCOUNT)
	binary.BigEndian.PutUint16(buf[8:10], h.NSCOUNT)
	binary.BigEndian.PutUint16(buf[10:12], h.ARCOUNT)

	return buf
}




func ParseQuestion(msg []byte, offset int)(Question , int, error){

	start :=offset
	q := Question{}
	labels := []string{}

	// parse qname
	for {
		length := int(msg[offset])
		if length==0{
			offset++
			break
		}

		offset++
		label := string(msg[offset + length])
		labels  = append(labels, label)
		offset += length
	}

	// build domain name
	q.Name=""
	for i, l:=range labels{
		if i>0{
			q.Name += "."
		}
		q.Name +=l
	}

	// qtype
	q.QType = binary.BigEndian.Uint16(msg[offset : offset+2])
	offset+=2

	// qClass
	q.QClass = binary.BigEndian.Uint16(msg[offset : offset+2])
	offset +=2

	q.Raw = msg[start : offset]
	return q, offset, nil
}

func WriteQuestion(q Question)[]byte{
	return q.Raw
}


func BuildARecord(nameOffset uint16, ipAddr string, ttl uint32) []byte{
	answer := make([]byte,0)

	pointer := uint16(0xC000)|nameOffset

	// name
	nameByte := make([]byte ,2)
	binary.BigEndian.PutUint16(nameByte,pointer)
	answer = append(answer, nameByte...)

	// type
	typeByte := make([]byte ,2)
	binary.BigEndian.PutUint16(typeByte, 1)
	answer = append(answer, typeByte...)

	// class
	classByte := make([]byte, 2)
	binary.BigEndian.PutUint16(classByte, 1)
	answer = append(answer, classByte...)

	// ttl
	ttlByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ttlByte,ttl)
	answer = append(answer, ttlByte...)

	// rdlength
	rdLen := make([]byte, 2)
	binary.BigEndian.PutUint16(rdLen,4)
	answer = append(answer, rdLen...)

	// rdata ->ip
	ip := net.ParseIP(ipAddr).To4()
	if ip == nil{
		return nil
	}
	answer = append(answer, ip...)
	return answer

}



