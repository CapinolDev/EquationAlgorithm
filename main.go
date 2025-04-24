package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func formatNumber(n float64) string {
	if n == float64(int64(n)) {
		return fmt.Sprintf("%d", int64(n))
	}
	return fmt.Sprintf("%.2f", n)
}

func formatSide(xCoeff, constant float64) string {
	result := ""

	if xCoeff != 0 {
		if xCoeff == 1 {
			result += "x"
		} else if xCoeff == -1 {
			result += "-x"
		} else {
			result += fmt.Sprintf("%sx", formatNumber(xCoeff))
		}
	}

	if constant != 0 {
		if constant > 0 && xCoeff != 0 {
			result += "+"
		}
		result += formatNumber(constant)
	}

	if result == "" {
		result = "0"
	}

	return result
}
func parseSide(expression string) (xCoeff, constant float64) {

	if expression[0] != '-' {
		expression = "+" + expression
	}

	re := regexp.MustCompile(`([+-]?)(\d*)(x?)`)

	matches := re.FindAllStringSubmatch(expression, -1)

	for _, match := range matches {
		sign := match[1]
		num := match[2]
		isX := match[3] == "x"

		if num == "" {
			num = "1"
		}
		val, _ := strconv.ParseFloat(num, 64)
		if sign == "-" {
			val = -val
		}

		if isX {
			xCoeff += val
		} else {
			constant += val
		}
	}

	return
}
func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hewwo :3")
	fmt.Print("Input left side: ")
	Scanner.Scan()
	leftSide := Scanner.Text()
	fmt.Print("Input right side: ")
	Scanner.Scan()
	rightSide := Scanner.Text()
	leftX, leftConst := parseSide(leftSide)
	rightX, rightConst := parseSide(rightSide)
	fmt.Printf("%s = %s\n", formatSide(leftX, leftConst), formatSide(rightX, rightConst))
	denominator := leftX - rightX
	numerator := rightConst - leftConst

	if denominator == 0 {
		if numerator == 0 {
			fmt.Println("Infinite solutions :3")
		} else {
			fmt.Println("No solution :<")
		}
	} else {
		x := numerator / denominator
		fmt.Printf("x = %s\n", formatNumber(x))
		fmt.Println("Done :3")
	}
}
