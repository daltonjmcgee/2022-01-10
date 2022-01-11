# 2022-01-10

### Technologies

- [Golang](https://go.dev/)

### Resources and details

This project comes out of my interest in building a super-light HTTP Server that works with GET methods and uses no dependencies. I want it to have some easy to use feature that allow you to build a website with HTML/CSS and JS when you don't need to use a database of any sort.

### Features
- URL -> `public/[filename].[html]` mapping. If you go to `website.com/hello` you'll be served the file from `./public/hello.html`. This works with subdirectories as well, e.g. `website.com/pages/hello` will serve `./public/pages/hello.html`.
