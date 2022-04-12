function solution(N, A) {
    //Create our counters
    let array = [...A];
    let counterMax = 0;
    let counterMaxTracker = counterMax
    let counters = new Array(N);
    //Loop 1;
    for (let index = 0; index < array.length; index++) {
        const X = array[index];
        const num = X - 1;
        //if true increase
        if (X <= N && X >= 1) {
            //decrease by 1 to match our counter index
            if(!counters[num]) //if empty
                counters[num] = counterMaxTracker;

            if (counters[num] < counterMaxTracker) //if less than the current max value
                counters[num] = counterMaxTracker + 1;
            else
                counters[num] += 1;
            //store our current max value
            if (counters[num] > counterMax) counterMax = counters[num];
        }
        //if true max all counters with max value
        if (X == N + 1)
            counterMaxTracker = counterMax;
    }
    //Loop 2
    //Now loop to update all values below our most recent max value
    for (let index = 0; index < counters.length; index++) {
        if(!counters[index]) //if empty
            counters[index] = counterMaxTracker;
    
        else if(counters[index] < counterMaxTracker)
            counters[index] = counterMaxTracker; 
    }
    return counters;
}

const testArray = [3, 4, 4, 6, 1, 4, 4];
console.log(solution(5, testArray))
//expect 7