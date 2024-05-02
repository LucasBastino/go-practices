package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

func suma(valores []int) int {
	totalSuma := 0
	for _, valor := range valores {
		totalSuma += valor
	}
	return totalSuma
}

func generarValores() []int {
	valores := []int{}
	for i := 0; i < 10000000; i++ {
		valores = append(valores, i)
	}
	return valores
}

func BenchmarkSuma(b *testing.B) {
	valores := generarValores()
	totalSuma := suma(valores)
	fmt.Println(totalSuma)
}

func BenchmarkSumaSync(b *testing.B) {
	valores := generarValores()
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

func BenchmarkSumaSyncAtomic(b *testing.B) {
	valores := generarValores()
	cantidadSubprocesos := runtime.GOMAXPROCS(0)
	wg := sync.WaitGroup{}
	wg.Add(cantidadSubprocesos)
	var totalSuma int64
	for i := 0; i < cantidadSubprocesos; i++ {
		indice := i
		go func() {
			defer wg.Done()
			inicio := indice * (len(valores) / cantidadSubprocesos)
			fin := (indice + 1) * (len(valores) / cantidadSubprocesos)
			parcialSuma := suma(valores[inicio:fin])

			atomic.AddInt64(&totalSuma, int64(parcialSuma))
		}()
	}
	wg.Wait()
	fmt.Println(totalSuma)
}
