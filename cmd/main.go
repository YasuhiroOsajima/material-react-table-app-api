package main

import "github.com/YasuhiroOsajima/material-react-table-app-api/internal/infrastructure"

func main() {
	router := infrastructure.NewRouter()
	router.Run()
}
