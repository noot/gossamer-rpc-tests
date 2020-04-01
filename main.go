package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

var tx = []byte{3, 16, 110, 111, 111, 116, 1, 64, 103, 111, 115, 115, 97, 109, 101, 114, 95, 105, 115, 95, 99, 111, 111, 108}

func callRPCNoParams(method string) error {
	client := &http.Client{}

	data := []byte(`{"jsonrpc":"2.0","method":"` + method + `","params":{},"id":1}`)
	buf := &bytes.Buffer{}
	_, err := buf.Write(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8545", buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
       	return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)

	return nil
}

func callRPC(method, params string) error {
	client := &http.Client{}

	data := []byte(`{"jsonrpc":"2.0","method":"` + method + `","params":"` + params + `","id":1}`)

	buf := &bytes.Buffer{}
	_, err := buf.Write(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://localhost:8545", buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
       	return err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", body)

	return nil
}

func main() {
	err := callRPCNoParams("system_Health")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// err := callRPCNoParams("system_NetworkState")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// err := callRPC("system_Peers")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	err = callRPC("author_SubmitExtrinsic", "0x" + fmt.Sprintf("%x", tx))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}