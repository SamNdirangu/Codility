//Perm Missing Elem
//Array A contains N integers of [1..(n+1)]
//find the missing integer
function solution(A) {
    let array = A.sort((a, b) => a - b);

    let missingInteger = 0;
    if (array.length < 1) return 1;
    console.log()
    for (let index = 0; index < array.length; index++) {
        const integer = index + 1;
        console.log(index)
        if (array[index] !== integer) {
            missingInteger = integer;
            break;
        }
    }

    if (missingInteger == 0) missingInteger = array[array.length - 1] + 1;

    return missingInteger;
}

console.log(solution([1, 2]))