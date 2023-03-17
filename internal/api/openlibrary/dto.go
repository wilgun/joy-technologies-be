package openlibrary

import "time"

type UserGetBookRequest struct {
	Subject string
}

type Author struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Availability struct {
	Status              string     `json:"status"`
	AvailableToBrowse   bool       `json:"available_to_browse"`
	AvailableToBorrow   bool       `json:"available_to_borrow"`
	AvailableToWaitlist bool       `json:"available_to_waitlist"`
	IsPrintDisabled     bool       `json:"is_printdisabled"`
	IsReadable          bool       `json:"is_readable"`
	IsLendable          bool       `json:"is_lendable"`
	IsPreviewable       bool       `json:"is_previewable"`
	Identifier          string     `json:"identifier"`
	ISBN                *string    `json:"isbn"`
	OCLC                *int       `json:"oclc"`
	OpenLibraryWork     string     `json:"openlibrary_work"`
	OpenLibraryEdition  string     `json:"openlibrary_edition"`
	LastLoanDate        *time.Time `json:"last_loan_date"`
	LastWaitlistDate    *time.Time `json:"last_waitlist_date"`
	IsRestricted        bool       `json:"is_restricted"`
	IsBrowseable        bool       `json:"is_browseable"`
	SRC                 string     `json:"__src__"`
}

type Work struct {
	Title             string   `json:"title"`
	CoverId           *int     `json:"cover_id"`
	CoverEditionKey   string   `json:"cover_edition_key"`
	Subject           []string `json:"subject"`
	IACollection      []string `json:"ia_collection"`
	LendingLibrary    bool     `json:"lendinglibrary"`
	PrintDisabled     bool     `json:"printdisabled"`
	LendingIdentifier string   `json:"lending_identifier"`
	Authors           []Author `json:"authors"`
	FirstPublishYear  int      `json:"first_publish_year"`
	IA                string   `json:"ia"`
	PublicScan        bool     `json:"public_scan"`
	HasFullText       bool     `json:"has_fulltext"`
	Availability      `json:"availability"`
}

type UserGetBookResponse struct {
	Name      string `json:"name"`
	WorkCount int    `json:"work_count"`
	Works     []Work `json:"works"`
}
