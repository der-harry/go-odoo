package odoo

// ThemeWebsitePage represents theme.website.page model.
type ThemeWebsitePage struct {
	CopyIds           *Relation `xmlrpc:"copy_ids,omitempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName       *String   `xmlrpc:"display_name,omitempty"`
	FooterVisible     *Bool     `xmlrpc:"footer_visible,omitempty"`
	HeaderColor       *String   `xmlrpc:"header_color,omitempty"`
	HeaderOverlay     *Bool     `xmlrpc:"header_overlay,omitempty"`
	HeaderVisible     *Bool     `xmlrpc:"header_visible,omitempty"`
	Id                *Int      `xmlrpc:"id,omitempty"`
	IsNewPageTemplate *Bool     `xmlrpc:"is_new_page_template,omitempty"`
	IsPublished       *Bool     `xmlrpc:"is_published,omitempty"`
	Url               *String   `xmlrpc:"url,omitempty"`
	ViewId            *Many2One `xmlrpc:"view_id,omitempty"`
	WebsiteIndexed    *Bool     `xmlrpc:"website_indexed,omitempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ThemeWebsitePages represents array of theme.website.page model.
type ThemeWebsitePages []ThemeWebsitePage

// ThemeWebsitePageModel is the odoo model name.
const ThemeWebsitePageModel = "theme.website.page"

// Many2One convert ThemeWebsitePage to *Many2One.
func (twp *ThemeWebsitePage) Many2One() *Many2One {
	return NewMany2One(twp.Id.Get(), "")
}

// CreateThemeWebsitePage creates a new theme.website.page model and returns its id.
func (c *Client) CreateThemeWebsitePage(twp *ThemeWebsitePage) (int64, error) {
	ids, err := c.CreateThemeWebsitePages([]*ThemeWebsitePage{twp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateThemeWebsitePage creates a new theme.website.page model and returns its id.
func (c *Client) CreateThemeWebsitePages(twps []*ThemeWebsitePage) ([]int64, error) {
	var vv []interface{}
	for _, v := range twps {
		vv = append(vv, v)
	}
	return c.Create(ThemeWebsitePageModel, vv, nil)
}

// UpdateThemeWebsitePage updates an existing theme.website.page record.
func (c *Client) UpdateThemeWebsitePage(twp *ThemeWebsitePage) error {
	return c.UpdateThemeWebsitePages([]int64{twp.Id.Get()}, twp)
}

// UpdateThemeWebsitePages updates existing theme.website.page records.
// All records (represented by ids) will be updated by twp values.
func (c *Client) UpdateThemeWebsitePages(ids []int64, twp *ThemeWebsitePage) error {
	return c.Update(ThemeWebsitePageModel, ids, twp, nil)
}

// DeleteThemeWebsitePage deletes an existing theme.website.page record.
func (c *Client) DeleteThemeWebsitePage(id int64) error {
	return c.DeleteThemeWebsitePages([]int64{id})
}

// DeleteThemeWebsitePages deletes existing theme.website.page records.
func (c *Client) DeleteThemeWebsitePages(ids []int64) error {
	return c.Delete(ThemeWebsitePageModel, ids)
}

// GetThemeWebsitePage gets theme.website.page existing record.
func (c *Client) GetThemeWebsitePage(id int64) (*ThemeWebsitePage, error) {
	twps, err := c.GetThemeWebsitePages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*twps)[0]), nil
}

// GetThemeWebsitePages gets theme.website.page existing records.
func (c *Client) GetThemeWebsitePages(ids []int64) (*ThemeWebsitePages, error) {
	twps := &ThemeWebsitePages{}
	if err := c.Read(ThemeWebsitePageModel, ids, nil, twps); err != nil {
		return nil, err
	}
	return twps, nil
}

// FindThemeWebsitePage finds theme.website.page record by querying it with criteria.
func (c *Client) FindThemeWebsitePage(criteria *Criteria) (*ThemeWebsitePage, error) {
	twps := &ThemeWebsitePages{}
	if err := c.SearchRead(ThemeWebsitePageModel, criteria, NewOptions().Limit(1), twps); err != nil {
		return nil, err
	}
	return &((*twps)[0]), nil
}

// FindThemeWebsitePages finds theme.website.page records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsitePages(criteria *Criteria, options *Options) (*ThemeWebsitePages, error) {
	twps := &ThemeWebsitePages{}
	if err := c.SearchRead(ThemeWebsitePageModel, criteria, options, twps); err != nil {
		return nil, err
	}
	return twps, nil
}

// FindThemeWebsitePageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsitePageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ThemeWebsitePageModel, criteria, options)
}

// FindThemeWebsitePageId finds record id by querying it with criteria.
func (c *Client) FindThemeWebsitePageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeWebsitePageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
