function solution(A) {
    //Find the number of pairs of passing cars
    let array = [...A];
    let Pcars = 0; //if 0

    let pairs = 0;
    for (let i = 0; i < array.length; i++) {
        if (array[i] == 0)
            Pcars++;
        else
            pairs += Pcars;
            
        if (pairs > 1000000000) {
            pairs = -1;
            break;
        }
        //console.log('pcars ' +Pcars + ' ::: ' +pairs);
    }

    return pairs;
}

const testArray = [0, 1, 0, 1, 1];
console.log(solution(testArray))
//expect 7