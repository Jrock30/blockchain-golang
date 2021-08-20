package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// MarshalText
/**
	MarshalText interface 구현
	시그니처가 틀리면 구현되지 않음 []byte, error
 */
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type URLDescription struct {
	URL 		URL	   `json:"url"`
	Method 		string `json:"method"`
	Description string `json:"description"`
	Payload string `json:"payload,omitempty"` //omitempty : 값이 있으면 보여주고 없으면 안보여주고
}

/**
Stringer interface
 - String 하나의 메소드만 구현시킴
 - 대문자로 시작하는 String 이어야 하고, 매개변수를 받지 않고, string 을 return 해주어야 한다.
 - Go 에서는 모든 interface 가 내재적으로 구현돼 있다. 이 말은 Go 한테 Stringer interface 라고 말해줄 필요가 없다는 것이다.
 - 아래 처럼 작성하면 URLDescription struct 를 오버라이드 한다고 보면 된다.
*/
//func (u URLDescription) String() string {
//	return "Hello I`m the URL Description"
//}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL: URL("/"),
			Method: "GET",
			Description: "See Documentation",
		},
		{
			URL			: URL("/blocks"),
			Method		: "POST",
			Description	: "Add A Block",
			Payload		: "data:string",
		},

	}
	fmt.Println(data)
	rw.Header().Add("Content-Type", "application/json") // header json
	// Marshal - 메모리형식으로 저장된 객체를 저장/송신 할 수 있도록 변환해 준다. (Go -> JSON)
	// UnMarshal - JSON -> GO
	//b, err := json.Marshal(data) // json 을 변환한 byte code 와 에러를 리턴함.
	// GO 는 에러를 콘솔에 뱉지 않으므로 에러를 보고 싶으면 아래처럼 해줘야함
	//utils.HandleErr(err) // 사용자 정의 유틸 만듦
	//fmt.Printf("%s", b) // Byte to String
	//fmt.Fprintf(rw,"%s", b) // writer
	json.NewEncoder(rw).Encode(data) // 위의 3줄과 같은 효과
}

/*
  Main Package (Entry Point)
*/
func main() {
	// 웹서버
	//explorer.Start()

	// API JSON
	http.HandleFunc("/", documentation)
	//fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
