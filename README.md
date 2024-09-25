# go-pacbiohifi

- Profiling of the pacbiohifi reads, making mers and selective filtering of the mers based on composition for better sort and filtering of the kmers and hapmers. 

- Running from the binary
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/go-pacbiohifi% \
./go-pacbiohifi -h
This is a pacbiohifi streamline reader, which will tell you about how your sequencing pacbiohifi looks.
You can give the reads from the fastq or you can give the pacbio bam file from the sequencing

Usage:
  flags [flags]

Flags:
  -h, --help               help for flags
  -i, --inputfile string   inputfile path to be given (default "path to the inputfile")
  -k, --kmer int           kmer to be used for the analysis (default 4)
  -d, --kmerremove float   kmer with this compositon to be removed (default 0.3)
  -o, --ouputfile string   outputfile to be given (default "path to the outputfile")
```
- Running from the github repository 
```
[gauravsablok@ultramarine]~/Desktop/codecreatede/golang/go-pacbiohifi% \
go run main.go -i ./sample-files/samplepacbiohifi.fastq -k 5 -d 0.5

```

Gaurav Sablok
