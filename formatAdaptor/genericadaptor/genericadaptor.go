package main

import (
	"errors"
	"os"
	"fmt"
	"encoding/csv"
)

type GenericAdaptor struct {

}

/* func (ga *GenericAdaptor) Dependencies() (dependencies []dependency.Dependency) {
    dependencies = []dependency.Dependency{}
    return
} */

func (ga *GenericAdaptor) Split(srcFile string, dstFiles []string) (err error) {
	inputFile, err := os.Open(srcFile)
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(inputFile)
	records, err := csvReader.ReadAll()

	if len(records[0]) > len(dstFiles) {
		return errors.New("Too many columns, not enough files specified!")
	}

	var i int
	var j int
	for j = 0; j < len(records[0]); j++ {
		outputFile, err := os.Create(dstFiles[j])
		if err != nil {
			return err
		}

		newColumn := make([][]string, len(records), len(records))
		for i = 0; i < len(records); i++ {
			newColumn[i] = make([]string, 1, 1)
			fmt.Printf("%v, %v\n", i, j)
			newColumn[i][0] = records[i][j]
		}

		csvWriter := csv.NewWriter(outputFile)
		csvWriter.WriteAll(newColumn)
	}

	return nil
}

/*func (ga *GenericAdaptor) Join(srcFiles []string, dstFile string) (datasetPath string, err error) {
	
}

func (ga *GenericAdaptor) Shape(path string) (ncols, nrows uint, err error) {

}

func (ga *GenericAdaptor) ToRaw(srcFile string, dstFile) (rawPath string, err error) {

}

func (ga *GenericAdaptor) FromRaw(srcFile string, dstFile) (rawPath string, err error) {

}*/

func main() {
	blank := make([]string, 5, 5)
	blank[0] = "/home/david/git/ProtoML/formatAdaptor/funcOutput1.csv"
	blank[1] = "/home/david/git/ProtoML/formatAdaptor/funcOutput2.csv"
	blank[2] = "/home/david/git/ProtoML/formatAdaptor/funcOutput3.csv"
	blank[3] = "/home/david/git/ProtoML/formatAdaptor/funcOutput4.csv"
	blank[4] = "/home/david/git/ProtoML/formatAdaptor/funcOutput5.csv"
	tester := GenericAdaptor { }
	err := tester.Split("/home/david/git/ProtoML/formatAdaptor/example.csv", blank)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
