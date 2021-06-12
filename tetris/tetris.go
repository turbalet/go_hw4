package tetris

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var types = map[string]int {
	"int": 8,
	"int32": 4,
	"int64": 8,
	"uint": 8,
	"uint32": 4,
	"float32": 4,
	"float64" : 8,
	"uint64": 8,
	"rune": 4,
	"byte": 1,
	"bool": 1,
	"string": 16,
	"*": 8,
	"complex64": 8,
	"complex128": 16,
	"complex": 16,
}

func getType(field string) string {
	splitted := strings.Split(field, " ")
	//fmt.Println(splitted[len(splitted) - 1])
	return splitted[len(splitted) - 1]
}

func optimize(fields []string) []string{

	fmt.Println(fields)
	var isPointer1, isPointer2 bool
	var type1, type2 string
	for i := 0; i < len(fields) - 1; i++ {
		for j := 0; j < len(fields) - i - 1; j++ {
			if strings.Contains(fields[j+1], "*") {
				isPointer2 = true
			} else if strings.Contains(fields[j], "*") {
				isPointer1 = true
			}
			if isPointer1 {
				type1 = "*"
				//fmt.Println(type1)
			} else {
				type1 = getType(strings.TrimSpace(fields[j]))
			}
			if isPointer2 {
				type2 = "*"
				//fmt.Println(type2)
			} else {
				type2 = getType(strings.TrimSpace(fields[j + 1]))
			}
			if types[type1] > types[type2] {
				fields[j], fields[j+1] = fields[j+1], fields[j]
			}
			isPointer2 = false
			isPointer1 = false
		}
	}
	fmt.Println(fields)
	return fields
}

func Tetris(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var imports[] string
	isImports := true
	// appends imports to "imports"
	for scanner.Scan() && isImports {
		if strings.Contains(scanner.Text(), "struct {") {
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), "}") {
					isImports = false
					break
				} else {
					imports = append(imports, scanner.Text())
				}
			}
		}
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	// copy "imports"
	oldFields := make([]string, len(imports))
	copy(oldFields, imports)
	// sort "imports"
	imports = optimize(imports)

	// replace old imports with sorted
	newContents := strings.Replace(string(content), strings.Join(oldFields, "\n"),strings.Join(imports, "\n") , -1)

	err = ioutil.WriteFile(path, []byte(newContents), 0)
	if err != nil {
		log.Fatal(err)
	}
}