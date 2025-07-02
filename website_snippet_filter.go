package odoo

// WebsiteSnippetFilter represents website.snippet.filter model.
type WebsiteSnippetFilter struct {
	ActionServerId   *Many2One `xmlrpc:"action_server_id,omitempty"`
	CanPublish       *Bool     `xmlrpc:"can_publish,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	FieldNames       *String   `xmlrpc:"field_names,omitempty"`
	FilterId         *Many2One `xmlrpc:"filter_id,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	IsPublished      *Bool     `xmlrpc:"is_published,omitempty"`
	Limit            *Int      `xmlrpc:"limit,omitempty"`
	ModelName        *String   `xmlrpc:"model_name,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	WebsiteId        *Many2One `xmlrpc:"website_id,omitempty"`
	WebsitePublished *Bool     `xmlrpc:"website_published,omitempty"`
	WebsiteUrl       *String   `xmlrpc:"website_url,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteSnippetFilters represents array of website.snippet.filter model.
type WebsiteSnippetFilters []WebsiteSnippetFilter

// WebsiteSnippetFilterModel is the odoo model name.
const WebsiteSnippetFilterModel = "website.snippet.filter"

// Many2One convert WebsiteSnippetFilter to *Many2One.
func (wsf *WebsiteSnippetFilter) Many2One() *Many2One {
	return NewMany2One(wsf.Id.Get(), "")
}

// CreateWebsiteSnippetFilter creates a new website.snippet.filter model and returns its id.
func (c *Client) CreateWebsiteSnippetFilter(wsf *WebsiteSnippetFilter) (int64, error) {
	ids, err := c.CreateWebsiteSnippetFilters([]*WebsiteSnippetFilter{wsf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteSnippetFilter creates a new website.snippet.filter model and returns its id.
func (c *Client) CreateWebsiteSnippetFilters(wsfs []*WebsiteSnippetFilter) ([]int64, error) {
	var vv []interface{}
	for _, v := range wsfs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteSnippetFilterModel, vv, nil)
}

// UpdateWebsiteSnippetFilter updates an existing website.snippet.filter record.
func (c *Client) UpdateWebsiteSnippetFilter(wsf *WebsiteSnippetFilter) error {
	return c.UpdateWebsiteSnippetFilters([]int64{wsf.Id.Get()}, wsf)
}

// UpdateWebsiteSnippetFilters updates existing website.snippet.filter records.
// All records (represented by ids) will be updated by wsf values.
func (c *Client) UpdateWebsiteSnippetFilters(ids []int64, wsf *WebsiteSnippetFilter) error {
	return c.Update(WebsiteSnippetFilterModel, ids, wsf, nil)
}

// DeleteWebsiteSnippetFilter deletes an existing website.snippet.filter record.
func (c *Client) DeleteWebsiteSnippetFilter(id int64) error {
	return c.DeleteWebsiteSnippetFilters([]int64{id})
}

// DeleteWebsiteSnippetFilters deletes existing website.snippet.filter records.
func (c *Client) DeleteWebsiteSnippetFilters(ids []int64) error {
	return c.Delete(WebsiteSnippetFilterModel, ids)
}

// GetWebsiteSnippetFilter gets website.snippet.filter existing record.
func (c *Client) GetWebsiteSnippetFilter(id int64) (*WebsiteSnippetFilter, error) {
	wsfs, err := c.GetWebsiteSnippetFilters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wsfs)[0]), nil
}

// GetWebsiteSnippetFilters gets website.snippet.filter existing records.
func (c *Client) GetWebsiteSnippetFilters(ids []int64) (*WebsiteSnippetFilters, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.Read(WebsiteSnippetFilterModel, ids, nil, wsfs); err != nil {
		return nil, err
	}
	return wsfs, nil
}

// FindWebsiteSnippetFilter finds website.snippet.filter record by querying it with criteria.
func (c *Client) FindWebsiteSnippetFilter(criteria *Criteria) (*WebsiteSnippetFilter, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.SearchRead(WebsiteSnippetFilterModel, criteria, NewOptions().Limit(1), wsfs); err != nil {
		return nil, err
	}
	return &((*wsfs)[0]), nil
}

// FindWebsiteSnippetFilters finds website.snippet.filter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSnippetFilters(criteria *Criteria, options *Options) (*WebsiteSnippetFilters, error) {
	wsfs := &WebsiteSnippetFilters{}
	if err := c.SearchRead(WebsiteSnippetFilterModel, criteria, options, wsfs); err != nil {
		return nil, err
	}
	return wsfs, nil
}

// FindWebsiteSnippetFilterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteSnippetFilterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteSnippetFilterModel, criteria, options)
}

// FindWebsiteSnippetFilterId finds record id by querying it with criteria.
func (c *Client) FindWebsiteSnippetFilterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteSnippetFilterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
