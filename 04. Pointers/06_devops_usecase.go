package main

import "fmt"

// ════════════════════════════════════════════════
// Server Configuration
// ════════════════════════════════════════════════
type ServerConfig struct {
	Name        string
	Environment string
	Port        int
	MemoryGB    int
	CPUCores    int
}

// ════════════════════════════════════════════════
// Display server configuration
// ════════════════════════════════════════════════
func displayConfig(config ServerConfig) {
	fmt.Printf("Server: %s | Env: %s | Port: %d | RAM: %dGB | CPU: %d cores\n",
		config.Name, config.Environment, config.Port, config.MemoryGB, config.CPUCores)
}

// ════════════════════════════════════════════════
// Update server resources (pointer)
// ════════════════════════════════════════════════
func updateResources(config *ServerConfig, ram int, cpu int) {
	config.MemoryGB = ram
	config.CPUCores = cpu
}

// ════════════════════════════════════════════════
// Change environment (pointer)
// ════════════════════════════════════════════════
func changeEnvironment(config *ServerConfig, newEnv string, newPort int) {
	config.Environment = newEnv
	config.Port = newPort
}

// ════════════════════════════════════════════════
// Get server info (multiple returns)
// ════════════════════════════════════════════════
func getServerInfo(config ServerConfig) (string, int, string) {
	name := config.Name
	totalResources := config.MemoryGB + config.CPUCores
	status := "Server info retrieved"

	return name, totalResources, status
}

// ════════════════════════════════════════════════
// Create config logger (closure)
// ════════════════════════════════════════════════
func createLogger(serverName string) func(string) {
	logCount := 0

	return func(message string) {
		logCount++
		fmt.Printf("[Log #%d] [%s] %s\n", logCount, serverName, message)
	}
}

func main() {
	fmt.Println("=== DevOps: Server Management System ===")

	// ════════════════════════════════════════════════
	// Initial server setup
	// ════════════════════════════════════════════════
	webServer := ServerConfig{
		Name:        "web-server-prod",
		Environment: "staging",
		Port:        8080,
		MemoryGB:    8,
		CPUCores:    4,
	}

	fmt.Println("--- Initial Configuration ---")
	displayConfig(webServer)

	// ════════════════════════════════════════════════
	// Create logger for this server (closure)
	// ════════════════════════════════════════════════
	log := createLogger(webServer.Name)
	log("Server initialized")

	// ════════════════════════════════════════════════
	// Update resources using pointer
	// ════════════════════════════════════════════════
	fmt.Println("\n--- Scaling Resources ---")
	log("Scaling up resources...")
	updateResources(&webServer, 16, 8)
	displayConfig(webServer)
	log("Resources scaled successfully")

	// ════════════════════════════════════════════════
	// Promote to production (pointer)
	// ════════════════════════════════════════════════
	fmt.Println("\n--- Promoting to Production ---")
	log("Changing environment to production")
	changeEnvironment(&webServer, "production", 443)
	displayConfig(webServer)
	log("Server promoted to production")

	// ════════════════════════════════════════════════
	// Get server info (multiple returns)
	// ════════════════════════════════════════════════
	fmt.Println("\n--- Server Summary ---")
	name, resources, status := getServerInfo(webServer)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Total Resources: %d\n", resources)
	fmt.Printf("Status: %s\n", status)

	// ════════════════════════════════════════════════
	// Multiple operations
	// ════════════════════════════════════════════════
	fmt.Println("\n--- Additional Operations ---")
	log("Performing health check")
	log("Backup completed")
	log("System ready")
}
