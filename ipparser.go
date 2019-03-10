package ipparser

import (
	"net"
)

/*IPv4Addr is ipv4 addres repr*/
type IPv4Addr struct {
	/*IPv4 addres repr*/
	ip net.IP
}

func (ip *IPv4Addr) toInt() {
	ip.ip.String()
}
