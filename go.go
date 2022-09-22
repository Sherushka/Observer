package awesomeProject1

import "fmt"

type Observable interface {
	sendAll()
	subscribe(observer Observer)
	unsubscribe(observer Observer)
}

type Observer interface {
	handleEvent([]string)
}

type Person struct {
	name string
}

func (person *Person) Person(name string) {
	person.name = name
}

func (person *Person) handleEvent(vacancies []string) {
	fmt.Println("Hi dear ", person.name)
	fmt.Println("Vacancies updated: ")
	for _, vacancy := range vacancies {
		fmt.Println(vacancy)
	}
}

type JobSite struct {
	subscribers []Observer
	vacancies   []string
}

func (job *JobSite) subscribe(observer Observer) {
	job.subscribers = append(job.subscribers, observer)
}

func (job *JobSite) AddVacancies(vacancy string) {
	job.vacancies = append(job.vacancies, vacancy)
	job.SendAll()
}

func (job *JobSite) SendAll() {
	for _, observer := range job.subscribers {
		observer.handleEvent(job.vacancies)
	}
}
