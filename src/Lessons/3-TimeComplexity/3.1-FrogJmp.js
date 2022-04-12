//FrogJump
//Get from X to greater/Equal than Y Each jump size is D
function solution(X,Y,D) {
    const distance = Y - X;
    const steps = Math.ceil(distance/D);

    return steps;
}

console.log(solution(10,10,30))