package files

import (
	"fmt"
	"os"
)

func WriteFiles(content []byte, name string) {
	file, err := os.Create(name)

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func ReadFiles(name string) ([]byte, error) {
	file, err := os.ReadFile(name)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("ReadFiles: ", string(file))
	return file, nil
}
