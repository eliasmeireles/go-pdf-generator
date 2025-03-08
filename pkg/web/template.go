package web

var Template = `
<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8"/>
    <link rel="icon" href="{{.WebProviderHost}}/favicon.ico"/>
    <meta name="viewport" content="width=device-width,initial-scale=1"/>
    <meta name="theme-color" content="#000000"/>
    <meta name="description" content="Web site created using create-react-app"/>
    <link rel="apple-touch-icon" href="{{.WebProviderHost}}/logo192.png"/>
    <link rel="manifest" href="{{.WebProviderHost}}/manifest.json"/>
    <script>window.postsData = {{.PostsData}}</script>
    <title>React App</title>
    <script defer="defer" src="{{.WebProviderHost}}/static/js/main.js"></script>
    <link href="{{.WebProviderHost}}/static/css/main.css" rel="stylesheet">
</head>
<body>
<noscript>You need to enable JavaScript to run this app.</noscript>
<h1>React App</h1>
<div id="root"></div>
</body>
</html>
`
