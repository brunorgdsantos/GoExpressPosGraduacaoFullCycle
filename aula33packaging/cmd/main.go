package main

import (
	"fmt"
	"github.com/brunorgdsantos/GoExpressPosGraduacaoFullCycle/aula33packging/math"
)

func main() {
	m := math.Math{A: 10, B: 20}
	fmt.Println(m.Add())
}
