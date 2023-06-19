## Code diagram

<img src="https://kroki.io/c4plantuml/svg/eNptUk1v2zAMvftXcEYPLlDUl54GDFgbYEOAFs2SDTsWqkTHQizR0Ee7YNh_HyXZqbP1RIniI9971GcfhAvRDNUHbeUQFcLq5mlFZiSLNlyP6akKOgwIu6MPaGBFXKS02DthoCMHtzafxXD06Kpqg86TbSJfrqD-waHmuCOD3BJedehBgEKvHWa4KHBfX1anud8iRmxeySnP2J8p1nBV35MUQZMF6rhHp5mUJBuEttruIZf_34ViGGPgNo_5kMh87xFKOnUKfBsdJRKMrorMpzuKVgl3bOQNIyaNX7RVWc_FoO3hU92HMPqPbbtnVfH5WpJpjerRGObjW7E0pr78XQGbl-miawwHbvzAAbZMhpOJ2lcahN0zj1R7puM5dpq4Yv0Iu-BQmGJsAbR36fUc1kzz3-jDqkd5KLguWpm8fG8UuyHRe1Rct5nPUNaQujknjoz7U1VbHKZNz8taW7Y1rSE9TclJ60bkPiz4RQtY3a-nqvI869uiUB6GedXCKvCBHHqQfbQHn_Bc2yVLF_NPa57w2uflLn5Xqj1ZspR4qxhwShQlPCUQiElpgi4Qb8aWYTLbOiM7R-bfcbO4vACf_z3-YopoJb5D89yRtX2hA-tPVqT_7mjgwr8hrEZz"></img>

Using this puml codeblock:
```
@startuml
!include C4_Component.puml

title System Code diagram for Anagramalyser

Person(user, "User", "Someone with a desire for anagrams")
ComponentQueue(words, "Words" ,"Location of a file containing words")
ComponentQueue(output, "Output", "The output of the program")

System_Boundary(c4, "AnagramFinder", $link="https://github.com/mdhemmings/anagramalyser"){
  Container(main, "Main Routine", "Golang")
  ComponentQueue(bufio, "IO Streamer", "Golang/Bufio")
  Component(anagram, "Anagram Checker", "function")
  ComponentQueue(processed, "Processed Words", "Array")
}

Rel(user, words, "Inputs")
Rel(words, main, "Passed in via CLI")
Rel(main, bufio, "Reads location and stores chunks in buffer")
Rel(user, output, "Reads list of anagrams")
Rel(anagram, processed, "Adds processed words into array")
Rel(processed, anagram, "Reads checked words from")
Rel(anagram, bufio, "Checks for existence of anagrams")
Rel(main, bufio, "Invokes and controls")
```

(If for whatever reason kroki is down, then refer to diagram.svg in docs folder)