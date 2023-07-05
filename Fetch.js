

var api = "https://jsonplaceholder.typicode.com/posts"

fetch(api)
    .then(function (res) {
        return res.json();
    })
    .then(function (posts) {
        console.log(posts)
    })