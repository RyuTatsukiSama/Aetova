package docSlices

import (
	"fmt"
	"reflect"
)

func Slice_main() {
	/* Define an array containing programming languages */
	var languages [9]string = [9]string{
		"C", "Lisp", "C++", "Java", "Python",
		"JS", "Ruby", "Go", "Rust",
	}

	/* define slices */
	classics := languages[0:3]  // alternatively languages[:3]
	modern := make([]string, 4) // len(modern) = 4
	modern = languages[3:7]     // include 3 exclude 7
	new := languages[7:9]       // alternatively languages[7:]

	fmt.Printf("classic languages: %v type %T\n", classics, classics)
	fmt.Printf("modern languages: %v type %T\n", modern, modern)
	fmt.Printf("new languages: %v type %T\n", new, new)

	allLangs := languages[:]                     //copy of the array
	fmt.Println(reflect.TypeOf(allLangs).Kind()) // slice

	/* Crea a slice containing web framworks */
	frameworks := []string{
		"React", "Vue", "Angular", "Svelte",
		"Laravel", "Django", "Flask", "Fiber",
	}

	jsFramework := frameworks[0:4:4]          // length 4 capacity 4
	frameworks = append(frameworks, "Meteor") // not possible with arrays

	fmt.Printf("all frameworks: %v type %T\n", frameworks, frameworks)
	fmt.Printf("js frameworks: %v type %T\n", jsFramework, jsFramework)
}
