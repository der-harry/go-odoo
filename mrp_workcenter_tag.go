package odoo

// MrpWorkcenterTag represents mrp.workcenter.tag model.
type MrpWorkcenterTag struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenterTags represents array of mrp.workcenter.tag model.
type MrpWorkcenterTags []MrpWorkcenterTag

// MrpWorkcenterTagModel is the odoo model name.
const MrpWorkcenterTagModel = "mrp.workcenter.tag"

// Many2One convert MrpWorkcenterTag to *Many2One.
func (mwt *MrpWorkcenterTag) Many2One() *Many2One {
	return NewMany2One(mwt.Id.Get(), "")
}

// CreateMrpWorkcenterTag creates a new mrp.workcenter.tag model and returns its id.
func (c *Client) CreateMrpWorkcenterTag(mwt *MrpWorkcenterTag) (int64, error) {
	ids, err := c.CreateMrpWorkcenterTags([]*MrpWorkcenterTag{mwt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenterTag creates a new mrp.workcenter.tag model and returns its id.
func (c *Client) CreateMrpWorkcenterTags(mwts []*MrpWorkcenterTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range mwts {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterTagModel, vv, nil)
}

// UpdateMrpWorkcenterTag updates an existing mrp.workcenter.tag record.
func (c *Client) UpdateMrpWorkcenterTag(mwt *MrpWorkcenterTag) error {
	return c.UpdateMrpWorkcenterTags([]int64{mwt.Id.Get()}, mwt)
}

// UpdateMrpWorkcenterTags updates existing mrp.workcenter.tag records.
// All records (represented by ids) will be updated by mwt values.
func (c *Client) UpdateMrpWorkcenterTags(ids []int64, mwt *MrpWorkcenterTag) error {
	return c.Update(MrpWorkcenterTagModel, ids, mwt, nil)
}

// DeleteMrpWorkcenterTag deletes an existing mrp.workcenter.tag record.
func (c *Client) DeleteMrpWorkcenterTag(id int64) error {
	return c.DeleteMrpWorkcenterTags([]int64{id})
}

// DeleteMrpWorkcenterTags deletes existing mrp.workcenter.tag records.
func (c *Client) DeleteMrpWorkcenterTags(ids []int64) error {
	return c.Delete(MrpWorkcenterTagModel, ids)
}

// GetMrpWorkcenterTag gets mrp.workcenter.tag existing record.
func (c *Client) GetMrpWorkcenterTag(id int64) (*MrpWorkcenterTag, error) {
	mwts, err := c.GetMrpWorkcenterTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mwts)[0]), nil
}

// GetMrpWorkcenterTags gets mrp.workcenter.tag existing records.
func (c *Client) GetMrpWorkcenterTags(ids []int64) (*MrpWorkcenterTags, error) {
	mwts := &MrpWorkcenterTags{}
	if err := c.Read(MrpWorkcenterTagModel, ids, nil, mwts); err != nil {
		return nil, err
	}
	return mwts, nil
}

// FindMrpWorkcenterTag finds mrp.workcenter.tag record by querying it with criteria.
func (c *Client) FindMrpWorkcenterTag(criteria *Criteria) (*MrpWorkcenterTag, error) {
	mwts := &MrpWorkcenterTags{}
	if err := c.SearchRead(MrpWorkcenterTagModel, criteria, NewOptions().Limit(1), mwts); err != nil {
		return nil, err
	}
	return &((*mwts)[0]), nil
}

// FindMrpWorkcenterTags finds mrp.workcenter.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterTags(criteria *Criteria, options *Options) (*MrpWorkcenterTags, error) {
	mwts := &MrpWorkcenterTags{}
	if err := c.SearchRead(MrpWorkcenterTagModel, criteria, options, mwts); err != nil {
		return nil, err
	}
	return mwts, nil
}

// FindMrpWorkcenterTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterTagModel, criteria, options)
}

// FindMrpWorkcenterTagId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
