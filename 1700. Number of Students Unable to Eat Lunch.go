package main

func CountStudents(students []int, sandwiches []int) int {

	st := make(map[int]int)

	for i := range students {
		st[students[i]]++
	}

	for {
		if len(students) == 0 {
			return 0
		}
		if ok := st[sandwiches[0]]; ok == 0 {
			return len(students)
		}
		if students[0] == sandwiches[0] {
			k := students[0]
			students = students[1:]
			sandwiches = sandwiches[1:]
			st[k]--
		} else {
			kaprizniy := students[0]
			students = students[1:]
			students = append(students, kaprizniy)
		}
	}

}
