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
	"net/http"
	"fmt"
	"log"
	"os"
	"strings"
	_ "github.com/go-sql-driver/mysql"
	// adding the mysql driver so that it can interact with HUGO as a frontend.
	// adding the support for the charts plotting
	/*
	 Task to do:
	 1. Add the support for the mapping of the HiC to the PacbioHifi
	 2. Extracting the graphs
	 3. Making a linkedlist for the graphs and implementing a LIFO on the sequence based nodes.
	 4. implementing the http/net module for redireting the reads and the alignments to the browsers and integrating a browser based plot for the visualization.
	 4. big check if any.
	 */
)

func main() {

	argsread := os.Args[1:]
	argswrite := os.Args[2:]
	argskmer := os.Args[3:]
	argsinput1 := os.Args[4:]
	argsinput2 := os.Args[5:]
	readbuffer := bufio.NewScanner(argsread)
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
		for j := 0; j <= len(sequences[i])-int(argskmer); j++ {
			seqtok = append(seqtok, string(sequences[i][j:j+int(argskmer)]))
		}
	}

	mapmer := make(map[string]string)

	// this will a map of the sequences to act as a getter.
	for i := range header {
			mapmer[string(header[i])] = string(sequences[i])
		}

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
		fmt.Println("The filtered kmers with the elective selection are %s and their length are %T:", string(filteredKmer[i]), len(string(filteredKmer[i])))
	}
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
	return fp.Sync()
	}

       // adding the support for the hybrid check for the assembly of the pacbiohifi and the illumina reads

	open1 := bufio.NewScanner(argsinput1)
	open2 := bufio.NewScanner(argsinput2)

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

 // last part
	// an API struct for the pacbio for the continous routing across the http.handlerequest
  type pacbio struct {
		name string
		idseq string
	}

	// function to make the json API for the http request.

	func (*pacbio) unravel () string {
		fileopen := os.OpenFile(os.Args[1:], flag int, perm os.FileMode)
		fileread := bufio.NewScanner(fileopen)
		for i := range fileread.Scan() {
			fileline := fileread.Scan.Text()
			idstore := []pacbio{}
			if strings.HasPrefix(fileline, "@") {
           seqstore = append(seqstore, seqstore{
						 name : fileline
					 })
			if fileline[0] == "A" || fileline[0] == "T" || fileline[0] == "G" || fileline[0] == "C" {
				seqstor<t_k€>Ã½require"cmp.utils.feedkeys".run(10)
				de := []pacbio{}
				seqstore = append(seqstore, seqstore{
					idseq : string(fileline)
				})
			}
      finalcom := []string{idstore, seqstore }
			}
		}
		return len(finalcom)
	}

  func getseq (w http.ResponseWriter, r *http.Request) {
		  varde, err := finalcom
		  if err != nil {
				panic (err)
				log.Fatal(Err.new("an empy struct is decalred"))
		  return
		 } else {
			 io.Write([]byte{"This is a pacbiohifi API from sequencing to sequence viewer across the cluster"})
		 }
	}
// this last part is still in debugging mode as writing a interface for the http
}
