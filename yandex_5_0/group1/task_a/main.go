package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

const (
	InvalidInputConvertionError = "invalid input: firstSubSlice cant be converted to int"
	InvalidInputParamLength     = "invalid input: slice length is not equal to 2"
)

func main() {
	reader, writer := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	vasya, err := readHumParams(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}
	maria, err := readHumParams(reader)
	if err != nil {
		writeResult(writer, err.Error())
		os.Exit(2)
	}

	vMinTree, vMaxTree := vasya.Diapason()
	mMinTree, mMaxTree := maria.Diapason()

	var totalTrees int
	if vMaxTree < mMinTree || mMaxTree < vMinTree {
		totalTrees = (vMaxTree - vMinTree + 1) + (mMaxTree - mMinTree + 1)
	} else {
		totalStart := min(vMinTree, mMinTree)
		totalEnd := max(vMaxTree, mMaxTree)
		totalTrees = totalEnd - totalStart + 1
	}

	writeResult(writer, strconv.Itoa(totalTrees))
}

type HumParams struct {
	TreeNum    int
	FinishDist int
}

func NewHumParams(firstSubSlice, secondSubSlice string) (*HumParams, error) {
	var treeNum, FinishDist int
	var err error

	treeNum, err = strconv.Atoi(firstSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConvertionError)
	}
	FinishDist, err = strconv.Atoi(secondSubSlice)
	if err != nil {
		return nil, errors.New(InvalidInputConvertionError)
	}

	return &HumParams{treeNum, FinishDist}, nil
}

func (hp *HumParams) Diapason() (Start int, Finish int) {
	return hp.TreeNum - hp.FinishDist, hp.TreeNum + hp.FinishDist
}

func writeResult(writer *bufio.Writer, result string) {
	_, _ = writer.WriteString(result)
}

func readHumParams(reader *bufio.Reader) (*HumParams, error) {
	line, _ := reader.ReadString('\n')
	line = strings.TrimRight(line, "\n\r")

	slice := strings.Split(line, " ")
	if len(slice) != 2 {
		return nil, errors.New(InvalidInputParamLength)
	}

	params, err := NewHumParams(slice[0], slice[1])
	if err != nil {
		return nil, err
	}

	return params, nil
}
