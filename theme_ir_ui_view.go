package odoo

// ThemeIrUiView represents theme.ir.ui.view model.
type ThemeIrUiView struct {
	Active        *Bool      `xmlrpc:"active,omitempty"`
	Arch          *String    `xmlrpc:"arch,omitempty"`
	ArchFs        *String    `xmlrpc:"arch_fs,omitempty"`
	CopyIds       *Relation  `xmlrpc:"copy_ids,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomizeShow *Bool      `xmlrpc:"customize_show,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	InheritId     *String    `xmlrpc:"inherit_id,omitempty"`
	Key           *String    `xmlrpc:"key,omitempty"`
	Mode          *Selection `xmlrpc:"mode,omitempty"`
	Name          *String    `xmlrpc:"name,omitempty"`
	Priority      *Int       `xmlrpc:"priority,omitempty"`
	Type          *String    `xmlrpc:"type,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ThemeIrUiViews represents array of theme.ir.ui.view model.
type ThemeIrUiViews []ThemeIrUiView

// ThemeIrUiViewModel is the odoo model name.
const ThemeIrUiViewModel = "theme.ir.ui.view"

// Many2One convert ThemeIrUiView to *Many2One.
func (tiuv *ThemeIrUiView) Many2One() *Many2One {
	return NewMany2One(tiuv.Id.Get(), "")
}

// CreateThemeIrUiView creates a new theme.ir.ui.view model and returns its id.
func (c *Client) CreateThemeIrUiView(tiuv *ThemeIrUiView) (int64, error) {
	ids, err := c.CreateThemeIrUiViews([]*ThemeIrUiView{tiuv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateThemeIrUiView creates a new theme.ir.ui.view model and returns its id.
func (c *Client) CreateThemeIrUiViews(tiuvs []*ThemeIrUiView) ([]int64, error) {
	var vv []interface{}
	for _, v := range tiuvs {
		vv = append(vv, v)
	}
	return c.Create(ThemeIrUiViewModel, vv, nil)
}

// UpdateThemeIrUiView updates an existing theme.ir.ui.view record.
func (c *Client) UpdateThemeIrUiView(tiuv *ThemeIrUiView) error {
	return c.UpdateThemeIrUiViews([]int64{tiuv.Id.Get()}, tiuv)
}

// UpdateThemeIrUiViews updates existing theme.ir.ui.view records.
// All records (represented by ids) will be updated by tiuv values.
func (c *Client) UpdateThemeIrUiViews(ids []int64, tiuv *ThemeIrUiView) error {
	return c.Update(ThemeIrUiViewModel, ids, tiuv, nil)
}

// DeleteThemeIrUiView deletes an existing theme.ir.ui.view record.
func (c *Client) DeleteThemeIrUiView(id int64) error {
	return c.DeleteThemeIrUiViews([]int64{id})
}

// DeleteThemeIrUiViews deletes existing theme.ir.ui.view records.
func (c *Client) DeleteThemeIrUiViews(ids []int64) error {
	return c.Delete(ThemeIrUiViewModel, ids)
}

// GetThemeIrUiView gets theme.ir.ui.view existing record.
func (c *Client) GetThemeIrUiView(id int64) (*ThemeIrUiView, error) {
	tiuvs, err := c.GetThemeIrUiViews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tiuvs)[0]), nil
}

// GetThemeIrUiViews gets theme.ir.ui.view existing records.
func (c *Client) GetThemeIrUiViews(ids []int64) (*ThemeIrUiViews, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.Read(ThemeIrUiViewModel, ids, nil, tiuvs); err != nil {
		return nil, err
	}
	return tiuvs, nil
}

// FindThemeIrUiView finds theme.ir.ui.view record by querying it with criteria.
func (c *Client) FindThemeIrUiView(criteria *Criteria) (*ThemeIrUiView, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.SearchRead(ThemeIrUiViewModel, criteria, NewOptions().Limit(1), tiuvs); err != nil {
		return nil, err
	}
	return &((*tiuvs)[0]), nil
}

// FindThemeIrUiViews finds theme.ir.ui.view records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrUiViews(criteria *Criteria, options *Options) (*ThemeIrUiViews, error) {
	tiuvs := &ThemeIrUiViews{}
	if err := c.SearchRead(ThemeIrUiViewModel, criteria, options, tiuvs); err != nil {
		return nil, err
	}
	return tiuvs, nil
}

// FindThemeIrUiViewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrUiViewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ThemeIrUiViewModel, criteria, options)
}

// FindThemeIrUiViewId finds record id by querying it with criteria.
func (c *Client) FindThemeIrUiViewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeIrUiViewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
