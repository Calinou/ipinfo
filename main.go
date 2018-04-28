// ipinfo: Retrieve information about IP addresses
//
// Copyright © 2018 Hugo Locurcio and contributors
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"encoding/json"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

type IpInfo struct {
	Ip       string `json:"ip"`       // IP address
	Hostname string `json:"hostname"` // Hostname
	City     string `json:"city"`     // City
	Region   string `json:"region"`   // Region
	Country  string `json:"country"`  // Country
	Loc      string `json:"loc"`      // Location ("latitude,longitude" in degrees)
	Postal   string `json:"postal"`   // Postal code
	Org      string `json:"org"`      // Organization
}

func main() {
	app := cli.NewApp()
	app.Name = "ipinfo"
	app.Version = "0.0.1"
	app.Usage = "Retrieve information about IP addresses"
	app.UsageText = app.Name + " [IP address]"

	app.Action = func(c *cli.Context) error {
		client := &http.Client{}
		req, err := client.Get("http://ipinfo.io/json")

		if err != nil {
			fmt.Println("Error: Requesting IP address information failed.")
			os.Exit(1)
		}

		defer req.Body.Close()
		ipInfo := IpInfo{}
		json.NewDecoder(req.Body).Decode(&ipInfo)

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
