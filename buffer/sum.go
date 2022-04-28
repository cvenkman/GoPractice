package main

/*
 * Задача: на стандартный ввод подаются целые числа
 * каждое число подается на стандартный ввод с новой строки (количество чисел не известно).
 * Ввод заканчивается с пустой строкой.
 * Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
 * !! без "fmt"
 */

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func main() {
	err := run()
	if err != nil {
		os.Exit(1)
	}
}

func run() error {
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		sum += num
	}

	io.WriteString(os.Stdout, strconv.Itoa(sum))
	io.WriteString(os.Stdout, "\n")
	return nil
}
