# findNumber
inspired by numberphile video: [youtube - numberphile -> number 277777788888899](https://www.youtube.com/watch?v=Wim9WJeDTHQ)

This program will find number of specified multiplication persistance (MP).
By multiplication persistance I mean that for example if we have number "923" ... it has
1) 9 * 2 * 3 = 54
2) 5 * 4 = 20
3) 2 * 0 = 0
.. MP of 3.
In this program you must specify MP and program returns first number of that specified MP.

For example:
```
go run number.go 4 // find and print number 77
```

keep in mind that biggest found multi. is 11 (number 277777788888899) [youtube -> numberphile](https://www.youtube.com/watch?v=Wim9WJeDTHQ)
so if you want to find for example 10 it will take several minutes ...

![screen](https://github.com/ino76/findNumber/blob/master/screen.png)

How to run/build:

1) If you doesnt have installed go programming language, you must go to [golang.org](https://golang.org/), download it and install it on your computer.

2) download this repository 

3) open terminal and go to this repository

4) build executable with
```
go build number.go
```
and then run it by (on windows, linux and max will have different extension)
```
./number.exe 4
```
or just run it like 
```
go run number.go 4
```

PS: I've added Julia, C and Python versions .. maybe I'll add JS version in future.
