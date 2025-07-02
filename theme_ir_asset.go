package odoo

// ThemeIrAsset represents theme.ir.asset model.
type ThemeIrAsset struct {
	Active      *Bool      `xmlrpc:"active,omitempty"`
	Bundle      *String    `xmlrpc:"bundle,omitempty"`
	CopyIds     *Relation  `xmlrpc:"copy_ids,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	Directive   *Selection `xmlrpc:"directive,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Key         *String    `xmlrpc:"key,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Path        *String    `xmlrpc:"path,omitempty"`
	Sequence    *Int       `xmlrpc:"sequence,omitempty"`
	Target      *String    `xmlrpc:"target,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ThemeIrAssets represents array of theme.ir.asset model.
type ThemeIrAssets []ThemeIrAsset

// ThemeIrAssetModel is the odoo model name.
const ThemeIrAssetModel = "theme.ir.asset"

// Many2One convert ThemeIrAsset to *Many2One.
func (tia *ThemeIrAsset) Many2One() *Many2One {
	return NewMany2One(tia.Id.Get(), "")
}

// CreateThemeIrAsset creates a new theme.ir.asset model and returns its id.
func (c *Client) CreateThemeIrAsset(tia *ThemeIrAsset) (int64, error) {
	ids, err := c.CreateThemeIrAssets([]*ThemeIrAsset{tia})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateThemeIrAsset creates a new theme.ir.asset model and returns its id.
func (c *Client) CreateThemeIrAssets(tias []*ThemeIrAsset) ([]int64, error) {
	var vv []interface{}
	for _, v := range tias {
		vv = append(vv, v)
	}
	return c.Create(ThemeIrAssetModel, vv, nil)
}

// UpdateThemeIrAsset updates an existing theme.ir.asset record.
func (c *Client) UpdateThemeIrAsset(tia *ThemeIrAsset) error {
	return c.UpdateThemeIrAssets([]int64{tia.Id.Get()}, tia)
}

// UpdateThemeIrAssets updates existing theme.ir.asset records.
// All records (represented by ids) will be updated by tia values.
func (c *Client) UpdateThemeIrAssets(ids []int64, tia *ThemeIrAsset) error {
	return c.Update(ThemeIrAssetModel, ids, tia, nil)
}

// DeleteThemeIrAsset deletes an existing theme.ir.asset record.
func (c *Client) DeleteThemeIrAsset(id int64) error {
	return c.DeleteThemeIrAssets([]int64{id})
}

// DeleteThemeIrAssets deletes existing theme.ir.asset records.
func (c *Client) DeleteThemeIrAssets(ids []int64) error {
	return c.Delete(ThemeIrAssetModel, ids)
}

// GetThemeIrAsset gets theme.ir.asset existing record.
func (c *Client) GetThemeIrAsset(id int64) (*ThemeIrAsset, error) {
	tias, err := c.GetThemeIrAssets([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tias)[0]), nil
}

// GetThemeIrAssets gets theme.ir.asset existing records.
func (c *Client) GetThemeIrAssets(ids []int64) (*ThemeIrAssets, error) {
	tias := &ThemeIrAssets{}
	if err := c.Read(ThemeIrAssetModel, ids, nil, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAsset finds theme.ir.asset record by querying it with criteria.
func (c *Client) FindThemeIrAsset(criteria *Criteria) (*ThemeIrAsset, error) {
	tias := &ThemeIrAssets{}
	if err := c.SearchRead(ThemeIrAssetModel, criteria, NewOptions().Limit(1), tias); err != nil {
		return nil, err
	}
	return &((*tias)[0]), nil
}

// FindThemeIrAssets finds theme.ir.asset records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAssets(criteria *Criteria, options *Options) (*ThemeIrAssets, error) {
	tias := &ThemeIrAssets{}
	if err := c.SearchRead(ThemeIrAssetModel, criteria, options, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAssetIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAssetIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ThemeIrAssetModel, criteria, options)
}

// FindThemeIrAssetId finds record id by querying it with criteria.
func (c *Client) FindThemeIrAssetId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeIrAssetModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
