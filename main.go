package main

import (
	"PracticalProject/config"
	"PracticalProject/routes"
	"fmt"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {

	config.InitDB()

	fmt.Println("Hello World")
	r := routes.SetupRouter()

	r.Run(":8888")
}
