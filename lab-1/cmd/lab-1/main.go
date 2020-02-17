package main

import "fmt"

// Table - таблица для шифровки
type Table [][]string

// CreateTable - конструктор таблицы
func CreateTable(alphabet []string) Table {
	var table = make(Table, len(alphabet)+1)

	for i := 0; i <= len(alphabet); i++ {
		table[i] = make([]string, len(alphabet)+1)

		for j := 0; j <= len(alphabet); j++ {
			if i == 0 && j == 0 {
				table[i][j] = ""
			} else if i == 0 {
				table[i][j] = alphabet[j-1]
			} else if j == 0 {
				table[i][j] = alphabet[i-1]
			} else {
				table[i][j] = alphabet[(i+j-1)%(len(alphabet))]
			}
		}
	}

	return table
}

func (t Table) String() string {
	format := "%-3s"
	result := ""

	for _, row := range t {
		for _, item := range row {
			result += fmt.Sprintf(format, item)
		}

		result += fmt.Sprintln()
	}

	return result
}

func encode(input string) string {
	return ""
}

func decode(input string) string {
	return ""
}

func main() {
	table := CreateTable([]string{"a", "b", "c", "d"})

	fmt.Println(table)
}
