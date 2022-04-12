//Binary Gap
//Time complexity 0^N
function solution(someNumber) {
	var numberBinary = someNumber.toString(2); //Convert to a binary format string

	var gapSize = 0; //Store our current largest gap
	var tempGapSize = 0; //Temporary store our working gapsize
	//Start our loop
	for (let index = 0; index < numberBinary.length; index++) {
		const bit = numBinary[index];
		if (bit == '1') {
			if (tempGapSize > gapSize) {
				//if our tempgapsize is bigger than our gapsize update
				gapSize = tempGapSize;
			}
			// Reset to zero
			tempGapSize = 0;
		} else {
			//Add another our size with one as the bit is another zero
			tempGapSize++;
		}
	}
	return gapSize;
}