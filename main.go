package main

import (
	"flag"
	"time"

	"github.com/labstack/gommon/log"

	"github.com/AccumulatedFinance/metrics-api/api"
	"github.com/AccumulatedFinance/metrics-api/evm"
	"github.com/AccumulatedFinance/metrics-api/store"
)

const API_PORT = 8083
const STACME_ADDRESS = "0x7AC168c81F4F3820Fa3F22603ce5864D6aB3C547"

func main() {

	var InfuraKey string
	flag.StringVar(&InfuraKey, "k", "", "Infura API key")
	flag.Parse()

	client, err := evm.NewEVMClient(InfuraKey)
	if err != nil {
		log.Fatal(err)
	}

	die := make(chan bool)
	go getStats(client, die)

	log.Fatal(api.StartAPI(API_PORT))

}

func getStats(client *evm.EVMClient, die chan bool) {

	for {

		select {
		default:

			stACME, err := client.GetERC20TotalSupply(STACME_ADDRESS)
			if err != nil {
				log.Info(err)
			}

			store.StACME = stACME

			log.Info(store.StACME)

			now := time.Now()
			store.UpdatedAt = &now

			time.Sleep(time.Duration(10) * time.Minute)

		case <-die:
			return
		}

	}

}
