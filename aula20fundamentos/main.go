package main

func main() {

	a := 10
	b := 20
	c := 30

	if a > b { //No GO não existe elseif, apenas if e else usados separadamente
		println(a)
	} else {
		println(b)
	}

	switch a {
	case 1:
		println(a)
	case 2:
		println(b)
	default:
		println(c)
	}
}
