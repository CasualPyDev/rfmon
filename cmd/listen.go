/*
Copyright © 2021 Andreas Kreisig

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
)

// listenCmd represents the listen command
var listenCmd = &cobra.Command{
	Use:   "listen",
	Short: "Starts the rfmon server and listens on predefined or default port.",
	Run: func(cmd *cobra.Command, args []string) {
		rfmonListener()
	},
}

func init() {
	rootCmd.AddCommand(listenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func rfmonListener() {

	service := ":56565"
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	errHandler(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	errHandler(err)

	fmt.Println("rfmon listening on port", service)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {

	for {
		buf := make([]byte, 1024)
		size, err := conn.Read(buf)
		errHandler(err)

		data := buf[:size]
		log.Println("Read new data from connection", data)
		conn.Write(data) // for now it's a simple echo server
	}
}

func errHandler(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
