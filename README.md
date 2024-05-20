# text-adventure
A text adventure written for CSCI-20 at Butte College

## Description
This is a text adventure game inspired by Lewis Carroll Puzzles and formal logic. You fall asleep at your computer and enter a dream world where you are asked to help solve some logic puzzles using tools such as Modus Ponens, Negation, Contrapositive, and AND statements. There are several rooms to explore, with descriptions for traveling and looking in different directions. There are puzzles in each room which are adapted from Lewis Carroll's book on formal logic, [Symbolic Logic](https://www.gutenberg.org/cache/epub/28696/pg28696-images.html#b1).

## How to Run
Simply clone the file and run with `go run .` to run the program.

## Overall Project Description
This was a somewhat ambitious project for an intro programming class, as it involved keeping track of many state changes and providing more user tools beyond those one might expect to find in a simple text adventure. The data structures built for the puzzles were the most difficult to get right, but I think I did a pretty good job here. Each time I added a new feature (modus ponens, contrapositive, etc), I needed to make sure the puzzle's state would update correctly, and allow the user to keep track of the changes they were making to the puzzle. This meant navigating golang's tools for changing data, while making sure I didn't leave any dangling pointers hanging around.

In addition to adding each of the above logic puzzle tools, I wanted to make the user experience as smooth as possible. Logic puzzles are confusing enough on their own without my interface adding more confusion, so I made sure to provide help messages for all of the logic puzzle tools. Additionally, I made sure that each of these tools had short-cut commands that would make the puzzle experience a bit more smooth, and tried to anticipate user confusion by having help messages pop up when the user entered empty input. The help messages are tailored to the state of the game - for example, while in a puzzle, one help message is shown, and while navigating the map, another help message is presented.

## Challenges
The biggest unexpected challenge I encountered was in what I had expected would be a relatively straightforward feature - printing slowly to terminal while allowing the user to speed up the text description if they entered input while the description was slowly printing. The goal was to print each character, one after the other, with a short time delay, to simulate the computer typing to the screen, but without forcing the user to wait for the text to print if they wanted to speed things up.

It turns out that there is no built in library which allows for golang to read from `stdin` while writing to `stdout`. Because each of these processes is independent and blocking on their threads, it required a multithreaded approach to solve. The code for my solution can be found in the [PrintSlow](https://github.com/josephleblanc/text-adventure/blob/main/myprint/print_slow.go) function. I solved the main problem of reading and writing from `stdin` and `stdout`, respectively, by making two threads, each with a function running inside. The function writing to `stdout` needs to receive data from the function reading from `stdin` so it can know to speed up if the user enters some input, so I needed to use channels that sent messages between threads and be careful not to cause any race conditions that might stall the program.

This feature came out pretty nicely, but I still was not able to completely solve one problem. That is, reading from `stdin` is a process that cannot be stopped early - it continues listening until an input is entered, which means the process writing to `stdout`, even if it manages to send a message to the other thread, cannot cancel its execution. The result is that the text is written slowly until the description ends, and then requires user input to continue. As this is a text adventure game where user input would be asked for at these points in any case, this did not result in too large a problem, but it did limit the user's ability to enter input requesting help or quiting in the middle of descriptions.

Overall, I learned a lot about how concurrent programming works in go through this feature, and was able to receive great advice both from an instructor at Butte College as well as help from some users on the golang slack channel.

## What I learned
I learned a lot about golang through this project 
- how to structure a project into different folders using golangs package system
- how to manage state with golang maps
- how to run concurrent programs with non-blocking messages sent across threads and no race conditions
- how to read and write from files
- usage of switch statements and select statements with default cases
- reading and writng user input

Additionally, I learned a little about how to design for the end user, for example
- how to write help messages for all user input
- allowing the user to view current state of a problem at all times
- providing help messages when the user may be confused
- adding features that just look pretty, like the slow print
On the whole, this was a good experience to learn about programming in golang.
