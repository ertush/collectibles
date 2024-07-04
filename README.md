# Collectibles
## Collection/List methods for some of the golang types
Have you ever wanted to have some handy methods that can operate on any of the collections types in go? here is the solution.

<hr>

## Methods
 - ```Index(<Tslice>, <val>)``` where Tslice is a slice of type ```T```. Returns an **Int** which is the index of val
 - ```Include(<Tslice>, <val>)``` where Tslice is a slice of type ```T```. Returns a **Boolean**
 - ```IsConsecutive(<Tslice>)``` where Tslice is a slice of type ```T```. Returns a **Boolean** ```true``` if elements in slice are consecutive
 - ```Compare(<Tslice1>, <Tslice2>, <func>)``` where Tslice1, Tslice2 are slices of type ```T``` and ```func``` a function to compare individual elements of the slices. Returns a **Boolean**
 - ```Map(<Tslice>, <func>)``` where Tslice is a slice of type ```T``` and func a function.Returns **T** applies ```func``` to each item in the slice.
 - ```Filter(<Tslice>, <func>)``` where Tslice is a slice of type ```T``` and ```func``` a function that has an expression . Returns a **slice of type T**
 - ```All(<Tslice>, <func>)``` Returns a **Boolean**. ```true``` if all elements in slice are not ```falsey``` i.e 0 or Null
 - ```Any(<Tslice>, <func>)``` Returns a **Boolean**. ```true``` if atleast 1 element in slice the is not ```truthy``` 

<hr>

## Example

```
package main

import (
        "fmt"
        "github.com/ertush/collectibles"
)


func square(b byte) byte {
	return b * b
}


func main() {

	var b *collectibles.Byt
        
        bs := []byte{1, 3, 5, 6, 7, 9}
              
        sq := b.Map(bs, square)
        
        fmt.Printf("\n\n%v, %v\n\n", sq, b.Include(sq, 36)) // [1 9 25 36 49 81], true
}

```
