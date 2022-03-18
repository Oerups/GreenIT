//Create class which will read the file and return the data
package main

import (
	"bufio"
	"os"
	"strings"
)

func getData(filename string) map[string][]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	regions := []string{}
	departments := []string{}
	municipalities := []string{}
	reader.ReadString('\n')

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		regions = append(regions, splited[3])
		departments = append(departments, splited[1])
		municipalities = append(municipalities, splited[0])
	}

	//create map with regions and departments and municipalities
	toReturn := make(map[string][]string)
	toReturn["regions"] = RemoveDuplicate(regions)
	toReturn["departments"] = RemoveDuplicate(departments)
	toReturn["municipalities"] = RemoveDuplicate(municipalities)

	return toReturn
}

func search(regions string, departments string, municipalities string) []map[string]string {
	if municipalities != "" {
		return findDataLine("data.csv", municipalities, 0)
	}

	if departments != "" {
		return findDataLine("data.csv", departments, 1)
	}

	if regions != "" {
		return findDataLine("data.csv", regions, 3)
	}

	return nil
}

func findDataLine(filename string, needle string, index int) []map[string]string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	headers := []string{}
	toReturn := []map[string]string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		splited := strings.Split(line, ";")
		if len(headers) == 0 {
			headers = splited
		}
		if strings.Contains(splited[index], needle) {
			//Create associative array from headers and splited
			data := make(map[string]string)
			for i := 0; i < len(headers); i++ {
				data[strings.Trim(headers[i], "\n\r")] = strings.Trim(splited[i], "\n\r")
			}
			toReturn = append(toReturn, data)
		}
	}
	return toReturn
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
