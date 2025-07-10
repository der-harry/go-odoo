package odoo

// DynamicFields represents dynamic.fields model.
type DynamicFields struct {
	AddFieldInTree      *Bool      `xmlrpc:"add_field_in_tree,omitempty"`
	Column1             *String    `xmlrpc:"column1,omitempty"`
	Column2             *String    `xmlrpc:"column2,omitempty"`
	CompanyDependent    *Bool      `xmlrpc:"company_dependent,omitempty"`
	CompleteName        *String    `xmlrpc:"complete_name,omitempty"`
	Compute             *String    `xmlrpc:"compute,omitempty"`
	Copied              *Bool      `xmlrpc:"copied,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CreatedFormViewId   *Many2One  `xmlrpc:"created_form_view_id,omitempty"`
	CreatedTreeViewId   *Many2One  `xmlrpc:"created_tree_view_id,omitempty"`
	CurrencyField       *String    `xmlrpc:"currency_field,omitempty"`
	Depends             *String    `xmlrpc:"depends,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Domain              *String    `xmlrpc:"domain,omitempty"`
	ExtraFeatures       *Bool      `xmlrpc:"extra_features,omitempty"`
	FieldDescription    *String    `xmlrpc:"field_description,omitempty"`
	FieldType           *Selection `xmlrpc:"field_type,omitempty"`
	FormViewId          *Many2One  `xmlrpc:"form_view_id,omitempty"`
	FormViewIds         *Relation  `xmlrpc:"form_view_ids,omitempty"`
	FormViewInheritId   *String    `xmlrpc:"form_view_inherit_id,omitempty"`
	GroupExpand         *Bool      `xmlrpc:"group_expand,omitempty"`
	Groups              *Relation  `xmlrpc:"groups,omitempty"`
	Help                *String    `xmlrpc:"help,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	Index               *Bool      `xmlrpc:"index,omitempty"`
	IsDynamicField      *Bool      `xmlrpc:"is_dynamic_field,omitempty"`
	IsVisibleInTreeView *Bool      `xmlrpc:"is_visible_in_tree_view,omitempty"`
	Model               *String    `xmlrpc:"model,omitempty"`
	ModelId             *Many2One  `xmlrpc:"model_id,omitempty"`
	Modules             *String    `xmlrpc:"modules,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	OnDelete            *Selection `xmlrpc:"on_delete,omitempty"`
	Position            *Selection `xmlrpc:"position,omitempty"`
	PositionFieldId     *Many2One  `xmlrpc:"position_field_id,omitempty"`
	Readonly            *Bool      `xmlrpc:"readonly,omitempty"`
	RefModelId          *Many2One  `xmlrpc:"ref_model_id,omitempty"`
	Related             *String    `xmlrpc:"related,omitempty"`
	RelatedFieldId      *Many2One  `xmlrpc:"related_field_id,omitempty"`
	Relation            *String    `xmlrpc:"relation,omitempty"`
	RelationField       *String    `xmlrpc:"relation_field,omitempty"`
	RelationFieldId     *Many2One  `xmlrpc:"relation_field_id,omitempty"`
	RelationTable       *String    `xmlrpc:"relation_table,omitempty"`
	Required            *Bool      `xmlrpc:"required,omitempty"`
	Sanitize            *Bool      `xmlrpc:"sanitize,omitempty"`
	SanitizeAttributes  *Bool      `xmlrpc:"sanitize_attributes,omitempty"`
	SanitizeForm        *Bool      `xmlrpc:"sanitize_form,omitempty"`
	SanitizeOverridable *Bool      `xmlrpc:"sanitize_overridable,omitempty"`
	SanitizeStyle       *Bool      `xmlrpc:"sanitize_style,omitempty"`
	SanitizeTags        *Bool      `xmlrpc:"sanitize_tags,omitempty"`
	Selectable          *Bool      `xmlrpc:"selectable,omitempty"`
	Selection           *String    `xmlrpc:"selection,omitempty"`
	SelectionField      *String    `xmlrpc:"selection_field,omitempty"`
	SelectionIds        *Relation  `xmlrpc:"selection_ids,omitempty"`
	Size                *Int       `xmlrpc:"size,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	Status              *Selection `xmlrpc:"status,omitempty"`
	Store               *Bool      `xmlrpc:"store,omitempty"`
	StripClasses        *Bool      `xmlrpc:"strip_classes,omitempty"`
	StripStyle          *Bool      `xmlrpc:"strip_style,omitempty"`
	Tracking            *Int       `xmlrpc:"tracking,omitempty"`
	Translate           *Bool      `xmlrpc:"translate,omitempty"`
	TreeFieldId         *Many2One  `xmlrpc:"tree_field_id,omitempty"`
	TreeFieldIds        *Relation  `xmlrpc:"tree_field_ids,omitempty"`
	TreeFieldPosition   *Selection `xmlrpc:"tree_field_position,omitempty"`
	TreeViewId          *Many2One  `xmlrpc:"tree_view_id,omitempty"`
	TreeViewIds         *Relation  `xmlrpc:"tree_view_ids,omitempty"`
	TreeViewInheritId   *String    `xmlrpc:"tree_view_inherit_id,omitempty"`
	Ttype               *Selection `xmlrpc:"ttype,omitempty"`
	WidgetId            *Many2One  `xmlrpc:"widget_id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// DynamicFieldss represents array of dynamic.fields model.
type DynamicFieldss []DynamicFields

// DynamicFieldsModel is the odoo model name.
const DynamicFieldsModel = "dynamic.fields"

// Many2One convert DynamicFields to *Many2One.
func (df *DynamicFields) Many2One() *Many2One {
	return NewMany2One(df.Id.Get(), "")
}

// CreateDynamicFields creates a new dynamic.fields model and returns its id.
func (c *Client) CreateDynamicFields(df *DynamicFields) (int64, error) {
	ids, err := c.CreateDynamicFieldss([]*DynamicFields{df})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDynamicFields creates a new dynamic.fields model and returns its id.
func (c *Client) CreateDynamicFieldss(dfs []*DynamicFields) ([]int64, error) {
	var vv []interface{}
	for _, v := range dfs {
		vv = append(vv, v)
	}
	return c.Create(DynamicFieldsModel, vv, nil)
}

// UpdateDynamicFields updates an existing dynamic.fields record.
func (c *Client) UpdateDynamicFields(df *DynamicFields) error {
	return c.UpdateDynamicFieldss([]int64{df.Id.Get()}, df)
}

// UpdateDynamicFieldss updates existing dynamic.fields records.
// All records (represented by ids) will be updated by df values.
func (c *Client) UpdateDynamicFieldss(ids []int64, df *DynamicFields) error {
	return c.Update(DynamicFieldsModel, ids, df, nil)
}

// DeleteDynamicFields deletes an existing dynamic.fields record.
func (c *Client) DeleteDynamicFields(id int64) error {
	return c.DeleteDynamicFieldss([]int64{id})
}

// DeleteDynamicFieldss deletes existing dynamic.fields records.
func (c *Client) DeleteDynamicFieldss(ids []int64) error {
	return c.Delete(DynamicFieldsModel, ids)
}

// GetDynamicFields gets dynamic.fields existing record.
func (c *Client) GetDynamicFields(id int64) (*DynamicFields, error) {
	dfs, err := c.GetDynamicFieldss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// GetDynamicFieldss gets dynamic.fields existing records.
func (c *Client) GetDynamicFieldss(ids []int64) (*DynamicFieldss, error) {
	dfs := &DynamicFieldss{}
	if err := c.Read(DynamicFieldsModel, ids, nil, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDynamicFields finds dynamic.fields record by querying it with criteria.
func (c *Client) FindDynamicFields(criteria *Criteria) (*DynamicFields, error) {
	dfs := &DynamicFieldss{}
	if err := c.SearchRead(DynamicFieldsModel, criteria, NewOptions().Limit(1), dfs); err != nil {
		return nil, err
	}
	return &((*dfs)[0]), nil
}

// FindDynamicFieldss finds dynamic.fields records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDynamicFieldss(criteria *Criteria, options *Options) (*DynamicFieldss, error) {
	dfs := &DynamicFieldss{}
	if err := c.SearchRead(DynamicFieldsModel, criteria, options, dfs); err != nil {
		return nil, err
	}
	return dfs, nil
}

// FindDynamicFieldsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDynamicFieldsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DynamicFieldsModel, criteria, options)
}

// FindDynamicFieldsId finds record id by querying it with criteria.
func (c *Client) FindDynamicFieldsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DynamicFieldsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
