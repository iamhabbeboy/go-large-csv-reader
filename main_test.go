package main

import (
	"encoding/json"
	"os"
	"reflect"
	"sync"
	"testing"
)

type TestData struct {
	Name string
	Age  int
}

func TestCreateFile(t *testing.T) {
	var wg sync.WaitGroup
	var mutex sync.Mutex

	data := []TestData{
		{Name: "John", Age: 20},
		{Name: "Doe", Age: 30},
		{Name: "Jane", Age: 40},
	}
	wg.Add(1)
	go CreateFile(1, data, &wg, &mutex)
	wg.Wait()

	filename := "files/test-1.txt"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("File %v does not exist", filename)
	}
	fopen, err := os.Open(filename)
	if err != nil {
		t.Errorf("Error opening file %v", filename)
	}
	defer fopen.Close()
	var jsonData []TestData
	err = json.NewDecoder(fopen).Decode(&jsonData)
	if err != nil {
		t.Errorf("Error decoding json file %v", err.Error())
	}
	if !reflect.DeepEqual(data, jsonData) {
		t.Errorf("Data does not match")
	}
}
