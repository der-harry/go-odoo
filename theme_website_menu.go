package odoo

// ThemeWebsiteMenu represents theme.website.menu model.
type ThemeWebsiteMenu struct {
	CopyIds             *Relation `xmlrpc:"copy_ids,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	MegaMenuClasses     *String   `xmlrpc:"mega_menu_classes,omitempty"`
	MegaMenuContent     *String   `xmlrpc:"mega_menu_content,omitempty"`
	Name                *String   `xmlrpc:"name,omitempty"`
	NewWindow           *Bool     `xmlrpc:"new_window,omitempty"`
	PageId              *Many2One `xmlrpc:"page_id,omitempty"`
	ParentId            *Many2One `xmlrpc:"parent_id,omitempty"`
	Sequence            *Int      `xmlrpc:"sequence,omitempty"`
	Url                 *String   `xmlrpc:"url,omitempty"`
	UseMainMenuAsParent *Bool     `xmlrpc:"use_main_menu_as_parent,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ThemeWebsiteMenus represents array of theme.website.menu model.
type ThemeWebsiteMenus []ThemeWebsiteMenu

// ThemeWebsiteMenuModel is the odoo model name.
const ThemeWebsiteMenuModel = "theme.website.menu"

// Many2One convert ThemeWebsiteMenu to *Many2One.
func (twm *ThemeWebsiteMenu) Many2One() *Many2One {
	return NewMany2One(twm.Id.Get(), "")
}

// CreateThemeWebsiteMenu creates a new theme.website.menu model and returns its id.
func (c *Client) CreateThemeWebsiteMenu(twm *ThemeWebsiteMenu) (int64, error) {
	ids, err := c.CreateThemeWebsiteMenus([]*ThemeWebsiteMenu{twm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateThemeWebsiteMenu creates a new theme.website.menu model and returns its id.
func (c *Client) CreateThemeWebsiteMenus(twms []*ThemeWebsiteMenu) ([]int64, error) {
	var vv []interface{}
	for _, v := range twms {
		vv = append(vv, v)
	}
	return c.Create(ThemeWebsiteMenuModel, vv, nil)
}

// UpdateThemeWebsiteMenu updates an existing theme.website.menu record.
func (c *Client) UpdateThemeWebsiteMenu(twm *ThemeWebsiteMenu) error {
	return c.UpdateThemeWebsiteMenus([]int64{twm.Id.Get()}, twm)
}

// UpdateThemeWebsiteMenus updates existing theme.website.menu records.
// All records (represented by ids) will be updated by twm values.
func (c *Client) UpdateThemeWebsiteMenus(ids []int64, twm *ThemeWebsiteMenu) error {
	return c.Update(ThemeWebsiteMenuModel, ids, twm, nil)
}

// DeleteThemeWebsiteMenu deletes an existing theme.website.menu record.
func (c *Client) DeleteThemeWebsiteMenu(id int64) error {
	return c.DeleteThemeWebsiteMenus([]int64{id})
}

// DeleteThemeWebsiteMenus deletes existing theme.website.menu records.
func (c *Client) DeleteThemeWebsiteMenus(ids []int64) error {
	return c.Delete(ThemeWebsiteMenuModel, ids)
}

// GetThemeWebsiteMenu gets theme.website.menu existing record.
func (c *Client) GetThemeWebsiteMenu(id int64) (*ThemeWebsiteMenu, error) {
	twms, err := c.GetThemeWebsiteMenus([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*twms)[0]), nil
}

// GetThemeWebsiteMenus gets theme.website.menu existing records.
func (c *Client) GetThemeWebsiteMenus(ids []int64) (*ThemeWebsiteMenus, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.Read(ThemeWebsiteMenuModel, ids, nil, twms); err != nil {
		return nil, err
	}
	return twms, nil
}

// FindThemeWebsiteMenu finds theme.website.menu record by querying it with criteria.
func (c *Client) FindThemeWebsiteMenu(criteria *Criteria) (*ThemeWebsiteMenu, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.SearchRead(ThemeWebsiteMenuModel, criteria, NewOptions().Limit(1), twms); err != nil {
		return nil, err
	}
	return &((*twms)[0]), nil
}

// FindThemeWebsiteMenus finds theme.website.menu records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsiteMenus(criteria *Criteria, options *Options) (*ThemeWebsiteMenus, error) {
	twms := &ThemeWebsiteMenus{}
	if err := c.SearchRead(ThemeWebsiteMenuModel, criteria, options, twms); err != nil {
		return nil, err
	}
	return twms, nil
}

// FindThemeWebsiteMenuIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeWebsiteMenuIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ThemeWebsiteMenuModel, criteria, options)
}

// FindThemeWebsiteMenuId finds record id by querying it with criteria.
func (c *Client) FindThemeWebsiteMenuId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeWebsiteMenuModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
