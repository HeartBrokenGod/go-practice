package mayuresh

type Author struct {
	FirstName string
	LastName  string
	Bio       string
}

type Post struct {
	Title   string
	Content string
	Author
}

//AuthorName is comment
func (a Author) AuthorName() string {
	return a.FirstName
}

/*
func main() {
	a := mayuresh.Author{
		FirstName: "mayuresh",
		LastName:  "thombare",
		Bio:       "writer",
	}
	p := mayuresh.Post{
		Title:   "myTitle",
		Content: "mycontent",
		Author:  a,
	}

	fmt.Println(a.AuthorName())
	fmt.Println(a, p, p.AuthorName())
}

mayuresh
{mayuresh thombare writer} {myTitle mycontent {mayuresh thombare writer}} mayuresh
*/
