package main

import "fmt"

const rsc string = "# * * Rossens System Center * * #"
const eto string = "# * * * * * * * * * * * * * * * #"

func main() {

    platform := "plasticlimit welcomes you"
    fmt.Println(rsc)
    fmt.Println(eto)
    fmt.Println(platform)

    fmt.Println("hello strange creature" + " that appears to have stepped out of a legend")

	i := 1
    	for i <= 3 {
        	fmt.Println(i)
        	i = i + 1
    	}

	for j := 7; j <= 9; j++ {
    	    fmt.Println(j)
    	}

    var chiffre int = 8

	if chiffre%4 == 0 {
    	    fmt.Println(chiffre,"is divisible by 4")
    	} else {
    	    fmt.Println(chiffre,"isn't divisible by 4")
    	} 	

   	
}
