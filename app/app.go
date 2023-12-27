package app

import (
	"block_chain/config"
	"block_chain/repository"
	"block_chain/service"
	. "block_chain/types"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/inconshreveable/log15"
)

type App struct {
	config     *config.Config
	service    *service.Service
	repository *repository.Repository

	log log15.Logger
}

func NewApp(config *config.Config, difficulty int64) {
	a := &App{
		config: config,
		log:    log15.New("module", "app"),
	}

	var err error

	if a.repository, err = repository.NewRepository(config); err != nil {
		panic(err)
	}

	a.service = service.NewService(a.repository, 1)

	sc := bufio.NewScanner(os.Stdin)
	useCase()

	for {
		sc.Scan()
		input := strings.Split(sc.Text(), " ")
		if err = a.inputValueAssessment(input); err != nil {
			a.log.Error("Failed to call cli", "err", err, "input", input)
			fmt.Println()
		}
	}
}

func (a *App) inputValueAssessment(input []string) error {
	msg := errors.New("Check Use Case")
	if len(input) == 0 {
		return msg
	} else {
		switch input[0] {
		case CreateWallet:
			fmt.Println("CreateWallet in Switch")
			if wallet := a.service.MakeWallet(); wallet == nil {
				panic("failed to create")
			} else {
				fmt.Println("Success To Create wallet")
				fmt.Println(wallet.PrivateKey)
				fmt.Println(wallet.PublicKey)
			}

		case TransferCoin:
			fmt.Println("TransferCoin in Switch")

		case OppsCoin:
			fmt.Println("OppsCoin in Switch")

		case "":
			fmt.Println()

		default:
			return errors.New("failed to find cli order")
		}
	}
	return nil
}

func useCase() {
	fmt.Println()
	fmt.Println("This is Opps Module For BlockChain Core With Mongo")
	fmt.Println()
	fmt.Println("Use Case")
	fmt.Println()
	fmt.Println("1. ", CreateWallet)
	fmt.Println("2. ", TransferCoin)
	fmt.Println("3. ", OppsCoin)
	fmt.Println()
}
