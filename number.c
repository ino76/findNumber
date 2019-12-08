#include <stdio.h>
#include <stdlib.h>
#include <inttypes.h>
#include <limits.h>

/*
    This program should iterate through numbers and try to multiplicate
    its digits between themselves while counting how many multiplications
    is possible until there is only one digit
*/

// function that return length of a number
int getLen(uint64_t number)
{
    int count = 0;

    while (number > 0)
    {
        number /= 10;
        count++;
    }
    return count;
}

// function that multiply digits of a number and return result
uint64_t multiplyDigits(uint64_t number)
{
    uint64_t mult = 1;

    while (number > 0)
    {
        mult *= number % 10;
        number /= 10;
    }

    return mult;
}

void printMulti(uint64_t number)
{
    printf("\n\n  The number is: %" PRIu64 "  \n\n", number);
}

// Function for finding a number ...
void findNumber(int findMultiplicationOf)
{

    uint64_t number = 0;
    int maxMulti = 0;

    while (number < UINT64_MAX)
    {
        uint64_t currentNumber = number;
        int currentMulti = 0;

        while (getLen(currentNumber) > 1)
        {
            currentNumber = multiplyDigits(currentNumber);
            currentMulti++;
        }

        if (currentMulti > maxMulti)
        {
            maxMulti = currentMulti;
        }

        // if we found multi high like we specified log message and break a loop
        if (maxMulti >= findMultiplicationOf)
        {
            printf("\nNumber %" PRIu64 " has multiplication persistance of %d.\n\n", number, maxMulti);
            printMulti(number);
            break;
        }

        number++;
    }
}

int main(int argc, char *argv[])
{
    if (argc == 2)
    {
        findNumber(atoi(argv[1]));
    }
    else
    {
        printf("You need to call this program with some number a argument.\n");
    }
    return 0;
}