package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	version := "latest"
	if len(os.Args) > 1 {
		version = os.Args[1]
	}

	url := "https://example.com/updates/bootstrap-" + version
	tmpPath := filepath.Join(os.TempDir(), "bootstrap")
	targetPath := "/home/sstanton/bootstrap"

	fmt.Println("Downloading:", url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Download failed:", err)
		return
	}
	defer resp.Body.Close()

	out, err := os.Create(tmpPath)
	if err != nil {
		fmt.Println("Failed to create temp file:", err)
		return
	}
	_, err = io.Copy(out, resp.Body)
	out.Close()
	if err != nil {
		fmt.Println("Failed to save file:", err)
		return
	}

	fmt.Println("Killing existing process...")
	exec.Command("pkill", "-f", "bootstrap").Run()
	time.Sleep(2 * time.Second)

	_ = os.Remove(targetPath)
	err = os.Rename(tmpPath, targetPath)
	if err != nil {
		fmt.Println("Failed to move new binary into place:", err)
		return
	}
	_ = os.Chmod(targetPath, 0755)

	fmt.Println("Restarting:", targetPath)
	cmd := exec.Command(targetPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	_ = cmd.Start()
}
