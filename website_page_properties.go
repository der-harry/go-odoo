package odoo

// WebsitePageProperties represents website.page.properties model.
type WebsitePageProperties struct {
	CanPublish                *Bool      `xmlrpc:"can_publish,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	DatePublish               *Time      `xmlrpc:"date_publish,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	GroupsId                  *Relation  `xmlrpc:"groups_id,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	IsHomepage                *Bool      `xmlrpc:"is_homepage,omitempty"`
	IsInMenu                  *Bool      `xmlrpc:"is_in_menu,omitempty"`
	IsNewPageTemplate         *Bool      `xmlrpc:"is_new_page_template,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	MenuIds                   *Relation  `xmlrpc:"menu_ids,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	OldUrl                    *String    `xmlrpc:"old_url,omitempty"`
	RedirectOldUrl            *Bool      `xmlrpc:"redirect_old_url,omitempty"`
	RedirectType              *Selection `xmlrpc:"redirect_type,omitempty"`
	TargetModelId             *Many2One  `xmlrpc:"target_model_id,omitempty"`
	Url                       *String    `xmlrpc:"url,omitempty"`
	Visibility                *Selection `xmlrpc:"visibility,omitempty"`
	VisibilityPasswordDisplay *String    `xmlrpc:"visibility_password_display,omitempty"`
	WebsiteId                 *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteIndexed            *Bool      `xmlrpc:"website_indexed,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// WebsitePagePropertiess represents array of website.page.properties model.
type WebsitePagePropertiess []WebsitePageProperties

// WebsitePagePropertiesModel is the odoo model name.
const WebsitePagePropertiesModel = "website.page.properties"

// Many2One convert WebsitePageProperties to *Many2One.
func (wpp *WebsitePageProperties) Many2One() *Many2One {
	return NewMany2One(wpp.Id.Get(), "")
}

// CreateWebsitePageProperties creates a new website.page.properties model and returns its id.
func (c *Client) CreateWebsitePageProperties(wpp *WebsitePageProperties) (int64, error) {
	ids, err := c.CreateWebsitePagePropertiess([]*WebsitePageProperties{wpp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsitePageProperties creates a new website.page.properties model and returns its id.
func (c *Client) CreateWebsitePagePropertiess(wpps []*WebsitePageProperties) ([]int64, error) {
	var vv []interface{}
	for _, v := range wpps {
		vv = append(vv, v)
	}
	return c.Create(WebsitePagePropertiesModel, vv, nil)
}

// UpdateWebsitePageProperties updates an existing website.page.properties record.
func (c *Client) UpdateWebsitePageProperties(wpp *WebsitePageProperties) error {
	return c.UpdateWebsitePagePropertiess([]int64{wpp.Id.Get()}, wpp)
}

// UpdateWebsitePagePropertiess updates existing website.page.properties records.
// All records (represented by ids) will be updated by wpp values.
func (c *Client) UpdateWebsitePagePropertiess(ids []int64, wpp *WebsitePageProperties) error {
	return c.Update(WebsitePagePropertiesModel, ids, wpp, nil)
}

// DeleteWebsitePageProperties deletes an existing website.page.properties record.
func (c *Client) DeleteWebsitePageProperties(id int64) error {
	return c.DeleteWebsitePagePropertiess([]int64{id})
}

// DeleteWebsitePagePropertiess deletes existing website.page.properties records.
func (c *Client) DeleteWebsitePagePropertiess(ids []int64) error {
	return c.Delete(WebsitePagePropertiesModel, ids)
}

// GetWebsitePageProperties gets website.page.properties existing record.
func (c *Client) GetWebsitePageProperties(id int64) (*WebsitePageProperties, error) {
	wpps, err := c.GetWebsitePagePropertiess([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wpps)[0]), nil
}

// GetWebsitePagePropertiess gets website.page.properties existing records.
func (c *Client) GetWebsitePagePropertiess(ids []int64) (*WebsitePagePropertiess, error) {
	wpps := &WebsitePagePropertiess{}
	if err := c.Read(WebsitePagePropertiesModel, ids, nil, wpps); err != nil {
		return nil, err
	}
	return wpps, nil
}

// FindWebsitePageProperties finds website.page.properties record by querying it with criteria.
func (c *Client) FindWebsitePageProperties(criteria *Criteria) (*WebsitePageProperties, error) {
	wpps := &WebsitePagePropertiess{}
	if err := c.SearchRead(WebsitePagePropertiesModel, criteria, NewOptions().Limit(1), wpps); err != nil {
		return nil, err
	}
	return &((*wpps)[0]), nil
}

// FindWebsitePagePropertiess finds website.page.properties records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePagePropertiess(criteria *Criteria, options *Options) (*WebsitePagePropertiess, error) {
	wpps := &WebsitePagePropertiess{}
	if err := c.SearchRead(WebsitePagePropertiesModel, criteria, options, wpps); err != nil {
		return nil, err
	}
	return wpps, nil
}

// FindWebsitePagePropertiesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePagePropertiesIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsitePagePropertiesModel, criteria, options)
}

// FindWebsitePagePropertiesId finds record id by querying it with criteria.
func (c *Client) FindWebsitePagePropertiesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePagePropertiesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
