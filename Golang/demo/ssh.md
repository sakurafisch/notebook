# ssh

```go
package main

import (
	"fmt"

	"golang.org/x/crypto/ssh"
)

func connectViaSsh(user, host string, password string) (*ssh.Client, *ssh.Session) {
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host, config)
	fmt.Println(err)

	session, err := client.NewSession()
	fmt.Println(err)

	return client, session
}

func main() {
	client, _ := connectViaSsh("root", "localhost:10022", "admin")
	client.Close()
}
```
