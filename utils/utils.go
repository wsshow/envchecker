package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"net"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func HashCode(s string) uint32 {
	return crc32.ChecksumIEEE([]byte(s))
}

func IsPathExist(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func CreatDir(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Chmod(dirPath, 0777)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(path string) ([]byte, error) {
	f, err := os.OpenFile(path, os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	return io.ReadAll(f)
}

func WriteFile(path string, data []byte) error {
	return os.WriteFile(path, data, 0600)
}

func NotExistToMkdir(dirPath string) {
	if !IsPathExist(dirPath) {
		CreatDir(dirPath)
	}
}

func Cmd(name string, arg ...string) (string, error) {
	result, err := exec.Command(name, arg...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(result)), nil
}

func CheckPort(host string, port int) bool {
	p := strconv.Itoa(port)
	addr := net.JoinHostPort(host, p)
	conn, err := net.DialTimeout("tcp", addr, 3*time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func RemoveWhitespace(s string) string {
	var rs []rune
	for _, r := range s {
		if r == ' ' || r == '\t' || r == '\n' || r == '\r' {
			continue
		}
		rs = append(rs, r)
	}
	return string(rs)
}

func TrimSuffix(s string) string {
	var rs []rune
	for _, r := range s {
		if r == '\n' || r == '\r' {
			continue
		}
		rs = append(rs, r)
	}
	return string(rs)
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func TrimUntilNum(s string) string {
	for index, length := 0, len(s); index < length; index++ {
		if IsDigit(s[index]) {
			return s[index:]
		}
	}
	return s
}

func GetFilenameFromUrl(url string) string {
	index := strings.LastIndex(url, "/")
	if index == -1 {
		return url
	}
	return url[index+1:]
}

func FindExecPath(name string) (execPath string, err error) {
	if runtime.GOOS == "windows" {
		execPath, err = Cmd("powershell", fmt.Sprintf("Get-Command -Name %s -ErrorAction SilentlyContinue", name))
	} else {
		execPath, err = Cmd("/bin/bash", fmt.Sprintf("which %s", name))
	}
	return
}

func GetFilename(filePath string) string {
	fileName := path.Base(filePath)
	fileExt := path.Ext(fileName)
	onlyFileName := strings.TrimSuffix(fileName, fileExt)
	return onlyFileName
}
