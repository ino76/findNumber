#=
    This program should iterate through numbers and try to multiplicate
    its digits between themselves while counting how many multiplications
    is possible until there is only one digit
=#


# Function for finding a number ...
function findNumber()

    findMultiplication = parse(UInt64, ARGS[1])
   
    number = UInt64(0)
    maxNumber = UInt64(typemax(UInt64))
    
    maxMulti = 0 
    
    while number < maxNumber    

        currentNumber = number
        currentMulti = 0
        
        while getLen(currentNumber) > 1
            currentNumber = multiplyDigits(currentNumber)
            currentMulti += 1
        end

        if currentMulti > maxMulti
            maxMulti = currentMulti
        end

        if maxMulti == findMultiplication
            println()
            println("Number $number has multiplication of $maxMulti.")
            println()
            break
        end

        number += 1
    end
end

# function that multiply digits of a number and return result
function multiplyDigits(number)
    return reduce(*, digits(number))
end

# function that return length of a number
function getLen(number)
    return length(digits(number))
end


@time findNumber()