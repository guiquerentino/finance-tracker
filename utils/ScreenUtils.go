package utils

import "C"
import (
	"os"
	"os/exec"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
