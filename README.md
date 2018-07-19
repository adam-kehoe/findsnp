# findsnp

A simple tool for looking up genotype calls in a 23andme raw data file. Given a 23andme data file, you can search by identifier (rsid or 23andme's internal id). You can also provide a list of identifiers to search SNPs in bulk.

## Examples

Simple SNP lookup
```
$ findsnp --snp=rs429358
```

SNP lookup showing negative orientation
```
$ findsnp --snp=rs30197 --negative
```

SNP lookup for a set of SNPs
```
$ findsnp -geneset="yourlist.txt"
```

To create a geneset, create a list of SNPs with one SNP per line:

```
rs17580
rs28929474
```