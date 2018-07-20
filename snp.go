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
TODO: support data format for genesets including RR for each SNP, and corresponding histogram/summary of risks
*/

// SNPs holds a map of each snp -> alleles
type SNPs struct {
	snp map[string]string
}

// findSNP returns an allele for a given individual SNP id
func (s *SNPs) findSNP(snpID string, negativeOrientation bool) (allele string) {
	allele, found := s.snp[snpID]
	switch {
	case found && negativeOrientation:
		return dnaComplement(allele)
	case found:
		return allele
	default:
		return ""
	}
}

// findSNPs returns a list of allele information for a list of SNPs
func (s *SNPs) findSNPs(geneset []string, negativeOrientation bool) (alleles []string) {
	for _, snp := range geneset {
		allele := s.findSNP(snp, negativeOrientation)
		if allele != "" {
			alleles = append(alleles, fmt.Sprintf("SNP:%s\t%s\n", snp, allele))
		} else {
			alleles = append(alleles, fmt.Sprintf("SNP:%s\tNot Found", snp))
		}
	}
	return
}

// loadSNPs loads the 23andme file into memory
func loadSNPs(path string) (snps SNPs, err error) {
	var file *os.File
	if path == "~/.dna/genome.txt" {
		file, err = os.Open(os.Getenv("HOME") + "/.dna/genome.txt")
	} else {
		file, err = os.Open(path)
	}

	check(err)
	defer file.Close()

	snps.snp = make(map[string]string)

	// TODO Check that the file matches 23andme format and raise error if needed
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "\t")
		// comments in 23andme files begin with an #
		// so we'll skip those and just add lines with data to the map
		if !strings.HasPrefix(data[0], "#") {
			snp := data[0]
			allele := data[3]
			snps.snp[snp] = allele
		}
	}
	if err := scanner.Err(); err != nil {
		return snps, err
	}
	return snps, nil
}

// loadGeneset loads a list of SNPs from disk
func loadGeneset(genesetPath string) (geneset []string) {
	// TODO: support multiple file formats by extension
	// for now, just a flat list of SNPs
	file, err := os.Open(genesetPath)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		geneset = append(geneset, scanner.Text())
	}
	return
}

// interactiveMode allows the user to search for SNPs one at a time
func interactiveMode(snps *SNPs) {
	fmt.Println("Interactive Mode\nType 'exit' to quit\nType 'negative' to switch to negative orientation\nType 'positive' to switch to positive orientation")
	negativeMode := false
	for {
		fmt.Print("Enter a SNP id> ")
		var snp string
		fmt.Scanf("%s", &snp)
		switch {
		case snp == "exit" || snp == "quit":
			os.Exit(0)
		case snp == "negative":
			fmt.Println("Switching to negative orientation")
			negativeMode = true
			continue
		case snp == "positive":
			fmt.Println("Switching to positive orientation")
			negativeMode = false
			continue
		default:
			allele := snps.findSNP(snp, negativeMode)
			if allele != "" {
				fmt.Printf("SNP: %s\t%s\n", snp, allele)
			} else {
				fmt.Printf("SNP: %s not found\n", snp)
			}
		}
	}

}

func main() {
	snpID := flag.String("snp", "", "SNP identifier to search for")
	negativeOrientation := flag.Bool("negative", false, "set for negative orientation")
	filePath := flag.String("filepath", "~/.dna/genome.txt", "the path to your 23andme data")
	genesetPath := flag.String("geneset", "", "path to a geneset list")
	interactive := flag.Bool("interactive", false, "starts interactive lookup mode")
	flag.Parse()

	snps, err := loadSNPs(*filePath)
	check(err)

	switch {
	case *snpID == "" && *genesetPath == "" && !*interactive:
		log.Fatal("Either a SNP or a geneset need to be set, or use interactive mode")
	case *snpID != "" && *genesetPath != "":
		log.Fatal("You must either select a SNP or a geneset, not both")
	case *snpID != "":
		allele := snps.findSNP(*snpID, *negativeOrientation)
		if allele != "" {
			fmt.Printf("SNP:\t%s\nAllele:\t%s\n", *snpID, allele)
		} else {
			fmt.Printf("SNP: %s not found", *snpID)
		}
	case *interactive:
		interactiveMode(&snps)
	case *genesetPath != "":
		geneset := loadGeneset(*genesetPath)
		alleles := snps.findSNPs(geneset, *negativeOrientation)
		for _, allele := range alleles {
			fmt.Print(allele)
		}
	}

}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

// dnaComplement returns the complement of a base pair
// it converts positive to negative orientation, or vice versa
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
