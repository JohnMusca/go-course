package main

import "fmt"

func main() {
    fmt.Println(bubble([]int{3, 2, 1, 5}))
}

func bubble(s []int) []int {

    for i := 0; i < len(s); i++ {
        for j := 0; j < len(s)-1-i; j++ {
			
            if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
    
    return s;
}
