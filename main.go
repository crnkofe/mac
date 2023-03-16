// Simple MAC converter
// Copyright (C) 2023 Simon Mihevc
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("First arg. must be an integer or a MAC")
		return
	}

	maybeMAC := os.Args[1]
	if strings.Contains(maybeMAC, ":") {
		mac, err := net.ParseMAC(maybeMAC)
		if err != nil {
			fmt.Printf("Failed parsing MAC string: %s\n", err.Error())
			return
		}
		for len(mac) < 8 {
			mac = append([]byte{0}, mac...)
		}
		fmt.Printf("%d", binary.BigEndian.Uint64(mac))
	} else {
		num, err := strconv.ParseUint(os.Args[1], 10, 64)
		if err != nil {
			fmt.Printf("Failed parsing input: %s\n", err.Error())
			return
		}
		raw := []byte(fmt.Sprintf("%012x", num))
		actualMAC := ""
		for i := 0; i < len(raw); i++ {
			if i > 0 && (i%2) == 0 {
				actualMAC = actualMAC + ":"
			}

			actualMAC = actualMAC + string(raw[i])
		}
		fmt.Println(actualMAC)
	}
}
