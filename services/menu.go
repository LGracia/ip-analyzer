package services

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"lgracia.com/ip-analyzer/models"
)

func ShowMenu() {
	is_running := true
	run(&is_running)
}

func menu() {
	fmt.Println("**********************")
	fmt.Println("*********MENU*********")
	fmt.Println("**********************")
	fmt.Println("1. Analyze IP")
	fmt.Println("2. Generate statistics")
	fmt.Println("3. Exit")
}

func ipMenu(c *models.Country) {
	fmt.Println("Write an IP")

	reader := bufio.NewReader(os.Stdin)
	ip, _, _ := reader.ReadLine()
	stringIP := string(ip)

	if !checkIPAddress(stringIP) {
		return
	}

	handleRequest(c, stringIP)
}

func checkIPAddress(ip string) bool {
    if net.ParseIP(ip) == nil {
        fmt.Printf("IP Address: %s - Invalid\n", ip)
		return false
    }

	return true
}

func processSelection(char rune, is_running *bool) {
	var c models.Country
	var s models.Statistic

	switch char {
	case '1':
		ipMenu(&c)
		saveStatistic(&s, &c)
		c.Show()
		break
	case '2':
		showStatistic(&s)
		break
	case '3':
		*is_running = false
		break
	default:
		fmt.Println("ERROR: Invalid selection")
		break
	}
}

func run(is_running *bool) {
	for *is_running {
		menu()

		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()

		if err != nil {
			fmt.Println(err)
		}

		processSelection(char, is_running)
	}
}
