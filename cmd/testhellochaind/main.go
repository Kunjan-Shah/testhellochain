package main

import (
	"os"
	"log"
	// "fmt"
	"github.com/joho/godotenv"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	"github.com/ignite/cli/ignite/pkg/cosmoscmd"
	"github.com/ignite/cli/ignite/pkg/xstrings"
	"testhellochain/app"
)

func main() {
	enverr := godotenv.Load("secret.env")
	if enverr != nil {
		log.Fatal(enverr)
	}
	apiKey, exists := os.LookupEnv("API_KEY")
	apiKey = apiKey
	if !exists {
		log.Fatal("API_KEY does not exist in secret.env")
	}
	rootCmd, _ := cosmoscmd.NewRootCmd(
		app.Name,
		app.AccountAddressPrefix,
		app.DefaultNodeHome,
		xstrings.NoDash(app.Name),
		app.ModuleBasics,
		app.New,
		// this line is used by starport scaffolding # root/arguments
	)
	if err := svrcmd.Execute(rootCmd, "", app.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
