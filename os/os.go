//author: liu.yang02@ucarinc.com
//date: 20190701

package vos

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"regexp"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/vogo/logger"
	vstrings "github.com/vogo/vogo/strings"
)

const (
	RootHome = "/root"
)

var (
	currentUserName string
)

// GetCurrentUserName get current user name
func GetCurrentUserName() string {
	if currentUserName == "" {
		u, err := user.Current()
		if err != nil {
			panic(err)
		}

		currentUserName = u.Username
	}

	return currentUserName
}

// PidExist whether pid exists
// see: https://stackoverflow.com/questions/15204162/check-if-a-process-exists-in-go-way
func PidExist(pid int) bool {
	p, err := os.FindProcess(pid)
	return err == nil && p.Signal(syscall.Signal(0)) == nil
}

// Kill process
func Kill(pid int) error {
	logger.Infof("kill process %d", pid)
	proc, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	if proc == nil {
		return nil
	}

	return proc.Kill()
}

// GetUserHome return user home.
func CurrUserHome() string {
	u, err := user.Current()
	if err == nil {
		return u.HomeDir
	}

	return RootHome
}

func GetUserHome(userName string) string {
	u, err := user.Lookup(userName)
	if err == nil {
		return u.HomeDir
	}

	return RootHome
}

func ExecShell(fullCommand string) ([]byte, error) {
	logger.Infof("exec: %s", fullCommand)
	cmd := exec.Command("/bin/sh", "-c", fullCommand)

	return cmd.CombinedOutput()
}

func ExecContext(ctx context.Context, fullCommand string) ([]byte, error) {
	logger.Infof("exec: %s", fullCommand)
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", fullCommand)

	return cmd.CombinedOutput()
}

// Shell execute shell without log
func Shell(fullCommand string) ([]byte, error) {
	cmd := exec.Command("/bin/sh", "-c", fullCommand)
	return cmd.CombinedOutput()
}

func SingleCommandResult(fullCommand string) (string, error) {
	output, err := ExecShell(fullCommand)
	if err != nil {
		return "", err
	}

	result := string(output)
	result = strings.Replace(result, "\r", "", -1)
	result = strings.Replace(result, "\n", "", -1)

	return result, nil
}

var (
	ErrPortNotFound = errors.New("port not found")
)

// GetPidByPort get process pid by port
// command example: netstat -anp|grep '8888 ' |grep 'LISTEN'|awk '{printf $7}'|cut -d/ -f1
func GetPidByPort(port int) (int, error) {
	fullCommand := fmt.Sprintf("lsof -iTCP:%d -sTCP:LISTEN -n -P |grep LISTEN | awk '{print $2}'", port)
	result, err := ExecShell(fullCommand)

	if err != nil {
		return 0, fmt.Errorf("command error: %+v", err)
	}

	if result == nil {
		return 0, fmt.Errorf("no process start at %d", port)
	}

	lines := bytes.Split(result, []byte{'\n'})
	pid := lines[len(lines)-1]

	if len(pid) == 0 && len(lines) > 1 {
		pid = lines[len(lines)-2]
	}

	logger.Infof("the pid of port %d is %s", port, pid)

	p, err := strconv.Atoi(string(pid))
	if err != nil {
		logger.Warnf("can't find pid for port %d, result: %s", port, pid)
		return -1, ErrPortNotFound
	}

	return p, nil
}

// GetProcessUser the the user of process by pid
// example: ps -o ruser -p 16787 | tail -1
func GetProcessUser(pid int) (string, error) {
	fullCommand := fmt.Sprintf("ps -o ruser -p %d | tail -1", pid)
	return SingleCommandResult(fullCommand)
}

func GetJavaHome(pid int) (string, error) {
	if !PidExist(pid) {
		return "", fmt.Errorf("no process for pid %d", pid)
	}

	fullCommand := fmt.Sprintf(`lsof -p %d \
|grep "/bin/java" \
|awk '{print $9}' \
|xargs ls -l \
|awk '{if($1~/^l/){print $11}else{print $9}}' \
|xargs ls -l \
|awk '{if($1~/^l/){print $11}else{print $9}}'`, pid)

	result, err := SingleCommandResult(fullCommand)
	if err != nil {
		return "", err
	}

	if result == "" {
		return "", errors.New("can't find java home")
	}

	if !strings.HasSuffix(result, "/bin/java") {
		return "", fmt.Errorf("can't get java home from path %s", result)
	}

	idx := strings.Index(result, "/bin/java")
	javaHome := result[:idx]

	return javaHome, nil
}

