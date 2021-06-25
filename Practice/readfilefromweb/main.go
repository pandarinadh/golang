package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/"

func main() {

	resp, err := http.Get(url)

	checkError(err)

	fmt.Printf("Response type: %T\n", resp)

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	content := string(bytes)
	//fmt.Print(content)

	users := userFromJson(content)
	//	file, err := os.Create("./fromString.txt")
	//	checkError(err)

	for _, user := range users {

		fmt.Println(user)
		//		_, err := io.WriteString(file,user.Title)
		checkError(err)

		//fmt.Println(length)
	}

}

func userFromJson(content string) []User {
	users := make([]User, 0, 20)
	//var users []User
	fmt.Print(content)
	decoder := json.NewDecoder(strings.NewReader(content))
	//fmt.Println(decoder)
	_, err := decoder.Token()
	checkError(err)

	//fmt.Println(decoder.Token())

	var user User
	for decoder.More() {
		err := decoder.Decode(&user)
		checkError(err)
		fmt.Println(user)
		users = append(users, user)
	}

	/*
		json.Unmarshal([]byte(content), &users)

		fmt.Printf("%v", users)
	*/

	return users

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	UserId, Id int
	Title      string
	Completed  bool
}
