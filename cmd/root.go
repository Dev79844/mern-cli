package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"path/filepath"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mern-cli",
	Short: "A CLI tool to set up a MERN project",
	Long:  `mern-cli is a CLI tool to set up a MERN (MongoDB, Express.js, React.js, Node.js) project with ease.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a node and react app",
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the name of the project:")
		projectName, _ := reader.ReadString('\n')
		projectName = strings.TrimSpace(projectName)
		fmt.Print("Which package manager do you want to use (npm/yarn)?")
		packageManager, _ := reader.ReadString('\n')
		packageManager = strings.TrimSpace(packageManager)

		// Build frontend
		installDepsFrontend(projectName)

		// Create backend directory
		backendDir := filepath.Join(projectName + "-backend")
		err := os.MkdirAll(backendDir, 0755)
		if err != nil {
			fmt.Println("Error creating backend directory:", err)
			return
		}

		// Install backend dependencies
		installDepsBackend(backendDir)

		// Initialize GitHub repository for backend
		cmdGit := exec.Command("git", "--version")
		if err := cmdGit.Run(); err == nil {
			cmdGitBackend := exec.Command("git", "init", backendDir)
			cmdGitBackend.Run()
		}
	},
}

func installDepsFrontend(projectName string) {
	cmdFrontend := exec.Command("npm", "create", "vite@latest", projectName+"-frontend", "--", "--template", "react")
	_, err := cmdFrontend.CombinedOutput()
	if err != nil {
		fmt.Println("Error initializing frontend:", err)
		return
	}
	fmt.Println("Initialized frontend")
}

func installDepsBackend(backendDir string) {
	err := os.Chdir(backendDir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	cmdBackend := exec.Command("npm", "init", "-y")
	if err := cmdBackend.Run(); err != nil {
		fmt.Println("Error initializing npm:", err)
		return
	}

	cmdInstall := exec.Command("npm", "install", "express", "dotenv", "morgan", "mongoose")
	if err := cmdInstall.Run(); err != nil {
		fmt.Println("Error installing packages:", err)
		return
	}
	fmt.Println("Initialized backend")
}

func init() {
	rootCmd.AddCommand(createCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
