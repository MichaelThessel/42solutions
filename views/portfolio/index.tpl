<div class="portfolio">
  <div class="image">
    <img src="/static/img/portfolio/{{.Portfolio.Image}}" />
  </div>

  <div class="meta">
    <div class="pager">
      <span class="prev">
        {{if .Prev}}
          <a href="/portfolio/{{.Prev}}"><i class="icon-forward"></i></a>
        {{else}}
          <span class="disabled"><i class="icon-forward"></i></span>
        {{end}}
      </span>

      <span class="next">
        {{if .Next}}
          <a href="/portfolio/{{.Next}}"><i class="icon-forward"></i></a>
        {{else}}
          <span class="disabled"><i class="icon-forward"></i></span>
        {{end}}
      </span>
    </div>

    <a class="link" target="_blank" href="http://{{.Portfolio.Domain}}">{{.Portfolio.Domain}}</a>
    <p class="description">{{.Portfolio.Text}}</p>
  </div>

</div>
