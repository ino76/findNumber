def findNumber():

    findMultiplicationOf = 9

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
            break

        number += 1


def getLen(n):
    return len(str(n))


def multiplyDigits(n):
    mult = 1
    for x in str(n):
        mult *= int(x)
    return mult


findNumber()
