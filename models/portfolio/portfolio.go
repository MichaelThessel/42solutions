package fortytwo

var portfolios = [...]Portfolio{
    Portfolio {  1, "Design, implementation, hosting.", "pop-700x687.png", "pop-tours.com", true },
    Portfolio {  2, "Open Source project, built in collaboration with others on GitHub.", "pwx-700x570.png", "pwx.io", true },
    Portfolio {  3, "Complete revamp, performance optimizations & feature additions.", "andrewchristian.com_-700x536.png", "www.andrewchristian.com", true },
    Portfolio {  4, "Complete revamp & feature additions.", "paperdivas.com_.au_-700x505.png", "paperdivas.com.au", true },
    Portfolio {  5, "Complete site revamp & feature additions.", "lovestruckinvitations.com_.au_-700x512.png", "lovestruckinvitations.com.au", true },
    Portfolio {  6, "Implemented NLM3 XML document conversion service.", "pkp-udev.lib_.sfu_.ca_-700x555.png", "pkp-udev.lib.sfu.ca", true },
    Portfolio {  7, "Member core development team.", "pkp.sfu_.ca_-700x565.png", "pkp.sfu.ca", true },
    Portfolio {  8, "Lead developer.", "karmaexchange.com_-700x530.jpg", "karmaexchange.com", false },
    Portfolio {  9, "Developed web fronted for CRM system.", "canadawide.com_-700x555.png", "canadawide.com", true },
    Portfolio { 10, "Functionality extensions & performance optimizations.", "prometheanworld.com_-700x467.png", "prometheanworld.com", true },
    Portfolio { 11, "Implementation of additional features & site maintenance", "coursework.info_-700x292.png", "coursework.info", true },
    Portfolio { 12, "Extensive extension.", "malaysia.com_-700x467.png", "malaysia.com", true },
    Portfolio { 13, "Lead developer. Complete site revamp & extension.", "thestudentroom.co_.uk_-700x526.png", "thestudentroom.co.uk", true },
    Portfolio { 14, "Implementation of web application to aid the migration to SAP.", "acg-world.com_-700x517.png", "acg-world.com", true },
    Portfolio { 15, "Implementation of framework to implement business logic for intranet and website.", "mpipks-dresden.mpg.de_-700x528.png", "mpipks-dresden.mpg.de", true },
}

type Portfolio struct {
    Id int
    Text string
    Image string
    Domain string
    DomainActive bool
}

func (p *Portfolio) Get(id int) bool {
    for i, portfolio := range portfolios {
        if portfolio.Id == id {
            *p = portfolios[i]

            return true
        }
    }

    return false
}

func (p *Portfolio) Next() int {
    next := p.Id + 1
    if next > len(portfolios) {
        next = 0
    }

    return next
}

func (p *Portfolio) Prev() int {
    return p.Id - 1
}
