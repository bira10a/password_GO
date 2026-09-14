package files

import (
	"fmt"
	"os"
)

func WriteFiles(content, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func ReadFiles(name string) {
	file, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("ReadFiles: ", string(file))
}
