package ior

import (
	"encoding/csv"
	"io"
	"strconv"
)

type Result struct {
	// Access is the type of operation (read or write)
	Access string
	// Bandwidth is in MiB/s
	Bandwidth float64
	IOPS      float64
	Latency   float64
	// BlockSize is in KiB
	BlockSize float64
	// TransferSize is in KiB
	TransferSize float64
	// OpenDuration is in seconds
	OpenDuration float64
	// WrRdDuration is in seconds
	WrRdDuration float64
	// CloseDuration is in seconds
	CloseDuration float64
	// TotalDuration is in seconds
	TotalDuration float64
	// Tasks is the number of processes
	Tasks uint64
	// Iteration is the iteration number
	Iteration uint64
}

type Reader struct {
	logs *csv.Reader
}

func NewReader(logs io.Reader) *Reader {
	return &Reader{
		logs: csv.NewReader(logs),
	}
}

func (r *Reader) ReadAsResult() (*Result, error) {
	rec, err := r.logs.Read()
	if err != nil {
		return nil, err
	}

	// If this is the header, re-read again. Header starts with "access".
	if rec[0] == "access" {
		rec, err = r.logs.Read()
		if err != nil {
			return nil, err
		}
	}

	access := rec[0]
	bw, err := strconv.ParseFloat(rec[1], 64)
	if err != nil {
		return nil, err
	}
	iops, err := strconv.ParseFloat(rec[2], 64)
	if err != nil {
		return nil, err
	}
	latency, err := strconv.ParseFloat(rec[3], 64)
	if err != nil {
		return nil, err
	}
	blockSize, err := strconv.ParseFloat(rec[4], 64)
	if err != nil {
		return nil, err
	}
	transferSize, err := strconv.ParseFloat(rec[5], 64)
	if err != nil {
		return nil, err
	}
	openDuration, err := strconv.ParseFloat(rec[6], 64)
	if err != nil {
		return nil, err
	}
	wrRdDuration, err := strconv.ParseFloat(rec[7], 64)
	if err != nil {
		return nil, err
	}
	closeDuration, err := strconv.ParseFloat(rec[8], 64)
	if err != nil {
		return nil, err
	}
	totalDuration, err := strconv.ParseFloat(rec[9], 64)
	if err != nil {
		return nil, err
	}
	tasks, err := strconv.ParseUint(rec[10], 10, 64)
	if err != nil {
		return nil, err
	}
	iter, err := strconv.ParseUint(rec[11], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Result{
		Access:        access,
		Bandwidth:     bw,
		IOPS:          iops,
		Latency:       latency,
		BlockSize:     blockSize,
		TransferSize:  transferSize,
		OpenDuration:  openDuration,
		WrRdDuration:  wrRdDuration,
		CloseDuration: closeDuration,
		TotalDuration: totalDuration,
		Tasks:         tasks,
		Iteration:     iter,
	}, nil
}

func ComputeAvgReadWrite(r *Reader) (read *Result, write *Result, err error) {
	var sumRead, sumWrite Result
	var countRead, countWrite int

	for {
		r, err := r.ReadAsResult()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, nil, err
		}

		var sum *Result
		if r.Access == "write" {
			countWrite++
			sum = &sumWrite
		} else if r.Access == "read" {
			countRead++
			sum = &sumRead
		}
		sum.Bandwidth += r.Bandwidth
		sum.BlockSize += r.BlockSize
		sum.CloseDuration += r.CloseDuration
		sum.IOPS += r.IOPS
		sum.Latency += r.Latency
		sum.OpenDuration += r.OpenDuration
		sum.Tasks += r.Tasks
		sum.TotalDuration += r.TotalDuration
		sum.TransferSize += r.TransferSize
		sum.WrRdDuration += r.WrRdDuration
		sum.Iteration = r.Iteration + 1
	}

	if countRead != 0 {
		read = &sumRead
		read.Access = "read"
		read.Bandwidth /= float64(countRead)
		read.BlockSize /= float64(countRead)
		read.CloseDuration /= float64(countRead)
		read.IOPS /= float64(countRead)
		read.Latency /= float64(countRead)
		read.OpenDuration /= float64(countRead)
		read.Tasks /= uint64(countRead)
		read.TotalDuration /= float64(countRead)
		read.TransferSize /= float64(countRead)
		read.WrRdDuration /= float64(countRead)
	}

	if countWrite != 0 {
		write = &sumWrite
		write.Access = "write"
		write.Bandwidth /= float64(countWrite)
		write.BlockSize /= float64(countWrite)
		write.CloseDuration /= float64(countWrite)
		write.IOPS /= float64(countWrite)
		write.Latency /= float64(countWrite)
		write.OpenDuration /= float64(countWrite)
		write.Tasks /= uint64(countWrite)
		write.TotalDuration /= float64(countWrite)
		write.TransferSize /= float64(countWrite)
		write.WrRdDuration /= float64(countWrite)
	}

	return read, write, nil
}
