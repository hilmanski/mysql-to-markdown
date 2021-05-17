## Convert mysql to Markdown

Go file to helps you convert mysql data to markdown.
My case is I want to move my blog posts to SSG (static site generator). 

## Run go file
//create folder named "blogs"

go run main.go

//it will create all you markdown files inside blogs folder 

## Feature
You can convert html format to markdown files

## Customize

Don't forget to customize with you own need. Inside the main.go file, you can see the structure which reflects your column inside database.

Later, you can also see what if there is null value, columns is different for each project, that's why you need to customize it.

