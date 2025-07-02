package odoo

// WebsiteVisitor represents website.visitor model.
type WebsiteVisitor struct {
	AccessToken            *String    `xmlrpc:"access_token,omitempty"`
	CountryFlag            *String    `xmlrpc:"country_flag,omitempty"`
	CountryId              *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Email                  *String    `xmlrpc:"email,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	IsConnected            *Bool      `xmlrpc:"is_connected,omitempty"`
	LangId                 *Many2One  `xmlrpc:"lang_id,omitempty"`
	LastConnectionDatetime *Time      `xmlrpc:"last_connection_datetime,omitempty"`
	LastVisitedPageId      *Many2One  `xmlrpc:"last_visited_page_id,omitempty"`
	LeadCount              *Int       `xmlrpc:"lead_count,omitempty"`
	LeadIds                *Relation  `xmlrpc:"lead_ids,omitempty"`
	Mobile                 *String    `xmlrpc:"mobile,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	PageCount              *Int       `xmlrpc:"page_count,omitempty"`
	PageIds                *Relation  `xmlrpc:"page_ids,omitempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerImage           *String    `xmlrpc:"partner_image,omitempty"`
	TimeSinceLastAction    *String    `xmlrpc:"time_since_last_action,omitempty"`
	Timezone               *Selection `xmlrpc:"timezone,omitempty"`
	VisitCount             *Int       `xmlrpc:"visit_count,omitempty"`
	VisitorPageCount       *Int       `xmlrpc:"visitor_page_count,omitempty"`
	WebsiteId              *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteTrackIds        *Relation  `xmlrpc:"website_track_ids,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// WebsiteVisitors represents array of website.visitor model.
type WebsiteVisitors []WebsiteVisitor

// WebsiteVisitorModel is the odoo model name.
const WebsiteVisitorModel = "website.visitor"

// Many2One convert WebsiteVisitor to *Many2One.
func (wv *WebsiteVisitor) Many2One() *Many2One {
	return NewMany2One(wv.Id.Get(), "")
}

// CreateWebsiteVisitor creates a new website.visitor model and returns its id.
func (c *Client) CreateWebsiteVisitor(wv *WebsiteVisitor) (int64, error) {
	ids, err := c.CreateWebsiteVisitors([]*WebsiteVisitor{wv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteVisitor creates a new website.visitor model and returns its id.
func (c *Client) CreateWebsiteVisitors(wvs []*WebsiteVisitor) ([]int64, error) {
	var vv []interface{}
	for _, v := range wvs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteVisitorModel, vv, nil)
}

// UpdateWebsiteVisitor updates an existing website.visitor record.
func (c *Client) UpdateWebsiteVisitor(wv *WebsiteVisitor) error {
	return c.UpdateWebsiteVisitors([]int64{wv.Id.Get()}, wv)
}

// UpdateWebsiteVisitors updates existing website.visitor records.
// All records (represented by ids) will be updated by wv values.
func (c *Client) UpdateWebsiteVisitors(ids []int64, wv *WebsiteVisitor) error {
	return c.Update(WebsiteVisitorModel, ids, wv, nil)
}

// DeleteWebsiteVisitor deletes an existing website.visitor record.
func (c *Client) DeleteWebsiteVisitor(id int64) error {
	return c.DeleteWebsiteVisitors([]int64{id})
}

// DeleteWebsiteVisitors deletes existing website.visitor records.
func (c *Client) DeleteWebsiteVisitors(ids []int64) error {
	return c.Delete(WebsiteVisitorModel, ids)
}

// GetWebsiteVisitor gets website.visitor existing record.
func (c *Client) GetWebsiteVisitor(id int64) (*WebsiteVisitor, error) {
	wvs, err := c.GetWebsiteVisitors([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wvs)[0]), nil
}

// GetWebsiteVisitors gets website.visitor existing records.
func (c *Client) GetWebsiteVisitors(ids []int64) (*WebsiteVisitors, error) {
	wvs := &WebsiteVisitors{}
	if err := c.Read(WebsiteVisitorModel, ids, nil, wvs); err != nil {
		return nil, err
	}
	return wvs, nil
}

// FindWebsiteVisitor finds website.visitor record by querying it with criteria.
func (c *Client) FindWebsiteVisitor(criteria *Criteria) (*WebsiteVisitor, error) {
	wvs := &WebsiteVisitors{}
	if err := c.SearchRead(WebsiteVisitorModel, criteria, NewOptions().Limit(1), wvs); err != nil {
		return nil, err
	}
	return &((*wvs)[0]), nil
}

// FindWebsiteVisitors finds website.visitor records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitors(criteria *Criteria, options *Options) (*WebsiteVisitors, error) {
	wvs := &WebsiteVisitors{}
	if err := c.SearchRead(WebsiteVisitorModel, criteria, options, wvs); err != nil {
		return nil, err
	}
	return wvs, nil
}

// FindWebsiteVisitorIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteVisitorIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteVisitorModel, criteria, options)
}

// FindWebsiteVisitorId finds record id by querying it with criteria.
func (c *Client) FindWebsiteVisitorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteVisitorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
