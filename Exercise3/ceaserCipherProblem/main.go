package main

import "fmt"

/*
 * Complete the 'caesarCipher' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts following parameters:
 *  1. STRING s
 *  2. INTEGER k
 */

func caesarCipher(s string, k int32) string {
    // Write your code here
    k%=26
    var result string
    for _,v := range s{
        if v>=65 && v<=90{
            a := v+k
            if a>90{
                a = 64 + (a-90)
            }
            result+=string(a)
        }else if v>=97 && v<=122{
            a := v+k
            if a>122{
                a = 96 + (a-122)
            }
            result+=string(a)
        }else{
            result+=string(v)
        }
        
    }
    return result

}

func main() {
    s := "There is nothing but water everywhere :- quoteCentral"
	k := 97

	fmt.Printf("Cipher text is: %s\n",caesarCipher(s,int32(k)))
}
