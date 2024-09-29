package main

import (
	"fmt"
	"github.com/fanfaronDo/music_api/internal/app"
	"github.com/fanfaronDo/music_api/pkg/config"
)

func main() {
	cnf := config.ConfigLoad()

	if err := app.Run(cnf); err != nil {
		fmt.Println(err)
	}
}
