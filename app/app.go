package app

import (
	"block_chain/config"
	"block_chain/global"
	"block_chain/repository"
	"block_chain/service"
	. "block_chain/types"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
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

	a.service = service.NewService(a.repository, difficulty)

	sc := bufio.NewScanner(os.Stdin)
	useCase()

	for {

		from := global.FROM()
		if from != "" {
			a.log.Info("Current Conneted Wallet", "from", from)
			fmt.Println()
		}

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

		from := global.FROM()

		switch input[0] {
		case "1":
			fmt.Println("CreateWallet in Switch")
			if wallet := a.service.MakeWallet(); wallet == nil {
				panic("failed to create")
			} else {
				fmt.Println()
				a.log.Info("Success To Create wallet", "pk", wallet.PrivateKey, "pu", wallet.PublicKey)
				fmt.Println()
			}

		case "2":
			if from != "" {
				a.log.Debug("Already Connected", "from", from)
			} else {
				if wallet, err := a.service.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed To Find Wallet PK is Nil", "pk", input[1])
					} else {
						a.log.Crit("Filed To Find Wallet", "pk", input[1], "err", err)
					}
					fmt.Println()
				} else {
					global.SetFrom(wallet.PublicKey)
					a.log.Info("Success To Connect Wallet", "from", wallet.PublicKey)
					fmt.Println()
				}
			}

		case "3":
			if from == "" {
				a.log.Debug("Connect Wallet First")
			} else {
				if wallet, err := a.service.GetWallet(input[1]); err != nil {
					if err == mongo.ErrNoDocuments {
						a.log.Debug("Failed To Find Wallet PK is Nil", "pk", input[1])
					} else {
						a.log.Crit("Filed To Find Wallet", "pk", input[1], "err", err)
					}
					fmt.Println()
				} else {
					global.SetFrom(wallet.PublicKey)
					fmt.Println()
					if from == wallet.PublicKey {
						a.log.Debug("Already The Same Address", "from", wallet.PublicKey)
					} else {
						a.log.Info("Success To Change Wallet", "from", wallet.PublicKey)
					}
					fmt.Println()
				}
			}

		case "4":
			if from == "" {
				fmt.Println()
				a.log.Debug("Connect Wallet First")
				fmt.Println()
			} else {
				fmt.Println()
				fmt.Println("TransferCoin in Switch")
				a.service.CreateBlock([]*Transaction{}, []byte{}, from)
				fmt.Println()
			}

		case "5":

			if from == "" {
				a.log.Debug("Connect Wallet First")
			} else {
				fmt.Println("OppsCoin in Switch")
				a.service.CreateBlock([]*Transaction{}, []byte{}, from)
			}

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
	fmt.Println("2. ", ConnectWallet, "<PK>")
	fmt.Println("3. ", ChangeWallet, "<PK>")
	fmt.Println("4. ", TransferCoin, "<To> <Amount>")
	fmt.Println("5. ", OppsCoin, "<To> <Amount>")
	fmt.Println()
}
