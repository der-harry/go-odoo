package odoo

// EventSaleReport represents event.sale.report model.
type EventSaleReport struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	EventDateBegin              *Time      `xmlrpc:"event_date_begin,omitempty"`
	EventDateEnd                *Time      `xmlrpc:"event_date_end,omitempty"`
	EventId                     *Many2One  `xmlrpc:"event_id,omitempty"`
	EventRegistrationCreateDate *Time      `xmlrpc:"event_registration_create_date,omitempty"`
	EventRegistrationId         *Many2One  `xmlrpc:"event_registration_id,omitempty"`
	EventRegistrationName       *String    `xmlrpc:"event_registration_name,omitempty"`
	EventRegistrationState      *Selection `xmlrpc:"event_registration_state,omitempty"`
	EventTicketId               *Many2One  `xmlrpc:"event_ticket_id,omitempty"`
	EventTicketPrice            *Float     `xmlrpc:"event_ticket_price,omitempty"`
	EventTypeId                 *Many2One  `xmlrpc:"event_type_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InvoicePartnerId            *Many2One  `xmlrpc:"invoice_partner_id,omitempty"`
	ProductId                   *Many2One  `xmlrpc:"product_id,omitempty"`
	SaleOrderDate               *Time      `xmlrpc:"sale_order_date,omitempty"`
	SaleOrderId                 *Many2One  `xmlrpc:"sale_order_id,omitempty"`
	SaleOrderLineId             *Many2One  `xmlrpc:"sale_order_line_id,omitempty"`
	SaleOrderPartnerId          *Many2One  `xmlrpc:"sale_order_partner_id,omitempty"`
	SaleOrderState              *Selection `xmlrpc:"sale_order_state,omitempty"`
	SaleOrderUserId             *Many2One  `xmlrpc:"sale_order_user_id,omitempty"`
	SalePrice                   *Float     `xmlrpc:"sale_price,omitempty"`
	SalePriceUntaxed            *Float     `xmlrpc:"sale_price_untaxed,omitempty"`
	SaleStatus                  *Selection `xmlrpc:"sale_status,omitempty"`
}

// EventSaleReports represents array of event.sale.report model.
type EventSaleReports []EventSaleReport

// EventSaleReportModel is the odoo model name.
const EventSaleReportModel = "event.sale.report"

// Many2One convert EventSaleReport to *Many2One.
func (esr *EventSaleReport) Many2One() *Many2One {
	return NewMany2One(esr.Id.Get(), "")
}

// CreateEventSaleReport creates a new event.sale.report model and returns its id.
func (c *Client) CreateEventSaleReport(esr *EventSaleReport) (int64, error) {
	ids, err := c.CreateEventSaleReports([]*EventSaleReport{esr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventSaleReport creates a new event.sale.report model and returns its id.
func (c *Client) CreateEventSaleReports(esrs []*EventSaleReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range esrs {
		vv = append(vv, v)
	}
	return c.Create(EventSaleReportModel, vv, nil)
}

// UpdateEventSaleReport updates an existing event.sale.report record.
func (c *Client) UpdateEventSaleReport(esr *EventSaleReport) error {
	return c.UpdateEventSaleReports([]int64{esr.Id.Get()}, esr)
}

// UpdateEventSaleReports updates existing event.sale.report records.
// All records (represented by ids) will be updated by esr values.
func (c *Client) UpdateEventSaleReports(ids []int64, esr *EventSaleReport) error {
	return c.Update(EventSaleReportModel, ids, esr, nil)
}

// DeleteEventSaleReport deletes an existing event.sale.report record.
func (c *Client) DeleteEventSaleReport(id int64) error {
	return c.DeleteEventSaleReports([]int64{id})
}

// DeleteEventSaleReports deletes existing event.sale.report records.
func (c *Client) DeleteEventSaleReports(ids []int64) error {
	return c.Delete(EventSaleReportModel, ids)
}

// GetEventSaleReport gets event.sale.report existing record.
func (c *Client) GetEventSaleReport(id int64) (*EventSaleReport, error) {
	esrs, err := c.GetEventSaleReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*esrs)[0]), nil
}

// GetEventSaleReports gets event.sale.report existing records.
func (c *Client) GetEventSaleReports(ids []int64) (*EventSaleReports, error) {
	esrs := &EventSaleReports{}
	if err := c.Read(EventSaleReportModel, ids, nil, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindEventSaleReport finds event.sale.report record by querying it with criteria.
func (c *Client) FindEventSaleReport(criteria *Criteria) (*EventSaleReport, error) {
	esrs := &EventSaleReports{}
	if err := c.SearchRead(EventSaleReportModel, criteria, NewOptions().Limit(1), esrs); err != nil {
		return nil, err
	}
	return &((*esrs)[0]), nil
}

// FindEventSaleReports finds event.sale.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventSaleReports(criteria *Criteria, options *Options) (*EventSaleReports, error) {
	esrs := &EventSaleReports{}
	if err := c.SearchRead(EventSaleReportModel, criteria, options, esrs); err != nil {
		return nil, err
	}
	return esrs, nil
}

// FindEventSaleReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventSaleReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventSaleReportModel, criteria, options)
}

// FindEventSaleReportId finds record id by querying it with criteria.
func (c *Client) FindEventSaleReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventSaleReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
