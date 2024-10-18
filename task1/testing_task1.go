package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func readData(fileName string) (string, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func runTest(input string) (string, error) {
	cmd := exec.Command("go", "run", "./task1.go")
	cmd.Stdin = strings.NewReader(input)

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func main() {
	directory := "./remove-digit/"

	files, err := ioutil.ReadDir(directory)
	if err != nil {
		fmt.Println("Упс... Ошибка чтения папки с тестам")
		return
	}

	for _, file := range files {
		testFile := directory + file.Name()

		if strings.HasSuffix(file.Name(), ".a") {
			continue
		}
		answerFile := testFile + ".a"

		input, err := readData(testFile)
		if err != nil {
			fmt.Println("Ошибка чтения теста:", err)
			continue
		}
		expect, err := readData(answerFile)
		if err != nil {
			fmt.Println("Ошибка чтения ответа:", err)
			continue
		}
		actual, err := runTest(input)
		if err != nil {
			fmt.Println("Ошибка выполнения программы:", err)
			continue
		}
		if strings.TrimSpace(actual) == strings.TrimSpace(expect) {
			fmt.Println(testFile, "пройден ^^")
		} else {
			fmt.Println(testFile, "не пройден :(")
			fmt.Println("Ожидалось:", expect)
			fmt.Println("Получено:", actual)
		}
	}
}
