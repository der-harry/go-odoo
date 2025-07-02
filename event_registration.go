package odoo

// EventRegistration represents event.registration model.
type EventRegistration struct {
	Active                      *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	Barcode                     *String     `xmlrpc:"barcode,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyName                 *String     `xmlrpc:"company_name,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	DateClosed                  *Time       `xmlrpc:"date_closed,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	Email                       *String     `xmlrpc:"email,omitempty"`
	EventBeginDate              *Time       `xmlrpc:"event_begin_date,omitempty"`
	EventDateRange              *String     `xmlrpc:"event_date_range,omitempty"`
	EventEndDate                *Time       `xmlrpc:"event_end_date,omitempty"`
	EventId                     *Many2One   `xmlrpc:"event_id,omitempty"`
	EventOrganizerId            *Many2One   `xmlrpc:"event_organizer_id,omitempty"`
	EventTicketId               *Many2One   `xmlrpc:"event_ticket_id,omitempty"`
	EventUserId                 *Many2One   `xmlrpc:"event_user_id,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	LeadCount                   *Int        `xmlrpc:"lead_count,omitempty"`
	LeadIds                     *Relation   `xmlrpc:"lead_ids,omitempty"`
	MailRegistrationIds         *Relation   `xmlrpc:"mail_registration_ids,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty"`
	PartnerId                   *Many2One   `xmlrpc:"partner_id,omitempty"`
	Phone                       *String     `xmlrpc:"phone,omitempty"`
	RegistrationAnswerChoiceIds *Relation   `xmlrpc:"registration_answer_choice_ids,omitempty"`
	RegistrationAnswerIds       *Relation   `xmlrpc:"registration_answer_ids,omitempty"`
	RegistrationProperties      interface{} `xmlrpc:"registration_properties,omitempty"`
	SaleOrderId                 *Many2One   `xmlrpc:"sale_order_id,omitempty"`
	SaleOrderLineId             *Many2One   `xmlrpc:"sale_order_line_id,omitempty"`
	SaleStatus                  *Selection  `xmlrpc:"sale_status,omitempty"`
	State                       *Selection  `xmlrpc:"state,omitempty"`
	UtmCampaignId               *Many2One   `xmlrpc:"utm_campaign_id,omitempty"`
	UtmMediumId                 *Many2One   `xmlrpc:"utm_medium_id,omitempty"`
	UtmSourceId                 *Many2One   `xmlrpc:"utm_source_id,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// EventRegistrations represents array of event.registration model.
type EventRegistrations []EventRegistration

// EventRegistrationModel is the odoo model name.
const EventRegistrationModel = "event.registration"

// Many2One convert EventRegistration to *Many2One.
func (er *EventRegistration) Many2One() *Many2One {
	return NewMany2One(er.Id.Get(), "")
}

// CreateEventRegistration creates a new event.registration model and returns its id.
func (c *Client) CreateEventRegistration(er *EventRegistration) (int64, error) {
	ids, err := c.CreateEventRegistrations([]*EventRegistration{er})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventRegistration creates a new event.registration model and returns its id.
func (c *Client) CreateEventRegistrations(ers []*EventRegistration) ([]int64, error) {
	var vv []interface{}
	for _, v := range ers {
		vv = append(vv, v)
	}
	return c.Create(EventRegistrationModel, vv, nil)
}

// UpdateEventRegistration updates an existing event.registration record.
func (c *Client) UpdateEventRegistration(er *EventRegistration) error {
	return c.UpdateEventRegistrations([]int64{er.Id.Get()}, er)
}

// UpdateEventRegistrations updates existing event.registration records.
// All records (represented by ids) will be updated by er values.
func (c *Client) UpdateEventRegistrations(ids []int64, er *EventRegistration) error {
	return c.Update(EventRegistrationModel, ids, er, nil)
}

// DeleteEventRegistration deletes an existing event.registration record.
func (c *Client) DeleteEventRegistration(id int64) error {
	return c.DeleteEventRegistrations([]int64{id})
}

// DeleteEventRegistrations deletes existing event.registration records.
func (c *Client) DeleteEventRegistrations(ids []int64) error {
	return c.Delete(EventRegistrationModel, ids)
}

// GetEventRegistration gets event.registration existing record.
func (c *Client) GetEventRegistration(id int64) (*EventRegistration, error) {
	ers, err := c.GetEventRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ers)[0]), nil
}

// GetEventRegistrations gets event.registration existing records.
func (c *Client) GetEventRegistrations(ids []int64) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.Read(EventRegistrationModel, ids, nil, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistration finds event.registration record by querying it with criteria.
func (c *Client) FindEventRegistration(criteria *Criteria) (*EventRegistration, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, NewOptions().Limit(1), ers); err != nil {
		return nil, err
	}
	return &((*ers)[0]), nil
}

// FindEventRegistrations finds event.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrations(criteria *Criteria, options *Options) (*EventRegistrations, error) {
	ers := &EventRegistrations{}
	if err := c.SearchRead(EventRegistrationModel, criteria, options, ers); err != nil {
		return nil, err
	}
	return ers, nil
}

// FindEventRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventRegistrationModel, criteria, options)
}

// FindEventRegistrationId finds record id by querying it with criteria.
func (c *Client) FindEventRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
