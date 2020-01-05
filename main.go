package main

func main() {

}

//fun+tab
//ctrl+shift+t test

func cashback(amount int)int  {
	//esli summa pokupki bolwe 3000
	//togda casback 5% ot pokupki
	const bound = 3000 // granica s kotoroy nachislyaetsya cashback
	const percent = 5
	result := 0

	if amount>=bound{
		result = amount * percent / 100
	}

	return result
}
