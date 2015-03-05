package main

import (
	"net"
	"reflect"
	"testing"
)

var parseCIDRTests = []struct {
	in    []string
	nList netList
	err   error
}{
	{[]string{"135.104.0.0/32"}, netList{{IP: net.IPv4(135, 104, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 255)}}, nil},
	{[]string{"135.104.0.0/32"}, netList{{IP: net.IPv4(135, 104, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 255)}}, nil},
	{[]string{"0.0.0.0/24"}, netList{{IP: net.IPv4(0, 0, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)}}, nil},
	{[]string{"135.104.0.0/24"}, netList{{IP: net.IPv4(135, 104, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)}}, nil},
	{[]string{"135.104.0.1/32"}, netList{{IP: net.IPv4(135, 104, 0, 1), Mask: net.IPv4Mask(255, 255, 255, 255)}}, nil},
	{[]string{"135.104.0.1/24"}, netList{{IP: net.IPv4(135, 104, 0, 0), Mask: net.IPv4Mask(255, 255, 255, 0)}}, nil},
	{[]string{"::1/128"}, netList{{IP: net.ParseIP("::1"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:ffff"))}}, nil},
	{[]string{"abcd:2345::/127"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff:ffff:ffff:ffff:fffe"))}}, nil},
	{[]string{"abcd:2345::/65"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff:8000::"))}}, nil},
	{[]string{"abcd:2345::/64"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:ffff::"))}}, nil},
	{[]string{"abcd:2345::/63"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff:fffe::"))}}, nil},
	{[]string{"abcd:2345::/33"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:8000::"))}}, nil},
	{[]string{"abcd:2345::/32"}, netList{{IP: net.ParseIP("abcd:2345::"), Mask: net.IPMask(net.ParseIP("ffff:ffff::"))}}, nil},
	{[]string{"abcd:2344::/31"}, netList{{IP: net.ParseIP("abcd:2344::"), Mask: net.IPMask(net.ParseIP("ffff:fffe::"))}}, nil},
	{[]string{"abcd:2300::/24"}, netList{{IP: net.ParseIP("abcd:2300::"), Mask: net.IPMask(net.ParseIP("ffff:ff00::"))}}, nil},
	{[]string{"abcd:2345::/24"}, netList{{IP: net.ParseIP("abcd:2300::"), Mask: net.IPMask(net.ParseIP("ffff:ff00::"))}}, nil},
	{[]string{"2001:DB8::/48"}, netList{{IP: net.ParseIP("2001:DB8::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff::"))}}, nil},
	{[]string{"2001:DB8::1/48"}, netList{{IP: net.ParseIP("2001:DB8::"), Mask: net.IPMask(net.ParseIP("ffff:ffff:ffff::"))}}, nil},
	{[]string{"192.168.1.1/255.255.255.0"}, nil, &net.ParseError{Type: "CIDR address", Text: "192.168.1.1/255.255.255.0"}},
	{[]string{"192.168.1.1/35"}, nil, &net.ParseError{Type: "CIDR address", Text: "192.168.1.1/35"}},
	{[]string{"2001:db8::1/-1"}, netList{}, &net.ParseError{Type: "CIDR address", Text: "2001:db8::1/-1"}},
	{[]string{""}, netList{}, &net.ParseError{Type: "CIDR address", Text: ""}},
}

func TestParseCIDRList(t *testing.T) {
	for _, tt := range parseCIDRTests {
		nList, err := parseCIDRList(tt.in)
		if !reflect.DeepEqual(err, tt.err) {
			t.Errorf("parseCIDRList(%q) = %v; want %v", tt.in, nList, tt.nList)
		}
		// if err == nil && (!tt.ip.Equal(ip) || !tt.net.IP.Equal(net.IP) || !reflect.DeepEqual(net.Mask, tt.net.Mask)) {
		// 	t.Errorf("ParseCIDR(%q) = %v, {%v, %v}; want %v, {%v, %v}", tt.in, ip, net.IP, net.Mask, tt.ip, tt.net.IP, tt.net.Mask)
		// }
	}
}

var netListContainsTests = []struct {
	ip    net.IP
	nList netList
	ok    bool
}{
	{net.IPv4(172, 16, 1, 1), netList{{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)}}, true},
	{net.IPv4(172, 24, 0, 1), netList{{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(13, 32)}}, false},
	{net.IPv4(192, 168, 0, 3), netList{{IP: net.IPv4(192, 168, 0, 0), Mask: net.IPv4Mask(0, 0, 255, 252)}}, true},
	{net.IPv4(192, 168, 0, 4), netList{{IP: net.IPv4(192, 168, 0, 0), Mask: net.IPv4Mask(0, 255, 0, 252)}}, false},
	{net.ParseIP("2001:db8:1:2::1"), netList{{IP: net.ParseIP("2001:db8:1::"), Mask: net.CIDRMask(47, 128)}}, true},
	{net.ParseIP("2001:db8:1:2::1"), netList{{IP: net.ParseIP("2001:db8:2::"), Mask: net.CIDRMask(47, 128)}}, false},
	{net.ParseIP("2001:db8:1:2::1"), netList{{IP: net.ParseIP("2001:db8:1::"), Mask: net.IPMask(net.ParseIP("ffff:0:ffff::"))}}, true},
	{net.ParseIP("2001:db8:1:2::1"), netList{{IP: net.ParseIP("2001:db8:1::"), Mask: net.IPMask(net.ParseIP("0:0:0:ffff::"))}}, false},
}

func TestIPNetContains(t *testing.T) {
	for _, tt := range netListContainsTests {
		if ok := tt.nList.Contains(tt.ip); ok != tt.ok {
			t.Errorf("netList(%v).Contains(%v) = %v, want %v", tt.nList, tt.ip, ok, tt.ok)
		}
	}
}
