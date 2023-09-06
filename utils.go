package main

import (
	"net"
	"os"
)

func SublimeContains(s, substr string) bool {
	rs, rsubstr := []rune(s), []rune(substr)
	if len(rsubstr) > len(rs) {
		return false
	}

	var ok = true
	var i, j = 0, 0
	for ; i < len(rsubstr); i++ {
		found := -1
		for ; j < len(rs); j++ {
			if rsubstr[i] == rs[j] {
				found = j
				break
			}
		}
		if found == -1 {
			ok = false
			break
		}
		j += 1
	}
	return ok
}

// getLocalIP returns the non loopback local IP of the host
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func isFile(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsRegular()
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.Mode().IsDir()
}
