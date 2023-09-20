package hw4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Editor() {
	editorSearch(inputEditor())
}

func inputEditor() []string {
	scn := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter Lines and press Ctrl+] then press Enter:")
	var lines []string
	for scn.Scan() {
		line := scn.Text()
		if len(line) == 1 {
			if line[0] == '\x1D' {
				break
			}
		}
		lines = append(lines, line)
	}

	if len(lines) > 0 {
		fmt.Println()
		fmt.Println("Result:")
		for _, line := range lines {
			fmt.Println(line)
		}
		fmt.Println()
	}

	if err := scn.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return lines
}

func editorSearch(lines []string) {
	var search_line string
	fmt.Println("Enter what to search:")
	fmt.Scanf("%s", &search_line)

	for idx, line := range lines {
		if strings.Contains(line, search_line) {
			fmt.Printf("Line %v contains '%v': %v\n", idx, search_line, line)
		}
	}
}

type journal struct {
	grade int
	name  string
}

var students = []string{"Patrik", "Sponge Bob", "Mr. Krabs", "Squidward", "Sandy"}

func inputGrades() []journal {
	var grade int
	studentsGrades := make([]journal, len(students))
	for x, student := range students {
		fmt.Printf("Enter grade for %v ", student)
		fmt.Scan(&grade)
		studentsGrades[x] = journal{grade: grade, name: student}
		grade = 0
	}
	return studentsGrades
}

func averageGrade(studentsGrades []journal) {
	fmt.Printf("%v\n", studentsGrades)
	sum := 0
	for _, student := range studentsGrades {
		sum += student.grade
	}
	avg := float64(sum) / float64(len(studentsGrades))
	fmt.Printf("Average = %v", avg)
}

func StudentsGrades() {
	averageGrade(inputGrades())
}
