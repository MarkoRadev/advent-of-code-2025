// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// func main() {
// 	dial := 50
// 	zeroCounts := 0
// 	file, err := os.Open("input.txt")

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		line := scanner.Text()
// 		if len(line) == 0 {
// 			continue
// 		}

// 		zeroClicks := 0
// 		direction := line[0]       
// 		distanceStr := line[1:]
// 		distance, err := (strconv.Atoi(distanceStr))
// 		if err != nil {
// 			panic(err)
// 		} else {
// 			zeroClicks = distance / 100
// 			distance = distance % 100
// 		}

// 		if direction == 'L' {

// 			if dial != 0 {
// 				if (dial - distance) > 0 {
// 					zeroCounts += zeroClicks
// 				} else if (dial - distance) < 0 {
// 					zeroCounts += zeroClicks + 1
// 				}
// 			} else {
// 				if (dial - distance) < 0 {
// 					zeroCounts += zeroClicks
// 				} else {
// 					zeroCounts = zeroCounts + zeroClicks - 1
// 				}
// 			}

// 			dial -= distance
// 			if dial < 0 {
// 				dial += 100
// 			}
// 		} else if direction == 'R' {

// 			if dial != 0 {
// 				if (dial + distance) < 100 {
// 					zeroCounts += zeroClicks
// 				} else if (dial + distance) > 100 {
// 					zeroCounts += zeroClicks + 1
// 				}
// 			} else {
// 				if (dial + distance) > 0 {
// 					zeroCounts += zeroClicks
// 				} else {
// 					zeroCounts = zeroCounts + zeroClicks - 1
// 				}
// 			}

// 			dial += distance
// 			if dial >= 100 {
// 				dial -= 100
// 			}
// 		}

// 		if dial == 0 {
// 			zeroCounts++
// 		}
// 	}

// 	if err := scanner.Err(); err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Password:", zeroCounts)
// }


package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	dial := 50
	zeroCounts := 0

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		direction := line[0]   
		distanceStr := line[1:]
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			panic(err)
		}

		if direction == 'L' {
			for i := 0; i < distance; i++ {
				dial--          
				if dial < 0 {   
					dial = 99
				}
				if dial == 0 {  
					zeroCounts++
				}
			}
		} else if direction == 'R' {
				for i := 0; i < distance; i++ {
					dial++         
					if dial >= 100 { 
						dial = 0
					}
					if dial == 0 {
						zeroCounts++
					}
				}
		} else {
			panic("invalid direction: " + string(direction))
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Password:", zeroCounts)
}
