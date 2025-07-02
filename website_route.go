package odoo

// WebsiteRoute represents website.route model.
type WebsiteRoute struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Path        *String   `xmlrpc:"path,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteRoutes represents array of website.route model.
type WebsiteRoutes []WebsiteRoute

// WebsiteRouteModel is the odoo model name.
const WebsiteRouteModel = "website.route"

// Many2One convert WebsiteRoute to *Many2One.
func (wr *WebsiteRoute) Many2One() *Many2One {
	return NewMany2One(wr.Id.Get(), "")
}

// CreateWebsiteRoute creates a new website.route model and returns its id.
func (c *Client) CreateWebsiteRoute(wr *WebsiteRoute) (int64, error) {
	ids, err := c.CreateWebsiteRoutes([]*WebsiteRoute{wr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteRoute creates a new website.route model and returns its id.
func (c *Client) CreateWebsiteRoutes(wrs []*WebsiteRoute) ([]int64, error) {
	var vv []interface{}
	for _, v := range wrs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteRouteModel, vv, nil)
}

// UpdateWebsiteRoute updates an existing website.route record.
func (c *Client) UpdateWebsiteRoute(wr *WebsiteRoute) error {
	return c.UpdateWebsiteRoutes([]int64{wr.Id.Get()}, wr)
}

// UpdateWebsiteRoutes updates existing website.route records.
// All records (represented by ids) will be updated by wr values.
func (c *Client) UpdateWebsiteRoutes(ids []int64, wr *WebsiteRoute) error {
	return c.Update(WebsiteRouteModel, ids, wr, nil)
}

// DeleteWebsiteRoute deletes an existing website.route record.
func (c *Client) DeleteWebsiteRoute(id int64) error {
	return c.DeleteWebsiteRoutes([]int64{id})
}

// DeleteWebsiteRoutes deletes existing website.route records.
func (c *Client) DeleteWebsiteRoutes(ids []int64) error {
	return c.Delete(WebsiteRouteModel, ids)
}

// GetWebsiteRoute gets website.route existing record.
func (c *Client) GetWebsiteRoute(id int64) (*WebsiteRoute, error) {
	wrs, err := c.GetWebsiteRoutes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// GetWebsiteRoutes gets website.route existing records.
func (c *Client) GetWebsiteRoutes(ids []int64) (*WebsiteRoutes, error) {
	wrs := &WebsiteRoutes{}
	if err := c.Read(WebsiteRouteModel, ids, nil, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRoute finds website.route record by querying it with criteria.
func (c *Client) FindWebsiteRoute(criteria *Criteria) (*WebsiteRoute, error) {
	wrs := &WebsiteRoutes{}
	if err := c.SearchRead(WebsiteRouteModel, criteria, NewOptions().Limit(1), wrs); err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// FindWebsiteRoutes finds website.route records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRoutes(criteria *Criteria, options *Options) (*WebsiteRoutes, error) {
	wrs := &WebsiteRoutes{}
	if err := c.SearchRead(WebsiteRouteModel, criteria, options, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRouteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRouteIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteRouteModel, criteria, options)
}

// FindWebsiteRouteId finds record id by querying it with criteria.
func (c *Client) FindWebsiteRouteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteRouteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
