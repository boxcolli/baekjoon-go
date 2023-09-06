# Accelerate your baekjoon solution
- define how your program scans in ```read()```
- define your solution in ```main()```
- fulfill testcases in /testcase.txt
- run ```go test .```

# Things you should do
## For minimum
- edit:
    - testcase.txt
    - testcase_test.go - ```var readResult input{}```
    - read.go - ```read()```
    - main.go - ```main()```
- run: ```$ go test main_test.go```

## Full test
- edit:
    - testcase.txt
    - testcas_test.go - ```var readResult input{}```
    - read.go - ```read()```
    - read_test.go - ```TestRead()```
    - main.go - ```main()```
- run: ```$ go test .```

# Useful commands
- ```$ go test .```
- ```$ go test my_test.go```
- ```$ go test [-v] my_test.go```
- ```$ go test [-run <RegexForTestFunction>] my_test.go```
