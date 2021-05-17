## Convert mysql to Markdown

Go file to helps you convert mysql data to markdown files.

Example usage: migrating blog posts to SSG (static site generator). After generating markdown files, you can copy paste it to your SSG folder


## How to
1. create folder named "blogs"

2. setup your database and table name @DB_NAME and @TABLE_NAME

3. Customize your column name at -type Blog struct-

4. Customize your markdown files at func createMarkdownFiles

5. go run main.go

//it will generate all you markdown files inside blogs folder 

## Feature
You can convert html format to markdown files

## Customize

Don't forget to customize with you own need. Inside the main.go file, you can see the structure which reflects your column inside database.

Later, you can also see what if there is null value, columns is different for each project, that's why you need to customize it.

