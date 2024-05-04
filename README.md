# Go-blog
Blog in golang with supports in markdown etc.

## Build from source
```
git clone https://github.com/Grewather/go-blog/
cd go-blog
go build .
```


## Usage
Go to/create ```.env``` fiie.
```
CODE="" # code for admin panel etc. should be long and strong
BLOG_TITLE="go-blog" # name of the blog
DSN="" # database connection string
```
Enter values you want here.
To change about go to /about/about.md and edit it as you want.
You also need to setup mysql database.
After that you can run it, and it should be working.

### admin panel (add,edit,delete)
go to /admin/ 

