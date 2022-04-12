
function solution(A) {
    //Check to see if Array A is a permutation, ie has all numbers 1 to N
    let array = [...A];
    const maxNum = array.length; //ideally the largest number should be the length of the array
    const expectedSum = ((1 + maxNum) * maxNum) / 2; //If we add all the numbers up in permutation series
    let permsTracker = new Array(maxNum).fill(true); //Create an array of similar in length
    
    let sum = 0; // we work to get the sum
   
    for (let index = 0; index < array.length; index++) {
        const num = array[index];
        if(permsTracker[num-1]){
            permsTracker[num-1] = false;
            sum += num;
        }
        //if array contains duplicate or large numbers outside the permutation range
        //The sum will be higher  and automatically fail
        if (sum > expectedSum) 
            break;
    }
    if(sum != expectedSum) return 0;
    return 1;
}
const testArray = [1, 4, 2,3];
console.log(solution(testArray))
//expect 7