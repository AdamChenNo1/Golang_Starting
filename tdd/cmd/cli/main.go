/*
 * File: /cmd/cli/main.go                                                      *
 * Project: tdd                                                                *
 * Created At: Friday, 2022/06/24 , 14:04:52                                   *
 * Author: elchn                                                               *
 * -----                                                                       *
 * Last Modified: Friday, 2022/06/24 , 16:12:05                                *
 * Modified By: elchn                                                          *
 * -----                                                                       *
 * HISTORY:                                                                    *
 * Date      	By	Comments                                                   *
 * ----------	---	---------------------------------------------------------  *
 */
package main

import (
	"bufio"
	"fmt"
	"go_start/tdd/config"
	"go_start/tdd/model"
	"go_start/tdd/store"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	fstore, err := store.FileSystemPlayerStoreFromFile(config.DbFileName)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := NewCLI(fstore, os.Stdin)

	game.PlayPoker()
}

type CLI struct {
	playStore model.PlayerStore
	in        *bufio.Scanner
}

func NewCLI(store model.PlayerStore, in io.Reader) *CLI {
	return &CLI{
		playStore: store,
		in:        bufio.NewScanner(in),
	}
}

func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.playStore.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()

	return cli.in.Text()
}
