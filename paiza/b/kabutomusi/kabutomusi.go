package main

import (
	"fmt"
)

type kabutomusi struct {
	coler      string
	position   int
	lightGroup []string
}

type light struct {
	isLeft    bool
	lightType string
}

func main() {
	var (
		n                int
		red, green, blue int
		lights           []light
		check            bool = false
	)

	fmt.Scan(&n)
	fmt.Scan(&red, &green, &blue)

	red_kabuto := kabutomusi{coler: "red", position: red, lightGroup: []string{"R", "Y", "M", "W"}}
	green_kabuto := kabutomusi{coler: "green", position: green, lightGroup: []string{"G", "Y", "C", "W"}}
	blue_kabuto := kabutomusi{coler: "blue", position: blue, lightGroup: []string{"B", "M", "C", "W"}}

	for i := 0; i < n; i++ {
		var x string
		var y string

		fmt.Scan(&x, &y)

		if x == "L" {
			lights = append(lights, light{isLeft: true, lightType: y})
			continue
		}
		lights = append(lights, light{isLeft: false, lightType: y})
	}

	for _, v := range lights {

		moveKabuto(&red_kabuto, v)
		moveKabuto(&green_kabuto, v)
		moveKabuto(&blue_kabuto, v)
	
		if positionCheck(red_kabuto.position, green_kabuto.position, blue_kabuto.position) {
			check = true
			break
		}
	}
	
	if check {
		fmt.Printf("%d\n", red_kabuto.position)
		return
	}
	fmt.Printf("no\n")
}

func moveKabuto(kabuto *kabutomusi, light light) {
	if contains(kabuto.lightGroup, light.lightType) {
		if light.isLeft {
			kabuto.position -= 1
		} else {
			kabuto.position += 1
		}
	}
}

func contains(x []string, y string) bool {
	for _, v := range x {
		if v == y {
			return true
		}
	}
	return false
}

func positionCheck(x, y, z int) bool {
	if x == y && x == z {
		return true
	}
	return false
}
