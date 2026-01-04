package main

import "fmt"

// ════════════════════════════════════════════════
// Server Configuration Structure
// ════════════════════════════════════════════════
type ServerConfig struct {
	Name     string
	CPUCores int
	MemoryGB int
	DiskGB   int
}

// ════════════════════════════════════════════════
// Upgrade server resources using POINTER
// ════════════════════════════════════════════════
func upgradeServer(config *ServerConfig, newCPU int, newRAM int) {
	// Modifies original server directly via pointer
	config.CPUCores = newCPU
	config.MemoryGB = newRAM
	// No return needed - original modified
}

// ════════════════════════════════════════════════
// Display server configuration
// ════════════════════════════════════════════════
func displayServer(config ServerConfig) {
	fmt.Printf("Server: %s | CPU: %d cores | RAM: %dGB | Disk: %dGB\n",
		config.Name, config.CPUCores, config.MemoryGB, config.DiskGB)
}

func main() {
	fmt.Println("=== DevOps: Server Configuration Management ===\n")

	// ════════════════════════════════════════════════
	// Initial server configuration
	// ════════════════════════════════════════════════
	prodServer := ServerConfig{
		Name:     "production-api",
		CPUCores: 4,
		MemoryGB: 8,
		DiskGB:   200,
	}

	fmt.Println("Initial configuration:")
	displayServer(prodServer)

	// ════════════════════════════════════════════════
	// First upgrade - high traffic detected
	// ════════════════════════════════════════════════
	fmt.Println("\n📈 Scaling up due to high traffic...")
	upgradeServer(&prodServer, 8, 16)
	displayServer(prodServer)

	// ════════════════════════════════════════════════
	// Second upgrade - more capacity needed
	// ════════════════════════════════════════════════
	fmt.Println("\n📈 Further scaling needed...")
	upgradeServer(&prodServer, 16, 32)
	displayServer(prodServer)

	fmt.Println("\n✅ Server scaled successfully!")
}
