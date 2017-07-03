package main

import (
	"fmt"
)

type Fetcher interface {
	Data() (string, string)
}

// type App8 struct {
// 	ID         string
// 	Project_id string
// 	Type       string
// }

// func (a Apps8) Data() string {
// 	return r.length * r.width
// }

type Creds struct {
	ID           string
	Project_id   string
	Ssh_username string
	Ssh_password string
	Ssh_key      string
}

func (c Creds) Data() (string, string) {
	id := c.ID
	// pid := c.Project_id
	ssh_u := c.Ssh_username
	// ssh_p := c.Ssh_password
	// ssh_k := c.Ssh_key

	return id, ssh_u
	//return id, pid, ssh_u, ssh_p, ssh_k
}

func main() {

	//c := Creds{ID: "1", Project_id: "jghj", Ssh_username: "aman", Ssh_password: "12345", Ssh_key: "fte445"}
	c := Creds{ID: "1", Ssh_username: "aman"}
	p := [...]Fetcher{c}
	for n, _ := range p {
		fmt.Println("data is ", p[n].Data())
	}

}

// c := Creds{Ssh_username: "aman"}
// q := Square{side: 5}
// qwe := [...]Fetcher{c}

// fmt.Println("Looping through shapes for area ...")
// for n, _ := range qwe {
// 	fmt.Println("Shape details: ", qwe[n])
// 	fmt.Println("data is ", qwe[n].Data())
// }
