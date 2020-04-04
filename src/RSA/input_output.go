package main

import(
    "os"
    "fmt"
    "io/ioutil"
    "bufio"
    "log"
)

/*******************************************************************************
* Does file reading from the user and reads the entire content of the file as  *
* ascii                                                                        *
*******************************************************************************/
func readFile(filename string) []uint8{
    // Reads the entire file and returns (content) an array of unsigined int8
    // Meaning each character in the file is already converted to the
    // ascii value
    content, err := ioutil.ReadFile("../TestFiles/" + filename)

    // Error Condition
    if err != nil {
        log.Fatal(err)
    }

    return content
}


/*******************************************************************************
* Writes out file as a String/Byte Array character by character                *
*******************************************************************************/
func writeFile(filename string, text []byte){
    // Writes out the entire file character by character
    err := ioutil.WriteFile("../OutFiles/" + filename, text, 600)

    // If Error Occurs
    if err != nil{
        fmt.Println(err)
    }
}

/*******************************************************************************
* Writes out the plaintext as a String character by character out to file      *
*******************************************************************************/
func writePlainText(filename string, text []string){
    // File to write out path
    file, err := os.Create("../OutFiles/" + filename)
    if err != nil{
        fmt.Println(err)
    }

    // Defer closing the file, so don't close the file yet ;)
    // Once everything is done this cheeky little line will run
    defer file.Close()
    
    // Open the File Writer
    w := bufio.NewWriter(file)

    for _, character := range text{
        
        // Writes the character
        fmt.Fprint(w, character)
    }
}
