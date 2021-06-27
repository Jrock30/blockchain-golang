package main

import (
	"fmt"
	"github.com/jrock30/coin-basic/blockchain"
	"log"
	"net/http"
	"text/template"
)

const port string = ":4000"

type homeData struct {
	PageTitle string // 대문자를 써주어야 랜더에서도 사용가능.
	Blocks []*blockchain.Block
}

// 하나는 point 하나는 아니다
func home (writer http.ResponseWriter, request *http.Request) {
	// 데이터를 출력하긴 하지만 console 이 아닌 Writer 에 출력 (단순하게 현재는 텍스트만 출력)
	//fmt.Fprintf(writer, "hello from home!!")

	// 템플릿 렌더링
	// GO는 따로 Exception 이 없기 때문에 직접 에러 처리를 해주어야한다.
	//tmpl, err := template.ParseFiles("templates/home.gohtml")
	//if err != nil {
	//	log.Fatal(err) // os.Exit(1)
	//}

	// 에러 처리 Must function 이 자동으로 해준다.
	//tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	// ParseGlob 패턴을 보낼수 있다.
	tmpl := template.Must(template.ParseGlob("templates/home.gohtml"))
	data := homeData {"Home", blockchain.GetBlockchain().AllBLocks()}
	tmpl.Execute(writer, data)

}

/*
  Main Package (Entry Point)
*/
func main() {
	// black chain START

	//chain := blockchain.GetBlockchain()
	//chain.AddBlock("Second Block")
	//chain.AddBlock("Third Block")
	//chain.AddBlock("Fourth Block")
	//for _, block := range chain.AllBLocks() {
	//	fmt.Printf("Data: %s\n", block.Data)
	//	fmt.Printf("Hash: %s\n", block.Hash)
	//	fmt.Printf("PrevHash: %s\n", block.PrevHash)
	//}

	// black chain END

	// web server start

	// 경로 패턴 및 handler function
	http.HandleFunc("/", home)
	fmt.Printf("Listening on http://localhost:%s\n", port)
	// 서버 오픈
	log.Fatal(http.ListenAndServe(port, nil))
	// web server end
}
