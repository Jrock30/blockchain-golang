package explorer

import (
	"fmt"
	"github.com/jrock30/coin-basic/blockchain"
	"log"
	"net/http"
	"text/template"
)

const (
	//port string        = ":4000"
	templateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	PageTitle string // 대문자를 써주어야 랜더에서도 사용가능.
	Blocks    []*blockchain.Block
}

// 하나는 point 하나는 아니다
func home(rw http.ResponseWriter, r *http.Request) {
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
	//tmpl := template.Must(template.ParseGlob("templates/pages/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBLocks()}
	//tmpl.Execute(writer, data)
	templates.ExecuteTemplate(rw, "home", data)

}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()                   // form 을 파싱한다
		data := r.Form.Get("blockData") // form 의 blockData 를 가져온다.
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/", http.StatusPermanentRedirect)
	}
}

func Start(port int) {
	handler := http.NewServeMux() // ServeMux 는 url(/) 과 함수(home) 를 연결해주는 역할을 한다.
	// .Must 는 error 에 대해서 우리가 직접 확인하지 않도록 해준
	// 패턴으로 가져온다 standard library(**/* 는 불가능하다.)
	templates = template.Must(template.ParseGlob(templateDir + "pages/*.gohtml"))
	// 두번째 패턴 가져오기
	templates = template.Must(templates.ParseGlob(templateDir + "partials/*.gohtml"))
	// 경로 패턴 및 handler function
	//http.HandleFunc("/", home)
	//http.HandleFunc("/add", add)
	/**
	DefaultServeMux 를 사용하는 것이 아닌 커스텀 ServeMux 를 사용
	*/
	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	fmt.Printf("Listening on http://localhost:%d\n", port)
	// 서버 오픈
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
