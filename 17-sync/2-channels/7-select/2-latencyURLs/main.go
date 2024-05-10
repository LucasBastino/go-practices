package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type result struct {
	url     string
	err     error
	latency time.Duration
}

func main() {
	// creo un canal que va a recibir y enviar los resultados de los GET
	resultsChan := make(chan result)
	// creo un canal que sirve de stopper, tras x cantidad de tiempo recibe una señal
	// (time.After directamente devuelve un canal de tipo Time)
	stopperChan := time.After(5 * time.Second)

	urls := []string{
		"https://amazon.com",
		"https://google.com",
		"https://nytimes.com",
		"https://wsj.com",
		"http://localhost:8085/wait",
	}

	for _, url := range urls {
		go checkURL(url, resultsChan)
	}

	// recibo un result del canal por cada url
	for range urls {
		select {
		// si se recibe bien o tiene un error hace esto ↓
		case result := <-resultsChan:
			if result.err != nil {
				log.Fatal(result.err.Error())
			} else {
				fmt.Printf("url: %s, %v\n", result.url, result.latency)
			}
			// si pasan mas de 5 segundos el stopper recibe la señal y ejecuta esto ↓
		case <-stopperChan:
			fmt.Println("timeout")
		}
	}

}

func checkURL(url string, resultsChan chan result) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// si hay error envia la url, el error y 0
		resultsChan <- result{url, err, 0}
	} else {
		// calcula el tiempo que paso desde que empezo (var start) hasta que entro a la URL
		latency := time.Since(start).Round(time.Millisecond)
		// envia la url, error nil y el tiempo transcurrido
		resultsChan <- result{url, nil, latency}
		resp.Body.Close()
	}

}

// se puede nombrar a los for, los while.. de esta manera:
// func syntax() {

// loop:
// 	for {
// 		select {
// 		case <-ch1:
// 			fmt.Printf("lalalala")
// 		case <-stopper:
// 			break loop

// 		}
// 	}
// }

// consejo:
// nunca poner un default en el select adentro de un loop, consume mucha CPU
