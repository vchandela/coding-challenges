## Challenge 1: Write Your Own wc tool

### Steps to test the code:
- cd ccwc
- go install -> Install package in GOPATH
- Now, try the following commands:
    - ccwc -c test.txt
    - ccwc -l test.txt
    - ccwc -w test.txt
    - ccwc -m test.txt
    - ccwc test.txt
    - cat test.txt | ccwc -l
    - ccwc -l < test.txt
    - cat test.txt | ccwc
    - ccwc < test.txt

### Key Concepts
- flag
- switch
- file open & close
- buffered reader (bufio) vs io.Reader (reads directly in memory) [[link](https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762)]
- error handling (compare diff ways)
- regex
- keyboard vs pipe/redirection [[link](https://rderik.com/blog/identify-if-output-goes-to-the-terminal-or-is-being-redirected-in-golang/)]
- Rule of visibility: When an identifier (including a function name) starts with an uppercase letter, it is exported and can be accessed from other packages. If it starts with a lowercase letter, it is unexported and can only be accessed within the same package.
- bytes.Buffer vs temp file: Tradeoff b/w memory and disk seeks
- testing package: Add unit tests in Go