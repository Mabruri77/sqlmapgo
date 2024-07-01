package code

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ReplaceStringStar(str string, newValue string) {
	for i, value := range str {
		if value == '*' {
			fmt.Println(str[0:i] + newValue + str[i+1:])
		}
	}
}

func GetArgs() {
	fmt.Println(os.Args[1])
}

func ReadRequest(payload string) {
	link := "http://testphp.vulnweb.com/userinfo.php"
	data := url.Values{}
	data.Set("uname", "test"+payload)
	data.Set("pass", "test")
	req, err := http.NewRequest("POST", link, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()
}
