package main

import "net"

type netList []net.IPNet

func parseCIDRList(s []string) (netList, error) {
	ns := make(netList, cap(s))
	for i, cidr := range s {
		_, n, err := net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}

		ns[i] = *n
	}

	return ns, nil
}

func (ns netList) Contains(ip net.IP) bool {
	if len(ns) == 0 {
		return true
	}

	for _, n := range ns {
		if n.Contains(ip) {
			return true
		}
	}

	return false
}
