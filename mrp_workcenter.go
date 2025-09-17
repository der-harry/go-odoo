package odoo

// MrpWorkcenter represents mrp.workcenter model.
type MrpWorkcenter struct {
	Active                         *Bool       `xmlrpc:"active,omitempty"`
	AlternativeWorkcenterIds       *Relation   `xmlrpc:"alternative_workcenter_ids,omitempty"`
	AnalyticDistribution           interface{} `xmlrpc:"analytic_distribution,omitempty"`
	AnalyticPrecision              *Int        `xmlrpc:"analytic_precision,omitempty"`
	BlockedTime                    *Float      `xmlrpc:"blocked_time,omitempty"`
	CapacityIds                    *Relation   `xmlrpc:"capacity_ids,omitempty"`
	Code                           *String     `xmlrpc:"code,omitempty"`
	Color                          *Int        `xmlrpc:"color,omitempty"`
	CompanyId                      *Many2One   `xmlrpc:"company_id,omitempty"`
	CostsHour                      *Float      `xmlrpc:"costs_hour,omitempty"`
	CostsHourAccountIds            *Relation   `xmlrpc:"costs_hour_account_ids,omitempty"`
	CreateDate                     *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                     *Many2One   `xmlrpc:"currency_id,omitempty"`
	DefaultCapacity                *Float      `xmlrpc:"default_capacity,omitempty"`
	DisplayName                    *String     `xmlrpc:"display_name,omitempty"`
	DistributionAnalyticAccountIds *Relation   `xmlrpc:"distribution_analytic_account_ids,omitempty"`
	ExpenseAccountId               *Many2One   `xmlrpc:"expense_account_id,omitempty"`
	HasMessage                     *Bool       `xmlrpc:"has_message,omitempty"`
	HasRoutingLines                *Bool       `xmlrpc:"has_routing_lines,omitempty"`
	Id                             *Int        `xmlrpc:"id,omitempty"`
	KanbanDashboardGraph           *String     `xmlrpc:"kanban_dashboard_graph,omitempty"`
	MessageAttachmentCount         *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds             *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter         *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError             *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                     *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower              *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction              *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter       *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds              *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	Name                           *String     `xmlrpc:"name,omitempty"`
	Note                           *String     `xmlrpc:"note,omitempty"`
	Oee                            *Float      `xmlrpc:"oee,omitempty"`
	OeeTarget                      *Float      `xmlrpc:"oee_target,omitempty"`
	OrderIds                       *Relation   `xmlrpc:"order_ids,omitempty"`
	Performance                    *Int        `xmlrpc:"performance,omitempty"`
	ProductiveTime                 *Float      `xmlrpc:"productive_time,omitempty"`
	RatingIds                      *Relation   `xmlrpc:"rating_ids,omitempty"`
	ResourceCalendarId             *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId                     *Many2One   `xmlrpc:"resource_id,omitempty"`
	RoutingLineIds                 *Relation   `xmlrpc:"routing_line_ids,omitempty"`
	Sequence                       *Int        `xmlrpc:"sequence,omitempty"`
	TagIds                         *Relation   `xmlrpc:"tag_ids,omitempty"`
	TimeEfficiency                 *Float      `xmlrpc:"time_efficiency,omitempty"`
	TimeIds                        *Relation   `xmlrpc:"time_ids,omitempty"`
	TimeStart                      *Float      `xmlrpc:"time_start,omitempty"`
	TimeStop                       *Float      `xmlrpc:"time_stop,omitempty"`
	Tz                             *Selection  `xmlrpc:"tz,omitempty"`
	WebsiteMessageIds              *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WorkcenterLoad                 *Float      `xmlrpc:"workcenter_load,omitempty"`
	WorkingState                   *Selection  `xmlrpc:"working_state,omitempty"`
	WorkorderCount                 *Int        `xmlrpc:"workorder_count,omitempty"`
	WorkorderLateCount             *Int        `xmlrpc:"workorder_late_count,omitempty"`
	WorkorderPendingCount          *Int        `xmlrpc:"workorder_pending_count,omitempty"`
	WorkorderProgressCount         *Int        `xmlrpc:"workorder_progress_count,omitempty"`
	WorkorderReadyCount            *Int        `xmlrpc:"workorder_ready_count,omitempty"`
	WriteDate                      *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkcenters represents array of mrp.workcenter model.
type MrpWorkcenters []MrpWorkcenter

// MrpWorkcenterModel is the odoo model name.
const MrpWorkcenterModel = "mrp.workcenter"

// Many2One convert MrpWorkcenter to *Many2One.
func (mw *MrpWorkcenter) Many2One() *Many2One {
	return NewMany2One(mw.Id.Get(), "")
}

// CreateMrpWorkcenter creates a new mrp.workcenter model and returns its id.
func (c *Client) CreateMrpWorkcenter(mw *MrpWorkcenter) (int64, error) {
	ids, err := c.CreateMrpWorkcenters([]*MrpWorkcenter{mw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkcenter creates a new mrp.workcenter model and returns its id.
func (c *Client) CreateMrpWorkcenters(mws []*MrpWorkcenter) ([]int64, error) {
	var vv []interface{}
	for _, v := range mws {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkcenterModel, vv, nil)
}

// UpdateMrpWorkcenter updates an existing mrp.workcenter record.
func (c *Client) UpdateMrpWorkcenter(mw *MrpWorkcenter) error {
	return c.UpdateMrpWorkcenters([]int64{mw.Id.Get()}, mw)
}

// UpdateMrpWorkcenters updates existing mrp.workcenter records.
// All records (represented by ids) will be updated by mw values.
func (c *Client) UpdateMrpWorkcenters(ids []int64, mw *MrpWorkcenter) error {
	return c.Update(MrpWorkcenterModel, ids, mw, nil)
}

// DeleteMrpWorkcenter deletes an existing mrp.workcenter record.
func (c *Client) DeleteMrpWorkcenter(id int64) error {
	return c.DeleteMrpWorkcenters([]int64{id})
}

// DeleteMrpWorkcenters deletes existing mrp.workcenter records.
func (c *Client) DeleteMrpWorkcenters(ids []int64) error {
	return c.Delete(MrpWorkcenterModel, ids)
}

// GetMrpWorkcenter gets mrp.workcenter existing record.
func (c *Client) GetMrpWorkcenter(id int64) (*MrpWorkcenter, error) {
	mws, err := c.GetMrpWorkcenters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// GetMrpWorkcenters gets mrp.workcenter existing records.
func (c *Client) GetMrpWorkcenters(ids []int64) (*MrpWorkcenters, error) {
	mws := &MrpWorkcenters{}
	if err := c.Read(MrpWorkcenterModel, ids, nil, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkcenter finds mrp.workcenter record by querying it with criteria.
func (c *Client) FindMrpWorkcenter(criteria *Criteria) (*MrpWorkcenter, error) {
	mws := &MrpWorkcenters{}
	if err := c.SearchRead(MrpWorkcenterModel, criteria, NewOptions().Limit(1), mws); err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// FindMrpWorkcenters finds mrp.workcenter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenters(criteria *Criteria, options *Options) (*MrpWorkcenters, error) {
	mws := &MrpWorkcenters{}
	if err := c.SearchRead(MrpWorkcenterModel, criteria, options, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkcenterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkcenterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkcenterModel, criteria, options)
}

// FindMrpWorkcenterId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkcenterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkcenterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
