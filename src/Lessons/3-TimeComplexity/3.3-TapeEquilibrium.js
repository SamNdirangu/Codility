function solution(A){
    let array = [...A]; //to get code pred
    let totalSum = 0;
    let totalDeductionSum = 0;
    array.forEach(num => totalSum+=num);


    let minDiff;
    for (let index = 0; index < array.length-1; index++) {
        totalDeductionSum += array[index];
        totalSum -= array[index]
        console.log(totalDeductionSum+' :: '+totalSum);
        const tempDiff = Math.abs(totalDeductionSum - totalSum);

        if(minDiff == undefined) minDiff = tempDiff;

        if(tempDiff == 0) {
            minDiff = tempDiff;
            break;
        }
        if(tempDiff < minDiff) minDiff = tempDiff;
    }
    return minDiff;
}
const testArray = [3,1,2,4,3];
console.log(solution(testArray))
//expect 1