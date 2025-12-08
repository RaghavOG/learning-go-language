package main

import "fmt"

// Run with: go run guide/examples/03_structs_interfaces.go
func main() {
	dev := Developer{
		Person: Person{Name: "Raghav", Age: 21},
		Skills: []string{"Go", "CLI", "APIs"},
	}

	dev.AddSkill("Testing")
	dev.Print()

	// Interface satisfaction is implicit
	var w Worker = dev
	w.Work()

	// Type assertion
	if d, ok := w.(Developer); ok {
		fmt.Println("asserted back to Developer with", len(d.Skills), "skills")
	}

	// Pointer vs value receivers
	counter := Counter{}
	counter.Inc()
	counter.Inc()
	fmt.Println("counter value:", counter.Value())
}

type Person struct {
	Name string
	Age  int
}

// Developer embeds Person; gains its fields/methods.
type Developer struct {
	Person
	Skills []string
}

// Value receiver (copy for read-only-ish)
func (d Developer) Print() {
	fmt.Printf("%s (%d) knows %v\n", d.Name, d.Age, d.Skills)
}

// Pointer receiver (mutates)
func (d *Developer) AddSkill(skill string) {
	d.Skills = append(d.Skills, skill)
}

type Worker interface {
	Work()
}

// Implicitly satisfies Worker.
func (d Developer) Work() {
	fmt.Println(d.Name, "is coding in Go")
}

type Counter struct {
	value int
}

func (c *Counter) Inc() { c.value++ }
func (c Counter) Value() int {
	return c.value
}

