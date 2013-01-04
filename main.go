// Copyright 2012 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/matttproud/prometheus/retrieval"
	"github.com/matttproud/prometheus/storage/metric/leveldb"
	"log"
	"os"
)

func main() {
	m, err := leveldb.NewLevelDBMetricPersistence("/tmp/metrics")
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	defer func() {
		m.Close()
	}()

	t := &retrieval.Target{
		Address: "http://localhost:8080/metrics.json",
	}

	for i := 0; i < 100000; i++ {
		c, err := t.Scrape()
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, s := range c {
			m.AppendSample(&s)
		}

		fmt.Printf("Finished %d\n", i)
	}
}