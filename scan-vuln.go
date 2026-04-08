package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	// Verifica se o usuário passou a rede
	if len(os.Args) < 2 {
		fmt.Println("Uso: ./scan-vuln 192.168.1")
		return
	}

	baseIP := os.Args[1]

	fmt.Println("Varrendo rede:", baseIP+".0/24")

	// Portas alvo
	portas := []string{
		"80", "22", "443", "23", "3389", "445",
		"135", "139", "21", "389", "3306",
		"5432", "1433", "6379", "8080",
	}

	for i := 1; i <= 254; i++ {

		target := fmt.Sprintf("%s.%d", baseIP, i)

		for _, porta := range portas {

			conn, err := net.DialTimeout("tcp", target+":"+porta, 500*time.Millisecond)

			if err == nil {
				fmt.Printf("[+] Host ativo: %s (porta %s aberta)\n", target, porta)
				conn.Close()
			}
		}
	}

	fmt.Println("\nScan finalizado.")
}

