package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"testing"
)

type testCase struct {
	Slice  []int
	Answer int
}

func getFileIndex(index int) string {
	fileIndex := ""
	if index > 9 {
		fileIndex = strconv.Itoa(index)
	} else {
		fileIndex = "0" + strconv.Itoa(index)
	}

	return fileIndex
}

func testTable() []testCase {
	table := make([]testCase, 0)
	for i := 1; i<11; i++ {
		tcase := testCase{}
		fileIndex := getFileIndex(i)

		fileOfTests, err := os.Open("tests/" + fileIndex)
		if err != nil {
			log.Fatal(err)
		}
		defer fileOfTests.Close()
		scannerOfTestFile := bufio.NewScanner(fileOfTests)
		scannerOfTestFile.Split(bufio.ScanWords)

		fileOfAnswersForTests, err := os.Open("tests/" + fileIndex +".a")
		if err != nil {
			log.Fatal(err)
		}
		defer fileOfAnswersForTests.Close()
		scannerOfAnswerFile := bufio.NewScanner(fileOfAnswersForTests)

		scannerOfTestFile.Scan()
		numTestCases, err := strconv.Atoi(scannerOfTestFile.Text())
		if err != nil {
			log.Fatal(err)
		}
		for idxOfTestCase := 0; idxOfTestCase < numTestCases; idxOfTestCase++ {
			scannerOfTestFile.Scan()
			numOfElementsInSlice, err := strconv.Atoi(scannerOfTestFile.Text())
			if err != nil {
				log.Fatal(err)
			}
			sliceOfInts := make([]int,0)
			// sliceOfStrings := strings.Split(scannerOfTestFile.Text(), " ")
			for j:= 0 ; j<numOfElementsInSlice; j++{
				scannerOfTestFile.Scan()
				sliceTemporary, err := strconv.Atoi(scannerOfTestFile.Text())
				
				if err!=nil{
					log.Println(numOfElementsInSlice)
					log.Println(i)
					log.Println(idxOfTestCase)
					log.Println(j)
					log.Println(sliceTemporary)
					log.Fatal(err)
				}
				sliceOfInts = append(sliceOfInts, sliceTemporary)
			}

			scannerOfAnswerFile.Scan()
			answer, err := strconv.Atoi(scannerOfAnswerFile.Text())
			if err != nil {
				log.Fatal(err)
			}
			tcase = testCase{
				Slice:  sliceOfInts,
				Answer: answer,
			}
		}
	
		table = append(table, tcase)
	}

	return table
}

func TestCalculateAmountToPay(t *testing.T) {
	testCase := testTable()
	for i, test := range testCase {
		if calculateAmountToPay(test.Slice) != test.Answer {
			t.Errorf("Test %v. %v is not equal %v", i, calculateAmountToPay(test.Slice), test.Answer)
		}
	}
}