// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"net/http"
	"time"
	"os"

	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"contrib.go.opencensus.io/exporter/stackdriver/propagation"

	"go.opencensus.io/stats/view"
)


func main() {
	var server = "http://"+os.Getenv("APPLICATION_SERVER")+":50030"
	// Register stats and trace exporters to export the collected data.
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("YOUR_PROJECT_ID"),
	})
	if err != nil {
		log.Fatal(err)
	}
	view.RegisterExporter(exporter)
	trace.RegisterExporter(exporter)

	// Always trace for this demo.
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	// Report stats at every second.
	view.SetReportingPeriod(1 * time.Second)

	client := &http.Client{
		Transport: &ochttp.Transport{
				// Use Google Cloud propagation format.
				Propagation: &propagation.HTTPFormat{},
		},
	}

	resp, err := client.Get(server)
	if err != nil {
		log.Printf("Failed to get response: %v", err)
	} else {
		resp.Body.Close()
	}

	time.Sleep(2 * time.Second) // Wait until stats are reported.
}
