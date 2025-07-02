package odoo

// WebsiteRobots represents website.robots model.
type WebsiteRobots struct {
	Content     *String   `xmlrpc:"content,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteRobotss represents array of website.robots model.
type WebsiteRobotss []WebsiteRobots

// WebsiteRobotsModel is the odoo model name.
const WebsiteRobotsModel = "website.robots"

// Many2One convert WebsiteRobots to *Many2One.
func (wr *WebsiteRobots) Many2One() *Many2One {
	return NewMany2One(wr.Id.Get(), "")
}

// CreateWebsiteRobots creates a new website.robots model and returns its id.
func (c *Client) CreateWebsiteRobots(wr *WebsiteRobots) (int64, error) {
	ids, err := c.CreateWebsiteRobotss([]*WebsiteRobots{wr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteRobots creates a new website.robots model and returns its id.
func (c *Client) CreateWebsiteRobotss(wrs []*WebsiteRobots) ([]int64, error) {
	var vv []interface{}
	for _, v := range wrs {
		vv = append(vv, v)
	}
	return c.Create(WebsiteRobotsModel, vv, nil)
}

// UpdateWebsiteRobots updates an existing website.robots record.
func (c *Client) UpdateWebsiteRobots(wr *WebsiteRobots) error {
	return c.UpdateWebsiteRobotss([]int64{wr.Id.Get()}, wr)
}

// UpdateWebsiteRobotss updates existing website.robots records.
// All records (represented by ids) will be updated by wr values.
func (c *Client) UpdateWebsiteRobotss(ids []int64, wr *WebsiteRobots) error {
	return c.Update(WebsiteRobotsModel, ids, wr, nil)
}

// DeleteWebsiteRobots deletes an existing website.robots record.
func (c *Client) DeleteWebsiteRobots(id int64) error {
	return c.DeleteWebsiteRobotss([]int64{id})
}

// DeleteWebsiteRobotss deletes existing website.robots records.
func (c *Client) DeleteWebsiteRobotss(ids []int64) error {
	return c.Delete(WebsiteRobotsModel, ids)
}

// GetWebsiteRobots gets website.robots existing record.
func (c *Client) GetWebsiteRobots(id int64) (*WebsiteRobots, error) {
	wrs, err := c.GetWebsiteRobotss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// GetWebsiteRobotss gets website.robots existing records.
func (c *Client) GetWebsiteRobotss(ids []int64) (*WebsiteRobotss, error) {
	wrs := &WebsiteRobotss{}
	if err := c.Read(WebsiteRobotsModel, ids, nil, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRobots finds website.robots record by querying it with criteria.
func (c *Client) FindWebsiteRobots(criteria *Criteria) (*WebsiteRobots, error) {
	wrs := &WebsiteRobotss{}
	if err := c.SearchRead(WebsiteRobotsModel, criteria, NewOptions().Limit(1), wrs); err != nil {
		return nil, err
	}
	return &((*wrs)[0]), nil
}

// FindWebsiteRobotss finds website.robots records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRobotss(criteria *Criteria, options *Options) (*WebsiteRobotss, error) {
	wrs := &WebsiteRobotss{}
	if err := c.SearchRead(WebsiteRobotsModel, criteria, options, wrs); err != nil {
		return nil, err
	}
	return wrs, nil
}

// FindWebsiteRobotsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteRobotsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteRobotsModel, criteria, options)
}

// FindWebsiteRobotsId finds record id by querying it with criteria.
func (c *Client) FindWebsiteRobotsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteRobotsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
