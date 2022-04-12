function solution(A, B, K) {
    //This functions gets the number of integers divisible by K between A and B
    //Seems a mathetical problem this one avoid loops as possible
    let divisibleIntegers = 0;

    if(A%K == 0){
        divisibleIntegers++;
        divisibleIntegers += Math.floor((B-A)/K);
    } else {
        let num = A + 1;
        let keepFinding = true;
        while(keepFinding){
            if(num%K == 0){
                divisibleIntegers++;
                divisibleIntegers += Math.floor((B-num)/K);
                break;
            }
            num++;
        }
    }
    return divisibleIntegers;
}

console.log(solution(5, 22, 2))
//expect 7