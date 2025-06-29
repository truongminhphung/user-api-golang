package main

import (
	"fmt"
	"user-api/global"
	"user-api/internal/initialize"
)

func main() {
	fmt.Println("User API Server is running...")
	r := initialize.Run()

	// Use the loaded configuration for server port
	serverAddr := fmt.Sprintf(":%d", global.Config.Server.Port)
	fmt.Printf("Server starting on port %d in %s mode\n", global.Config.Server.Port, global.Config.Server.Mode)

	r.Run(serverAddr)

	fmt.Printf("Visit http://localhost:%d to access the API\n", global.Config.Server.Port)
	fmt.Println("Press Ctrl+C to stop the server")
}
