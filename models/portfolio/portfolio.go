package fortytwo

var portfolios = [...]Portfolio{
	Portfolio{1, "Founder & Developer of real time application level intrusion detection system", "iogly-700x658.png", []string{"Golang", "jQuery", "SASS"}, "iogly.com", true},
	Portfolio{2, "Design, implementation, hosting", "pop-700x687.png", []string{"JS", "Angular2", "Burbon Neat", "SASS"}, "pop-tours.com", true},
	Portfolio{3, "Lead developer of open source project", "pwx-700x570.png", []string{"PHP", "MySQL", "Silex", "Bootstrap"}, "github.com/MichaelThessel/pwx", true},
	Portfolio{4, "Complete revamp, performance optimizations & feature additions", "andrewchristian.com_-700x536.jpg", []string{"PHP", "MySQL", "Memcache", "Redis", "Magento", "LESS"}, "www.andrewchristian.com", true},
	Portfolio{5, "Complete revamp, performance optimizations & feature additions", "paperdivas.com_.au_-700x505.png", []string{"PHP", "MySQL", "Memcache", "Redis", "Magento", "LESS"}, "paperdivas.com.au", true},
	Portfolio{6, "Complete revamp, performance optimizations & feature additions", "lovestruckinvitations.com_.au_-700x512.png", []string{"PHP", "MySQL", "Memcache", "Redis", "Magento", "LESS"}, "lovestruckinvitations.com.au", true},
	Portfolio{7, "Implementation of NLM3 XML document conversion service", "pkp-udev.lib_.sfu_.ca_-700x555.png", []string{"PHP", "MySQL", "Zend Framework 2", "Bourbon Neat", "SASS"}, "pkp-udev.lib.sfu.ca", true},
	Portfolio{8, "Member core development team", "pkp.sfu_.ca_-700x565.png", []string{"PHP", "MySQL", "Custom Framework"}, "pkp.sfu.ca", true},
	Portfolio{9, "Development of Groupon clone", "karmaexchange.com_-700x530.jpg", []string{"PHP", "MySQL", "Drupal 7"}, "karmaexchange.com", false},
	Portfolio{10, "Development of web fronted for CRM system", "canadawide.com_-700x555.png", []string{"PHP", "MySQL", "Zend Framework 1", "Bootstrap"}, "canadawide.com", true},
	Portfolio{11, "Functionality extensions & performance optimizations", "prometheanworld.com_-700x467.png", []string{"PHP", "MySQL", "Box UK CMS"}, "prometheanworld.com", true},
	Portfolio{12, "Implementation of additional features & site maintenance", "coursework.info_-700x292.png", []string{"Perl", "MySQL", "Gossamer Links Framework"}, "coursework.info", true},
	Portfolio{13, "Extensive feature additions & maintenance", "malaysia.com_-700x467.png", []string{"Perl", "MySQL", "Gossamer Links Framework"}, "malaysia.com", true},
	Portfolio{14, "Complete site revamp, performance optimizations & feature additions", "thestudentroom.co_.uk_-700x526.png", []string{"PHP", "MySQL", "Memcache", "vBulletin"}, "thestudentroom.co.uk", true},
	Portfolio{15, "Implementation of web application to aid the migration to SAP", "acg-world.com_-700x517.png", []string{"Ruby", "MySQL", "Rails"}, "acg-world.com", true},
	Portfolio{16, "Implementation of framework to implement business logic for intranet and website", "mpipks-dresden.mpg.de_-700x528.png", []string{"PHP", "MySQL", "Custom Framework"}, "mpipks-dresden.mpg.de", true},
}

type Portfolio struct {
	Id           int
	Text         string
	Image        string
	Technologies []string
	Domain       string
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
