package main

import "fmt"

/*
 * Complete the 'camelcase' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING s as parameter.
 */

func camelcase(s string) int32 {
    // Write your code here
    var wordCount int32
    
    for _,v := range s{
        if v>=65 && v<=90{
            wordCount++
        }
    }
    return wordCount+1
}

func main() {
    s := "theQuickBrownFoxJumpedOverTheLazyDog"
	fmt.Printf("Word count is: %d\n",camelcase(s))
}

