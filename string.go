package main

import "fmt"
import "strings"

func main() {
	fmt.Println(strings.Contains("Jaka Kelana", "Jaka")) 
	fmt.Println(strings.Split("Jaka Kelana", " ")) 
	fmt.Println(strings.ToLower("Jaka Kelana")) 
	fmt.Println(strings.ToUpper("Jaka Kelana")) 
	fmt.Println(strings.TrimSpace("   Jaka Kelana   ")) 
	fmt.Println(strings.Replace("Jaka Kelana", "Jaka", "Eko" , 1))
	fmt.Println(strings.ReplaceAll("Jaka Kelana", "Jaka", "Eko"))
	fmt.Println(strings.Trim("   Jaka Kelana   ", " "))
}