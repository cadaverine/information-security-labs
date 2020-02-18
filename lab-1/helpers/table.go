package helpers

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
				table[i][j] = " "
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

func (t *Table) FindRowIndex(char string, columnIndex int) int {
	for i := 0; i < len(*t); i++ {
		if (*t)[i][columnIndex] == char {
			return i
		}
	}

	return -1
}

func (t *Table) FindColumnIndex(char string, rowIndex int) int {
	for j := 0; j < len(*t); j++ {
		if (*t)[rowIndex][j] == char {
			return j
		}
	}

	return -1
}
