package templating

type Course struct {
	Number string
	Name   string
	Units  string
}

type Semester struct {
	Term    string
	Courses []Course
}

type Year struct {
	AcaYear string
	Fall    Semester
	Spring  Semester
	Summer  Semester
}

//years := []year{
//{
//AcaYear: "2020-2021",
//Fall: semester{
//Term: "Fall",
//Courses: []course{
//course{"CSCI-40", "Introduction to Programming in Go", "4"},
//course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
//course{"CSCI-140", "Mobile Apps Using Go", "4"},
//},
//},
//Spring: semester{
//Term: "Spring",
//Courses: []course{
//course{"CSCI-50", "Advanced Go", "5"},
//course{"CSCI-190", "Advanced Web Programming with Go", "5"},
//course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
//},
//},
//},
//{
//AcaYear: "2021-2022",
//Fall: semester{
//Term: "Fall",
//Courses: []course{
//course{"CSCI-40", "Introduction to Programming in Go", "4"},
//course{"CSCI-130", "Introduction to Web Programming with Go", "4"},
//course{"CSCI-140", "Mobile Apps Using Go", "4"},
//},
//},
//Spring: semester{
//Term: "Spring",
//Courses: []course{
//course{"CSCI-50", "Advanced Go", "5"},
//course{"CSCI-190", "Advanced Web Programming with Go", "5"},
//course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
//},
//},
//},
//}
//
//err := tpl.ExecuteTemplate(os.Stdout, "myschedule.gohtml",years)
//if err != nil {
//log.Fatalln(err)
//}
