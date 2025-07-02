package odoo

// WebsiteRewrite represents website.rewrite model.
type WebsiteRewrite struct {
	Active       *Bool      `xmlrpc:"active,omitempty"`
	CreateDate   *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	Name         *String    `xmlrpc:"name,omitempty"`
	RedirectType *Selection `xmlrpc:"redirect_type,omitempty"`
	RouteId      *Many2One  `xmlrpc:"route_id,omitempty"`
	Sequence     *Int       `xmlrpc:"sequence,omitempty"`
	UrlFrom      *String    `xmlrpc:"url_from,omitempty"`
	UrlTo        *String    `xmlrpc:"url_to,omitempty"`
	WebsiteId    *Many2One  `xmlrpc:"website_id,omitempty"`
	WriteDate    *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// WebsiteRewrites represents array of website.rewrite model.
type WebsiteRewrites []WebsiteRewrite

// WebsiteRewriteModel is the odoo model name.
const WebsiteRewriteModel = "website.rewrite"

// Many2One convert WebsiteRewrite to *Many2One.
func (wr *WebsiteRewrite) Many2One() *Many2One {
	return NewMany2One(wr.Id.Get(), "")
}

// CreateWebsiteRewrite creates a new website.rewrite model and returns its id.
func (c *Client) CreateWebsiteRewrite(wr *WebsiteRewrite) (int64, error) {
	ids, err := c.CreateWebsiteRewrites([]*WebsiteRewrite{wr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteRewrite creates a new website.rewrite model and returns its id.
func (c *Client) CreateWebsiteRewrites(wrs []*WebsiteRewrite) ([]int64, error) {
	var vv []interface{}
	for _, v := range wrs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteRewriteModel, vv, nil)
}

// UpdateWebsiteRewrite updates an existing website.rewrite record.
func (c *Client) UpdateWebsiteRewrite(wr *WebsiteRewrite) error {
	return c.UpdateWebsiteRewrites([]int64{wr.Id.Get()}, wr)
}

// UpdateWebsiteRewrites updates existing website.rewrite records.
// All records (represented by ids) will be updated by wr values.
func (c *Client) UpdateWebsiteRewrites(ids []int64, wr *WebsiteRewrite) error {
	return c.Update(WebsiteRewriteModel, ids, wr, nil)
}

// DeleteWebsiteRewrite deletes an existing website.rewrite record.
func (c *Client) DeleteWebsiteRewrite(id int64) error {
	return c.DeleteWebsiteRewrites([]int64{id})
}

// DeleteWebsiteRewrites deletes existing website.rewrite records.
func (c *Client) DeleteWebsiteRewrites(ids []int64) error {
	return c.Delete(WebsiteRewriteModel, ids)
}

// GetWebsiteRewrite gets website.rewrite existing record.
func (c *Client) GetWebsiteRewrite(id int64) (*WebsiteRewrite, error) {
	wrs, err := c.GetWebsiteRewrites([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// GetWebsiteRewrites gets website.rewrite existing records.
func (c *Client) GetWebsiteRewrites(ids []int64) (*WebsiteRewrites, error) {
	wrs := &WebsiteRewrites{}
	if err := c.Read(WebsiteRewriteModel, ids, nil, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRewrite finds website.rewrite record by querying it with criteria.
func (c *Client) FindWebsiteRewrite(criteria *Criteria) (*WebsiteRewrite, error) {
	wrs := &WebsiteRewrites{}
	if err := c.SearchRead(WebsiteRewriteModel, criteria, NewOptions().Limit(1), wrs); err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// FindWebsiteRewrites finds website.rewrite records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRewrites(criteria *Criteria, options *Options) (*WebsiteRewrites, error) {
	wrs := &WebsiteRewrites{}
	if err := c.SearchRead(WebsiteRewriteModel, criteria, options, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRewriteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRewriteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteRewriteModel, criteria, options)
}

// FindWebsiteRewriteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteRewriteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteRewriteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
