package main

import "fmt"

func main(){
	basic()
	group()
	notValue()
	short()
}

func basic(){
	var i int = 1
	fmt.Println(i)
}

func group(){
	var (
		i int = 2
		s string = "hoge"
	)
	fmt.Println(i, s)
}

func notValue(){
	var (
		i int
		f bool
	)
	fmt.Println(i, f)
}

func short(){
	f := 1.2
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", f)
}

