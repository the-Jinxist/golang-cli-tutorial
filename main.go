/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/the-Jinxist/golang-cli-tutorial/cmd"
	"github.com/the-Jinxist/golang-cli-tutorial/config"
)

func main() {

	config.InitDB()

	cmd.Execute()
}
