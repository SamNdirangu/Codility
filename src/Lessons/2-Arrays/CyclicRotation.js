//Cyclic Rotation
function solution(A,K){
    //An array to be shifted right by K times
    let rotationTimes = K;
    let integerArray = A;
    //if k is zero or less return the original array no change
    if(rotationTimes <= 0 || rotationTimes == integerArray.length)
        return integerArray;

    if(rotationTimes > integerArray.length) {
        //Avoid unneccessary long loop
        rotationTimes = rotationTimes % integerArray.length;
    }
    for (let index = 0; index < rotationTimes; index++) {
        const temp = integerArray.pop()  
        integerArray.unshift(temp);
    }

    return integerArray;
}

console.log(solution([3, 8, 9, 7, 6],1))