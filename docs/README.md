# Anagramalyser

Anagramalyser is a Go application for finding anagrams within a text file.
<br><br>
## Installation
<hr>

Either manually build:

```bash
go build -o anagramalyser ./cmd
```
(for windows add .exe to anagramalyser).

Or use one of the releases from Github available [here](https://github.com/mdhemmings/anagramalyser/releases)
<br><br>
## Usage
<hr>

### Simple usage:
```bash
> .\anagramalyser.exe -file .\testfile.txt
Found anagrams for abc are [cba]
Found anagrams for cba are [abc]
```

### Advanced usage:
```bash
.\anagramalyser -h
Usage of anagramalyser:
  -buffersize int
        size of buffer to use, defaults to 2048*1024 (default 2097152)
  -file string
        location of file to analyse (default "./testfile.txt")
  -showempty
        set to true to show empty anagram matches
```

### Testing:
```bash
> go test -v ./...
=== RUN   TestMain
-> Building ...
OS is windows
-> Running test ...
--- PASS: TestMain (0.28s)
PASS
ok      github.com/mdhemmings/anagramalyser/cmd 0.529s
=== RUN   TestCheckAnagrams
1
1
--- PASS: TestCheckAnagrams (0.00s)
=== RUN   TestCheckIsAnagram
--- PASS: TestCheckIsAnagram (0.00s)
PASS
ok      github.com/mdhemmings/anagramalyser/cmd/anagram 0.212s
```

## Contributing
<hr>
Pull requests are welcome. Please make sure to update tests as appropriate.
<br><br>

## Licence
<hr>

[MIT](https://choosealicense.com/licenses/mit/)