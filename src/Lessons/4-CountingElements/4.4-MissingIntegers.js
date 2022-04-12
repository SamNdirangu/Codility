function solution(A) {
    //Find missing integers within the array
    let array = [... A.sort((a, b) => a - b)];

    let missingInteger = 1;
    for (let index = 0; index < array.length; index++) {
        const integer = array[index];

        if(integer > 0){
            if(integer != missingInteger){
                if(array[index-1] == missingInteger)
                    missingInteger++;
                else
                    break;
            }
        }     
    }
    if(missingInteger == array[array.length-1]) missingInteger++;
    return missingInteger;
}

const testArray = [-2,-1];
console.log(solution(testArray))
//expect 7