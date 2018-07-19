package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
TODO: support a list of SNPs to search for
TODO: support data format for genesets including RR for each SNP, and corresponding histogram/summary of risks
*/

func findSNP(snpID, path string, negativeOrientation bool) (snp string, err error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "\t")
		snp := data[0]
		if snpID == snp {
			allele := data[3]
			if negativeOrientation {
				allele = dnaComplement(allele)
				fmt.Println("Showing Negative Orientation")
			}
			return fmt.Sprintf("SNP:\t%s\nAlleles:\t%s", snp, allele), nil
		}
	}
	return "", nil
}
func main() {
	snpID := flag.String("snp", "", "SNP identifier to search for")
	negativeOrientation := flag.Bool("negative", false, "set for negative orientation")
	filePath := flag.String("filepath", "~/.dna/genome.txt", "the path to your 23andme data")
	// listPath := flag.String("listpath", "", "path to a list of alleles to search for")

	flag.Parse()

	if *snpID == "" {
		snpID = &os.Args[1]
	}

	result, err := findSNP(*snpID, *filePath, *negativeOrientation)
	check(err)
	if result != "" {
		fmt.Println(result)
	} else {
		fmt.Printf("SNP %s not found\n", *snpID)
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func dnaComplement(dna string) string {
	var reverseOrientation string
	complements := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	for _, char := range dna {
		reverseOrientation += complements[string(char)]
	}
	return reverseOrientation
}
