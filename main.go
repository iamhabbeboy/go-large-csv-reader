package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	f, err := os.Open("files/sales.csv")
	if err != nil {
		fmt.Println(err)
	}

	fileReader := csv.NewReader(f)
	fileReader.FieldsPerRecord = -1
	records, err := fileReader.ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	total := len(records)
	fmt.Println(total)
	limit := 500
	perList := total / limit
	for i := 0; i < int(perList); i++ {
		wg.Add(1)
		start := i * limit
		end := start + limit
		if end > total {
			end = total
		}
		go createFile(i, records[start:end], &wg, &mutex)
		fmt.Println(start, end)
		fmt.Println("=====================================")
		time.Sleep(1 * time.Second)
	}
	wg.Wait()
}

func createFile(index int, data any, wg *sync.WaitGroup, mut *sync.Mutex) {
	defer wg.Done()
	mut.Lock()
	fopen, err := os.Create(fmt.Sprintf("files/test-%v.txt", index))
	if err != nil {
		fmt.Println(err)
	}
	defer fopen.Close()
	j, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fopen.WriteString(string(j))
	mut.Unlock()
}
