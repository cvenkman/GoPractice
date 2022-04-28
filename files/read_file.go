package rf

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	
	reader := bufio.NewReader(file)
	
	buf := make([]byte, 10)
	n, err := reader.Read(buf) // читаем в buf 10 байт из ранее открытого файла
	if err != nil && err != io.EOF {
		// io.EOF не совсем ошибка - это состояние, указывающее, что файл прочитан до конца
		panic(err)
	}
	fmt.Printf("прочитано %d байт: %s\n", n, buf) // прочитано 10 байт: bufio ...
	
	s, err := reader.ReadString('\n') // читаем данные до разрыва абзаца ('\n')
	fmt.Printf("%s\n", s)         // ... здесь будет строка
}