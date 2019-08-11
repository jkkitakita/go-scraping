package domain

type (
	// Site is a struct of a site to be scraped
	Site struct {
		Domain     string     `yaml:"domain"`
		Category   Category   `yaml:"category"`
		Pagination Pagination `yaml:"pagination"`
		Element    Element    `yaml:"element"`
	}

	// Category is a struct of a category of the site to be scraped
	Category struct {
		Name string `yaml:"name"`
		Path string `yaml:"path"`
	}

	// Pagination is a struct of a pagination of the site to be scraped
	Pagination struct {
		PerPage PerPage `yaml:"per_page"`
		Page    Page    `yaml:"page"`
	}

	// Element is a struct of elements of the site to be scraped
	Element struct {
		Name   string `yaml:"name"`
		Price  string `yaml:"price"`
		Detail Detail `yaml:"detail"`
	}

	// Detail is a struct of elements of the site to be scraped
	Detail struct {
		Table  string `yaml:"table"`
		Column string `yaml:"column"`
		Value  string `yaml:"value"`
	}

	// PerPage is a struct of per page of the site to be scraped
	PerPage struct {
		Suffix string `yaml:"suffix"`
		Number int    `yaml:"number"`
	}

	// Page is a struct of page of the site to be scraped
	Page struct {
		Suffix string `yaml:"suffix"`
		Offset int    `yaml:"offset"`
		Limit  int    `yaml:"limit"`
	}
)
