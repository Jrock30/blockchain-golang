package main

import "github.com/jrock30/coin-basic/cli"

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

	cli.Start()

}
