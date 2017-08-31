package main

import (
	"errors"
	"fmt"
	"net"
)

func main() {

	fmt.Println("Will attempt to print out ip addresses..")

	fmt.Println("Attempting to call ExternalIPv4()")
	ip, err := ExternalIPv4()

	if err != nil {

		fmt.Println("ExternalIPv4() resulted in error..")
		fmt.Printf("Err: %s\n", err.Error())
	} else {

		fmt.Println("ExternalIPv4() was successful (no error), Result:")
		fmt.Printf("%+v\n", ip)
	}

	fmt.Println("----------------------------------------")

}

// ExternalIPv4 Return the IPv4 external address of this device.
// Note external does not necessarily mean WAN IP. On most networks it will be the LAN IP of device as opposed
// to internal localhost address (127.0.0.1)
func ExternalIPv4() (net.IP, error) {

	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, err
	}

	fmt.Printf("result of net.Interfaces(): %+v\n", ifaces)
	fmt.Println("--")

	fmt.Println("Entering loop through ifaces")

	for i, iface := range ifaces {

		fmt.Println("--")
		fmt.Printf("iface %d - %+v", i, iface)
		fmt.Println("--")

		if iface.Flags&net.FlagUp == 0 {

			fmt.Println("iface is down - skipping out of loop")

			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {

			fmt.Println("loopback interface - skipping out of loop")

			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}

		fmt.Println("--")
		fmt.Printf("result of iface.Addrs: %+v\n", addrs)
		fmt.Println("--")

		fmt.Println("Entering loop addrs")
		for i, addr := range addrs {

			fmt.Println("--")
			fmt.Printf("addr %d - %+v", i, addr)
			fmt.Println("--")

			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {

				fmt.Println("ip is loopback - skipping out of loop")

				continue
			}
			ip = ip.To4()
			if ip == nil {

				fmt.Println("not an ipv4 address - skipping out of loop")

				continue // not an ipv4 address
			}

			fmt.Println("Should have just gotten an ipv4 address and returning that :)")
			return ip, nil
		}
	}

	fmt.Println("Will return an error :(")

	return nil, errors.New("Device does not appear to be network connected.")
}
