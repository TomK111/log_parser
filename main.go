package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var (
		sum     map[string]int
		domains []string
		total   int
	)

	sum = make(map[string]int)

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {
		fields := strings.Fields(in.Text())
		if len(fields) != 2 {
			fmt.Println("Wrong Input.")
			return
		}

		domain := fields[0]

		visits, _ := strconv.Atoi(fields[1])

		if _, ok := sum[domain]; !ok {
			domains = append(domains, domain)
		}
		total += visits
		sum[domain] += visits
	}

	fmt.Printf("%-30s %10s\n", "DOMAIN", "VISITS")
	fmt.Println(strings.Repeat("-", 45))

	sort.Strings(domains)
	for _, domain := range domains {
		visits := sum[domain]
		fmt.Printf("%-30s %10d\n", domain, visits)
	}

	fmt.Printf("%-30s %10d\n", "TOTAL", total)
}
