import sys


def findNumber():

    findMultiplicationOf = int(sys.argv[1])

    number = 0
    maxMulti = 0

    while number < 99999999999999:
        currentNumber = number
        currentMulti = 0

        while getLen(currentNumber) > 1:
            currentNumber = multiplyDigits(currentNumber)
            currentMulti += 1

        if currentMulti > maxMulti:
            maxMulti = currentMulti

        if maxMulti == findMultiplicationOf:
            print(
                f"Number { number } has multiplication persistance of { maxMulti }.")
            printMulti(number)
            break

        number += 1


def printMulti(n):
    print("----------------------------------")
    print(f"n:  {n}")
    print("----------------------------------")
    currentMulti = 1
    while getLen(n) > 1:
        n = multiplyDigits(n)
        print(f"{currentMulti}.  {n}")
        currentMulti += 1


def getLen(n):
    return len(str(n))


def multiplyDigits(n):
    mult = 1
    for x in str(n):
        mult *= int(x)
    return mult


findNumber()
