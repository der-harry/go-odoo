package odoo

// ThemeIrAttachment represents theme.ir.attachment model.
type ThemeIrAttachment struct {
	CopyIds     *Relation `xmlrpc:"copy_ids,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Key         *String   `xmlrpc:"key,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Url         *String   `xmlrpc:"url,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ThemeIrAttachments represents array of theme.ir.attachment model.
type ThemeIrAttachments []ThemeIrAttachment

// ThemeIrAttachmentModel is the odoo model name.
const ThemeIrAttachmentModel = "theme.ir.attachment"

// Many2One convert ThemeIrAttachment to *Many2One.
func (tia *ThemeIrAttachment) Many2One() *Many2One {
	return NewMany2One(tia.Id.Get(), "")
}

// CreateThemeIrAttachment creates a new theme.ir.attachment model and returns its id.
func (c *Client) CreateThemeIrAttachment(tia *ThemeIrAttachment) (int64, error) {
	ids, err := c.CreateThemeIrAttachments([]*ThemeIrAttachment{tia})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateThemeIrAttachment creates a new theme.ir.attachment model and returns its id.
func (c *Client) CreateThemeIrAttachments(tias []*ThemeIrAttachment) ([]int64, error) {
	var vv []interface{}
	for _, v := range tias {
		vv = append(vv, v)
	}
	return c.Create(ThemeIrAttachmentModel, vv, nil)
}

// UpdateThemeIrAttachment updates an existing theme.ir.attachment record.
func (c *Client) UpdateThemeIrAttachment(tia *ThemeIrAttachment) error {
	return c.UpdateThemeIrAttachments([]int64{tia.Id.Get()}, tia)
}

// UpdateThemeIrAttachments updates existing theme.ir.attachment records.
// All records (represented by ids) will be updated by tia values.
func (c *Client) UpdateThemeIrAttachments(ids []int64, tia *ThemeIrAttachment) error {
	return c.Update(ThemeIrAttachmentModel, ids, tia, nil)
}

// DeleteThemeIrAttachment deletes an existing theme.ir.attachment record.
func (c *Client) DeleteThemeIrAttachment(id int64) error {
	return c.DeleteThemeIrAttachments([]int64{id})
}

// DeleteThemeIrAttachments deletes existing theme.ir.attachment records.
func (c *Client) DeleteThemeIrAttachments(ids []int64) error {
	return c.Delete(ThemeIrAttachmentModel, ids)
}

// GetThemeIrAttachment gets theme.ir.attachment existing record.
func (c *Client) GetThemeIrAttachment(id int64) (*ThemeIrAttachment, error) {
	tias, err := c.GetThemeIrAttachments([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*tias)[0]), nil
}

// GetThemeIrAttachments gets theme.ir.attachment existing records.
func (c *Client) GetThemeIrAttachments(ids []int64) (*ThemeIrAttachments, error) {
	tias := &ThemeIrAttachments{}
	if err := c.Read(ThemeIrAttachmentModel, ids, nil, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAttachment finds theme.ir.attachment record by querying it with criteria.
func (c *Client) FindThemeIrAttachment(criteria *Criteria) (*ThemeIrAttachment, error) {
	tias := &ThemeIrAttachments{}
	if err := c.SearchRead(ThemeIrAttachmentModel, criteria, NewOptions().Limit(1), tias); err != nil {
		return nil, err
	}
	return &((*tias)[0]), nil
}

// FindThemeIrAttachments finds theme.ir.attachment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAttachments(criteria *Criteria, options *Options) (*ThemeIrAttachments, error) {
	tias := &ThemeIrAttachments{}
	if err := c.SearchRead(ThemeIrAttachmentModel, criteria, options, tias); err != nil {
		return nil, err
	}
	return tias, nil
}

// FindThemeIrAttachmentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindThemeIrAttachmentIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ThemeIrAttachmentModel, criteria, options)
}

// FindThemeIrAttachmentId finds record id by querying it with criteria.
func (c *Client) FindThemeIrAttachmentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ThemeIrAttachmentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
