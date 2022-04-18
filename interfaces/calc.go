/* https://stepik.org/lesson/350788/step/10?unit=334666 */

package main

import (
	"fmt"           // пакет используется для проверки ответа, не удаляйте его
    "errors"
)

func readTask() (interface{}, interface{}, interface{}) {
	var a, b float64
	a = 32.23
	b = 2.6
	return a, b, "*"
}

func isTypeFloat(value interface{}) error {
	switch t := value.(type) {
	case float64:
		return nil
	default:
		fmt.Print("value=", value, ": ")
		fmt.Printf("%T\n", t)
		return errors.New("error in isTypeInt")
	}
}

func doOperation(value1, value2 float64, operation string) (float64, error) {
	switch operation {
	case "*":
		return value1 * value2, nil
	case "+":
		return value1 + value2, nil
	case "-":
		return value1 - value2, nil
	case "/":
		return value1 / value2, nil
	default:
		return -1, errors.New("неизвестная операция")
	}

}

func checkOperation(operation interface{}) error {
	switch operation.(type) {
	case string:
		return nil
	default:
		return errors.New("неизвестная операция")
	}
}

func main() {
	value1, value2, operation := readTask() // исходные данные получаются с помощью этой функции
                                            // все полученные значения имеют тип пустого интерфейса

	if isTypeFloat(value1) != nil || isTypeFloat(value2) != nil {
		return
	}

	if err := checkOperation(operation); err != nil {
		fmt.Println(err)
		return
	}

	value, err := doOperation(value1.(float64), value2.(float64), operation.(string))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%.4f\n", value)

}