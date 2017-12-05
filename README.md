# Eliza-Go

## Author: Matthew Shiel 

## Student ID: G00338622

This is my implementation of an Eliza chatbot in Go, based on the Eliza natural language processing program. This is for my 3rd year Data Representation and Querying module in GMIT. https://data-representation.github.io/problems/project.html

## How to Run The Chatbot

*Assumes that Git and Go are installed along with the prerequisites.*
**If not, they can be found from https://golang.org/dl/ and https://git-scm.com/downloads**

**1. Clone the Repository**
```bash
> git clone https://github.com/mattshiel/eliza-go.git
```
**2. Change Directory to the Folder**

```bash
Open the terminal/command line and navigate into the folder 
eg. > cd eliza-go
```

**2. Compile the Chatbot**

```bash
> go build server.go
```

**3. Run the Chatbot**

```bash
To run the chatbot enter './' followed by the executable produced
For Mac:
> ./server
For Windows:
> ./server.go.exe

Alternatively:
> go run server.go
```

**4. Open the localhost**
```bash
Go to your browser and type:
> 127.0.0.1:8080
```

## Design Components

The primary design components that went into this project were as follows, a web server to serve a HTML as the root, Javascript to get the user input with ajax and finally all functionality and responses are contained in the eliza package.


## Problems and Features

To give Eliza the impression she was thinking I implemented a delay to her responses, adding to this the delay will also randomise between 0.6 and 4 seconds per response. 

Eliza crafts her responses by examining the user's text input for a keyword. When the keyword is found it goes through a pre-processing procedure. This includes things like transforming all letters to lower case and trimming unnecessary whitespace. After this the word is tested against keywords in the 'responses.go' file. If a regex group is matched, the program gets the first match, the matched group then makes up a fraction of the response for added realism.
If Eliza does not find a match then there are a number of fallback default response she will chose from. All responses are chosen randomly.

A problem I encountered with this method is that selecting answers from the map of responses may be slightly non-deterministic, and the map will be out of order. In Weizenbaum's original design of Eliza he used a method of decomposition where all keywords held a certain value or 'precedence number'. This weighted approach is definitely another approach I could have taken when designing my program and something I might change in the future.


## References

This implementation of Eliza gave me many ideas and provided me with the main method of implementing my Eliza functionality https://github.com/kennysong/goeliza

I consulted both the the Golang, Bootstrap and JQuery documentation frequently over the course of my assignment.
