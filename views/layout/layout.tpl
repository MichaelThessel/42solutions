<!DOCTYPE html>
<html>
  <head>
    <title>42solutions.io</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=yes"/>
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Raleway">
  </head>
  <body>
    <header>
      <div class="container">
        <div class="logo"><a href="/">42Solutions</a></div>
        <nav>
          <ul>
            <li><a href="/">{{if eq .ActiveNav "main"}}&mdash;{{else}}//{{end}} Home</a></li>
            <li><a href="/portfolio">{{if eq .ActiveNav "portfolio"}}&mdash;{{else}}//{{end}} Portfolio</a></li>
            <li><a href="/contact">{{if eq .ActiveNav "contact"}}&mdash;{{else}}//{{end}} Contact</a></li>
          </ul>
        </nav>
      </div>
    </header>
    <div class="container">
      <div class="content">
        {{.LayoutContent}}
      </div>
    </div>
    <script src="/static/js/script.min.js"></script>
  </body>
</html>
