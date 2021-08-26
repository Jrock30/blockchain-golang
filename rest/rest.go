package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jrock30/coin-basic/blockchain"
	"github.com/jrock30/coin-basic/utils"
	"log"
	"net/http"
	"strconv"
)

//const port string = ":4000"
var port string

type url string

// MarshalText
/**
MarshalText interface 구현
시그니처가 틀리면 구현되지 않음 []byte, error
*/
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` //omitempty : 값이 있으면 보여주고 없으면 안보여주고
}

type addBlockBody struct {
	Message string
}

type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
}

/**
Stringer interface
 - String 하나의 메소드만 구현시킴
 - 대문자로 시작하는 String 이어야 하고, 매개변수를 받지 않고, string 을 return 해주어야 한다.
 - Go 에서는 모든 interface 가 내재적으로 구현돼 있다. 이 말은 Go 한테 Stringer interface 라고 말해줄 필요가 없다는 것이다.
 - 아래 처럼 작성하면 urlDescription struct 를 오버라이드 한다고 보면 된다.
*/
//func (u urlDescription) String() string {
//	return "Hello I`m the URL Description"
//}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{height}"),
			Method:      "GET",
			Description: "See A Block",
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

func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application/json")                // header json
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBLocks()) // Encode 가 Marshal 의 일을 해주고, 결과를 ResponseWrite 에 작성해준다.
	case "POST":
		var addBlockBody addBlockBody // addBlockBody(변수명) AddBlockBody(struct)
		/**
		 	json 을 decode 하고 addBlockBody 에 넣어준다. , & 포인터를 넣어주어야한다.
			utils.HandleErr 을 사용함으로써 에러처리를 한다.
		*/
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		fmt.Println(addBlockBody)
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
		//mux 를 쓰지 않으면 다른 http method 가 들어오면 아래의 처리를 해주어야함.
		//default:
		//	rw.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["height"])
	utils.HandleErr(err)                                  // error 핸들링을 직접 해주어야한다.
	block, err := blockchain.GetBlockchain().GetBlock(id) // 블록 하나를 가져온다.
	encoder := json.NewEncoder(rw)
	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})

	} else {
		encoder.Encode(block) // json 인코딩
	}
}

func Start(aPort int) {
	router := mux.NewRouter() // gorilla mux (URL Mapping 을 더 다양하게 사용할 수 있음)
	port = fmt.Sprintf(":%d", aPort)
	// API JSON
	//http.HandleFunc("/", documentation)
	//http.HandleFunc("/blocks", blocks)
	/**
	DefaultServeMux 를 사용하는 것이 아닌 커스텀 ServeMux 를 사용
	*/
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/block/{height:[0-9]+}", block).Methods("GET") // gorilla mux 라이브러리를 사용하면 param 값을 {} 및 정규식도 받을 수 있다. 기본 라이브러리는 안됨.
	fmt.Printf("Listening on http://localhost%s", port)
	fmt.Println()
	//log.Fatal(http.ListenAndServe(port, nil))
	log.Fatal(http.ListenAndServe(port, router))
}
