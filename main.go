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
	"github.com/fatih/color"
	"github.com/urfave/cli"
	"log"
	"net/http"
	"os"
)

type IpInfo struct {
	Ip       string            `json:"ip"`       // IP address
	Hostname string            `json:"hostname"` // Hostname
	City     string            `json:"city"`     // City
	Region   string            `json:"region"`   // Region
	Country  string            `json:"country"`  // Country
	Loc      string            `json:"loc"`      // Location ("latitude,longitude" in degrees)
	Postal   string            `json:"postal"`   // Postal code
	Org      string            `json:"org"`      // Organization
	Error    map[string]string `json:"error"`    // Error (if any)
}

func main() {
	app := cli.NewApp()
	app.Name = "ipinfo"
	app.Version = "0.1.0"
	app.Usage = "Retrieve information about IP addresses"
	app.UsageText = app.Name + " [IP address]"

	app.Action = func(c *cli.Context) error {
		numArgs := len(c.Args())

		var url string

		if numArgs == 0 {
			// Use the user's IP address (determined by ipinfo.io)
			url = "https://ipinfo.io/json"
		} else if numArgs == 1 {
			// Use the IP address provided on the command line
			url = "https://ipinfo.io/" + c.Args().Get(0) + "/json"
		} else {
			// Invalid number of arguments supplied
			fmt.Fprintf(
				color.Output,
				color.HiRedString("Error:")+
					" Not enough arguments supplied; expected 0 or 1 arguments (got %d).\n"+
					"Usage: "+app.UsageText+"\n",
				numArgs)
			os.Exit(1)
		}

		client := &http.Client{}
		req, err := client.Get(url)

		if err != nil {
			fmt.Fprintln(
				color.Output,
				color.HiRedString("Error:"),
				"Requesting IP address information failed.",
			)
			os.Exit(1)
		}

		// Decode the returned JSON into a struct
		defer req.Body.Close()
		ipInfo := IpInfo{}
		json.NewDecoder(req.Body).Decode(&ipInfo)

		if len(ipInfo.Error) == 0 {
			// Success; display the result
			fmt.Fprintln(
				color.Output,
				"\n    IP address  ", color.HiCyanString(ipInfo.Ip),
				"\n      Hostname  ", color.HiCyanString(ipInfo.Hostname),
				"\n          City  ", color.HiCyanString(ipInfo.City),
				"\n        Region  ", color.HiCyanString(ipInfo.Region),
				"\n       Country  ", color.HiCyanString(ipInfo.Country),
				"\n      Location  ", color.HiCyanString(ipInfo.Loc),
				"\n   Postal code  ", color.HiCyanString(ipInfo.Postal),
				"\n  Organization  ", color.HiCyanString(ipInfo.Org),
			)
		} else {
			// ipinfo.io returned an error
			fmt.Fprintln(
				color.Output,
				color.HiRedString("Error:"),
				"Invalid response from ipinfo.io:",
				ipInfo.Error["title"],
				"-",
				ipInfo.Error["message"],
			)
			os.Exit(1)
		}

		return nil
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
