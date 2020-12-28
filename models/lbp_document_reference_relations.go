package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type LbpDocumentReferenceRelations struct {
	Id                   int       `orm:"column(id);auto"`
	DocumentId           int       `orm:"column(document_id)" description:"引用文档Id"`
	ReferenceDocumentId  int       `orm:"column(reference_document_id)" description:"被引用文档Id,文档为全文引用时，document_paragraph_id为0"`
	ReferenceParagraphId int       `orm:"column(reference_paragraph_id)" description:"被引用文档段落Id"`
	Content              string    `orm:"column(content);size(1024)" description:"单行控件内容"`
	ReferenceInternalId  string    `orm:"column(reference_internal_id);size(255)" description:"控件唯一ID"`
	ReferenceType        int       `orm:"column(reference_type)" description:"控件类型 多行、单行"`
	SyncType             int       `orm:"column(sync_type)" description:"是否随文档同步变化"`
	CreatedAt            time.Time `orm:"column(created_at);type(timestamp);null"`
	UpdatedAt            time.Time `orm:"column(updated_at);type(timestamp);null"`
	DeletedAt            time.Time `orm:"column(deleted_at);type(timestamp);null"`
}

func (t *LbpDocumentReferenceRelations) TableName() string {
	return "lbp_document_reference_relations"
}

func init() {
	orm.RegisterModel(new(LbpDocumentReferenceRelations))
}

// AddLbpDocumentReferenceRelations insert a new LbpDocumentReferenceRelations into database and returns
// last inserted Id on success.
func AddLbpDocumentReferenceRelations(m *LbpDocumentReferenceRelations) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLbpDocumentReferenceRelationsById retrieves LbpDocumentReferenceRelations by Id. Returns error if
// Id doesn't exist
func GetLbpDocumentReferenceRelationsById(id int) (v *LbpDocumentReferenceRelations, err error) {
	o := orm.NewOrm()
	v = &LbpDocumentReferenceRelations{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLbpDocumentReferenceRelations retrieves all LbpDocumentReferenceRelations matches certain condition. Returns empty list if
// no records exist
func GetAllLbpDocumentReferenceRelations(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LbpDocumentReferenceRelations))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []LbpDocumentReferenceRelations
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateLbpDocumentReferenceRelations updates LbpDocumentReferenceRelations by Id and returns error if
// the record to be updated doesn't exist
func UpdateLbpDocumentReferenceRelationsById(m *LbpDocumentReferenceRelations) (err error) {
	o := orm.NewOrm()
	v := LbpDocumentReferenceRelations{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLbpDocumentReferenceRelations deletes LbpDocumentReferenceRelations by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLbpDocumentReferenceRelations(id int) (err error) {
	o := orm.NewOrm()
	v := LbpDocumentReferenceRelations{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LbpDocumentReferenceRelations{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
