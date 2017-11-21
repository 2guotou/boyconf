package boyconf

import (
	"net"
	"regexp"
)

//GetLocalIP 获取本地可用 IP 地址, 局域网内地址
func GetLocalIP() (ip string) {
	if ks, err := net.InterfaceAddrs(); err == nil {
		ex := `(\d+\.\d+\.\d+\.\d+)\/\d+`
		re := regexp.MustCompile(ex)
		for _, k := range ks {
			if match, _ := regexp.MatchString(ex, k.String()); k.Network() == "ip+net" && match {
				ip = re.FindStringSubmatch(k.String())[1]
				if ip != "127.0.0.1" && ip != "0.0.0.0" {
					return
				}
			}
		}
	}
	return
}
