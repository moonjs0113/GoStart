package GoLangBasic

func Package() {
	println("하나의 폴더에 있는 여러 go 파일들을 하나의 패키지라고 한다.")
	println("폴더명 = 패키지명")
	println("함수, 구조체, 인터페이스 등의 첫글자가")
	println("대문자면 public")
	println("소문자면 private")
}

func privateFunc() {
	println("같은 패키지 내에서만 사용가능한 함수")
	println("패키지를 import 해도 못 쓴다.")
}

func PublicFunc() {
	println("패키지 밖에서도 사용가능한 함수")
	println("패키지를 import 하면 쓸 수 있다.")
}
