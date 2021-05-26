/*
 * File         : port-scanner.go
 * Project      : Hacking With Go
 * Created Date : Wednesday, May 26th 2021, 1:44:41 PM
 * Author       : Pramod Devireddy
 *
 * Last Modified: Wednesday, 26th May 2021 2:16:04 pm
 * Modified By  : Pramod Devireddy
 *
 * Copyright (c)2021 - Pramod Devireddy
 * ***************************************************************
 */

package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()

			address := "scanme.nmap.org:" + strconv.Itoa(j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Println(j, "port open")
		}(i)
	}
	wg.Wait()
}
