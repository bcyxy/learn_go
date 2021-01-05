package testpkg

import (
	"flag"
	"fmt"
	"log"
	"regexp"
	"time"

	expect "github.com/google/goexpect"
	"golang.org/x/crypto/ssh"
)

const (
	timeout = 1 * time.Second
)

var (
	addr  = flag.String("address", "127.0.0.1:22", "address of telnet server")
	user  = flag.String("user", "yxy", "username to use")
	pass1 = flag.String("pass1", "666666", "password to use")
	pass2 = flag.String("pass2", "666666", "alternate password to use")
)

func myExpect(gExpPtr *expect.GExpect) {
	var expMod = regexp.MustCompile(`asdf`)
	for {
		out, match, err := gExpPtr.Expect(expMod, timeout)
		if err != nil {
			//fmt.Printf("err:%v|%v\n", err, out)
			continue
		}
		fmt.Printf("==%v\n", out)
		fmt.Printf("--%v\n", match)
		fmt.Printf("\n")
	}
}

//TestSSH xxx
func TestSSH() {
	flag.Parse()

	sshClt, err := ssh.Dial("tcp", *addr, &ssh.ClientConfig{
		User:            *user,
		Auth:            []ssh.AuthMethod{ssh.Password(*pass1)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("ssh.Dial(%q) failed: %v", *addr, err)
	}
	defer sshClt.Close()

	gExpPtr, _, err := expect.SpawnSSH(sshClt, timeout)
	if err != nil {
		log.Fatal(err)
	}
	defer gExpPtr.Close()

	go myExpect(gExpPtr)

	gExpPtr.Send("ls\n")
	time.Sleep(time.Duration(1) * time.Second)
	gExpPtr.Send("echo 'asdfzxcv'\n")
	time.Sleep(time.Duration(1) * time.Second)
}
