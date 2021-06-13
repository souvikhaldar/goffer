package webfuzz

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func Fuzz(ip, port, command string, poolSize int) error {
	conn, err := net.DialTimeout("tcp", ip+":"+port, 1000000*time.Microsecond)
	if err != nil {
		return fmt.Errorf("Faied to dial: %v", err)
	}
	defer conn.Close()

	fmt.Println("Command: ", command)

	crashed := make(chan int)
	pool := make(chan int, poolSize)
	fmt.Println("Pool size: ", poolSize)

	done := make(chan bool)

	go func() {
		cd := <-crashed
		log.Println("***********")
		log.Printf("crashed at:%d bytes\n", cd)
		log.Println("***********")
		done <- true
		return

	}()

	notCrashed := true
	for n := 100; notCrashed; n += 100 {
		pool <- 1
		go func(n int, pool chan int) {
			defer func() {
				<-pool
			}()

			log.Println("Fuzzing: ", n)
			garbage := strings.Repeat("A", n)

			if err := conn.SetDeadline(time.Now().Add(1000000 * time.Microsecond)); err != nil {
				fmt.Printf("Can't set write deadline: %s", err)
				return
			}

			if _, err := conn.Write([]byte(command + garbage)); err != nil && notCrashed {
				notCrashed = false
				crashed <- n
				fmt.Printf("Can't write more because: %s", err)
				return
			}
			rcv := make([]byte, 2048)
			_, err = conn.Read(rcv)
			if err != nil && notCrashed {
				//log.Println()
				notCrashed = false
				crashed <- n
				fmt.Printf("Could not read because: %s", err)
				return
			}

		}(n, pool)

	}

	<-done
	return nil
}

func FuzzContent(ip, port, command, content string, poolSize int) error {

	conn, err := net.DialTimeout("tcp", ip+":"+port, 1000000*time.Microsecond)
	if err != nil {
		return fmt.Errorf("Faied to dial: %v", err)
	}
	defer conn.Close()

	fmt.Println("Command: ", command)

	if err := conn.SetDeadline(time.Now().Add(1000000 * time.Microsecond)); err != nil {
		fmt.Printf("Can't set write deadline: %s", err)
		return err
	}

	if _, err := conn.Write([]byte(command + content)); err != nil {
		fmt.Printf("Can't write more because: %s", err)
		return err
	}
	rcv := make([]byte, 2048)
	_, err = conn.Read(rcv)
	if err != nil {
		//log.Println()
		fmt.Printf("Could not read because: %s", err)
		return err
	}
	return nil
}
