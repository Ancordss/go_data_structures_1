package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		key := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > key {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = key
	}
}

func main() {
	// Genera números aleatorios
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		numbers[i] = rand.Intn(20000001) - 10000000
	}

	// Crea un archivo de texto y escribe los números
	file, err := os.Create("numeros.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, num := range numbers {
		_, err := file.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Se han generado y almacenado en el archivo 'numeros.txt' un millón de números aleatorios.")

	// Ordena los números del archivo
	fmt.Println("Leyendo y ordenando los números del archivo...")

	inputFile, err := os.Open("numeros.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		nums = append(nums, num)
	}

	insertionSort(nums)

	// Escribe los números ordenados en un nuevo archivo
	outputFile, err := os.Create("numeros_ordenados.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()
}
