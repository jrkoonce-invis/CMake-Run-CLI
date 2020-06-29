package main

import (
	"bufio"
	"fmt"
	// "os"
	"os/exec"
	"strings"
)

func main() {

	args := []string{"-rf", "build"}
	err := exec.Command("rm", args...).Run()
	if err != nil {
		fmt.Printf("%s", err)
	}

	err = exec.Command("mkdir", "build").Run()
	if err != nil {
		fmt.Printf("%s", err)
	}

	cmd := exec.Command("cmake", "../")
	cmd.Dir = "./build"
	err = cmd.Run()
	if err != nil {
		fmt.Printf("%s", err)
	}

	args = []string{"--build", "./build"}
	err = exec.Command("cmake", args...).Run()
	if err != nil {
		fmt.Printf("%s", err)
	}

	/*
		temp, _ := os.Getwd()
		workingDir := strings.Split(temp, "/")
		directory := workingDir[len(workingDir)-1]
	*/

	out, err := exec.Command("cat", "CMakeLists.txt").Output()
	index := strings.Index(string(out), "project(")
	offset := strings.Index(string(out)[index:], ")")
	project := string(out)[index+8 : index+offset]

	cmd = exec.Command("./build/" + project)
	stdout, err := cmd.StdoutPipe()
	cmd.Start()

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	cmd.Wait()

	fmt.Println("Program Finished!")
}
