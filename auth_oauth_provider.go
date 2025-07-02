package odoo

// AuthOauthProvider represents auth.oauth.provider model.
type AuthOauthProvider struct {
	AuthEndpoint       *String   `xmlrpc:"auth_endpoint,omitempty"`
	Body               *String   `xmlrpc:"body,omitempty"`
	ClientId           *String   `xmlrpc:"client_id,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	CssClass           *String   `xmlrpc:"css_class,omitempty"`
	DataEndpoint       *String   `xmlrpc:"data_endpoint,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Enabled            *Bool     `xmlrpc:"enabled,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	Name               *String   `xmlrpc:"name,omitempty"`
	Scope              *String   `xmlrpc:"scope,omitempty"`
	Sequence           *Int      `xmlrpc:"sequence,omitempty"`
	ValidationEndpoint *String   `xmlrpc:"validation_endpoint,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AuthOauthProviders represents array of auth.oauth.provider model.
type AuthOauthProviders []AuthOauthProvider

// AuthOauthProviderModel is the odoo model name.
const AuthOauthProviderModel = "auth.oauth.provider"

// Many2One convert AuthOauthProvider to *Many2One.
func (aop *AuthOauthProvider) Many2One() *Many2One {
	return NewMany2One(aop.Id.Get(), "")
}

// CreateAuthOauthProvider creates a new auth.oauth.provider model and returns its id.
func (c *Client) CreateAuthOauthProvider(aop *AuthOauthProvider) (int64, error) {
	ids, err := c.CreateAuthOauthProviders([]*AuthOauthProvider{aop})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthOauthProvider creates a new auth.oauth.provider model and returns its id.
func (c *Client) CreateAuthOauthProviders(aops []*AuthOauthProvider) ([]int64, error) {
	var vv []interface{}
	for _, v := range aops {
		vv = append(vv, v)
	}
	return c.Create(AuthOauthProviderModel, vv, nil)
}

// UpdateAuthOauthProvider updates an existing auth.oauth.provider record.
func (c *Client) UpdateAuthOauthProvider(aop *AuthOauthProvider) error {
	return c.UpdateAuthOauthProviders([]int64{aop.Id.Get()}, aop)
}

// UpdateAuthOauthProviders updates existing auth.oauth.provider records.
// All records (represented by ids) will be updated by aop values.
func (c *Client) UpdateAuthOauthProviders(ids []int64, aop *AuthOauthProvider) error {
	return c.Update(AuthOauthProviderModel, ids, aop, nil)
}

// DeleteAuthOauthProvider deletes an existing auth.oauth.provider record.
func (c *Client) DeleteAuthOauthProvider(id int64) error {
	return c.DeleteAuthOauthProviders([]int64{id})
}

// DeleteAuthOauthProviders deletes existing auth.oauth.provider records.
func (c *Client) DeleteAuthOauthProviders(ids []int64) error {
	return c.Delete(AuthOauthProviderModel, ids)
}

// GetAuthOauthProvider gets auth.oauth.provider existing record.
func (c *Client) GetAuthOauthProvider(id int64) (*AuthOauthProvider, error) {
	aops, err := c.GetAuthOauthProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aops)[0]), nil
}

// GetAuthOauthProviders gets auth.oauth.provider existing records.
func (c *Client) GetAuthOauthProviders(ids []int64) (*AuthOauthProviders, error) {
	aops := &AuthOauthProviders{}
	if err := c.Read(AuthOauthProviderModel, ids, nil, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAuthOauthProvider finds auth.oauth.provider record by querying it with criteria.
func (c *Client) FindAuthOauthProvider(criteria *Criteria) (*AuthOauthProvider, error) {
	aops := &AuthOauthProviders{}
	if err := c.SearchRead(AuthOauthProviderModel, criteria, NewOptions().Limit(1), aops); err != nil {
		return nil, err
	}
	return &((*aops)[0]), nil
}

// FindAuthOauthProviders finds auth.oauth.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthOauthProviders(criteria *Criteria, options *Options) (*AuthOauthProviders, error) {
	aops := &AuthOauthProviders{}
	if err := c.SearchRead(AuthOauthProviderModel, criteria, options, aops); err != nil {
		return nil, err
	}
	return aops, nil
}

// FindAuthOauthProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthOauthProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AuthOauthProviderModel, criteria, options)
}

// FindAuthOauthProviderId finds record id by querying it with criteria.
func (c *Client) FindAuthOauthProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthOauthProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
