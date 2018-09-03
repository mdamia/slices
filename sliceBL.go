package sliceBL

func addPerson(people []string, p1 string, p2 string, p3 string) []string {
	people = append(people, p1)
	people = append(people, p2)
	people = append(people, p3)
	return people
}
