package odoo

// WebsitePagePropertiesBase represents website.page.properties.base model.
type WebsitePagePropertiesBase struct {
	CanPublish    *Bool     `xmlrpc:"can_publish,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	IsHomepage    *Bool     `xmlrpc:"is_homepage,omitempty"`
	IsInMenu      *Bool     `xmlrpc:"is_in_menu,omitempty"`
	IsPublished   *Bool     `xmlrpc:"is_published,omitempty"`
	MenuIds       *Relation `xmlrpc:"menu_ids,omitempty"`
	TargetModelId *String   `xmlrpc:"target_model_id,omitempty"`
	Url           *String   `xmlrpc:"url,omitempty"`
	WebsiteId     *Many2One `xmlrpc:"website_id,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsitePagePropertiesBases represents array of website.page.properties.base model.
type WebsitePagePropertiesBases []WebsitePagePropertiesBase

// WebsitePagePropertiesBaseModel is the odoo model name.
const WebsitePagePropertiesBaseModel = "website.page.properties.base"

// Many2One convert WebsitePagePropertiesBase to *Many2One.
func (wppb *WebsitePagePropertiesBase) Many2One() *Many2One {
	return NewMany2One(wppb.Id.Get(), "")
}

// CreateWebsitePagePropertiesBase creates a new website.page.properties.base model and returns its id.
func (c *Client) CreateWebsitePagePropertiesBase(wppb *WebsitePagePropertiesBase) (int64, error) {
	ids, err := c.CreateWebsitePagePropertiesBases([]*WebsitePagePropertiesBase{wppb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsitePagePropertiesBase creates a new website.page.properties.base model and returns its id.
func (c *Client) CreateWebsitePagePropertiesBases(wppbs []*WebsitePagePropertiesBase) ([]int64, error) {
	var vv []interface{}
	for _, v := range wppbs {
		vv = append(vv, v)
	}
	return c.Create(WebsitePagePropertiesBaseModel, vv, nil)
}

// UpdateWebsitePagePropertiesBase updates an existing website.page.properties.base record.
func (c *Client) UpdateWebsitePagePropertiesBase(wppb *WebsitePagePropertiesBase) error {
	return c.UpdateWebsitePagePropertiesBases([]int64{wppb.Id.Get()}, wppb)
}

// UpdateWebsitePagePropertiesBases updates existing website.page.properties.base records.
// All records (represented by ids) will be updated by wppb values.
func (c *Client) UpdateWebsitePagePropertiesBases(ids []int64, wppb *WebsitePagePropertiesBase) error {
	return c.Update(WebsitePagePropertiesBaseModel, ids, wppb, nil)
}

// DeleteWebsitePagePropertiesBase deletes an existing website.page.properties.base record.
func (c *Client) DeleteWebsitePagePropertiesBase(id int64) error {
	return c.DeleteWebsitePagePropertiesBases([]int64{id})
}

// DeleteWebsitePagePropertiesBases deletes existing website.page.properties.base records.
func (c *Client) DeleteWebsitePagePropertiesBases(ids []int64) error {
	return c.Delete(WebsitePagePropertiesBaseModel, ids)
}

// GetWebsitePagePropertiesBase gets website.page.properties.base existing record.
func (c *Client) GetWebsitePagePropertiesBase(id int64) (*WebsitePagePropertiesBase, error) {
	wppbs, err := c.GetWebsitePagePropertiesBases([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wppbs)[0]), nil
}

// GetWebsitePagePropertiesBases gets website.page.properties.base existing records.
func (c *Client) GetWebsitePagePropertiesBases(ids []int64) (*WebsitePagePropertiesBases, error) {
	wppbs := &WebsitePagePropertiesBases{}
	if err := c.Read(WebsitePagePropertiesBaseModel, ids, nil, wppbs); err != nil {
		return nil, err
	}
	return wppbs, nil
}

// FindWebsitePagePropertiesBase finds website.page.properties.base record by querying it with criteria.
func (c *Client) FindWebsitePagePropertiesBase(criteria *Criteria) (*WebsitePagePropertiesBase, error) {
	wppbs := &WebsitePagePropertiesBases{}
	if err := c.SearchRead(WebsitePagePropertiesBaseModel, criteria, NewOptions().Limit(1), wppbs); err != nil {
		return nil, err
	}
	return &((*wppbs)[0]), nil
}

// FindWebsitePagePropertiesBases finds website.page.properties.base records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePagePropertiesBases(criteria *Criteria, options *Options) (*WebsitePagePropertiesBases, error) {
	wppbs := &WebsitePagePropertiesBases{}
	if err := c.SearchRead(WebsitePagePropertiesBaseModel, criteria, options, wppbs); err != nil {
		return nil, err
	}
	return wppbs, nil
}

// FindWebsitePagePropertiesBaseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePagePropertiesBaseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsitePagePropertiesBaseModel, criteria, options)
}

// FindWebsitePagePropertiesBaseId finds record id by querying it with criteria.
func (c *Client) FindWebsitePagePropertiesBaseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePagePropertiesBaseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
