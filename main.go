package main

import (
	"flag"
	"fmt"
	"os"
)

func usage() {
	fmt.Printf("Please use following commands:\n\n")
	fmt.Printf("explorer:   Start the HTML Explorer\n")
	fmt.Printf("rest    :   Start the REST API (recommended)\n")
}

/*
  Main Package (Entry Point)
*/
func main() {
	/**
	port 가 달라도 URL 이 같으면 에러가 발생한다.
	Multiplexer 가 url 을 지켜보고 내가 원하는 함수를 실행해준다.
	Custom Multiplexer 를 만들수 있다. (각 각의 URL Handler 를 사용할 수 있다.)
	Start 들을 참고 할 것
	*/
	// WEB SERVER go 루틴
	//go explorer.Start(3000)

	// REST API SERVER
	//rest.Start(4000)

	// 커맨드를 받아서 실행

	fmt.Println(os.Args[2:])

	if len(os.Args) < 2 {
		usage()
		os.Exit(0) // 프로그램 종료
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)

	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		fmt.Println("Start REST API")
		rest.Parse(os.Args[2:])
	default:
		usage()
	}

	fmt.Println(*portFlag)
}
