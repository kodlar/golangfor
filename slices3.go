package main

import "fmt"


func main(){
	langs := getLangs()
	
	//for  i, element := range langs {
	for  _, element := range langs {
		//fmt.Println(langs[i])
		fmt.Println(element)
	}
	
}

func getLangs()[]string{
	
	var langs []string
	langs = append(langs, "Go")
	langs = append(langs, "Go2")
    langs = append(langs, "Go3")
    langs = append(langs, "Go4")
	return langs
	
}