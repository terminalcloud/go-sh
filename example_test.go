package sh

import (
	"fmt"
)

// How to use sh

// Run different bash commands in various ways
func Example_a() {
	exitCode, err := WaitCode(RunBash("exit 10"))
	fmt.Println(exitCode, err)
	// output:
	// 10 <nil>

}

func Example_b() {
	exitCode, err := WaitCode(RunBash("exit 0"))
	fmt.Println(exitCode, err)
	// output:
	// 0 <nil>
}

func Example_c() {
	exitCode, err := WaitCode(RunCmd("bash", "-c", "exit 45"))
	fmt.Println(exitCode, err)
	// output:
	// 45 <nil>
}

func Example_d() {
	exitCode, _ := WaitCode(RunCmd("non-exisistent", "-c", "exit 45"))
	fmt.Println(exitCode)
	// output:
	// -1
}
