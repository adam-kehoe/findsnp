# findsnp

A simple tool for looking up genotype calls in the 23andme raw data file. Given a 23andme data file, you can search by identifier (rsid or 23andme's internal id).

This is an experiment with the Go language, and is not intended for production use.

## Example

Simple SNP lookup
```
$ findsnp rs429358
```

SNP lookup showing negative orientation
```
$ findsnp --snp=rs30197 --negative
```
