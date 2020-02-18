package main

import (
	"fmt"
	"strings"

	"github.com/cadaverine/information-security-labs/lab-1/helpers"
)

const (
	alphabet = "a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z"
)

func supplementString(input string, key string) string {
	keyLen := len(key)

	if keyLen == 0 {
		return ""
	}

	result := []rune{}
	keyArr := []rune(key)

	for i := range input {
		result = append(result, keyArr[i%keyLen])
	}

	return string(result)
}

func encode(input, key string, table *helpers.Table) string {
	if len(key) == 0 {
		return ""
	}

	length := len(input)
	result := make([]string, length)

	inpArr := strings.Split(input, "")
	keyArr := strings.Split(supplementString(input, key), "")

	for k := 0; k < length; k++ {
		i := table.FindRowIndex(inpArr[k], 0)
		j := table.FindRowIndex(keyArr[k], 0)

		result[k] = (*table)[i][j]
	}

	return strings.Join(result, "")
}

func decode(encoded, key string, table *helpers.Table) string {
	if len(key) == 0 {
		return ""
	}

	length := len(encoded)
	result := make([]string, length)

	inpArr := strings.Split(encoded, "")
	keyArr := strings.Split(supplementString(encoded, key), "")

	for k := 0; k < length; k++ {
		i := table.FindColumnIndex(keyArr[k], 0)
		j := table.FindRowIndex(inpArr[k], i)

		result[k] = (*table)[0][j]
	}

	return strings.Join(result, "")
}

func main() {
	key := "abc"
	input := "test input string"

	table := helpers.CreateTable(strings.Split(alphabet, ", "))

	fmt.Printf("Таблица для шифрования:\n\n%v\n", table)
	fmt.Printf("Ключ: '%s'\nВходящая строка: '%s'\n", key, input)

	encoded := encode(input, key, &table)
	fmt.Printf("Зашифрованная строка: '%s'\n", encoded)

	decoded := decode(encoded, key, &table)
	fmt.Printf("Расшифрованная строка: '%s'\n", decoded)
}
