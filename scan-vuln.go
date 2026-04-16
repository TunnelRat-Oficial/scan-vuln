package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// gera todos os IPs da rede
func hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string

	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ipCopy := make(net.IP, len(ip))
		copy(ipCopy, ip)
		ips = append(ips, ipCopy.String())
	}

	// remove network e broadcast
	if len(ips) > 2 {
		return ips[1 : len(ips)-1], nil
	}

	return ips, nil
}

// incrementa IP
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}

func main() {

	var network string

	// argumento ou input
	if len(os.Args) >= 2 {
		network = os.Args[1]
	} else {
		fmt.Print("Digite a rede (ex: 192.168.1.0/24): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		network = strings.TrimSpace(input)
	}

	fmt.Println("Varrendo rede:", network)

	ips, err := hosts(network)
	if err != nil {
		fmt.Println("Erro na rede:", err)
		return
	}

	portas := []string{
		"80", "22", "443", "23", "3389", "445",
		"135", "139", "21", "389", "3306",
		"5432", "1433", "6379", "8080",
	}

	for _, ip := range ips {

		for _, porta := range portas {

			conn, err := net.DialTimeout("tcp", ip+":"+porta, 500*time.Millisecond)

			if err == nil {
				fmt.Printf("[+] %s:%s aberto\n", ip, porta)
				conn.Close()
			}
		}
	}

	fmt.Println("\nScan finalizado.")
	fmt.Println("Pressione ENTER para sair...")
	fmt.Scanln()
}
