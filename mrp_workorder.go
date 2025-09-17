package odoo

// MrpWorkorder represents mrp.workorder model.
type MrpWorkorder struct {
	AllowWorkorderDependencies *Bool      `xmlrpc:"allow_workorder_dependencies,omitempty"`
	Barcode                    *String    `xmlrpc:"barcode,omitempty"`
	BlockedByWorkorderIds      *Relation  `xmlrpc:"blocked_by_workorder_ids,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	Consumption                *Selection `xmlrpc:"consumption,omitempty"`
	CostsHour                  *Float     `xmlrpc:"costs_hour,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFinished               *Time      `xmlrpc:"date_finished,omitempty"`
	DateStart                  *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	Duration                   *Float     `xmlrpc:"duration,omitempty"`
	DurationExpected           *Float     `xmlrpc:"duration_expected,omitempty"`
	DurationPercent            *Int       `xmlrpc:"duration_percent,omitempty"`
	DurationUnit               *Float     `xmlrpc:"duration_unit,omitempty"`
	FinishedLotId              *Many2One  `xmlrpc:"finished_lot_id,omitempty"`
	HasWorksheet               *Bool      `xmlrpc:"has_worksheet,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	IsPlanned                  *Bool      `xmlrpc:"is_planned,omitempty"`
	IsProduced                 *Bool      `xmlrpc:"is_produced,omitempty"`
	IsUserWorking              *Bool      `xmlrpc:"is_user_working,omitempty"`
	JsonPopover                *String    `xmlrpc:"json_popover,omitempty"`
	LastWorkingUserId          *Relation  `xmlrpc:"last_working_user_id,omitempty"`
	LeaveId                    *Many2One  `xmlrpc:"leave_id,omitempty"`
	MoAnalyticAccountLineIds   *Relation  `xmlrpc:"mo_analytic_account_line_ids,omitempty"`
	MoveFinishedIds            *Relation  `xmlrpc:"move_finished_ids,omitempty"`
	MoveLineIds                *Relation  `xmlrpc:"move_line_ids,omitempty"`
	MoveRawIds                 *Relation  `xmlrpc:"move_raw_ids,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NeededByWorkorderIds       *Relation  `xmlrpc:"needed_by_workorder_ids,omitempty"`
	OperationId                *Many2One  `xmlrpc:"operation_id,omitempty"`
	OperationNote              *String    `xmlrpc:"operation_note,omitempty"`
	ProductId                  *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductTracking            *Selection `xmlrpc:"product_tracking,omitempty"`
	ProductUomId               *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductionAvailability     *Selection `xmlrpc:"production_availability,omitempty"`
	ProductionBomId            *Many2One  `xmlrpc:"production_bom_id,omitempty"`
	ProductionDate             *Time      `xmlrpc:"production_date,omitempty"`
	ProductionId               *Many2One  `xmlrpc:"production_id,omitempty"`
	ProductionState            *Selection `xmlrpc:"production_state,omitempty"`
	Progress                   *Float     `xmlrpc:"progress,omitempty"`
	QtyProduced                *Float     `xmlrpc:"qty_produced,omitempty"`
	QtyProducing               *Float     `xmlrpc:"qty_producing,omitempty"`
	QtyProduction              *Float     `xmlrpc:"qty_production,omitempty"`
	QtyRemaining               *Float     `xmlrpc:"qty_remaining,omitempty"`
	QtyReportedFromPreviousWo  *Float     `xmlrpc:"qty_reported_from_previous_wo,omitempty"`
	ScrapCount                 *Int       `xmlrpc:"scrap_count,omitempty"`
	ScrapIds                   *Relation  `xmlrpc:"scrap_ids,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	ShowJsonPopover            *Bool      `xmlrpc:"show_json_popover,omitempty"`
	State                      *Selection `xmlrpc:"state,omitempty"`
	TimeIds                    *Relation  `xmlrpc:"time_ids,omitempty"`
	WcAnalyticAccountLineIds   *Relation  `xmlrpc:"wc_analytic_account_line_ids,omitempty"`
	WorkcenterId               *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkingState               *Selection `xmlrpc:"working_state,omitempty"`
	WorkingUserIds             *Relation  `xmlrpc:"working_user_ids,omitempty"`
	Worksheet                  *String    `xmlrpc:"worksheet,omitempty"`
	WorksheetGoogleSlide       *String    `xmlrpc:"worksheet_google_slide,omitempty"`
	WorksheetType              *Selection `xmlrpc:"worksheet_type,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpWorkorders represents array of mrp.workorder model.
type MrpWorkorders []MrpWorkorder

// MrpWorkorderModel is the odoo model name.
const MrpWorkorderModel = "mrp.workorder"

// Many2One convert MrpWorkorder to *Many2One.
func (mw *MrpWorkorder) Many2One() *Many2One {
	return NewMany2One(mw.Id.Get(), "")
}

// CreateMrpWorkorder creates a new mrp.workorder model and returns its id.
func (c *Client) CreateMrpWorkorder(mw *MrpWorkorder) (int64, error) {
	ids, err := c.CreateMrpWorkorders([]*MrpWorkorder{mw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpWorkorder creates a new mrp.workorder model and returns its id.
func (c *Client) CreateMrpWorkorders(mws []*MrpWorkorder) ([]int64, error) {
	var vv []interface{}
	for _, v := range mws {
		vv = append(vv, v)
	}
	return c.Create(MrpWorkorderModel, vv, nil)
}

// UpdateMrpWorkorder updates an existing mrp.workorder record.
func (c *Client) UpdateMrpWorkorder(mw *MrpWorkorder) error {
	return c.UpdateMrpWorkorders([]int64{mw.Id.Get()}, mw)
}

// UpdateMrpWorkorders updates existing mrp.workorder records.
// All records (represented by ids) will be updated by mw values.
func (c *Client) UpdateMrpWorkorders(ids []int64, mw *MrpWorkorder) error {
	return c.Update(MrpWorkorderModel, ids, mw, nil)
}

// DeleteMrpWorkorder deletes an existing mrp.workorder record.
func (c *Client) DeleteMrpWorkorder(id int64) error {
	return c.DeleteMrpWorkorders([]int64{id})
}

// DeleteMrpWorkorders deletes existing mrp.workorder records.
func (c *Client) DeleteMrpWorkorders(ids []int64) error {
	return c.Delete(MrpWorkorderModel, ids)
}

// GetMrpWorkorder gets mrp.workorder existing record.
func (c *Client) GetMrpWorkorder(id int64) (*MrpWorkorder, error) {
	mws, err := c.GetMrpWorkorders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// GetMrpWorkorders gets mrp.workorder existing records.
func (c *Client) GetMrpWorkorders(ids []int64) (*MrpWorkorders, error) {
	mws := &MrpWorkorders{}
	if err := c.Read(MrpWorkorderModel, ids, nil, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkorder finds mrp.workorder record by querying it with criteria.
func (c *Client) FindMrpWorkorder(criteria *Criteria) (*MrpWorkorder, error) {
	mws := &MrpWorkorders{}
	if err := c.SearchRead(MrpWorkorderModel, criteria, NewOptions().Limit(1), mws); err != nil {
		return nil, err
	}
	return &((*mws)[0]), nil
}

// FindMrpWorkorders finds mrp.workorder records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkorders(criteria *Criteria, options *Options) (*MrpWorkorders, error) {
	mws := &MrpWorkorders{}
	if err := c.SearchRead(MrpWorkorderModel, criteria, options, mws); err != nil {
		return nil, err
	}
	return mws, nil
}

// FindMrpWorkorderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpWorkorderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpWorkorderModel, criteria, options)
}

// FindMrpWorkorderId finds record id by querying it with criteria.
func (c *Client) FindMrpWorkorderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpWorkorderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
