package main

import "fmt"

func main() {
	/*[string] is the format of the key. int, is the format of the value.
	  +this is basically key == 10
	*/
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	if name, ok := elements["B"]; ok {
		fmt.Println(name)
	}
}
