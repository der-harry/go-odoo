package odoo

// MrpUnbuild represents mrp.unbuild model.
type MrpUnbuild struct {
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	BomId                       *Many2One  `xmlrpc:"bom_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ConsumeLineIds              *Relation  `xmlrpc:"consume_line_ids,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	HasTracking                 *Selection `xmlrpc:"has_tracking,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	LocationDestId              *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationId                  *Many2One  `xmlrpc:"location_id,omitempty"`
	LotId                       *Many2One  `xmlrpc:"lot_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoBomId                     *Many2One  `xmlrpc:"mo_bom_id,omitempty"`
	MoId                        *Many2One  `xmlrpc:"mo_id,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	ProduceLineIds              *Relation  `xmlrpc:"produce_line_ids,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty                  *Float     `xmlrpc:"product_qty,omitempty"`
	ProductUomId                *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpUnbuilds represents array of mrp.unbuild model.
type MrpUnbuilds []MrpUnbuild

// MrpUnbuildModel is the odoo model name.
const MrpUnbuildModel = "mrp.unbuild"

// Many2One convert MrpUnbuild to *Many2One.
func (mu *MrpUnbuild) Many2One() *Many2One {
	return NewMany2One(mu.Id.Get(), "")
}

// CreateMrpUnbuild creates a new mrp.unbuild model and returns its id.
func (c *Client) CreateMrpUnbuild(mu *MrpUnbuild) (int64, error) {
	ids, err := c.CreateMrpUnbuilds([]*MrpUnbuild{mu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpUnbuild creates a new mrp.unbuild model and returns its id.
func (c *Client) CreateMrpUnbuilds(mus []*MrpUnbuild) ([]int64, error) {
	var vv []interface{}
	for _, v := range mus {
		vv = append(vv, v)
	}
	return c.Create(MrpUnbuildModel, vv, nil)
}

// UpdateMrpUnbuild updates an existing mrp.unbuild record.
func (c *Client) UpdateMrpUnbuild(mu *MrpUnbuild) error {
	return c.UpdateMrpUnbuilds([]int64{mu.Id.Get()}, mu)
}

// UpdateMrpUnbuilds updates existing mrp.unbuild records.
// All records (represented by ids) will be updated by mu values.
func (c *Client) UpdateMrpUnbuilds(ids []int64, mu *MrpUnbuild) error {
	return c.Update(MrpUnbuildModel, ids, mu, nil)
}

// DeleteMrpUnbuild deletes an existing mrp.unbuild record.
func (c *Client) DeleteMrpUnbuild(id int64) error {
	return c.DeleteMrpUnbuilds([]int64{id})
}

// DeleteMrpUnbuilds deletes existing mrp.unbuild records.
func (c *Client) DeleteMrpUnbuilds(ids []int64) error {
	return c.Delete(MrpUnbuildModel, ids)
}

// GetMrpUnbuild gets mrp.unbuild existing record.
func (c *Client) GetMrpUnbuild(id int64) (*MrpUnbuild, error) {
	mus, err := c.GetMrpUnbuilds([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mus)[0]), nil
}

// GetMrpUnbuilds gets mrp.unbuild existing records.
func (c *Client) GetMrpUnbuilds(ids []int64) (*MrpUnbuilds, error) {
	mus := &MrpUnbuilds{}
	if err := c.Read(MrpUnbuildModel, ids, nil, mus); err != nil {
		return nil, err
	}
	return mus, nil
}

// FindMrpUnbuild finds mrp.unbuild record by querying it with criteria.
func (c *Client) FindMrpUnbuild(criteria *Criteria) (*MrpUnbuild, error) {
	mus := &MrpUnbuilds{}
	if err := c.SearchRead(MrpUnbuildModel, criteria, NewOptions().Limit(1), mus); err != nil {
		return nil, err
	}
	return &((*mus)[0]), nil
}

// FindMrpUnbuilds finds mrp.unbuild records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpUnbuilds(criteria *Criteria, options *Options) (*MrpUnbuilds, error) {
	mus := &MrpUnbuilds{}
	if err := c.SearchRead(MrpUnbuildModel, criteria, options, mus); err != nil {
		return nil, err
	}
	return mus, nil
}

// FindMrpUnbuildIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpUnbuildIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpUnbuildModel, criteria, options)
}

// FindMrpUnbuildId finds record id by querying it with criteria.
func (c *Client) FindMrpUnbuildId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpUnbuildModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
