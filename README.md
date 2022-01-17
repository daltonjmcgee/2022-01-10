# 2022-01-10

### Technologies

- [Golang](https://go.dev/)

### Resources and details

This project comes out of my interest in building a super-light HTTP Server that works with GET methods and uses no dependencies. I want it to have some easy to use feature that allow you to build a website with HTML/CSS and JS when you don't need to use a database of any sort.

### Features
- URL -> `public/[filename].[html]` mapping. If you go to `website.com/hello` you'll be served the file from `./public/hello.html`. This works with subdirectories as well, e.g. `website.com/pages/hello` will serve `./public/pages/hello.html`.
- Static files being served from the /static/ folder.
- Dynamic filenames that map to a key/value pair in a noSQL JSON file. e.g. [id].html will look for an "id" tag in all of the entries in the JSON "database". [title].html will look for the 'title' tag in all entries. First entry found is what is returned. In theory IDs should be a UUID so there is never a possible conflict.

@TODO
- Add config file
- Make compatible with Svelte.
