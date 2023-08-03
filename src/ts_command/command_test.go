package ts_command

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

// 执行程序返回 standard output and standard error
func TestCommandCombinedOutput(t *testing.T) {
	cmd := exec.Command("ls", "-lah", "ttttt")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

// 执行程序返回standard output
func TestCommandOutput(t *testing.T) {
	out, err := exec.Command("date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
}

// 用buffer接受输出,输入
func TestSetInputOutputSource(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	var stdin, stdout, stderr bytes.Buffer
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := stdout.String(), stderr.String()
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}

// print error and output to screen.
func TestPrintOutput(t *testing.T) {
	cmd := exec.Command("ls", "-lah")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

// cmd.Run() 阻塞等待命令执行结束
// cmd.Start() 不会等待命令完成
func TestExecAsync(t *testing.T) {
	var stdoutBuf, stderrBuf bytes.Buffer
	cmd := exec.Command("bash", "-c", "for i in 1 2 3 4;do echo $i;sleep 2;done")
	stdoutIn, _ := cmd.StdoutPipe()
	stderrIn, _ := cmd.StderrPipe()
	stdout := io.MultiWriter(os.Stdout, &stdoutBuf)
	stderr := io.MultiWriter(os.Stderr, &stderrBuf)
	err := cmd.Start()
	if err != nil {
		log.Fatalf("cmd.Start() failed with '%s'\n", err)
	}

	var errStdout, errStderr error
	go func() {
		_, errStdout = io.Copy(stdout, stdoutIn)
	}()
	go func() {
		_, errStderr = io.Copy(stderr, stderrIn)
	}()

	err = cmd.Wait()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	if errStdout != nil || errStderr != nil {
		log.Fatal("failed to capture stdout or stderr\n")
	}
	outStr, errStr := stdoutBuf.String(), stderrBuf.String()
	fmt.Printf("\nout:\n%s\nerr:\n%s\n", outStr, errStr)
}

// 执行时带上环境变量
func TestRunWithEnv(t *testing.T) {
	cmd := exec.Command("bash", "-c", "$programToExecute")
	additionalEnv := "programToExecute=ls"
	newEnv := append(os.Environ(), additionalEnv)
	cmd.Env = newEnv
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("%s", out)
}

// 预先检查命令是否存在, equals which command.
func TestCheckLsExists(t *testing.T) {
	path, err := exec.LookPath("ls")
	if err != nil {
		fmt.Printf("didn't find 'ls' executable\n")
	} else {
		fmt.Printf("'ls' executable is in '%s'\n", path)
	}
}

// 两个命令依次执行，管道通信
func TestPipBetweenCommands(t *testing.T) {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")
	r, w := io.Pipe()
	c1.Stdout = w
	c2.Stdin = r
	var b2 bytes.Buffer
	c2.Stdout = &b2
	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
	io.Copy(os.Stdout, &b2)
}

// 两个命令依次执行，管道通信 2
func TestPipBetweenCommands2(t *testing.T) {
	c1 := exec.Command("ls")
	c2 := exec.Command("wc", "-l")
	c2.Stdin, _ = c1.StdoutPipe()
	c2.Stdout = os.Stdout
	_ = c2.Start()
	_ = c1.Run()
	_ = c2.Wait()
}

// 不能直接用管道符，在 command 中
func TestPipeBetweenCommands3(t *testing.T) {
	c := exec.Command("ls", "|", "wc", "-l")
	c.Stdout = os.Stdout
	_ = c.Run()
}

// 不嫌丑可以用bash -c
func TestPipeBetweenCommands4(t *testing.T) {
	cmd := "cat /proc/cpuinfo | egrep '^model name' | uniq | awk '{print substr($0, index($0,$4))}'"
	out, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		fmt.Printf("Failed to execute command: %s", cmd)
	}
	fmt.Println(string(out))
}

// 按行读取输出内容
func TestReadByLine(t *testing.T) {
	cmd := exec.Command("ls", "-la")
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	reader := bufio.NewReader(stdout)
	for {
		//一次获取一行,_ 获取当前行是否被读完
		line, err := reader.ReadString('\n')
		// output, _, err := outputBuf.ReadLine()
		line = strings.TrimSpace(line)
		// 判断是否到文件的结尾了否则出错
		if err != nil || io.EOF == err {
			break
		}
		log.Println(line)
	}
	cmd.Wait()
}

// 获得exit code
func TestGetExitCode(t *testing.T) {
	name, args := "ls", []string{"-alrth", "/ssss"}
	defaultFailedCode := -1
	var stdout, stderr string
	var exitCode int

	log.Println("run command:", name, args)

	var outbuf, errbuf bytes.Buffer
	cmd := exec.Command(name, args...)
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf

	err := cmd.Run()
	stdout = outbuf.String()
	stderr = errbuf.String()

	if err != nil {
		// try to get the exit code
		if _, ok := err.(*exec.ExitError); ok {
		// if exitError, ok := err.(*exec.ExitError); ok {
			log.Println("is exit error")
			// ws := exitError.Sys().(syscall.WaitStatus)
			ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Printf("Could not get exit code for failed program: %v, %v", name, args)
			exitCode = defaultFailedCode
			if stderr == "" {
				stderr = err.Error()
			}
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}
	log.Printf("command result, stdout: %v, stderr: %v, exitCode: %v", stdout, stderr, exitCode)
}
