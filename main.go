package main

import (
    "strconv"
    "os"
    "io"
    "strings"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    md "github.com/JohannesKaufmann/html-to-markdown"
)

//Custom your own column here
//use sql.NullString for optional value
type Blog struct {
    Id string
    Title string
	Slug string
	Body string
    Excerpt string
	Tags sql.NullString
    Featured_image sql.NullString
    Published bool
    Markdown bool
    Created_at string
}

func main() {

    //Connect mysql
    db, err := sqlx.Connect("mysql", "root:root@tcp(127.0.0.1:3306)/hugodb?parseTime=true")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    //Get all data from DB
    var blogs []Blog
    err = db.Select(&blogs, "SELECT * FROM blogs")
    if err != nil {
        panic(err.Error()) 
    }

    for _, blog := range blogs {

        // For meta data add double quote
        _title := "\"" + blog.Title + "\""
        _slug := "\"" + blog.Slug + "\""
        _description := "\"" + getMetaSubject(blog.Excerpt, blog.Body) + "\""
        _created_at := "\"" + blog.Created_at + "\""

        _draft := false
        if blog.Published == true {
            _draft = true
        }

        //For possible null values
        _tags := ""
        if blog.Tags.Valid {
            _tags  = "\"" + blog.Tags.String + "\""
        }
        //For possible null values
        _featured_image := ""
        if blog.Featured_image.Valid  {
            _featured_image = "\"" + blog.Featured_image.String + "\""
        }

        _fileName := blog.Slug
        _subject := getBody(blog.Body, blog.Markdown)

        //markdown file structure, 
            //you can customize it
        text := "+++\n" +
                "title = "+ _title + "\n" +
                "slug = "+ _slug + "\n" +
                "description = "+ _description + "\n" +
                "published_date = "+ _created_at + "\n" +
                "featured_image = "+  _featured_image + "\n" +
                "tags = "+ _tags + "\n" +
                "draft = "+ strconv.FormatBool(_draft) + "\n" +
                "+++ \n" +
                _subject

        if err := WriteStringToFile(_fileName+".md", text); err != nil {
          panic(err)
        }
    }   

}

func WriteStringToFile(filepath, s string) error {
	fo, err := os.Create("./blogs/"+filepath)
	if err != nil {
		return err
	}
	defer fo.Close()

	_, err = io.Copy(fo, strings.NewReader(s))
	if err != nil {
		return err
	}
	return nil
}

//Check if body is html turn it to markdown
func getBody(text string, isMarkdown bool) string {
    if(isMarkdown == true) {
        return text
     }   
    
    //for makrdown needs
    converter := md.NewConverter("", true, nil)
    markdown, err := converter.ConvertString(text)
    if err != nil {
        panic(err.Error()) 
    }
    
    return markdown
}

//Create excerpt if not exists
func getMetaSubject(excerpt string, body string) string {
    if(excerpt != "") {
        return excerpt
    }

    body = strings.Replace(body, "\n", "", -1)
    body = strings.Replace(body, "\"", "'", -1) //prevent if any double quote in meta
    if len(body) < 120 {
        return body
    } else {
        rune := []rune(body)
        return string(rune[0:120])
    }
}