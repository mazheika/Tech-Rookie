package main 
import (  
"os"  
"fmt"
"sort"
) 
func main() { 
buffer := make([]byte, 1024)  
file, _ := os.Open("data.txt")  
var err error  
line := ""  
output := "" // Read loop  
for ; err == nil; _, err = file.Read(buffer) {  
	// appending each byte to a line  
	if buffer[0] != '\n' {  
		line += string(buffer[0])  
		// Line is done, let's sort it  
	} else {  
	// Collect words for each line 
	word := ""  
	var words []string  
	for _, c := range line {  
		if c != ' ' {  
			word += string(c)  
		} else {  
			words = append(words, word)  
			word = ""  
		}  
	sort.Strings(words)
	line = "" // write words to the output buffer  for _, word := range words {  
	for _, c := range words {  
		output += string(c)  
	}
	output +=  " "  
	}
	output += "\n"  
	} // create a file
	file, err = os.Create("res.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, c := range []byte(output) {
		file.Write([]byte{c})
	}
}
}

/*if os.IsNotExist(err) {
	file, err := os.Create("res.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
} else {
	fmt.Println("File already exists!", "res.txt")
	return
}
*/