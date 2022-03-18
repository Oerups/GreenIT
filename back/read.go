//Create class which will read the file and return the data
package main

import (
	"bufio"
	"os"
	"strings"
)

func getCommune(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	communes := []string{}
	line, err := reader.ReadString('\n')

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		communes = append(communes, splited[0])
	}

	return RemoveDuplicate(communes)
}

func getRegion(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	regions := []string{}
	line, err := reader.ReadString('\n')

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		regions = append(regions, splited[3])
	}

	return RemoveDuplicate(regions)
}

func getDepartements(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	departements := []string{}
	line, err := reader.ReadString('\n')

	for {
		line, err = reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		departements = append(departements, splited[1])
	}

	return RemoveDuplicate(departements)
}

func findDataLine(filename string, city string) map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	headers := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		if len(headers) == 0 {
			headers = splited
		}
		if strings.Contains(splited[0], city) {
			//Create associative array from headers and splited
			data := make(map[string]string)
			for i := 0; i < len(headers); i++ {
				data[strings.Trim(headers[i], "\n\r")] = strings.Trim(splited[i], "\n\r")
			}
			return data
		}
	}
	return nil
}

func RemoveDuplicate(array []string) []string {
	m := make(map[string]string)
	for _, x := range array {
		m[x] = x
	}
	var ClearedArr []string
	for x, _ := range m {
		ClearedArr = append(ClearedArr, x)
	}
	return ClearedArr
}
