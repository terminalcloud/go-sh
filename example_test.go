package sh

import (
	"fmt"
	"time"
)

// How to use sh

// Capture a bash command's exit code
func Example_a() {
	exitCode, err := WaitCode(RunBash("exit 10"))
	fmt.Println(exitCode, err)
	// output:
	// 10 <nil>

}

// Capture a bash command's exit code
func Example_b() {
	exitCode, err := WaitCode(RunBash("exit 0"))
	fmt.Println(exitCode, err)
	// output:
	// 0 <nil>
}

// Capture a program's exit code
func Example_c() {
	exitCode, err := WaitCode(RunCmd("bash", "-c", "exit 45"))
	fmt.Println(exitCode, err)
	// output:
	// 45 <nil>
}

// Capture other errors
func Example_d() {
	exitCode, _ := WaitCode(RunCmd("non-exisistent", "-c", "exit 45"))
	fmt.Println(exitCode)
	// output:
	// -1
}

// Timeout a long-running programs
func Example_e() {
	cmd, err := RunBash("sleep 5")
	exitCode, err := WaitCodeTimeout(cmd, err, 1*time.Second)
	fmt.Println(exitCode, err)
	// output:
	// -1 process timed out
}
