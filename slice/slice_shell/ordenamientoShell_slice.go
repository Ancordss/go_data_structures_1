package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	// Genera números aleatorios
	rand.Seed(time.Now().UnixNano())

	nums := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		nums[i] = rand.Intn(20000001) - 10000000
	}

	// Crea un archivo de texto y escribe los números
	file, err := os.Create("numeros.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for _, num := range nums {
		_, err := file.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Se han generado y almacenado en el archivo 'numeros.txt' un millón de números aleatorios.")

	// Ordena los números del archivo
	fmt.Println("Leyendo y ordenando los números del archivo...")

	start := time.Now()

	inputFile, err := os.Open("numeros.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer inputFile.Close()

	var numsFromFile []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		numsFromFile = append(numsFromFile, num)
	}

	shellSort(numsFromFile)

	// Escribe los números ordenados en un nuevo archivo
	outputFile, err := os.Create("numeros_ordenados.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	for _, num := range numsFromFile {
		_, err := outputFile.WriteString(strconv.Itoa(num) + "\n")
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Los números se han ordenado y almacenado en el archivo 'numeros_ordenados.txt'.")

	elapsed := time.Since(start)
	fmt.Printf("El tiempo de ejecución de la función de ordenamiento fue de %s.\n", elapsed)
}

func shellSort(nums []int) {
	// Genera la secuencia de intervalos de Shell
	n := len(nums)
	interval := 1
	for interval < n/3 {
		interval = interval*3 + 1
	}

	// Aplica el algoritmo de Shell
	for interval >= 1 {
		for i := interval; i < n; i++ {
			for j := i; j >= interval && nums[j] < nums[j-interval]; j -= interval {
				nums[j], nums[j-interval] = nums[j-interval], nums[j]
			}
		}
		interval /= 3
	}
}
