package main

import "fmt"

func main() {
	/*[string] is the format of the key. int, is the format of the value.
	  +this is basically key == 10
	*/
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	if name, ok := elements["B"]; ok {
		fmt.Println(name)
	}
}
