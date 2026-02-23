# GoNotes 2:39pm 06122025

Variables in Go are passed by value (except for a few data types we haven't covered yet). "Pass by value" means that when a variable is passed into a function, that function receives a copy of the variable. The function is unable to mutate the caller's original data.

# 07122025

A function can return a value that the caller doesn't care about. We can explicitly ignore variables by using an underscore, or more precisely, the blank identifier _. In Go, the blank identifier isn't just a convention; it's a real language feature that completely discards the value.

# 09122025

The example above is much easier to read and understand. When writing code, it’s important to try to reduce the cognitive load on the reader by reducing the number of entities they need to think about at any given time.

A guard clause is an early return from a function when a given condition is met.
Go supports first-class and higher-order functions, which are just fancy ways of saying "functions as values". Functions are just another type -- like ints and strings and bools.


# 05012026

Ranges in Go iterate over elements of a dataset - gopher noobs (YouTube)
In Go, the for range statement is a built-in mechanism used to iterate over elements in various data structures like arrays, slices, maps, strings, and channels.
You can omit the index or value using the blank identifier _ if you only need one of the values: