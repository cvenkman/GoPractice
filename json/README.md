Формат json имеет типы данных. Четыре примитивных: строка, число, логический, null; и два структурных типа: объект и массив. Вот пример json кода с четырьмя полями разных типов:
```json
{
	"name":"qwerty",
	"price":258.25,
	"active":true,
	"description":null,
}
```

Json содержит массив объектов с двумя полями числового типа:
```json
[
	{"id":1,"price":2.58},
	{"id":2,"price":7.15}
]
```

### Аннотирование структур
```go
type myStruct struct {
	// при кодировании / декодировании будет использовано то же имя (Age),
	// но если значение поля равно 0 (пустое значение: false, nil, пустой слайс и пр.),
	// то при кодировании оно будет опущено, т.к. установлен тег "omitempty"
	Age int `json:",omitempty"`
	
	// при кодировании / декодировании поле всегда игнорируется
	Status bool `json:"-"`
}
```

Неэкспортируемые поля (имена которых начинаются со строчной буквы) не участвуют в кодировании / декодировании.


### Типы Encoder и Decoder
```go
type testStruct struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	src = testStruct{Name: "John Connor", Age: 35} // структура с данными
	dst = testStruct{}                             // структура без данных
	buf = new(bytes.Buffer)                        // буфер для чтения и записи
)

enc := json.NewEncoder(buf)
dec := json.NewDecoder(buf)

enc.Encode(src)
dec.Decode(&dst)

fmt.Print(dst) // {John Connor 35}
```

Во-первых, создали объекты Encoder и Decoder с помощью функций NewEncoder и NewDecoder соответственно. Каждая из этих функций получила в качестве аргумента буфер, который удовлетворяет одновременно и интерфейсу io.Reader, и интерфейсу io.Writer, соответственно мы смогли сначала записать в него данные, а затем их прочитать, используя методы Encode и Decode соответственно.

[десериализация JSON с неправильной типизацией](https://habr.com/ru/post/502176/)