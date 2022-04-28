package main

import (
	"fmt"
	"log"
	"os"

	"bufio"
	"io"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func run() error {
	file, err := os.Create("file.txt")
	if err != nil { 
		return err
	}
	defer file.Close()

	text := "Hello Golang!\n"

	/* with standart read/write */
	write(file, text)
	read(file)

	/* with bufio */
	// writeBuff(file, text)
	// readBuff(file) // не работает

	return nil
}

func write(file *os.File, data string) error {
    _, err := file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}

func read(file *os.File) error {
	file, err := os.Open(file.Name())
    if err != nil{
        fmt.Println(err) 
        os.Exit(1) 
    }
    defer file.Close()

    data := make([]byte, 64)
    for {
        n, err := file.Read(data)
        if err == io.EOF {
            break
        }
        fmt.Print(string(data[:n]))
    }
	return nil
}

func writeBuff(file *os.File, data string) error {
	writer := bufio.NewWriter(file) // создаем поток вывода

	_, err := file.WriteString(data)
	if err != nil {
		return err
	}

	writer.Flush() // сбрасываем данные из буфера в файл
	return nil
}

func readBuff(file *os.File) error {
	reader := bufio.NewReader(file)

	for{
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Print("end")
			break
		}
		if err != nil {
			return err
		}
		fmt.Print(line)
	}	

	return nil
}