package main

import (
	"encoding/json"
	"fmt"
	"github.com/jrock30/coin-basic/utils"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL string
	Method string
	Description string
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: "/",
			Method: "GET",
			Description: "See Documentation",
		},
	}
	// Marshal - 메모리형식으로 저장된 객체를 저장/송신 할 수 있도록 변환해 준다. (Go -> JSON)
	// UnMarshal - JSON -> GO
	b, err := json.Marshal(data) // json 을 변환한 byte code 와 에러를 리턴함.
	// GO 는 에러를 콘솔에 뱉지 않으므로 에러를 보고 싶으면 아래처럼 해줘야함
	utils.HandleErr(err) // 사용자 정의 유틸 만듦
	fmt.Printf("%s", b) // Byte to String
}

/*
  Main Package (Entry Point)
*/
func main() {
	// 웹서버
	//explorer.Start()

	// API JSON
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
