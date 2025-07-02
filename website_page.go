package odoo

// WebsitePage represents website.page model.
type WebsitePage struct {
	Active                    *Bool      `xmlrpc:"active,omitempty"`
	Arch                      *String    `xmlrpc:"arch,omitempty"`
	ArchBase                  *String    `xmlrpc:"arch_base,omitempty"`
	ArchDb                    *String    `xmlrpc:"arch_db,omitempty"`
	ArchFs                    *String    `xmlrpc:"arch_fs,omitempty"`
	ArchPrev                  *String    `xmlrpc:"arch_prev,omitempty"`
	ArchUpdated               *Bool      `xmlrpc:"arch_updated,omitempty"`
	CanPublish                *Bool      `xmlrpc:"can_publish,omitempty"`
	ControllerPageIds         *Relation  `xmlrpc:"controller_page_ids,omitempty"`
	CreateDate                *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomizeShow             *Bool      `xmlrpc:"customize_show,omitempty"`
	DatePublish               *Time      `xmlrpc:"date_publish,omitempty"`
	DisplayName               *String    `xmlrpc:"display_name,omitempty"`
	FirstPageId               *Many2One  `xmlrpc:"first_page_id,omitempty"`
	FooterVisible             *Bool      `xmlrpc:"footer_visible,omitempty"`
	GroupsId                  *Relation  `xmlrpc:"groups_id,omitempty"`
	HeaderColor               *String    `xmlrpc:"header_color,omitempty"`
	HeaderOverlay             *Bool      `xmlrpc:"header_overlay,omitempty"`
	HeaderTextColor           *String    `xmlrpc:"header_text_color,omitempty"`
	HeaderVisible             *Bool      `xmlrpc:"header_visible,omitempty"`
	Id                        *Int       `xmlrpc:"id,omitempty"`
	InheritChildrenIds        *Relation  `xmlrpc:"inherit_children_ids,omitempty"`
	InheritId                 *Many2One  `xmlrpc:"inherit_id,omitempty"`
	IsHomepage                *Bool      `xmlrpc:"is_homepage,omitempty"`
	IsInMenu                  *Bool      `xmlrpc:"is_in_menu,omitempty"`
	IsNewPageTemplate         *Bool      `xmlrpc:"is_new_page_template,omitempty"`
	IsPublished               *Bool      `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized            *Bool      `xmlrpc:"is_seo_optimized,omitempty"`
	IsVisible                 *Bool      `xmlrpc:"is_visible,omitempty"`
	Key                       *String    `xmlrpc:"key,omitempty"`
	MenuIds                   *Relation  `xmlrpc:"menu_ids,omitempty"`
	Mode                      *Selection `xmlrpc:"mode,omitempty"`
	Model                     *String    `xmlrpc:"model,omitempty"`
	ModelDataId               *Many2One  `xmlrpc:"model_data_id,omitempty"`
	ModelId                   *Many2One  `xmlrpc:"model_id,omitempty"`
	Name                      *String    `xmlrpc:"name,omitempty"`
	PageIds                   *Relation  `xmlrpc:"page_ids,omitempty"`
	Priority                  *Int       `xmlrpc:"priority,omitempty"`
	SeoName                   *String    `xmlrpc:"seo_name,omitempty"`
	ThemeTemplateId           *Many2One  `xmlrpc:"theme_template_id,omitempty"`
	Track                     *Bool      `xmlrpc:"track,omitempty"`
	Type                      *Selection `xmlrpc:"type,omitempty"`
	Url                       *String    `xmlrpc:"url,omitempty"`
	ViewId                    *Many2One  `xmlrpc:"view_id,omitempty"`
	Visibility                *Selection `xmlrpc:"visibility,omitempty"`
	VisibilityPassword        *String    `xmlrpc:"visibility_password,omitempty"`
	VisibilityPasswordDisplay *String    `xmlrpc:"visibility_password_display,omitempty"`
	WarningInfo               *String    `xmlrpc:"warning_info,omitempty"`
	WebsiteId                 *Many2One  `xmlrpc:"website_id,omitempty"`
	WebsiteIndexed            *Bool      `xmlrpc:"website_indexed,omitempty"`
	WebsiteMetaDescription    *String    `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords       *String    `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg          *String    `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle          *String    `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished          *Bool      `xmlrpc:"website_published,omitempty"`
	WebsiteUrl                *String    `xmlrpc:"website_url,omitempty"`
	WriteDate                 *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId                     *String    `xmlrpc:"xml_id,omitempty"`
}

// WebsitePages represents array of website.page model.
type WebsitePages []WebsitePage

// WebsitePageModel is the odoo model name.
const WebsitePageModel = "website.page"

// Many2One convert WebsitePage to *Many2One.
func (wp *WebsitePage) Many2One() *Many2One {
	return NewMany2One(wp.Id.Get(), "")
}

// CreateWebsitePage creates a new website.page model and returns its id.
func (c *Client) CreateWebsitePage(wp *WebsitePage) (int64, error) {
	ids, err := c.CreateWebsitePages([]*WebsitePage{wp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateWebsitePage creates a new website.page model and returns its id.
func (c *Client) CreateWebsitePages(wps []*WebsitePage) ([]int64, error) {
	var vv []interface{}
	for _, v := range wps {
		vv = append(vv, v)
	}
	return c.Create(WebsitePageModel, vv, nil)
}

// UpdateWebsitePage updates an existing website.page record.
func (c *Client) UpdateWebsitePage(wp *WebsitePage) error {
	return c.UpdateWebsitePages([]int64{wp.Id.Get()}, wp)
}

// UpdateWebsitePages updates existing website.page records.
// All records (represented by ids) will be updated by wp values.
func (c *Client) UpdateWebsitePages(ids []int64, wp *WebsitePage) error {
	return c.Update(WebsitePageModel, ids, wp, nil)
}

// DeleteWebsitePage deletes an existing website.page record.
func (c *Client) DeleteWebsitePage(id int64) error {
	return c.DeleteWebsitePages([]int64{id})
}

// DeleteWebsitePages deletes existing website.page records.
func (c *Client) DeleteWebsitePages(ids []int64) error {
	return c.Delete(WebsitePageModel, ids)
}

// GetWebsitePage gets website.page existing record.
func (c *Client) GetWebsitePage(id int64) (*WebsitePage, error) {
	wps, err := c.GetWebsitePages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*wps)[0]), nil
}

// GetWebsitePages gets website.page existing records.
func (c *Client) GetWebsitePages(ids []int64) (*WebsitePages, error) {
	wps := &WebsitePages{}
	if err := c.Read(WebsitePageModel, ids, nil, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebsitePage finds website.page record by querying it with criteria.
func (c *Client) FindWebsitePage(criteria *Criteria) (*WebsitePage, error) {
	wps := &WebsitePages{}
	if err := c.SearchRead(WebsitePageModel, criteria, NewOptions().Limit(1), wps); err != nil {
		return nil, err
	}
	return &((*wps)[0]), nil
}

// FindWebsitePages finds website.page records by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePages(criteria *Criteria, options *Options) (*WebsitePages, error) {
	wps := &WebsitePages{}
	if err := c.SearchRead(WebsitePageModel, criteria, options, wps); err != nil {
		return nil, err
	}
	return wps, nil
}

// FindWebsitePageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindWebsitePageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(WebsitePageModel, criteria, options)
}

// FindWebsitePageId finds record id by querying it with criteria.
func (c *Client) FindWebsitePageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(WebsitePageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
