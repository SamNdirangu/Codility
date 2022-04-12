function solution(X, A) {
    //Find how many seconds to get to X in array A;
    let array = [...A];
    let seconds = -1;
    let stepTracker = ((1 + X) * X) / 2;
    let stepsTaken = new Array(X).fill(false); //loop 1;

    //loop 2
    for (let index = 0; index < array.length; index++) {
        const element = array[index];
        if (element <= X && !stepsTaken[element - 1]) {
            stepsTaken[element - 1] = true;
            stepTracker -= element;
        }
        if (stepTracker == 0) {
            seconds = index;
            break;
        }
    }
    return seconds;
}
const testArray = [1, 3, 1, 3, 1, 2, 1, 3];
console.log(solution(3, testArray))
//expect 7
