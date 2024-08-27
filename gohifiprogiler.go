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
	_ "github.com/go-sql-driver/mysql"
	// adding the mysql driver so that it can interact with HUGO as a frontend.
	// adding the support for the charts plotting
)

func main() {

	// adding here also cobra cli and flags and help menu 

	argsread := os.Args[1:]
	argswrite := os.Args[2:]
	argskmer := os.Args[3:]
	argsinput1 := os.Args[4:]
	argsinput2 := os.Args[5:]

	if err != nil {
		panic (err)
		log.Fatal(err.Error())
		return
	}
	readfile, err := os.Open(argsread)
	if err != nil {
		panic(err.Error())
		log.Fatal(err)
		return
	}
	argsinput1 := os.Args[1:]
	argsinput2 := os.Args[2:]
	argsoutput1 := os.Args[3:]
	argsoutput2 := os.Args[4:]
	if argsinput1 || argsinput2 || argsinput3 || argsinput4 == nil {
		panic(err)
		log.Fatal(err.Error.new("The arguments cant be empty"))
	}

	readbuffer := bufio.NewScanner(readfile)
	header := []string{}
	sequences := []string{}
	header := []string{}
	sequences := []string{}
	for readbuffer.Scan() {
		line := readbuffer.Text()
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			sequences = append(sequences, line)
		}
		if string(line[0]) == "@" {
			header = append(header, line)
		}
	}

	seqtok := []string{}
	for i := range sequences {
		// this will prepare all the mers from all the sequences and not the sequence specific.
		for j := 0; j <= len(sequences[i])-2; j++ {
			seqtok = append(seqtok, string(sequences[i][j:j+2]))
		}
	}

	mapmer := make(map[string]string)
	// this will a map of the sequences to act as a getter. 
	for i := range header {
			mapmer[string(header[i])] = string(sequences[i])
		}

		// a anonymous function that will act as a callback in getting the density plot. 
	seqCount := []int{}
	seqHeaders := []string{}
	seqgcCount := []int{}
	seqplot := []int{}
	for i,j := range mapmer {
			seqHeaders = append(seqHeaders,i)
			seqCount = append(seqCount, (strings.Count(mapmer[i], "A")+ strings.Count(mapmer[i], "T")+strings.Count(mapmer[i], "G")+strings.Count(mapmer[i], "C")))
		    seqgcCount = append(seqgcCount, strings.Count(mapmer[i], "G")+strings.Count(mapmer[i], "C"))
		}
	for i := range  seqgcCount {
			seqplot = append(seqplot, int(seqgcCount[i])/int(seqCount[i]))
		}

	kmercomp := []int{}
	for i := range seqtok {
		hold := string(seqtok[i])
		kmercomp = append(kmercomp, strings.Count(hold, "A")+strings.Count(hold, "T")+strings.Count(hold, "G")+strings.Count(hold, "C"))
	}

	kmercompGC := []int{}
	for i := range seqtok {
		hold := string(seqtok[i])
		kmercomp = append(kmercomp, strings.Count(hold, "G")+strings.Count(hold, "C"))
	}

	kmerclassify := []string{}
	kmerfilter := []int{}
	for i := range seqtok {
		kmerclassify = append(kmerclassify, seqtok[i])
		kmerfilter = append(kmerfilter, int(kmercompGC[i])/int(kmercomp[i]))
	}

	filteredKmer := []string{}
	for i := range kmerclassify {
		// mention the threshold for the filtering of the low kmers. 
		// default value set to 10 and fyne GO application allows to select based on the plotting graph.
		if kmerfilter[i] < 10 {
			filteredKmer = append(filteredKmer, kmerclassify[i])
		}
	}

	for i := range filteredKmer {
		fmt.Println("The filtered kmers with the elective selection are %s 
		                   and their length are %T:", string(filteredKmer[i]), len(string(filteredKmer[i])))
	}
	
	func savemers(argswrite string, kmercomp []byte) error {
		// adding a file save function to save each information before the final build release. 
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
		if err != nil {
		return err
		}
		defer file.Close()
		_, err = file.Write(data)
		if err != nil {
		return err
		}
		return fp.Sync() // fsync
		}
	
	func savefilteredKmer(argskmer string, kmercomp []byte) error {
		// adding a file save function to save each information before the final build release. 
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
		if err != nil {
		return err
		}
		defer file.Close()
		_, err = file.Write(data)
		if err != nil {
		return err
		}
		return fp.Sync() // fsync
		}
	}

	/*
	  adding the support for the hybrid check for the assembly of the pacbiohifi and the illumina reads
	 */
		
	open1 := bufio.NewScanner(readinput1)
	open2 := bufio.NewScanner(readinput2)

	openinputH1 := []string{}
	openinputSeq1 := []string{}
	openinputH2 := []string{}
	openinputSeq2 := []string{}

	for open1.Scan() {
		line := open1.Text()
		if strings.HasPrefix(line, "@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH1, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq1 = append(openinputSeq1, line)
		}
	}

	for open2.Scan() {
		line := open2.Text()
		if strings.HasPrefix(line, "@") || strings.Contains(string(line), "length") {
			openinputH1 = append(openinputH2, strings.Split(line, " ")[0])
		}
		if string(line[0]) == "A" || string(line[0]) == "T" || string(line[0]) == "G" || string(line[0]) == "C" {
			openinputSeq2 = append(openinputSeq2, line)
		}
	}

	balancedH1 := []string{}
	balancedH2 := []string{}
	balancedH1head := []string{}
	balancedH2head := []string{}

	for i := range openinputH1 {
		if string(openinputH1[i]) == string(openinputH2[i]) {
			balancedH1 = append(balancedH1, openinputH1[i])
			balancedH2 = append(balancedH2, openinputH2[i])
			balancedH1head = append(balancedH1head, openinputSeq1[i])
			balancedH2head = append(balancedH2head, openinputSeq2[i])
		}
	}

	checklen := len(balancedH1) == len(balancedH1)
	if checklen != true {
		panic(err.Error())
		log.Fatal(Error.new("The balancer is uneven and the reads cant be assembled"))

	} else {
		fmt.Println("The balancer is equal")
	}
}
