package goipaddress

import (
	"net"
	"regexp"
	"strconv"
	"strings"
)

/*IPv4Range contains the range of IP addresses*/
type IPv4Range []string

/*IPv4Addr its a full IPv4 addres representation*/
type IPv4Addr struct {
	AddrIP net.IP
	IntIP  int64
}

/*IPv4Network represents an IPs network*/
type IPv4Network struct {
	AddrIP  string
	IPrange IPv4Range
}

// ToInt convert IPv4 addres to its integer representation
func ToInt(ipAddr string) int64 {
	strIP := ipAddr
	binIP := ""
	chunks := strings.Split(strIP, ".")

	for _, elem := range chunks {
		intElem, _ := strconv.Atoi(elem)
		binIPOctet := strconv.FormatInt(int64(intElem), 2)
		paddedOctet := strings.Repeat("0", 8-len(binIPOctet)) + binIPOctet
		binIP += paddedOctet
	}

	convToInt, _ := strconv.ParseInt(binIP, 2, 64)
	return convToInt
}

// FromInt convert int64 representation of IPv4 address to its string view
func FromInt(ipInt int64) string {
	binIP := strconv.FormatInt(ipInt, 2)
	binIP = strings.Repeat("0", 32-len(binIP)) + binIP
	convertedIP := ""
	for i, j := 0, 8; i <= 24; i += 8 {
		oct, _ := strconv.ParseInt(binIP[i:j], 2, 64)
		convertedIP += strconv.Itoa(int(oct)) + "."
		if j < 32 {
			j += 8
		}
	}
	return convertedIP[0 : len(convertedIP)-1]
}

// IPv4create fills fields of IPv4Addr struct with appropriate values
func IPv4create(ipAddr string) IPv4Addr {
	var ip IPv4Addr
	if isValid(ipAddr) {
		ip.AddrIP = net.ParseIP(ipAddr)
		ip.IntIP = ToInt(ipAddr)
	}
	return ip
}

/*IPv4NetworkCreate creates new IPv4Network instance*/
func IPv4NetworkCreate(ipAddr string) IPv4Network {
	var rang IPv4Network
	rang.AddrIP = ipAddr
	rang.IPrange = parseAster(ipAddr)
	return rang
}

func isValid(ip string) bool {
	splited := strings.Split(ip, ".")
	re, _ := regexp.Compile(`(\d){1,3}\.(\d){1,3}\.(\d){1,3}\.(\d){1,3}`)
	if re.Match([]byte(ip)) {
		for _, val := range splited {
			ival, err := strconv.Atoi(val)
			if err != nil {
				return false
			}
			if ival > 255 || ival < 0 {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func parseAster(ipAddr string) IPv4Range {
	var retRange IPv4Range
	for i := 1; i < 255; i++ {
		//var temp []string
		si := strconv.Itoa(i)
		repl := strings.Replace(ipAddr, "*", si, 1)
		retRange = append(retRange, repl)
	}
	return retRange
}

/*
func parseHyphen(ipAddr string) IPv4Range {

}
*/
