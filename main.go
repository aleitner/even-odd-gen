package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	var err error
	var outputFile *os.File

	packageName := flag.String("package", "main", "name of package the created file belongs to")
	functionName := flag.String("function", "main", "name of function that checks if a number is even or odd")
	num := flag.Int("n", 10, "Highest number to be checked if even or odd")
	outputPath := flag.String("output", "", "file to output code to")
	flag.Parse()

	// Default to stdout
	if *outputPath == "" {
		outputFile = os.Stdout
	} else {
		outputFile, err = os.Create(*outputPath)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// Write the package
	fmt.Fprintf(outputFile,"package %s\n\n", *packageName)

	// Write the imports
	fmt.Fprintln(outputFile,"import (")
	fmt.Fprintln(outputFile, "\t\"flag\"")
	fmt.Fprintln(outputFile, "\t\"fmt\"")
	fmt.Fprintln(outputFile, "\t\"math/rand\"")
	fmt.Fprintln(outputFile, "\t\"time\"")
	fmt.Fprintln(outputFile, ")\n")

	// Write the function
	fmt.Fprintf(outputFile,"func %s() {\n", *functionName)

	// Seed the guess
	fmt.Fprintln(outputFile,"\trand.Seed(time.Now().UnixNano())\n")

	// Parse the options
	fmt.Fprintln(outputFile,"\tnum := flag.Int(\"n\", 0, \"an int\")")
	fmt.Fprintln(outputFile, "\tflag.Parse()")
	fmt.Fprintln(outputFile,"\ti := *num\n")

	// Write the even/odd checks
	fmt.Fprintf(outputFile,"\t")

	for i:= 0; i <= *num; i++ {
		fmt.Fprintf(outputFile,"if i == %d {\n", i)

		if i % 2 == 0 {
			fmt.Fprintln(outputFile,"\t\tfmt.Println(\"Even\")")
		} else {
			fmt.Fprintln(outputFile,"\t\tfmt.Println(\"Odd\")")
		}

		fmt.Fprintf(outputFile,"\t} else ")
	}

	// The number is too great so let's just guess
	fmt.Fprintln(outputFile, "{")
	fmt.Fprintln(outputFile, "\t\tif rand.Intn(2) == 0 {")
	fmt.Fprintln(outputFile,"\t\t\tfmt.Println(\"idk... Even?\")")
	fmt.Fprintln(outputFile, "\t\t} else {")
	fmt.Fprintln(outputFile,"\t\t\tfmt.Println(\"idk... Odd?\")")
	fmt.Fprintln(outputFile, "\t\t}")
	fmt.Fprintln(outputFile, "\t}")


	// Close the function
	fmt.Fprintln(outputFile,"}")

	err = outputFile.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
}