package lobsters

import "fmt"
import "os"
import "testing"


func TestPosts(t *testing.T) {
	file, err := os.Open("rss/test/programming.rss")
	if err != nil {
		return
	}
	defer file.Close()

	posts, err := DecodePosts(file)

	for _, post := range posts {
		fmt.Println(post.PostDate.UTC())
	}
}
