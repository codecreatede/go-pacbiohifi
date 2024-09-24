/*

Author Gaurav Sablok
Universitat Potsdam Library
Date: 2024-8-7

A pacbiohifi go profiler that takes the pacbiohifi reads, makes the mers according to the length,
filters the mers according to the critera. This is supported witht a desktop application to see that
if your reads have high profiled genomic mers that could hinder the graph assembly.
Added the support for the fasta reads from Illumina also and also adding the support for the native checks of
the pacbiohifi reads to the illumina reads either coming from HiC

*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

var (
	inputfile  string
	outputfile string
	kmer       int
	kmerremove float64
)

var rootCmd = &cobra.Command{
	Use:  "flags",
	Long: "This is a pacbiohifi streamline reader, which will tell you about how your sequencing pacbiohifi looks. You can give the reads from the fastq or you can give the pacbio bam file from the sequencing",
	Run:  flagFunc,
}

func init() {
	rootCmd.Flags().
		StringVarP(&inputfile, "inputfile", "i", "path to the inputfile", "inputfile path to be given")
	rootCmd.Flags().
		StringVarP(&outputfile, "ouputfile", "o", "path to the outputfile", "outputfile to be given")
	rootCmd.Flags().IntVarP(&kmer, "kmer", "k", 4, "kmer to be used for the analysis")
	rootCmd.Flags().
		Float64VarP(&kmerremove, "kmerremove", "d", 0.3, "kmer with this compositon to be removed")
}

func flagFunc(cmd *cobra.Command, args []string) {
	readOpen, err := os.Open(inputfile)
	if err != nil {
		log.Fatal(err)
	}
	readbuffer := bufio.NewScanner(readOpen)
	header := []string{}
	sequences := []string{}

	for readbuffer.Scan() {
		line := readbuffer.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" ||
			string(line[0]) == "C" {
			sequences = append(sequences, line)
		}
		if string(line[0]) == "@" {
			header = append(header, line)
		}
	}

	seqtok := []string{}

	for i := range sequences {
		for j := 0; j <= len(sequences[i])-int(kmer); j++ {
			seqtok = append(seqtok, string(sequences[i][j:j+int(kmer)]))
		}
	}

	kmercomp := []int{}
	for i := range seqtok {
		storestring := strings.Count(
			string(seqtok[i]),
			"A",
		) + strings.Count(
			string(seqtok[i]),
			"T",
		) + strings.Count(
			string(seqtok[i]),
			"G",
		) + strings.Count(
			string(seqtok[i]),
			"C",
		)
		kmercomp = append(kmercomp, storestring)
	}

	kmerGC := []int{}
	for i := range seqtok {
		storestring := strings.Count(
			string(seqtok[i]),
			"G",
		) + strings.Count(
			string(seqtok[i]),
			"C",
		)
		kmerGC = append(kmerGC, storestring)
	}

	kmerProfile := []float64{}

	for i := range kmercomp {
		kmerProfile = append(kmerProfile, float64(kmerGC[i])/float64(kmercomp[i]))
	}

	filteredKmer := []string{}

	for i := range kmerProfile {
		if kmerProfile[i] <= kmerremove {
			continue
		} else {
			filteredKmer = append(filteredKmer, seqtok[i])
		}
	}

	for i := range filteredKmer {
		fmt.Println(filteredKmer[i])
	}

	file, err := os.Create("allprofiledKmers.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for i := range seqtok {
		_, err := file.WriteString(
			seqtok[i] + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	file1, err := os.Create("filteredKmer.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file1.Close()

	for i := range filteredKmer {
		_, err := file1.WriteString(filteredKmer[i] + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
