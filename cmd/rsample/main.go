/*rsample
 * command line tool for randomly sampling a fixed number of lines
 * from Stdin or from an arbitrary number of files using reservoire sampling.
 * Fabian Peters 2017
 */
package main

import (
	"bufio"
	"flag"
	"io"
	"os"
	"time"

	"github.com/vrtsig/reservoire/pkg/reservoire"
)

func check(err error) {
	if err != nil {
		os.Stderr.WriteString("ERROR: " + err.Error())
		os.Exit(1)
	}
}

func main() {
	// command line flags
	n := flag.Int("n", 1, "number of samples to be drawn.")
	seed := flag.Int64("seed", time.Now().UnixNano(), "random seed. If not specified, "+
		"time.Now() is used.")
	hasHeader := flag.Bool("header", false, "if true, the first line of "+
		"the input is preserved and does not count towards n. In case of multiple "+
		"input files, only the first header is preserved, all other headers are skipped.")
	outfile := flag.String("outfile", "", "output file. Leave empty to write to stdout.")
	flag.Parse()

	// can write either to a file or to stdout
	var out io.Writer
	if *outfile != "" {
		outFile, err := os.Create(*outfile)
		check(err)
		defer outFile.Close()
		out = outFile
	} else {
		out = os.Stdout
	}

	// create new reservoire
	r, err := reservoire.NewStringReservoire(*n, *seed)
	check(err)

	files := flag.Args()
	var header string
	// if files are stated, read them one by one
	if len(files) > 0 {
		for _, f := range files {
			h, err := addFile(f, &r, *hasHeader)
			check(err)
			if *hasHeader && header == "" {
				header = h
			}
		}
	} else { // otherwise try to read from stdin
		h, err := addStdin(&r, *hasHeader)
		check(err)
		header = h
	}

	// write output
	if *hasHeader {
		io.WriteString(out, header+"\n")
	}
	for _, s := range r.GetAll() {
		io.WriteString(out, s+"\n")
	}

}

// adds stdin to a reservoire and returns header
func addStdin(r *reservoire.StringReservoire, hasHeader bool) (string, error) {
	// check if stdin is empty
	stats, err := os.Stdin.Stat()
	if err != nil {
		return "", err
	}
	if stats.Size() == 0 {
		return "", nil
	}
	return addReader(os.Stdin, r, hasHeader)
}

// adds a single file to a reservoire and returns header
func addFile(fileName string, r *reservoire.StringReservoire, hasHeader bool) (string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer f.Close()
	return addReader(f, r, hasHeader)
}

// adds the contents of an io.Reader to a reservoire and returns header
func addReader(f io.Reader, r *reservoire.StringReservoire, hasHeader bool) (string, error) {
	var header string
	scanner := bufio.NewScanner(f)

	if hasHeader {
		scanner.Scan()
		header = scanner.Text()
	}

	for scanner.Scan() {
		r.Add(scanner.Text())
	}

	return header, scanner.Err()
}
