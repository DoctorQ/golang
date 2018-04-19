package main
import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}



func main(){
        var book1 Books
                var book2 Books

                book1.title = "Go语言"
                book1.author = "www.test.com"
                book1.subject = "1"
                book1.book_id = 123


                book2.title = "Go语言1"
                book2.author = "www.t1est.com"
                book2.subject = "2"
                book2.book_id = 1233

                /* 打印 book1 信息 */
                fmt.Printf( "Book 1 title : %s\n", book1.title)
                fmt.Printf( "Book 1 author : %s\n", book1.author)
                fmt.Printf( "Book 1 subject : %s\n", book1.subject)
                fmt.Printf( "Book 1 book_id : %d\n", book1.book_id)

                /* 打印 book2 信息 */
                fmt.Printf( "Book 2 title : %s\n", book2.title)
                fmt.Printf( "Book 2 author : %s\n", book2.author)
                fmt.Printf( "Book 2 subject : %s\n", book2.subject)
                fmt.Printf( "Book 2 book_id : %d\n", book2.book_id)
                var b_p *Books
    b_p = &book1
   fmt.Printf("title:%s\n",b_p.title) 

}
