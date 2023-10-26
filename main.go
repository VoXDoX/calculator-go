package main

import (
	"Calculator/funcions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение: ")
	input, _ := reader.ReadString('\n')
	str := strings.TrimSpace(input)

	calculate, err := funcions.ParseExpression(str)
	if err != nil {
		fmt.Println("опять все сломалось, ", err)
		return
	}
	condiner, err := calculate.WeConsider()
	if err != nil {
		fmt.Println("вышла ошибка, сломалось, ", err)
		return
	}
	fmt.Println("Результат: ", condiner)

}
