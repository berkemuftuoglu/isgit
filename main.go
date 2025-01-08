package main

import (
	"flag"
	"fmt"
	"os/exec"
	"strings"
)

// isGitRepo checks if the directory is a Git repository.
func isGitRepo(path string) bool {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--is-inside-work-tree")
	err := cmd.Run()
	return err == nil
}

// getCurrentBranch returns the current branch name or provides a user-friendly error.
func getCurrentBranch(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "rev-parse", "--abbrev-ref", "HEAD")
	output, err := cmd.Output()
	if err != nil {
		if strings.Contains(err.Error(), "128") {
			return "", fmt.Errorf("the repository has no commits or no active branch")
		}
		return "", fmt.Errorf("an error occurred while fetching the branch name: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// getLastCommit retrieves the last commit message, hash, and date or handles no commit scenarios.
func getLastCommit(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "log", "-1", "--pretty=format:%h %s (%ci)")
	output, err := cmd.Output()
	if err != nil {
		if strings.Contains(err.Error(), "128") {
			return "", fmt.Errorf("the repository has no commits yet")
		}
		return "", fmt.Errorf("an error occurred while fetching the last commit: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// getRemoteURL fetches the remote repository URL or handles cases with no remote.
func getRemoteURL(path string) (string, error) {
	cmd := exec.Command("git", "-C", path, "remote", "get-url", "origin")
	output, err := cmd.Output()
	if err != nil {
		if strings.Contains(err.Error(), "128") {
			return "", fmt.Errorf("no remote repository configured for this repository")
		}
		return "", fmt.Errorf("an error occurred while fetching the remote URL: %v", err)
	}
	return strings.TrimSpace(string(output)), nil
}

func printHelp() {
	fmt.Println(`Usage: isgit [OPTIONS]

Checks if a directory is a Git repository and provides repository details.

Options:
  --path <path>      Path to check for a Git repository (default: current directory)
  --details          Show detailed repository information (branch, last commit, remote)
  --remote           Show only the remote repository details
  --help             Display this help message`)
}

func main() {
	// Define flags
	pathFlag := flag.String("path", ".", "Path to check for a Git repository")
	detailsFlag := flag.Bool("details", false, "Show detailed repository information")
	remoteFlag := flag.Bool("remote", false, "Show remote repository details only")
	helpFlag := flag.Bool("help", false, "Display this help message")

	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	path := strings.TrimSpace(*pathFlag)

	// Check if the directory is a Git repository
	if !isGitRepo(path) {
		fmt.Println("No, this is not a Git repository.")
		return
	}

	if !*detailsFlag && !*remoteFlag {
		fmt.Println("Yes, this is a Git repository.")
	}

	// Handle --remote flag
	if *remoteFlag {
		remoteURL, err := getRemoteURL(path)
		if err != nil {
			fmt.Println("Remote:", err)
		} else {
			fmt.Println("Remote:", remoteURL)
		}
		return
	}

	// Handle --details flag
	if *detailsFlag {
		branch, err := getCurrentBranch(path)
		if err != nil {
			fmt.Println("Error fetching branch name:", err)
		} else {
			fmt.Println("Current branch:", branch)
		}

		lastCommit, err := getLastCommit(path)
		if err != nil {
			fmt.Println("Error fetching last commit:", err)
		} else {
			fmt.Println("Last commit:", lastCommit)
		}

		remoteURL, err := getRemoteURL(path)
		if err != nil {
			fmt.Println("Remote:", err)
		} else {
			fmt.Println("Remote:", remoteURL)
		}
	}
}

