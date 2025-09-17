package odoo

// MrpProduction represents mrp.production model.
type MrpProduction struct {
	ActivityCalendarEventId               *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                  *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration           *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                 *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                           *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                         *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                       *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                      *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                        *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                        *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllMoveIds                            *Relation  `xmlrpc:"all_move_ids,omitempty"`
	AllMoveRawIds                         *Relation  `xmlrpc:"all_move_raw_ids,omitempty"`
	AllowWorkorderDependencies            *Bool      `xmlrpc:"allow_workorder_dependencies,omitempty"`
	BackorderSequence                     *Int       `xmlrpc:"backorder_sequence,omitempty"`
	BomId                                 *Many2One  `xmlrpc:"bom_id,omitempty"`
	CompanyId                             *Many2One  `xmlrpc:"company_id,omitempty"`
	ComponentsAvailability                *String    `xmlrpc:"components_availability,omitempty"`
	ComponentsAvailabilityState           *Selection `xmlrpc:"components_availability_state,omitempty"`
	Consumption                           *Selection `xmlrpc:"consumption,omitempty"`
	CreateDate                            *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                             *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateDeadline                          *Time      `xmlrpc:"date_deadline,omitempty"`
	DateFinished                          *Time      `xmlrpc:"date_finished,omitempty"`
	DateStart                             *Time      `xmlrpc:"date_start,omitempty"`
	DelayAlertDate                        *Time      `xmlrpc:"delay_alert_date,omitempty"`
	DeliveryCount                         *Int       `xmlrpc:"delivery_count,omitempty"`
	DisplayName                           *String    `xmlrpc:"display_name,omitempty"`
	Duration                              *Float     `xmlrpc:"duration,omitempty"`
	DurationExpected                      *Float     `xmlrpc:"duration_expected,omitempty"`
	ExtraCost                             *Float     `xmlrpc:"extra_cost,omitempty"`
	FinishedMoveLineIds                   *Relation  `xmlrpc:"finished_move_line_ids,omitempty"`
	ForecastedIssue                       *Bool      `xmlrpc:"forecasted_issue,omitempty"`
	HasAnalyticAccount                    *Bool      `xmlrpc:"has_analytic_account,omitempty"`
	HasMessage                            *Bool      `xmlrpc:"has_message,omitempty"`
	Id                                    *Int       `xmlrpc:"id,omitempty"`
	IsDelayed                             *Bool      `xmlrpc:"is_delayed,omitempty"`
	IsLocked                              *Bool      `xmlrpc:"is_locked,omitempty"`
	IsOutdatedBom                         *Bool      `xmlrpc:"is_outdated_bom,omitempty"`
	IsPlanned                             *Bool      `xmlrpc:"is_planned,omitempty"`
	JsonPopover                           *String    `xmlrpc:"json_popover,omitempty"`
	LocationDestId                        *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationFinalId                       *Many2One  `xmlrpc:"location_final_id,omitempty"`
	LocationSrcId                         *Many2One  `xmlrpc:"location_src_id,omitempty"`
	LotProducingId                        *Many2One  `xmlrpc:"lot_producing_id,omitempty"`
	MessageAttachmentCount                *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                    *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                       *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                    *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                            *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                     *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                     *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter              *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                     *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveByproductIds                      *Relation  `xmlrpc:"move_byproduct_ids,omitempty"`
	MoveDestIds                           *Relation  `xmlrpc:"move_dest_ids,omitempty"`
	MoveFinishedIds                       *Relation  `xmlrpc:"move_finished_ids,omitempty"`
	MoveRawIds                            *Relation  `xmlrpc:"move_raw_ids,omitempty"`
	MrpProductionBackorderCount           *Int       `xmlrpc:"mrp_production_backorder_count,omitempty"`
	MrpProductionChildCount               *Int       `xmlrpc:"mrp_production_child_count,omitempty"`
	MrpProductionSourceCount              *Int       `xmlrpc:"mrp_production_source_count,omitempty"`
	MyActivityDateDeadline                *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                  *String    `xmlrpc:"name,omitempty"`
	NeverProductTemplateAttributeValueIds *Relation  `xmlrpc:"never_product_template_attribute_value_ids,omitempty"`
	OrderpointId                          *Many2One  `xmlrpc:"orderpoint_id,omitempty"`
	Origin                                *String    `xmlrpc:"origin,omitempty"`
	PickingIds                            *Relation  `xmlrpc:"picking_ids,omitempty"`
	PickingTypeId                         *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	Priority                              *Selection `xmlrpc:"priority,omitempty"`
	ProcurementGroupId                    *Many2One  `xmlrpc:"procurement_group_id,omitempty"`
	ProductDescriptionVariants            *String    `xmlrpc:"product_description_variants,omitempty"`
	ProductId                             *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductQty                            *Float     `xmlrpc:"product_qty,omitempty"`
	ProductTmplId                         *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductTracking                       *Selection `xmlrpc:"product_tracking,omitempty"`
	ProductUomCategoryId                  *Many2One  `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId                          *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	ProductUomQty                         *Float     `xmlrpc:"product_uom_qty,omitempty"`
	ProductVariantAttributes              *Relation  `xmlrpc:"product_variant_attributes,omitempty"`
	ProductionCapacity                    *Float     `xmlrpc:"production_capacity,omitempty"`
	ProductionLocationId                  *Many2One  `xmlrpc:"production_location_id,omitempty"`
	ProjectId                             *Many2One  `xmlrpc:"project_id,omitempty"`
	PropagateCancel                       *Bool      `xmlrpc:"propagate_cancel,omitempty"`
	PurchaseOrderCount                    *Int       `xmlrpc:"purchase_order_count,omitempty"`
	QtyProduced                           *Float     `xmlrpc:"qty_produced,omitempty"`
	QtyProducing                          *Float     `xmlrpc:"qty_producing,omitempty"`
	RatingIds                             *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReservationState                      *Selection `xmlrpc:"reservation_state,omitempty"`
	ReserveVisible                        *Bool      `xmlrpc:"reserve_visible,omitempty"`
	SaleLineId                            *Many2One  `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderCount                        *Int       `xmlrpc:"sale_order_count,omitempty"`
	ScrapCount                            *Int       `xmlrpc:"scrap_count,omitempty"`
	ScrapIds                              *Relation  `xmlrpc:"scrap_ids,omitempty"`
	SearchDateCategory                    *Selection `xmlrpc:"search_date_category,omitempty"`
	ShowAllocation                        *Bool      `xmlrpc:"show_allocation,omitempty"`
	ShowFinalLots                         *Bool      `xmlrpc:"show_final_lots,omitempty"`
	ShowLock                              *Bool      `xmlrpc:"show_lock,omitempty"`
	ShowLotIds                            *Bool      `xmlrpc:"show_lot_ids,omitempty"`
	ShowProduce                           *Bool      `xmlrpc:"show_produce,omitempty"`
	ShowProduceAll                        *Bool      `xmlrpc:"show_produce_all,omitempty"`
	ShowValuation                         *Bool      `xmlrpc:"show_valuation,omitempty"`
	State                                 *Selection `xmlrpc:"state,omitempty"`
	UnbuildCount                          *Int       `xmlrpc:"unbuild_count,omitempty"`
	UnbuildIds                            *Relation  `xmlrpc:"unbuild_ids,omitempty"`
	UnreserveVisible                      *Bool      `xmlrpc:"unreserve_visible,omitempty"`
	UseCreateComponentsLots               *Bool      `xmlrpc:"use_create_components_lots,omitempty"`
	UserId                                *Many2One  `xmlrpc:"user_id,omitempty"`
	ValidProductTemplateAttributeLineIds  *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	WarehouseId                           *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds                     *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkcenterId                          *Many2One  `xmlrpc:"workcenter_id,omitempty"`
	WorkorderIds                          *Relation  `xmlrpc:"workorder_ids,omitempty"`
	WriteDate                             *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                              *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MrpProductions represents array of mrp.production model.
type MrpProductions []MrpProduction

// MrpProductionModel is the odoo model name.
const MrpProductionModel = "mrp.production"

// Many2One convert MrpProduction to *Many2One.
func (mp *MrpProduction) Many2One() *Many2One {
	return NewMany2One(mp.Id.Get(), "")
}

// CreateMrpProduction creates a new mrp.production model and returns its id.
func (c *Client) CreateMrpProduction(mp *MrpProduction) (int64, error) {
	ids, err := c.CreateMrpProductions([]*MrpProduction{mp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMrpProduction creates a new mrp.production model and returns its id.
func (c *Client) CreateMrpProductions(mps []*MrpProduction) ([]int64, error) {
	var vv []interface{}
	for _, v := range mps {
		vv = append(vv, v)
	}
	return c.Create(MrpProductionModel, vv, nil)
}

// UpdateMrpProduction updates an existing mrp.production record.
func (c *Client) UpdateMrpProduction(mp *MrpProduction) error {
	return c.UpdateMrpProductions([]int64{mp.Id.Get()}, mp)
}

// UpdateMrpProductions updates existing mrp.production records.
// All records (represented by ids) will be updated by mp values.
func (c *Client) UpdateMrpProductions(ids []int64, mp *MrpProduction) error {
	return c.Update(MrpProductionModel, ids, mp, nil)
}

// DeleteMrpProduction deletes an existing mrp.production record.
func (c *Client) DeleteMrpProduction(id int64) error {
	return c.DeleteMrpProductions([]int64{id})
}

// DeleteMrpProductions deletes existing mrp.production records.
func (c *Client) DeleteMrpProductions(ids []int64) error {
	return c.Delete(MrpProductionModel, ids)
}

// GetMrpProduction gets mrp.production existing record.
func (c *Client) GetMrpProduction(id int64) (*MrpProduction, error) {
	mps, err := c.GetMrpProductions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// GetMrpProductions gets mrp.production existing records.
func (c *Client) GetMrpProductions(ids []int64) (*MrpProductions, error) {
	mps := &MrpProductions{}
	if err := c.Read(MrpProductionModel, ids, nil, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMrpProduction finds mrp.production record by querying it with criteria.
func (c *Client) FindMrpProduction(criteria *Criteria) (*MrpProduction, error) {
	mps := &MrpProductions{}
	if err := c.SearchRead(MrpProductionModel, criteria, NewOptions().Limit(1), mps); err != nil {
		return nil, err
	}
	return &((*mps)[0]), nil
}

// FindMrpProductions finds mrp.production records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductions(criteria *Criteria, options *Options) (*MrpProductions, error) {
	mps := &MrpProductions{}
	if err := c.SearchRead(MrpProductionModel, criteria, options, mps); err != nil {
		return nil, err
	}
	return mps, nil
}

// FindMrpProductionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMrpProductionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MrpProductionModel, criteria, options)
}

// FindMrpProductionId finds record id by querying it with criteria.
func (c *Client) FindMrpProductionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MrpProductionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
