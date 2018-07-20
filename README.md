# findsnp

A simple tool for looking up genotype calls in a 23andme raw data file. Given a 23andme data file, you can search by identifier (rsid or 23andme's internal id). You can also provide a list of identifiers to search SNPs in bulk, or search for individual SNPs in an interactive mode.

This tool was built to speed the process of looking up SNPs. It can also be useful for bulk queries for a family; for example, you can use the tool to check the status of multiple individuals against a common list of SNPs.

## Examples

Simple SNP lookup:
```
$ findsnp --snp=rs429358
```

SNP lookup showing negative orientation
```
$ findsnp --snp=rs30197 --negative
```

SNP lookup for a set of SNPs:
```
$ findsnp -geneset="yourlist.txt"
```

To create a geneset, create a list of SNPs with one SNP per line:

```
rs17580
rs28929474
```

Interactive mode:

```
$ findsnp -interactive
```

## Configuration

findsnp looks for a 23andme data file under ~/.dna/genome.txt by default. You can set this path with the -filepath argument.