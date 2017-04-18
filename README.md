# Reservoire Sampling
A simple implementation of [reservoire sampling](https://en.wikipedia.org/wiki/Reservoir_sampling) in `go`. Reservoire sampling can be used to uniformly sample a fixed number of elements from a set of unknown size.

## Usage:
Given a file `test.txt` with
```
header
line 1
line 2
line 3
line 4
line 5
```
the binary can be used as a line filter or with file names as parameters
```
> cat test.txt | rsample -n=3 -header -seed=0
header
line 1
line 2
line 4

> rsample -n=4 -header -outfile=out.txt -seed=1 test.txt
> cat out.txt
header
line 1
line 5
line 3
line 4
```
## Parameters:
```
> rsample -help

Usage of rsample:
  -header
    	if true, the first line of the input is preserved and does not count towards n. In case of multiple input files, only the first header is preserved, all other headers are skipped.
  -n int
    	number of samples to be drawn (default 1)
  -outfile string
    	output file. Leave empty to write to stdout.
  -seed int
    	random seed. If not specified, time.Now() is used.
```