// ReadAllJavaProcessEnv
func ReadAllJavaProcessEnv() []map[string]string {
	var processes []map[string]string

	result, err := ExecShell("ps -o pid,cmd -e |grep java |grep -v grep")

	if err != nil {
		return nil
	}

	lines := bytes.Split(result, []byte{'\n'})
	count := len(lines)

	if count <= 0 {
		return nil
	}

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		for line[0] == ' ' {
			line = line[1:]
		}

		index := bytes.Index(line, []byte{' '})
		if index <= 0 {
			continue
		}

		pid := line[:index]
		javaProc := line[index+1:]

		if len(javaProc) > 0 {
			for javaProc[0] == ' ' {
				javaProc = javaProc[1:]
			}

			procEnv := ReadProcEnv(pid)
			procEnv["java_process"] = string(javaProc)

			processes = append(processes, procEnv)
		}
	}

	return processes
}

func ReadProcEnv(pid []byte) map[string]string {
	env := make(map[string]string)
	environData, err := ioutil.ReadFile(fmt.Sprintf("/proc/%s/environ", pid))

	if err != nil {
		return env
	}

	environ := bytes.Split(environData, []byte{0x0})

	for _, e := range environ {
		items := strings.SplitN(string(e), "=", 2)
		if len(items) > 1 && vstrings.ContainsAny(items[0], "JAVA", "JRE", "PATH", "CATALINA", "USER", "HOME") {
			env[items[0]] = items[1]
		}
	}

	return env
}

var (
	ignoreLoadEnvs = []string{"JAVA_OPTS", "CLASSPATH"}
)

func isLoadIgnoreEnv(e string) bool {
	for _, env := range ignoreLoadEnvs {
		if env == e {
			return true
		}
	}

	return false
}

func LoadUserEnv() {
	profiles := getUserEnvProfiles()
	for _, profile := range profiles {
		if _, err := os.Stat(profile); err != nil {
			continue
		}

		loadEnvFromProfile(profile)
	}

	adjustPathEnv()
}

func adjustPathEnv() {
	addEnvPathBin("/bin")
	addEnvPathBin("/sbin")
	addEnvPathBin("/usr/bin")
	addEnvPathBin("/usr/sbin")
	addEnvPathBin("/usr/local/bin")
	addEnvPathBin("/usr/local/sbin")
}

func addEnvPathBin(bin string) {
	path := os.Getenv("PATH")
	if !EnvPathContains(path, bin) {
		if err := os.Setenv("PATH", path+EnvValueSplit+bin); err != nil {
			logger.Warnf("set env error: %v", err)
		}
	}
}

func EnvPathContains(path, bin string) bool {
	return strings.HasPrefix(path, bin+EnvValueSplit) ||
		strings.Contains(path, EnvValueSplit+bin+EnvValueSplit) ||
		strings.HasSuffix(path, EnvValueSplit+bin)
}

func loadEnvFromProfile(profile string) {
	logger.Infof("load env from %s", profile)

	commandStr := fmt.Sprintf("source %s && env", profile)
	ctx, cancel := context.WithCancel(context.Background())
	cmd := exec.CommandContext(ctx, "/bin/sh", "-c", commandStr)

	go func() {
		time.Sleep(time.Second * 2)
		cancel()
	}()

	result, err := cmd.CombinedOutput()
	if err != nil {
		logger.Debugf("load env error: %v", err)
		return
	}

	reg := regexp.MustCompile(`[A-Za-z0-9]+=.*`)
	lines := bytes.Split(result, []byte{'\n'})

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		if !reg.Match(line) {
			continue
		}

		index := bytes.Index(line, []byte{'='})

		key := string(line[:index])
		if isLoadIgnoreEnv(key) {
			continue
		}

		logger.Debugf("set env: %s", line)
		err := os.Setenv(key, string(line[index+1:]))

		if err != nil {
			logger.Errorf("failed to set env: %v", err)
		}
	}
}
