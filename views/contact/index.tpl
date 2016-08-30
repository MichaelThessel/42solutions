<div class="contact">
  {{if .IsValid}}
    <div class="message info">Thanks for getting in touch. I will get back to you shortly.</div>
  {{else}}
    <div class="contact-form">
      <form action="/contact" method="post" novalidate>
        <div class="field required">
          <label for="name">Name</label>
          <input type="text" name="name" id="name" value="{{.Message.Name}}" />
          {{if .Errors.name}}
            <ul class="errors">{{range $i, $e := .Errors.name}}<li>{{$e}}</li>{{end}}</ul>
          {{end}}
        </div>
        <div class="field required">
          <label for="email">Email</label>
          <input type="email" name="email" id="email" value="{{.Message.Email}}" />
          {{if .Errors.email}}
            <ul class="errors">{{range $i, $e := .Errors.email}}<li>{{$e}}</li>{{end}}</ul>
          {{end}}
        </div>
        <div class="field">
          <label for="telephone">Tel</label>
          <input type="tel" name="telephone" id="telephone" value="{{.Message.Telephone}}" />
          {{if .Errors.telephone}}
            <ul class="errors">{{range $i, $e := .Errors.telephone}}<li>{{$e}}</li>{{end}}</ul>
          {{end}}
        </div>
        <div class="field required">
          <label for="message">Message</label>
          <textarea name="message" id="message">{{.Message.Message}}</textarea>
          {{if .Errors.message}}
            <ul class="errors">{{range $i, $e := .Errors.message}}<li>{{$e}}</li>{{end}}</ul>
          {{end}}
        </div>
        <button type="submit" class="button"><span>Send</span></button>
      </form>
    </div>

    <div class="address">
      <address>
        42Solutions @ (co)space<br />
        2nd floor, 202 Strickland St.<br />
        Whitehorse, YT Y1A 2J8<br />
        Canada<br /><br />

        Tel: +1 867 332 9173<br />
        Email: <a href="mailto:michael@42solutions.io">michael@42solutions.io</a>
      </address>
    </div>
  {{end}}
</div>
