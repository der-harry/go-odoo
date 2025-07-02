package odoo

// LunchAlert represents lunch.alert model.
type LunchAlert struct {
	Active             *Bool      `xmlrpc:"active,omitempty"`
	AvailableToday     *Bool      `xmlrpc:"available_today,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	CronId             *Many2One  `xmlrpc:"cron_id,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	Fri                *Bool      `xmlrpc:"fri,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	LocationIds        *Relation  `xmlrpc:"location_ids,omitempty"`
	Message            *String    `xmlrpc:"message,omitempty"`
	Mode               *Selection `xmlrpc:"mode,omitempty"`
	Mon                *Bool      `xmlrpc:"mon,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	NotificationMoment *Selection `xmlrpc:"notification_moment,omitempty"`
	NotificationTime   *Float     `xmlrpc:"notification_time,omitempty"`
	Recipients         *Selection `xmlrpc:"recipients,omitempty"`
	Sat                *Bool      `xmlrpc:"sat,omitempty"`
	Sun                *Bool      `xmlrpc:"sun,omitempty"`
	Thu                *Bool      `xmlrpc:"thu,omitempty"`
	Tue                *Bool      `xmlrpc:"tue,omitempty"`
	Tz                 *Selection `xmlrpc:"tz,omitempty"`
	Until              *Time      `xmlrpc:"until,omitempty"`
	Wed                *Bool      `xmlrpc:"wed,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// LunchAlerts represents array of lunch.alert model.
type LunchAlerts []LunchAlert

// LunchAlertModel is the odoo model name.
const LunchAlertModel = "lunch.alert"

// Many2One convert LunchAlert to *Many2One.
func (la *LunchAlert) Many2One() *Many2One {
	return NewMany2One(la.Id.Get(), "")
}

// CreateLunchAlert creates a new lunch.alert model and returns its id.
func (c *Client) CreateLunchAlert(la *LunchAlert) (int64, error) {
	ids, err := c.CreateLunchAlerts([]*LunchAlert{la})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateLunchAlert creates a new lunch.alert model and returns its id.
func (c *Client) CreateLunchAlerts(las []*LunchAlert) ([]int64, error) {
	var vv []interface{}
	for _, v := range las {
		vv = append(vv, v)
	}
	return c.Create(LunchAlertModel, vv, nil)
}

// UpdateLunchAlert updates an existing lunch.alert record.
func (c *Client) UpdateLunchAlert(la *LunchAlert) error {
	return c.UpdateLunchAlerts([]int64{la.Id.Get()}, la)
}

// UpdateLunchAlerts updates existing lunch.alert records.
// All records (represented by ids) will be updated by la values.
func (c *Client) UpdateLunchAlerts(ids []int64, la *LunchAlert) error {
	return c.Update(LunchAlertModel, ids, la, nil)
}

// DeleteLunchAlert deletes an existing lunch.alert record.
func (c *Client) DeleteLunchAlert(id int64) error {
	return c.DeleteLunchAlerts([]int64{id})
}

// DeleteLunchAlerts deletes existing lunch.alert records.
func (c *Client) DeleteLunchAlerts(ids []int64) error {
	return c.Delete(LunchAlertModel, ids)
}

// GetLunchAlert gets lunch.alert existing record.
func (c *Client) GetLunchAlert(id int64) (*LunchAlert, error) {
	las, err := c.GetLunchAlerts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*las)[0]), nil
}

// GetLunchAlerts gets lunch.alert existing records.
func (c *Client) GetLunchAlerts(ids []int64) (*LunchAlerts, error) {
	las := &LunchAlerts{}
	if err := c.Read(LunchAlertModel, ids, nil, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindLunchAlert finds lunch.alert record by querying it with criteria.
func (c *Client) FindLunchAlert(criteria *Criteria) (*LunchAlert, error) {
	las := &LunchAlerts{}
	if err := c.SearchRead(LunchAlertModel, criteria, NewOptions().Limit(1), las); err != nil {
		return nil, err
	}
	return &((*las)[0]), nil
}

// FindLunchAlerts finds lunch.alert records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchAlerts(criteria *Criteria, options *Options) (*LunchAlerts, error) {
	las := &LunchAlerts{}
	if err := c.SearchRead(LunchAlertModel, criteria, options, las); err != nil {
		return nil, err
	}
	return las, nil
}

// FindLunchAlertIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLunchAlertIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(LunchAlertModel, criteria, options)
}

// FindLunchAlertId finds record id by querying it with criteria.
func (c *Client) FindLunchAlertId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LunchAlertModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
