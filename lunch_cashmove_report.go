package odoo

// LunchCashmoveReport represents lunch.cashmove.report model.
type LunchCashmoveReport struct {
	Amount      *Float    `xmlrpc:"amount,omitempty"`
	CurrencyId  *Many2One `xmlrpc:"currency_id,omitempty"`
	Date        *Time     `xmlrpc:"date,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
}

// LunchCashmoveReports represents array of lunch.cashmove.report model.
type LunchCashmoveReports []LunchCashmoveReport

// LunchCashmoveReportModel is the odoo model name.
const LunchCashmoveReportModel = "lunch.cashmove.report"

// Many2One convert LunchCashmoveReport to *Many2One.
func (lcr *LunchCashmoveReport) Many2One() *Many2One {
	return NewMany2One(lcr.Id.Get(), "")
}

// CreateLunchCashmoveReport creates a new lunch.cashmove.report model and returns its id.
func (c *Client) CreateLunchCashmoveReport(lcr *LunchCashmoveReport) (int64, error) {
	ids, err := c.CreateLunchCashmoveReports([]*LunchCashmoveReport{lcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchCashmoveReport creates a new lunch.cashmove.report model and returns its id.
func (c *Client) CreateLunchCashmoveReports(lcrs []*LunchCashmoveReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range lcrs {
		vv = append(vv, v)
	}
	return c.Create(LunchCashmoveReportModel, vv, nil)
}

// UpdateLunchCashmoveReport updates an existing lunch.cashmove.report record.
func (c *Client) UpdateLunchCashmoveReport(lcr *LunchCashmoveReport) error {
	return c.UpdateLunchCashmoveReports([]int64{lcr.Id.Get()}, lcr)
}

// UpdateLunchCashmoveReports updates existing lunch.cashmove.report records.
// All records (represented by ids) will be updated by lcr values.
func (c *Client) UpdateLunchCashmoveReports(ids []int64, lcr *LunchCashmoveReport) error {
	return c.Update(LunchCashmoveReportModel, ids, lcr, nil)
}

// DeleteLunchCashmoveReport deletes an existing lunch.cashmove.report record.
func (c *Client) DeleteLunchCashmoveReport(id int64) error {
	return c.DeleteLunchCashmoveReports([]int64{id})
}

// DeleteLunchCashmoveReports deletes existing lunch.cashmove.report records.
func (c *Client) DeleteLunchCashmoveReports(ids []int64) error {
	return c.Delete(LunchCashmoveReportModel, ids)
}

// GetLunchCashmoveReport gets lunch.cashmove.report existing record.
func (c *Client) GetLunchCashmoveReport(id int64) (*LunchCashmoveReport, error) {
	lcrs, err := c.GetLunchCashmoveReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lcrs)[0]), nil
}

// GetLunchCashmoveReports gets lunch.cashmove.report existing records.
func (c *Client) GetLunchCashmoveReports(ids []int64) (*LunchCashmoveReports, error) {
	lcrs := &LunchCashmoveReports{}
	if err := c.Read(LunchCashmoveReportModel, ids, nil, lcrs); err != nil {
		return nil, err
	}
	return lcrs, nil
}

// FindLunchCashmoveReport finds lunch.cashmove.report record by querying it with criteria.
func (c *Client) FindLunchCashmoveReport(criteria *Criteria) (*LunchCashmoveReport, error) {
	lcrs := &LunchCashmoveReports{}
	if err := c.SearchRead(LunchCashmoveReportModel, criteria, NewOptions().Limit(1), lcrs); err != nil {
		return nil, err
	}
	return &((*lcrs)[0]), nil
}

// FindLunchCashmoveReports finds lunch.cashmove.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoveReports(criteria *Criteria, options *Options) (*LunchCashmoveReports, error) {
	lcrs := &LunchCashmoveReports{}
	if err := c.SearchRead(LunchCashmoveReportModel, criteria, options, lcrs); err != nil {
		return nil, err
	}
	return lcrs, nil
}

// FindLunchCashmoveReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchCashmoveReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchCashmoveReportModel, criteria, options)
}

// FindLunchCashmoveReportId finds record id by querying it with criteria.
func (c *Client) FindLunchCashmoveReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchCashmoveReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
