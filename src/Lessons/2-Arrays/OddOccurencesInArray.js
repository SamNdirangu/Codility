//Odd Occurences in Array
function solution(A) {
    //Return integer within Array with no pair
    let integerArray = A;
    //if array contains only one number return the number
    if(integerArray.length == 1) return A[0];
    //First sort our array
    integerArray.sort();
    //Loop through
    let oddNumber = 0;
    for (let index = 1; index < integerArray.length; index+=2) {
        const number = integerArray[index];
        const prevNumber = integerArray[index-1];

        if(number != prevNumber){
            oddNumber = prevNumber;
            break;
        }
    }
    //incase the last number in the array is the odd one
    if(oddNumber == 0) oddNumber = integerArray[integerArray.length-1];
    return oddNumber;
}

console.log(solution([2, 2, 3, 3, 4,4,6,6,7]))