package main

import (
	"fmt"
	"log"	
    "os"
)

func main() {    
    resume, err := os.ReadFile("../resume/resume")
    if err != nil {
        log.Fatal(err)
    }
    jobDescription, err := os.ReadFile("../jobDescription/jobDescription") 
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(resume))
    fmt.Println(string(jobDescription))
}
