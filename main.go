package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const ollamaURL = "http://localhost:11434/api/generate"

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
	Done     bool   `json:"done"`
}

func main() {
	model := flag.String("model", "gemma3:1b", "Ollama model name")
	flag.Parse()

	diff, err := getGitDiff()
	if err != nil {
		exitWithError("Failed to get git diff", err)
	}

	if diff == "" {
		fmt.Println("No changes to commit")
		return
	}

	commitMsg, err := generateCommitMessage(*model, diff)
	if err != nil {
		exitWithError("Failed to generate commit message", err)
	}

	fmt.Printf("Generated commit message:\n\n%s\n\n", commitMsg)

	if err := createCommit(commitMsg); err != nil {
		exitWithError("Failed to create commit", err)
	}

	if err := pushToRemote(); err != nil {
		exitWithError("Failed to push changes", err)
	}

	fmt.Println("Successfully committed and pushed changes!")
}

func getGitDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--staged")
	output, err := cmd.Output()
	return string(output), err
}

func generateCommitMessage(model, diff string) (string, error) {
	prompt := fmt.Sprintf(`Generate a concise git commit message in active voice following conventional commits standard. 
Focus on the change purpose rather than implementation details. 

Changes:
%s

Commit message:`, diff)

	requestData := OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, _ := json.Marshal(requestData)
	resp, err := http.Post(ollamaURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var responseBuilder strings.Builder
	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk OllamaResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil {
			return "", err
		}
		responseBuilder.WriteString(chunk.Response)
		if chunk.Done {
			break
		}
	}

	return cleanCommitMessage(responseBuilder.String()), nil
}

func cleanCommitMessage(msg string) string {
	msg = strings.TrimSpace(msg)
	if strings.Contains(msg, "\"") {
		parts := strings.Split(msg, "\"")
		if len(parts) > 1 {
			msg = parts[1]
		}
	}
	return strings.Split(msg, "\n")[0]
}

func createCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

func pushToRemote() error {
	cmd := exec.Command("git", "push")
	return cmd.Run()
}

func exitWithError(msg string, err error) {
	fmt.Fprintf(os.Stderr, "ERROR: %s: %v\n", msg, err)
	os.Exit(1)
}