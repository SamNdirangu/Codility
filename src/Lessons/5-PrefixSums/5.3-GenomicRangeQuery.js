function solution(S, P, Q) {
    //Find the minimum factor
    //A = 1 C = 2 G = 3 T = 4
    let arrayP = [...P];
    let arrayQ = [...Q];
    let genomeString = String(S);
    let impactFactors = [];
    arrayP.forEach((num, index) => {
        let factor = 4;

        for (let i = num; i <= arrayQ[index]; i++) {
            const letter = genomeString[i];
            if (letter == 'A') {
                factor = 1;
                break;
            }
            switch (letter) {
                case 'C': if (factor > 2) factor = 2;
                    break;
                case 'G': if (factor > 3) factor = 3;
                    break;
                default:
                    break;
            }
        }
        impactFactors.push(factor);
    });
    return impactFactors;
}

console.log(solution('A', [0], [0]))
//expect 7