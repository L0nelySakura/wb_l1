package main

var justString string

func someFunc() {
    v := createHugeString(1 << 10)
    tmp := make([]byte, 100)
    copy(tmp, v[:100])
    justString = string(tmp)
	// Раньше justString ссылался на v, из-за этого она не могла быть собрана сборщиком мусора
	// теперь же, когда переменные независимы, v будет собрана сборщиком мусора
}

func main() {
    someFunc()
}