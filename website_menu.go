package odoo

// WebsiteMenu represents website.menu model.
type WebsiteMenu struct {
	ChildId          *Relation `xmlrpc:"child_id,omitempty"`
	ControllerPageId *Many2One `xmlrpc:"controller_page_id,omitempty"`
	CreateDate       *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName      *String   `xmlrpc:"display_name,omitempty"`
	GroupIds         *Relation `xmlrpc:"group_ids,omitempty"`
	Id               *Int      `xmlrpc:"id,omitempty"`
	IsMegaMenu       *Bool     `xmlrpc:"is_mega_menu,omitempty"`
	IsVisible        *Bool     `xmlrpc:"is_visible,omitempty"`
	MegaMenuClasses  *String   `xmlrpc:"mega_menu_classes,omitempty"`
	MegaMenuContent  *String   `xmlrpc:"mega_menu_content,omitempty"`
	Name             *String   `xmlrpc:"name,omitempty"`
	NewWindow        *Bool     `xmlrpc:"new_window,omitempty"`
	PageId           *Many2One `xmlrpc:"page_id,omitempty"`
	ParentId         *Many2One `xmlrpc:"parent_id,omitempty"`
	ParentPath       *String   `xmlrpc:"parent_path,omitempty"`
	Sequence         *Int      `xmlrpc:"sequence,omitempty"`
	ThemeTemplateId  *Many2One `xmlrpc:"theme_template_id,omitempty"`
	Url              *String   `xmlrpc:"url,omitempty"`
	WebsiteId        *Many2One `xmlrpc:"website_id,omitempty"`
	WriteDate        *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One `xmlrpc:"write_uid,omitempty"`
}

// WebsiteMenus represents array of website.menu model.
type WebsiteMenus []WebsiteMenu

// WebsiteMenuModel is the odoo model name.
const WebsiteMenuModel = "website.menu"

// Many2One convert WebsiteMenu to *Many2One.
func (wm *WebsiteMenu) Many2One() *Many2One {
	return NewMany2One(wm.Id.Get(), "")
}

// CreateWebsiteMenu creates a new website.menu model and returns its id.
func (c *Client) CreateWebsiteMenu(wm *WebsiteMenu) (int64, error) {
	ids, err := c.CreateWebsiteMenus([]*WebsiteMenu{wm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsiteMenu creates a new website.menu model and returns its id.
func (c *Client) CreateWebsiteMenus(wms []*WebsiteMenu) ([]int64, error) {
	var vv []interface{}
	for _, v := range wms {
		vv = append(vv, v)
	}
	return c.Create(WebsiteMenuModel, vv, nil)
}

// UpdateWebsiteMenu updates an existing website.menu record.
func (c *Client) UpdateWebsiteMenu(wm *WebsiteMenu) error {
	return c.UpdateWebsiteMenus([]int64{wm.Id.Get()}, wm)
}

// UpdateWebsiteMenus updates existing website.menu records.
// All records (represented by ids) will be updated by wm values.
func (c *Client) UpdateWebsiteMenus(ids []int64, wm *WebsiteMenu) error {
	return c.Update(WebsiteMenuModel, ids, wm, nil)
}

// DeleteWebsiteMenu deletes an existing website.menu record.
func (c *Client) DeleteWebsiteMenu(id int64) error {
	return c.DeleteWebsiteMenus([]int64{id})
}

// DeleteWebsiteMenus deletes existing website.menu records.
func (c *Client) DeleteWebsiteMenus(ids []int64) error {
	return c.Delete(WebsiteMenuModel, ids)
}

// GetWebsiteMenu gets website.menu existing record.
func (c *Client) GetWebsiteMenu(id int64) (*WebsiteMenu, error) {
	wms, err := c.GetWebsiteMenus([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wms)[0]), nil
}

// GetWebsiteMenus gets website.menu existing records.
func (c *Client) GetWebsiteMenus(ids []int64) (*WebsiteMenus, error) {
	wms := &WebsiteMenus{}
	if err := c.Read(WebsiteMenuModel, ids, nil, wms); err != nil {
		return nil, err
	}
	return wms, nil
}

// FindWebsiteMenu finds website.menu record by querying it with criteria.
func (c *Client) FindWebsiteMenu(criteria *Criteria) (*WebsiteMenu, error) {
	wms := &WebsiteMenus{}
	if err := c.SearchRead(WebsiteMenuModel, criteria, NewOptions().Limit(1), wms); err != nil {
		return nil, err
	}
	return &((*wms)[0]), nil
}

// FindWebsiteMenus finds website.menu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMenus(criteria *Criteria, options *Options) (*WebsiteMenus, error) {
	wms := &WebsiteMenus{}
	if err := c.SearchRead(WebsiteMenuModel, criteria, options, wms); err != nil {
		return nil, err
	}
	return wms, nil
}

// FindWebsiteMenuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsiteMenuIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsiteMenuModel, criteria, options)
}

// FindWebsiteMenuId finds record id by querying it with criteria.
func (c *Client) FindWebsiteMenuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsiteMenuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
