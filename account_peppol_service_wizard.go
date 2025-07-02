package odoo

// AccountPeppolServiceWizard represents account_peppol.service.wizard model.
type AccountPeppolServiceWizard struct {
	CreateDate  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName *String     `xmlrpc:"display_name,omitempty"`
	EdiUserId   *Many2One   `xmlrpc:"edi_user_id,omitempty"`
	Id          *Int        `xmlrpc:"id,omitempty"`
	ServiceIds  *Relation   `xmlrpc:"service_ids,omitempty"`
	ServiceInfo *String     `xmlrpc:"service_info,omitempty"`
	ServiceJson interface{} `xmlrpc:"service_json,omitempty"`
	WriteDate   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// AccountPeppolServiceWizards represents array of account_peppol.service.wizard model.
type AccountPeppolServiceWizards []AccountPeppolServiceWizard

// AccountPeppolServiceWizardModel is the odoo model name.
const AccountPeppolServiceWizardModel = "account_peppol.service.wizard"

// Many2One convert AccountPeppolServiceWizard to *Many2One.
func (asw *AccountPeppolServiceWizard) Many2One() *Many2One {
	return NewMany2One(asw.Id.Get(), "")
}

// CreateAccountPeppolServiceWizard creates a new account_peppol.service.wizard model and returns its id.
func (c *Client) CreateAccountPeppolServiceWizard(asw *AccountPeppolServiceWizard) (int64, error) {
	ids, err := c.CreateAccountPeppolServiceWizards([]*AccountPeppolServiceWizard{asw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPeppolServiceWizard creates a new account_peppol.service.wizard model and returns its id.
func (c *Client) CreateAccountPeppolServiceWizards(asws []*AccountPeppolServiceWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range asws {
		vv = append(vv, v)
	}
	return c.Create(AccountPeppolServiceWizardModel, vv, nil)
}

// UpdateAccountPeppolServiceWizard updates an existing account_peppol.service.wizard record.
func (c *Client) UpdateAccountPeppolServiceWizard(asw *AccountPeppolServiceWizard) error {
	return c.UpdateAccountPeppolServiceWizards([]int64{asw.Id.Get()}, asw)
}

// UpdateAccountPeppolServiceWizards updates existing account_peppol.service.wizard records.
// All records (represented by ids) will be updated by asw values.
func (c *Client) UpdateAccountPeppolServiceWizards(ids []int64, asw *AccountPeppolServiceWizard) error {
	return c.Update(AccountPeppolServiceWizardModel, ids, asw, nil)
}

// DeleteAccountPeppolServiceWizard deletes an existing account_peppol.service.wizard record.
func (c *Client) DeleteAccountPeppolServiceWizard(id int64) error {
	return c.DeleteAccountPeppolServiceWizards([]int64{id})
}

// DeleteAccountPeppolServiceWizards deletes existing account_peppol.service.wizard records.
func (c *Client) DeleteAccountPeppolServiceWizards(ids []int64) error {
	return c.Delete(AccountPeppolServiceWizardModel, ids)
}

// GetAccountPeppolServiceWizard gets account_peppol.service.wizard existing record.
func (c *Client) GetAccountPeppolServiceWizard(id int64) (*AccountPeppolServiceWizard, error) {
	asws, err := c.GetAccountPeppolServiceWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*asws)[0]), nil
}

// GetAccountPeppolServiceWizards gets account_peppol.service.wizard existing records.
func (c *Client) GetAccountPeppolServiceWizards(ids []int64) (*AccountPeppolServiceWizards, error) {
	asws := &AccountPeppolServiceWizards{}
	if err := c.Read(AccountPeppolServiceWizardModel, ids, nil, asws); err != nil {
		return nil, err
	}
	return asws, nil
}

// FindAccountPeppolServiceWizard finds account_peppol.service.wizard record by querying it with criteria.
func (c *Client) FindAccountPeppolServiceWizard(criteria *Criteria) (*AccountPeppolServiceWizard, error) {
	asws := &AccountPeppolServiceWizards{}
	if err := c.SearchRead(AccountPeppolServiceWizardModel, criteria, NewOptions().Limit(1), asws); err != nil {
		return nil, err
	}
	return &((*asws)[0]), nil
}

// FindAccountPeppolServiceWizards finds account_peppol.service.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolServiceWizards(criteria *Criteria, options *Options) (*AccountPeppolServiceWizards, error) {
	asws := &AccountPeppolServiceWizards{}
	if err := c.SearchRead(AccountPeppolServiceWizardModel, criteria, options, asws); err != nil {
		return nil, err
	}
	return asws, nil
}

// FindAccountPeppolServiceWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPeppolServiceWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPeppolServiceWizardModel, criteria, options)
}

// FindAccountPeppolServiceWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountPeppolServiceWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPeppolServiceWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
