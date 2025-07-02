package odoo

// EventEvent represents event.event model.
type EventEvent struct {
	Active                           *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId          *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline             *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration      *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon            *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                      *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                    *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                  *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                 *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                   *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                   *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AddressId                        *Many2One   `xmlrpc:"address_id,omitempty"`
	AddressInline                    *String     `xmlrpc:"address_inline,omitempty"`
	AddressSearch                    *Many2One   `xmlrpc:"address_search,omitempty"`
	BadgeFormat                      *Selection  `xmlrpc:"badge_format,omitempty"`
	BadgeImage                       *String     `xmlrpc:"badge_image,omitempty"`
	CompanyId                        *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryId                        *Many2One   `xmlrpc:"country_id,omitempty"`
	CreateDate                       *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                        *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                       *Many2One   `xmlrpc:"currency_id,omitempty"`
	DateBegin                        *Time       `xmlrpc:"date_begin,omitempty"`
	DateBeginLocated                 *String     `xmlrpc:"date_begin_located,omitempty"`
	DateEnd                          *Time       `xmlrpc:"date_end,omitempty"`
	DateEndLocated                   *String     `xmlrpc:"date_end_located,omitempty"`
	DateTz                           *Selection  `xmlrpc:"date_tz,omitempty"`
	Description                      *String     `xmlrpc:"description,omitempty"`
	DisplayName                      *String     `xmlrpc:"display_name,omitempty"`
	EventMailIds                     *Relation   `xmlrpc:"event_mail_ids,omitempty"`
	EventRegistrationsOpen           *Bool       `xmlrpc:"event_registrations_open,omitempty"`
	EventRegistrationsSoldOut        *Bool       `xmlrpc:"event_registrations_sold_out,omitempty"`
	EventRegistrationsStarted        *Bool       `xmlrpc:"event_registrations_started,omitempty"`
	EventTicketIds                   *Relation   `xmlrpc:"event_ticket_ids,omitempty"`
	EventTypeId                      *Many2One   `xmlrpc:"event_type_id,omitempty"`
	GeneralQuestionIds               *Relation   `xmlrpc:"general_question_ids,omitempty"`
	HasLeadRequest                   *Bool       `xmlrpc:"has_lead_request,omitempty"`
	HasMessage                       *Bool       `xmlrpc:"has_message,omitempty"`
	Id                               *Int        `xmlrpc:"id,omitempty"`
	IsFinished                       *Bool       `xmlrpc:"is_finished,omitempty"`
	IsOneDay                         *Bool       `xmlrpc:"is_one_day,omitempty"`
	IsOngoing                        *Bool       `xmlrpc:"is_ongoing,omitempty"`
	KanbanState                      *Selection  `xmlrpc:"kanban_state,omitempty"`
	KanbanStateLabel                 *String     `xmlrpc:"kanban_state_label,omitempty"`
	Lang                             *Selection  `xmlrpc:"lang,omitempty"`
	LeadCount                        *Int        `xmlrpc:"lead_count,omitempty"`
	LeadIds                          *Relation   `xmlrpc:"lead_ids,omitempty"`
	LegendBlocked                    *String     `xmlrpc:"legend_blocked,omitempty"`
	LegendDone                       *String     `xmlrpc:"legend_done,omitempty"`
	LegendNormal                     *String     `xmlrpc:"legend_normal,omitempty"`
	MessageAttachmentCount           *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds               *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                  *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter           *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError               *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                       *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter         *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MyActivityDateDeadline           *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                             *String     `xmlrpc:"name,omitempty"`
	Note                             *String     `xmlrpc:"note,omitempty"`
	OrganizerId                      *Many2One   `xmlrpc:"organizer_id,omitempty"`
	QuestionIds                      *Relation   `xmlrpc:"question_ids,omitempty"`
	RegistrationIds                  *Relation   `xmlrpc:"registration_ids,omitempty"`
	RegistrationPropertiesDefinition interface{} `xmlrpc:"registration_properties_definition,omitempty"`
	SaleOrderLinesIds                *Relation   `xmlrpc:"sale_order_lines_ids,omitempty"`
	SalePriceSubtotal                *Float      `xmlrpc:"sale_price_subtotal,omitempty"`
	SeatsAvailable                   *Int        `xmlrpc:"seats_available,omitempty"`
	SeatsLimited                     *Bool       `xmlrpc:"seats_limited,omitempty"`
	SeatsMax                         *Int        `xmlrpc:"seats_max,omitempty"`
	SeatsReserved                    *Int        `xmlrpc:"seats_reserved,omitempty"`
	SeatsTaken                       *Int        `xmlrpc:"seats_taken,omitempty"`
	SeatsUsed                        *Int        `xmlrpc:"seats_used,omitempty"`
	SpecificQuestionIds              *Relation   `xmlrpc:"specific_question_ids,omitempty"`
	StageId                          *Many2One   `xmlrpc:"stage_id,omitempty"`
	StartSaleDatetime                *Time       `xmlrpc:"start_sale_datetime,omitempty"`
	TagIds                           *Relation   `xmlrpc:"tag_ids,omitempty"`
	TicketInstructions               *String     `xmlrpc:"ticket_instructions,omitempty"`
	UseBarcode                       *Bool       `xmlrpc:"use_barcode,omitempty"`
	UserId                           *Many2One   `xmlrpc:"user_id,omitempty"`
	WebsiteMessageIds                *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                        *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                         *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// EventEvents represents array of event.event model.
type EventEvents []EventEvent

// EventEventModel is the odoo model name.
const EventEventModel = "event.event"

// Many2One convert EventEvent to *Many2One.
func (ee *EventEvent) Many2One() *Many2One {
	return NewMany2One(ee.Id.Get(), "")
}

// CreateEventEvent creates a new event.event model and returns its id.
func (c *Client) CreateEventEvent(ee *EventEvent) (int64, error) {
	ids, err := c.CreateEventEvents([]*EventEvent{ee})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateEventEvent creates a new event.event model and returns its id.
func (c *Client) CreateEventEvents(ees []*EventEvent) ([]int64, error) {
	var vv []interface{}
	for _, v := range ees {
		vv = append(vv, v)
	}
	return c.Create(EventEventModel, vv, nil)
}

// UpdateEventEvent updates an existing event.event record.
func (c *Client) UpdateEventEvent(ee *EventEvent) error {
	return c.UpdateEventEvents([]int64{ee.Id.Get()}, ee)
}

// UpdateEventEvents updates existing event.event records.
// All records (represented by ids) will be updated by ee values.
func (c *Client) UpdateEventEvents(ids []int64, ee *EventEvent) error {
	return c.Update(EventEventModel, ids, ee, nil)
}

// DeleteEventEvent deletes an existing event.event record.
func (c *Client) DeleteEventEvent(id int64) error {
	return c.DeleteEventEvents([]int64{id})
}

// DeleteEventEvents deletes existing event.event records.
func (c *Client) DeleteEventEvents(ids []int64) error {
	return c.Delete(EventEventModel, ids)
}

// GetEventEvent gets event.event existing record.
func (c *Client) GetEventEvent(id int64) (*EventEvent, error) {
	ees, err := c.GetEventEvents([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ees)[0]), nil
}

// GetEventEvents gets event.event existing records.
func (c *Client) GetEventEvents(ids []int64) (*EventEvents, error) {
	ees := &EventEvents{}
	if err := c.Read(EventEventModel, ids, nil, ees); err != nil {
		return nil, err
	}
	return ees, nil
}

// FindEventEvent finds event.event record by querying it with criteria.
func (c *Client) FindEventEvent(criteria *Criteria) (*EventEvent, error) {
	ees := &EventEvents{}
	if err := c.SearchRead(EventEventModel, criteria, NewOptions().Limit(1), ees); err != nil {
		return nil, err
	}
	return &((*ees)[0]), nil
}

// FindEventEvents finds event.event records by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEvents(criteria *Criteria, options *Options) (*EventEvents, error) {
	ees := &EventEvents{}
	if err := c.SearchRead(EventEventModel, criteria, options, ees); err != nil {
		return nil, err
	}
	return ees, nil
}

// FindEventEventIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindEventEventIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(EventEventModel, criteria, options)
}

// FindEventEventId finds record id by querying it with criteria.
func (c *Client) FindEventEventId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(EventEventModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
