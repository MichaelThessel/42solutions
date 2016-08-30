var testimonials = {
    el: null,
    id: null,
    current: 0,

    testimonialEl: null,
    decoratorEl: null,
    nameEl: null,
    textEl: null,

    testimonials: [
        {name: "Jack", text: "Michael is one of the most talented and detail-oriented programmers I've ever come across in almost 22 years online ..."},
        {name: "Rebecca", text: "I have been EXTREMELY happy with all the web development work you have done for us ..."},
        {name: "Robert", text: "By the way, I have nothing but glowing reviews for Michael and his level of expertise."},
    ],

    init: function(id) {
        this.id = id;
        this.el = document.getElementById(this.id);

        if (!this.el) { return; }

        this.decoratorEl = document.createElement('i');
        this.testimonialEl = document.createElement('p');
        this.textEl = document.createTextNode("");
        this.nameEl = document.createElement('span');

        this.testimonialEl.className = 'quote';
        this.decoratorEl.className = 'icon-bubble2';
        this.nameEl.className = 'name';

        this.el.appendChild(this.testimonialEl);

        this.testimonialEl.appendChild(this.decoratorEl);
        this.testimonialEl.appendChild(this.textEl);
        this.testimonialEl.appendChild(this.nameEl);

        this.run();
    },

    run: function() {
        this.update();
        setInterval(this.update.bind(this), 5000);
    },


    update: function() {
        this.current++;
        if (this.current >= this.testimonials.length) {
            this.current = 0;
        }

        this.textEl.nodeValue = this.testimonials[this.current].text;
        this.nameEl.innerHTML = this.testimonials[this.current].name;
    }
};
