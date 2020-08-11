package one

func RestoreIpAddresses(s string) []string {
	ips := make([]string, 0)
	address(s, 0, 4, nil, &ips)
	return ips
}

func address(s string, i, n int, ip []byte, ips *[]string) {
	if i == len(s) {
		if n == 0 {
			*ips = append(*ips, string(ip[:len(ip)-1]))
		}
		return
	}
	pn := 0
	if s[i] == '0' {
		// 只能独立一组
		if len(s)-i-1 <= (n-1)*3 {
			address(s, i+1, n-1, append(append(ip, s[i]), '.'), ips)
		}
		return
	}
	for ; i < len(s); i++ {
		pn = pn*10 + int(s[i]-'0')
		if pn > 255 {
			return
		}
		ip = append(ip, s[i])
		if len(s)-i-1 > (n-1)*3 {
			continue
		}
		address(s, i+1, n-1, append(ip, '.'), ips)
	}
}
