import React from "react";
import ReactDOM from "react-dom";

export class Testimonial extends React.Component {
    constructor() {
        super();

        this.testimonials = [
            {name: "Jack", text: "Michael is one of the most talented and detail-oriented programmers I've ever come across in almost 22 years online ..."},
            {name: "Rebecca", text: "I have been EXTREMELY happy with all the web development work you have done for us ..."},
            {name: "Robert", text: "By the way, I have nothing but glowing reviews for Michael and his level of expertise."},
        ];

        this.current = 0;
        this.state = {testimonial: this.testimonials[this.current]};

        setInterval(() => {
            this.changeState()
        }, 10000);
    }

    changeState() {
        this.current++;
        if (this.current >= this.testimonials.length) {
            this.current = 0;
        }

        this.setState({testimonial: this.testimonials[this.current]});
    }

    render() {
        return (
            <div class="testimonial quote">
              <i class="icon-bubble"></i>
              <p>{'"' + this.state.testimonial.text + '"'}<span class="name">{" -" + this.state.testimonial.name}</span></p>
            </div>
        );
    }
}

const t = document.getElementById('testimonials');
ReactDOM.render(<Testimonial/>, t);
