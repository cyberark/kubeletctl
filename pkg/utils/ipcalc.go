package utils

import "net"

// https://gist.github.com/kotakanbe/d3059af990252ba89a82

func getHosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	// remove network address and broadcast address
	lenIPs := len(ips)
	var ipAddresses []string

	switch {
	case lenIPs < 2:
		ipAddresses = ips

	default:
		// Shouldn't be panic here because we are checking the lenIPs before
		ipAddresses = ips[1 : len(ips)-1]
	}

	return ipAddresses, nil
}

//  http://play.golang.org/p/m8TNTtygK0
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
