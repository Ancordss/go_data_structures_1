package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type HashTable struct {
	size  int
	table map[int]bool
}

func (h *HashTable) add(key int) {
	h.table[key] = true
}

func (h *HashTable) contains(key int) bool {
	_, ok := h.table[key]
	return ok
}

func main() {
	// Genera números aleatorios
	rand.Seed(time.Now().UnixNano())

	table := &HashTable{size: 0, table: make(map[int]bool)}
	for i := 0; i < 1000000; i++ {
		num := rand.Intn(20000001) - 10000000
		table.add(num)
	}

	// Crea un archivo de texto y escribe los números
	file, err := os.Create("numeros.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	for num := range table.table {
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

	var nums []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		nums = append(nums, num)
	}

	shellSort(nums)

	// Escribe los números ordenados en un nuevo archivo
	outputFile, err := os.Create("numeros_ordenados.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer outputFile.Close()

	for _, num := range nums {
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

// variación del algoritmo de inserción que utiliza una secuencia de brechas para comparar y mover los elementos más distantes. Esto permite que los elementos grandes se muevan más rápidamente a su posición correcta en la lista
func shellSort(nums []int) {
	n := len(nums)
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := nums[i]
			j := i
			for j >= gap && nums[j-gap] > temp {
				nums[j] = nums[j-gap]
				j -= gap
			}
			nums[j] = temp
		}
	}
}
