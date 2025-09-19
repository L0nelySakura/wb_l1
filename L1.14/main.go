package main

import (
    "fmt"
)

func simpleTypeDetector(v interface{}) string {
    switch v.(type) {
    case int:
        return "int"
    case string:
        return "string"
    case bool:
        return "bool"
    case chan int, chan string, chan bool, chan interface{}:
        return "channel"
    default:
        return "unknown"
    }
}

func main() {
    items := []interface{}{100, "world", false, make(chan int)}
    
    for _, item := range items {
        fmt.Printf("%v is %s\n", item, simpleTypeDetector(item))
    }
}