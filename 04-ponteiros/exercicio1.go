func trocaValor(){
	x := 10
	y := 20
	p1 := &x
	p2 := &y
	fmt.Println("Valores antes da troca: ", x, y)
	*p1,*p2 = *p2, *p1
	fmt.Println("Valores antes da troca: ", x, y)
}