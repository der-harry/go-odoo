package odoo

// WebsiteTrack represents website.track model.
type WebsiteTrack struct {
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	PageId        *Many2One `xmlrpc:"page_id,omitempty"`
	Url           *String   `xmlrpc:"url,omitempty"`
	VisitDatetime *Time     `xmlrpc:"visit_datetime,omitempty"`
	VisitorId     *Many2One `xmlrpc:"visitor_id,omitempty"`
}

// WebsiteTracks represents array of website.track model.
type WebsiteTracks []WebsiteTrack

// WebsiteTrackModel is the odoo model name.
const WebsiteTrackModel = "website.track"

// Many2One convert WebsiteTrack to *Many2One.
func (wt *WebsiteTrack) Many2One() *Many2One {
	return NewMany2One(wt.Id.Get(), "")
}

// CreateWebsiteTrack creates a new website.track model and returns its id.
func (c *Client) CreateWebsiteTrack(wt *WebsiteTrack) (int64, error) {
	ids, err := c.CreateWebsiteTracks([]*WebsiteTrack{wt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteTrack creates a new website.track model and returns its id.
func (c *Client) CreateWebsiteTracks(wts []*WebsiteTrack) ([]int64, error) {
	var vv []interface{}
	for _, v := range wts {
		vv = append(vv, v)
	}
	return c.Create(WebsiteTrackModel, vv, nil)
}

// UpdateWebsiteTrack updates an existing website.track record.
func (c *Client) UpdateWebsiteTrack(wt *WebsiteTrack) error {
	return c.UpdateWebsiteTracks([]int64{wt.Id.Get()}, wt)
}

// UpdateWebsiteTracks updates existing website.track records.
// All records (represented by ids) will be updated by wt values.
func (c *Client) UpdateWebsiteTracks(ids []int64, wt *WebsiteTrack) error {
	return c.Update(WebsiteTrackModel, ids, wt, nil)
}

// DeleteWebsiteTrack deletes an existing website.track record.
func (c *Client) DeleteWebsiteTrack(id int64) error {
	return c.DeleteWebsiteTracks([]int64{id})
}

// DeleteWebsiteTracks deletes existing website.track records.
func (c *Client) DeleteWebsiteTracks(ids []int64) error {
	return c.Delete(WebsiteTrackModel, ids)
}

// GetWebsiteTrack gets website.track existing record.
func (c *Client) GetWebsiteTrack(id int64) (*WebsiteTrack, error) {
	wts, err := c.GetWebsiteTracks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wts)[0]), nil
}

// GetWebsiteTracks gets website.track existing records.
func (c *Client) GetWebsiteTracks(ids []int64) (*WebsiteTracks, error) {
	wts := &WebsiteTracks{}
	if err := c.Read(WebsiteTrackModel, ids, nil, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebsiteTrack finds website.track record by querying it with criteria.
func (c *Client) FindWebsiteTrack(criteria *Criteria) (*WebsiteTrack, error) {
	wts := &WebsiteTracks{}
	if err := c.SearchRead(WebsiteTrackModel, criteria, NewOptions().Limit(1), wts); err != nil {
		return nil, err
	}
	return &((*wts)[0]), nil
}

// FindWebsiteTracks finds website.track records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteTracks(criteria *Criteria, options *Options) (*WebsiteTracks, error) {
	wts := &WebsiteTracks{}
	if err := c.SearchRead(WebsiteTrackModel, criteria, options, wts); err != nil {
		return nil, err
	}
	return wts, nil
}

// FindWebsiteTrackIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteTrackIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteTrackModel, criteria, options)
}

// FindWebsiteTrackId finds record id by querying it with criteria.
func (c *Client) FindWebsiteTrackId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteTrackModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
