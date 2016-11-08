package main

import (
	"errors"
	"fmt"
	"log"
)

func Power(base float64, exponent int) (float64, error) {
	if equal(base, 0.0) && exponent < 0 {
		return 0.0, errors.New("Invalid Input")
	}
	absExponent := uint(exponent)
	if exponent < 0 {
		absExponent = uint(-exponent)
	}
	result := PowerWithUnsignedExponent(base, absExponent)
	if exponent < 0 {
		result = 1.0 / result
	}
	return result, nil
}

func PowerWithUnsignedExponent(base float64, exponent uint) float64 {
	if exponent == 0 {
		return float64(1)
	}
	if exponent == 1 {
		return base
	}
	result := PowerWithUnsignedExponent(base, exponent>>1)
	result *= result
	if exponent&1 == 1 {
		result *= base
	}
	return result
}

func equal(num1, num2 float64) bool {
	if num1-num2 > -0.0000001 && num1-num2 < 0.0000001 {
		return true
	}
	return false
}

func main() {
	result, err := Power(2, 4)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println(result)
}
