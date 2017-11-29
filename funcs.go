package boyconf

import (
	"errors"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
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

//GetCPULoad 获取 cpu 的负载
//i 1=15min 2=5min 3=1min
func GetCPULoad(i int) (avg float64, err error) {
	if i < 1 || i > 3 {
		err = errors.New("i out of range, 1=15min 2=5min 3=1min")
		return
	}
	var out []byte
	if out, err = exec.Command("uptime").Output(); err == nil {
		loads := strings.Fields(string(out))
		if l := len(loads); l > 3 {
			avg, err = strconv.ParseFloat(loads[l-i], 64)
		}
	}
	return
}
