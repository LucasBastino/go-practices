package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func suma(valores []int) int {
	totalSuma := 0
	for _, valor := range valores {
		totalSuma += valor
	}
	return totalSuma
}

func BenchmarkSuma(b *testing.B) {
	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}
	totalSuma := suma(valores)
	fmt.Println(totalSuma)
}

func BenchmarkSumaSync(b *testing.B) {
	valores := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36}
	cantidadSubprocesos := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	mt := sync.Mutex{}
	wg.Add(cantidadSubprocesos)
	totalSuma := 0
	for i := 0; i < cantidadSubprocesos; i++ {
		indice := i
		go func() {
			defer wg.Done()
			inicio := indice * (len(valores) / cantidadSubprocesos)
			fin := (indice + 1) * (len(valores) / cantidadSubprocesos)
			parcialSuma := suma(valores[inicio:fin])

			mt.Lock()
			totalSuma += parcialSuma
			mt.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(totalSuma)
}
