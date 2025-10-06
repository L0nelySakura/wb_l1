package main


import (
    "math/big"
    "fmt"
)


type BigCalculator struct {
    a *big.Int
    b *big.Int
}

func NewBigCalculator(a, b *big.Int) *BigCalculator {
	return &BigCalculator{
		a: new(big.Int).Set(a),
		b: new(big.Int).Set(b),
	}
}

func (bc *BigCalculator) Add() *big.Int {
	result := new(big.Int)
	return result.Add(bc.a, bc.b)
}

func (bc *BigCalculator) Subtract() *big.Int {
	result := new(big.Int)
	return result.Sub(bc.a, bc.b)
}

func (bc *BigCalculator) Multiply() *big.Int {
	result := new(big.Int)
	return result.Mul(bc.a, bc.b)
}

func (bc *BigCalculator) Divide() (*big.Int, error) {
	result := new(big.Int)
    if bc.b.Sign() == 0 {
        return nil, fmt.Errorf("деление на 0")
    }
	return result.Div(bc.a, bc.b), nil
}


func (bc *BigCalculator) SetValues(a, b *big.Int) {
	bc.a.Set(a)
	bc.b.Set(b)
}


func main() {
	a := new(big.Int)
	b := new(big.Int)
	
	a.SetString("1000000000000000000", 10) // 10^18
	b.SetString("50000", 10)  // 5 × 10^17
	
	// Создаем калькулятор
	calculator := NewBigCalculator(a, b)
	
	fmt.Println("Исходные значения:", calculator)
	fmt.Println()
	
	// Выполняем операции
	fmt.Printf("Сложение: %s + %s = %s\n", 
		calculator.a.String(), calculator.b.String(), calculator.Add().String())
	
	fmt.Printf("Вычитание: %s - %s = %s\n", 
		calculator.a.String(), calculator.b.String(), calculator.Subtract().String())
	
	fmt.Printf("Умножение: %s × %s = %s\n", 
		calculator.a.String(), calculator.b.String(), calculator.Multiply().String())
	

    val, err := calculator.Divide()
    if err == nil {
        fmt.Printf("Деление: %s ÷ %s = %s\n", 
            calculator.a.String(), calculator.b.String(), val)
    } else {
        fmt.Printf("Ошибка: %s.", err)
    }
}
