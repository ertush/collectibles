# Collectibles
## Collection/List methods for some of the golang types
Have you ever wanted to have some handy methods that can operate any collections types in go? here is the solution.

<hr>

## Methods
 - ```Index(<Tslice>, <val>)``` where Tslice is a slice of type ```T``. Returns an **Int** which is the index of val
 - ```Include(<Tslice>, <val>)``` where Tslice is a slice of type ```T```. Returns a **Boolean**
 - ```IsConsecutive(<Tslice>)``` where Tslice is a slice of type ```T```. Returns a **Boolean** ```true``` if elements in slice are consecutive
 - ```Compare(<Tslice1>, <Tslice2>, <func>)``` where Tslice1, Tslice2 are slices of type ```T``` and ```func``` a function to compare individual elements of the slices. Returns a **Boolean**
 - ```Map(<Tslice>, <func>)``` where Tslice is a slice of type ```T``` and func a function.Returns **Null** applies ```func``` to each item in the slice.
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


func main() {
	
	bs := []byte{1, 3, 5, 6, 7, 9}
	var b *collectibles.Byt
	
	
	fmt.Printf("%T, %v", b, b.Include(bs, 5)) // *collectibles.Byt, true
}

```