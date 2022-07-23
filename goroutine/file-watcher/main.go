package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const watchedPath = "./black-eye"

func main() {
	// runtime.GOMAXPROCS(4)

	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1)
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)

			if err := f.Close(); err != nil {
				log.Fatal("Failed to close file")
			}

			if err := os.Remove(filePath); err != nil {
				log.Fatal("Failed to remove file")
			}

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ = strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received Invoice '%v' for $%.2f and submitted for processing\n", invoice.Number, invoice.Amount)
				}
			}(string(data))
		}

		if err := d.Close(); err != nil {
			log.Fatal("Failed to close directory")
		}

		time.Sleep(100 * time.Millisecond)
	}
}

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
