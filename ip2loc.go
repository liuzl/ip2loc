package ip2loc

import (
	"bytes"
	"encoding/binary"
	"errors"
	"net"
)

var (
	ErrInvalidIp  = errors.New("invalid ip address")
	ErrIpNotFound = errors.New("ip address not found")
)

type IpRange struct {
	Begin uint32
	End   uint32
	Data  string
}

func Find(ip string) (string, error) {
	p, err := StrToInt(ip)
	if err != nil {
		return "", err
	}
	var low, high uint32 = 0, uint32(len(ipData) - 1)
	for low < high-1 {
		m := (low + high) / 2
		if p >= ipData[m].Begin && p <= ipData[m].End {
			return ipData[m].Data, nil
		}
		if p < ipData[m].Begin {
			high = m
		} else {
			low = m
		}
	}
	return "", ErrIpNotFound
}

func StrToInt(ip string) (uint32, error) {
	netIP := net.ParseIP(ip)
	if netIP == nil {
		return 0, ErrInvalidIp
	}
	ipv4 := netIP.To4()
	if ipv4 == nil {
		return 0, ErrInvalidIp
	}
	var long uint32
	binary.Read(bytes.NewBuffer(ipv4), binary.BigEndian, &long)
	return long, nil
}
