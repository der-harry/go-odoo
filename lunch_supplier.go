package odoo

// LunchSupplier represents lunch.supplier model.
type LunchSupplier struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
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
	AutomaticEmailTime          *Float     `xmlrpc:"automatic_email_time,omitempty"`
	AvailableLocationIds        *Relation  `xmlrpc:"available_location_ids,omitempty"`
	AvailableToday              *Bool      `xmlrpc:"available_today,omitempty"`
	City                        *String    `xmlrpc:"city,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CronId                      *Many2One  `xmlrpc:"cron_id,omitempty"`
	Delivery                    *Selection `xmlrpc:"delivery,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	Email                       *String    `xmlrpc:"email,omitempty"`
	EmailFormatted              *String    `xmlrpc:"email_formatted,omitempty"`
	Fri                         *Bool      `xmlrpc:"fri,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
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
	Moment                      *Selection `xmlrpc:"moment,omitempty"`
	Mon                         *Bool      `xmlrpc:"mon,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	OrderDeadlinePassed         *Bool      `xmlrpc:"order_deadline_passed,omitempty"`
	PartnerId                   *Many2One  `xmlrpc:"partner_id,omitempty"`
	Phone                       *String    `xmlrpc:"phone,omitempty"`
	RecurrencyEndDate           *Time      `xmlrpc:"recurrency_end_date,omitempty"`
	ResponsibleId               *Many2One  `xmlrpc:"responsible_id,omitempty"`
	Sat                         *Bool      `xmlrpc:"sat,omitempty"`
	SendBy                      *Selection `xmlrpc:"send_by,omitempty"`
	ShowConfirmButton           *Bool      `xmlrpc:"show_confirm_button,omitempty"`
	ShowOrderButton             *Bool      `xmlrpc:"show_order_button,omitempty"`
	StateId                     *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                      *String    `xmlrpc:"street,omitempty"`
	Street2                     *String    `xmlrpc:"street2,omitempty"`
	Sun                         *Bool      `xmlrpc:"sun,omitempty"`
	Thu                         *Bool      `xmlrpc:"thu,omitempty"`
	ToppingIds1                 *Relation  `xmlrpc:"topping_ids_1,omitempty"`
	ToppingIds2                 *Relation  `xmlrpc:"topping_ids_2,omitempty"`
	ToppingIds3                 *Relation  `xmlrpc:"topping_ids_3,omitempty"`
	ToppingLabel1               *String    `xmlrpc:"topping_label_1,omitempty"`
	ToppingLabel2               *String    `xmlrpc:"topping_label_2,omitempty"`
	ToppingLabel3               *String    `xmlrpc:"topping_label_3,omitempty"`
	ToppingQuantity1            *Selection `xmlrpc:"topping_quantity_1,omitempty"`
	ToppingQuantity2            *Selection `xmlrpc:"topping_quantity_2,omitempty"`
	ToppingQuantity3            *Selection `xmlrpc:"topping_quantity_3,omitempty"`
	Tue                         *Bool      `xmlrpc:"tue,omitempty"`
	Tz                          *Selection `xmlrpc:"tz,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	Wed                         *Bool      `xmlrpc:"wed,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
	ZipCode                     *String    `xmlrpc:"zip_code,omitempty"`
}

// LunchSuppliers represents array of lunch.supplier model.
type LunchSuppliers []LunchSupplier

// LunchSupplierModel is the odoo model name.
const LunchSupplierModel = "lunch.supplier"

// Many2One convert LunchSupplier to *Many2One.
func (ls *LunchSupplier) Many2One() *Many2One {
	return NewMany2One(ls.Id.Get(), "")
}

// CreateLunchSupplier creates a new lunch.supplier model and returns its id.
func (c *Client) CreateLunchSupplier(ls *LunchSupplier) (int64, error) {
	ids, err := c.CreateLunchSuppliers([]*LunchSupplier{ls})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchSupplier creates a new lunch.supplier model and returns its id.
func (c *Client) CreateLunchSuppliers(lss []*LunchSupplier) ([]int64, error) {
	var vv []interface{}
	for _, v := range lss {
		vv = append(vv, v)
	}
	return c.Create(LunchSupplierModel, vv, nil)
}

// UpdateLunchSupplier updates an existing lunch.supplier record.
func (c *Client) UpdateLunchSupplier(ls *LunchSupplier) error {
	return c.UpdateLunchSuppliers([]int64{ls.Id.Get()}, ls)
}

// UpdateLunchSuppliers updates existing lunch.supplier records.
// All records (represented by ids) will be updated by ls values.
func (c *Client) UpdateLunchSuppliers(ids []int64, ls *LunchSupplier) error {
	return c.Update(LunchSupplierModel, ids, ls, nil)
}

// DeleteLunchSupplier deletes an existing lunch.supplier record.
func (c *Client) DeleteLunchSupplier(id int64) error {
	return c.DeleteLunchSuppliers([]int64{id})
}

// DeleteLunchSuppliers deletes existing lunch.supplier records.
func (c *Client) DeleteLunchSuppliers(ids []int64) error {
	return c.Delete(LunchSupplierModel, ids)
}

// GetLunchSupplier gets lunch.supplier existing record.
func (c *Client) GetLunchSupplier(id int64) (*LunchSupplier, error) {
	lss, err := c.GetLunchSuppliers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lss)[0]), nil
}

// GetLunchSuppliers gets lunch.supplier existing records.
func (c *Client) GetLunchSuppliers(ids []int64) (*LunchSuppliers, error) {
	lss := &LunchSuppliers{}
	if err := c.Read(LunchSupplierModel, ids, nil, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindLunchSupplier finds lunch.supplier record by querying it with criteria.
func (c *Client) FindLunchSupplier(criteria *Criteria) (*LunchSupplier, error) {
	lss := &LunchSuppliers{}
	if err := c.SearchRead(LunchSupplierModel, criteria, NewOptions().Limit(1), lss); err != nil {
		return nil, err
	}
	return &((*lss)[0]), nil
}

// FindLunchSuppliers finds lunch.supplier records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchSuppliers(criteria *Criteria, options *Options) (*LunchSuppliers, error) {
	lss := &LunchSuppliers{}
	if err := c.SearchRead(LunchSupplierModel, criteria, options, lss); err != nil {
		return nil, err
	}
	return lss, nil
}

// FindLunchSupplierIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchSupplierIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchSupplierModel, criteria, options)
}

// FindLunchSupplierId finds record id by querying it with criteria.
func (c *Client) FindLunchSupplierId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchSupplierModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
