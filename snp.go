package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"terminal"
)

var snp_id string
var negative_orientation bool

func main() {
	flag.StringVar(&snp_id, "snp", "", "The SNP identifier to search for")
	flag.BoolVar(&negative_orientation, "negative", false, "Set for negative orientation")
	flag.Parse()

	if snp_id == "" {
		snp_id = os.Args[1]
	}

	// refactor for other environments/flexibility
	file, err := os.Open(os.Getenv("HOME") + "/.dna/genome.txt")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), "\t")
		// refactor to support lists of SNPs
		snp := data[0]
		if snp_id == snp {
			allele := data[3]
			if negative_orientation {
				allele = dna_complement(allele)
				terminal.Stdout.Color("g").Colorf("@{!r}Showing Negative Orientation").Nl()
			}
			terminal.Stdout.Color("g").Colorf("@{!b}SNP:     %s\nAlleles: %s", snp, allele).Nl()
			os.Exit(0)
		}
	}
	terminal.Stdout.Colorf("@{!r}SNP not found").Nl()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func dna_complement(dna string) string {
	var reverse_orientation string
	complements := map[string]string{
		"A": "T",
		"T": "A",
		"C": "G",
		"G": "C",
	}
	for _, char := range dna {
		reverse_orientation += complements[string(char)]
	}
	return reverse_orientation
}
