package main

import "fmt"

type person struct {
	first string
}

// mongo, postg are maps that take int and store person struct
type mongo map[int]person
type postg map[int]person

// func (receiver) name(parameters)
func (m mongo) save(n int, p person) {
	m[n] = p
}

// func (receiver) name(parameters) returntype
func (m mongo) retrieve(n int) person {
	return m[n]
}

func (pg postg) save(n int, p person) {
	pg[n] = p
}
func (pg postg) retrieve(n int) person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retrieve(n int) person
}

type personService struct {
	a accessor
}

func (ps personService) get(n int) (person, error) {
	p := ps.a.retrieve(n)
	if p.first == "" {
		return person{}, fmt.Errorf("no person with n %d", n)
	}
	return p, nil
}

func put(a accessor, n int, p person) {
	a.save(n, p)
}
func get(a accessor, n int) person {
	return a.retrieve(n)
}

func main() {
	dbm := mongo{}
	dbp := postg{}

	p1 := person{
		first: "goama",
	}
	p2 := person{
		first: "kana",
	}

	ps := personService{
		a: dbm,
	}

	put(dbm, 1, p1)
	put(dbm, 2, p2)
	fmt.Println("2 records saved to mongo")
	fmt.Println(get(dbm, 1))
	fmt.Println(get(dbm, 2))
	put(dbp, 1, p1)
	put(dbp, 2, p2)
	fmt.Println("2 records saved to postgres")
	fmt.Println(get(dbp, 1))
	fmt.Println(get(dbp, 2))

	fmt.Println(ps.get(1))
	fmt.Println(ps.get(2))
	fmt.Println(ps.get(3))
}
