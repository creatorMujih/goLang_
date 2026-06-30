package main

import "fmt"
	
func main(){
	details := data{name: "Joe", cluster: "three", contact: 07065631765}
	fmt.Println("cluster: " + details.cluster)
	fmt.Println(details)

	var farmTools tool
		farmTools = tool{name: "Hoe", class: "simple farm tool", number: 100}
		fmt.Println("name: " + farmTools.name)
		fmt.Println(farmTools) 
}

type data struct {
	name string
	cluster string
	contact int
}

type tool struct{
	name string
	class string
	number int
}

