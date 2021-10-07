package cli

import (
	"flag"
	"fmt"
	"github.com/jrock30/coin-basic/explorer"
	"github.com/jrock30/coin-basic/rest"
	"os"
)

func usage() {
	fmt.Printf("Please use following flags:\n\n")
	fmt.Printf("-port:   Set the PORT of the server\n")
	fmt.Printf("-mode:   Choose between 'html' and 'rest'\n\n")
	os.Exit(0) // 프로그램 종료
}

func Start() {
	// 커맨드를 받아서 실행
	//fmt.Println(os.Args[2:])

	// 기본 값 rest, 4000 으로 그냥 실행할 것 이면 아래 조건 생략
	if len(os.Args) == 1 {
		usage()
	}

	// corba 같은 것 으로 flag 를 더 유용하게 사용할 수 있다. 기본은 flag 가 정의 되지 않음
	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	default:
		usage()
	}
	// go run main.go -mode html -port 5000
	fmt.Println(*port, *mode)
}