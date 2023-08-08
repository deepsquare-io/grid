package result

import (
	"bufio"
	"io"
	"strconv"
	"strings"
	"time"
)

var CSVHeader = []string{
	"ProblemSize",
	"NB",
	"P",
	"Q",
	"Time",
	"Gflops",
	"Refine",
	"Iter",
	"Gflops_wrefinement",
}

type Result struct {
	ProblemSize       uint64
	NB                uint64
	P                 uint64
	Q                 uint64
	Time              time.Duration
	Gflops            float64
	Refine            float64
	Iterations        uint64
	GflopsWRefinement float64
}

type Reader struct {
	logs *bufio.Scanner
}

func NewReader(logs io.Reader) *Reader {
	return &Reader{
		logs: bufio.NewScanner(logs),
	}
}

func (r *Reader) ReadAsResult() (*Result, error) {
	// Search for valid line or EOF
	var line string
	for {
		if !r.logs.Scan() {
			return nil, io.EOF
		}
		line = strings.TrimSpace(r.logs.Text())
		if strings.HasPrefix(line, "HPL_AI") {
			break
		}
	}

	// Split the line into fields
	fields := strings.Fields(line)

	// Extract the required values

	problemSize, err := strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return nil, err
	}
	nb, err := strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return nil, err
	}
	p, err := strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return nil, err
	}
	q, err := strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return nil, err
	}
	timeF64, err := strconv.ParseFloat(fields[6], 64)
	if err != nil {
		return nil, err
	}
	time := time.Duration(timeF64 * float64(time.Second))
	gflops, err := strconv.ParseFloat(fields[7], 64)
	if err != nil {
		return nil, err
	}
	refine, err := strconv.ParseFloat(fields[8], 64)
	if err != nil {
		return nil, err
	}
	iter, err := strconv.ParseUint(fields[9], 10, 64)
	if err != nil {
		return nil, err
	}
	gflopsWRefinement, err := strconv.ParseFloat(fields[10], 64)
	if err != nil {
		return nil, err
	}

	// Write the extracted values to the CSV file
	return &Result{
		ProblemSize:       problemSize,
		NB:                nb,
		P:                 p,
		Q:                 q,
		Time:              time,
		Gflops:            gflops,
		Refine:            refine,
		Iterations:        iter,
		GflopsWRefinement: gflopsWRefinement,
	}, nil
}

func FindMaxGflopsResult(r *Reader) (res *Result, err error) {
	var maxGflops float64

	for {
		r, err := r.ReadAsResult()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		if r.Gflops >= maxGflops {
			maxGflops = r.Gflops
			res = r
		}
	}

	return res, nil
}

func ComputeAvgGflopsResult(r *Reader) (avg float64, err error) {
	var sumGlops float64
	var totalGflops int

	for {
		result, err := r.ReadAsResult()
		if err != nil {
			if err == io.EOF {
				break
			}
			return avg, err
		}
		totalGflops++
		sumGlops += result.Gflops
	}

	return sumGlops / float64(totalGflops), nil
}
