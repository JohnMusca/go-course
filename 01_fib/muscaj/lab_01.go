package main

import "fmt"

func main() {

    prev_value:= 0;
    next_value:= 1;

    for i:= 0; i<7; i++ {
        prev_value, next_value = fibonanci(prev_value, next_value);
        fmt.Printf("previous=%d next=%d\n",prev_value, next_value);
    }
}

//returns an int
func fibonanci(prev int, next int) (int, int) {

    new_prev:= next;
    new_next:= prev + next;

    return new_prev, new_next;
}
