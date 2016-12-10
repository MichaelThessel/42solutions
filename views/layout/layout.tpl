<!DOCTYPE html>
<html>
  <head>
    <title>42solutions.io</title>
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=yes"/>
    <link rel="stylesheet" type="text/css" href="/static/css/style.css">
    <link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Raleway">
    <link rel="shortcut icon" href="static/favicon.ico" type="image/x-icon">
    <link rel="icon" href="static/favicon.ico" type="image/x-icon">
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
    <script src="/static/js/script-min.js"></script>
    <script>
      (function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
      (i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
      m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
      })(window,document,'script','https://www.google-analytics.com/analytics.js','ga');
      ga('create', 'UA-88812654-1', 'auto');
      ga('send', 'pageview');
    </script>
  </body>
</html>
