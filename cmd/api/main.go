package main

import "cadigo-api/app"

func main() {
	err := app.NewApp()
	if err != nil {
		panic(err)
	}
}
