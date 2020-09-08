// +build linux darwin aix

package ofctrl

import "net"

func DialUnixOrNamedPipe(address string) (net.Conn, error) {
	return net.Dial("unix", address)
}
