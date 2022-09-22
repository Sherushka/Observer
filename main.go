package awesomeProject1

func main() {
	hhkz := JobSite{}
	Shiro := Person{name: "Shiro"}
	hhkz.AddVacancies("Senior HTML Developer")
	hhkz.subscribe(&Shiro)
	hhkz.AddVacancies("Java Junior Developer")
}
