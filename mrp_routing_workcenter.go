package odoo

// MrpRoutingWorkcenter represents mrp.routing.workcenter model.
type MrpRoutingWorkcenter struct {
	Active                                      *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId                     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration                 *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllowOperationDependencies                  *Bool      `xmlrpc:"allow_operation_dependencies,omitempty"`
	BlockedByOperationIds                       *Relation  `xmlrpc:"blocked_by_operation_ids,omitempty"`
	BomId                                       *Many2One  `xmlrpc:"bom_id,omitempty"`
	BomProductTemplateAttributeValueIds         *Relation  `xmlrpc:"bom_product_template_attribute_value_ids,omitempty"`
	CompanyId                                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	MessageAttachmentCount                      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter                    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline                      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                        *String    `xmlrpc:"name,omitempty"`
	NeededByOperationIds                        *Relation  `xmlrpc:"needed_by_operation_ids,omitempty"`
	Note                                        *String    `xmlrpc:"note,omitempty"`
	PossibleBomProductTemplateAttributeValueIds *Relation  `xmlrpc:"possible_bom_product_template_attribute_value_ids,omitempty"`
	RatingIds                                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	Sequence                                    *Int       `xmlrpc:"sequence,omitempty"`
	TimeComputedOn                              *String    `xmlrpc:"time_computed_on,omitempty"`
	TimeCycle                                   *Float     `xmlrpc:"time_cycle,omitempty"`
	TimeCycleManual                             *Float     `xmlrpc:"time_cycle_manual,omitempty"`
	TimeMode                                    *Selection `xmlrpc:"time_mode,omitempty"`
	TimeModeBatch                               *Int       `xmlrpc:"time_mode_batch,omitempty"`
	WebsiteMessageIds                           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkcenterId                                *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkorderCount                              *Int       `xmlrpc:"workorder_count,omitempty"`
	WorkorderIds                                *Relation  `xmlrpc:"workorder_ids,omitempty"`
	Worksheet                                   *String    `xmlrpc:"worksheet,omitempty"`
	WorksheetGoogleSlide                        *String    `xmlrpc:"worksheet_google_slide,omitempty"`
	WorksheetType                               *Selection `xmlrpc:"worksheet_type,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpRoutingWorkcenters represents array of mrp.routing.workcenter model.
type MrpRoutingWorkcenters []MrpRoutingWorkcenter

// MrpRoutingWorkcenterModel is the odoo model name.
const MrpRoutingWorkcenterModel = "mrp.routing.workcenter"

// Many2One convert MrpRoutingWorkcenter to *Many2One.
func (mrw *MrpRoutingWorkcenter) Many2One() *Many2One {
	return NewMany2One(mrw.Id.Get(), "")
}

// CreateMrpRoutingWorkcenter creates a new mrp.routing.workcenter model and returns its id.
func (c *Client) CreateMrpRoutingWorkcenter(mrw *MrpRoutingWorkcenter) (int64, error) {
	ids, err := c.CreateMrpRoutingWorkcenters([]*MrpRoutingWorkcenter{mrw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpRoutingWorkcenter creates a new mrp.routing.workcenter model and returns its id.
func (c *Client) CreateMrpRoutingWorkcenters(mrws []*MrpRoutingWorkcenter) ([]int64, error) {
	var vv []interface{}
	for _, v := range mrws {
		vv = append(vv, v)
	}
	return c.Create(MrpRoutingWorkcenterModel, vv, nil)
}

// UpdateMrpRoutingWorkcenter updates an existing mrp.routing.workcenter record.
func (c *Client) UpdateMrpRoutingWorkcenter(mrw *MrpRoutingWorkcenter) error {
	return c.UpdateMrpRoutingWorkcenters([]int64{mrw.Id.Get()}, mrw)
}

// UpdateMrpRoutingWorkcenters updates existing mrp.routing.workcenter records.
// All records (represented by ids) will be updated by mrw values.
func (c *Client) UpdateMrpRoutingWorkcenters(ids []int64, mrw *MrpRoutingWorkcenter) error {
	return c.Update(MrpRoutingWorkcenterModel, ids, mrw, nil)
}

// DeleteMrpRoutingWorkcenter deletes an existing mrp.routing.workcenter record.
func (c *Client) DeleteMrpRoutingWorkcenter(id int64) error {
	return c.DeleteMrpRoutingWorkcenters([]int64{id})
}

// DeleteMrpRoutingWorkcenters deletes existing mrp.routing.workcenter records.
func (c *Client) DeleteMrpRoutingWorkcenters(ids []int64) error {
	return c.Delete(MrpRoutingWorkcenterModel, ids)
}

// GetMrpRoutingWorkcenter gets mrp.routing.workcenter existing record.
func (c *Client) GetMrpRoutingWorkcenter(id int64) (*MrpRoutingWorkcenter, error) {
	mrws, err := c.GetMrpRoutingWorkcenters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mrws)[0]), nil
}

// GetMrpRoutingWorkcenters gets mrp.routing.workcenter existing records.
func (c *Client) GetMrpRoutingWorkcenters(ids []int64) (*MrpRoutingWorkcenters, error) {
	mrws := &MrpRoutingWorkcenters{}
	if err := c.Read(MrpRoutingWorkcenterModel, ids, nil, mrws); err != nil {
		return nil, err
	}
	return mrws, nil
}

// FindMrpRoutingWorkcenter finds mrp.routing.workcenter record by querying it with criteria.
func (c *Client) FindMrpRoutingWorkcenter(criteria *Criteria) (*MrpRoutingWorkcenter, error) {
	mrws := &MrpRoutingWorkcenters{}
	if err := c.SearchRead(MrpRoutingWorkcenterModel, criteria, NewOptions().Limit(1), mrws); err != nil {
		return nil, err
	}
	return &((*mrws)[0]), nil
}

// FindMrpRoutingWorkcenters finds mrp.routing.workcenter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpRoutingWorkcenters(criteria *Criteria, options *Options) (*MrpRoutingWorkcenters, error) {
	mrws := &MrpRoutingWorkcenters{}
	if err := c.SearchRead(MrpRoutingWorkcenterModel, criteria, options, mrws); err != nil {
		return nil, err
	}
	return mrws, nil
}

// FindMrpRoutingWorkcenterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpRoutingWorkcenterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpRoutingWorkcenterModel, criteria, options)
}

// FindMrpRoutingWorkcenterId finds record id by querying it with criteria.
func (c *Client) FindMrpRoutingWorkcenterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpRoutingWorkcenterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
