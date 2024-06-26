package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type ItemUom struct {
	bun.BaseModel    `bun:"table:item_uom"`
	UnitOfMeasure    string          `bun:"unit_of_measure,type:varchar(8)"`
	DeleteFlag       string          `bun:"delete_flag,type:char"`
	DateCreated      time.Time       `bun:"date_created,type:datetime"`
	DateLastModified time.Time       `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy string          `bun:"last_maintained_by,type:varchar(30),default:(user_name(null))"`
	UnitSize         float64         `bun:"unit_size,type:decimal(19,4)"`
	SellingUnit      sql.NullString  `bun:"selling_unit,type:char,nullzero"`
	PurchasingUnit   sql.NullString  `bun:"purchasing_unit,type:char,nullzero"`
	InvMastUid       int32           `bun:"inv_mast_uid,type:int"`
	CreatedBy        sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	ItemUomUid       int32           `bun:"item_uom_uid,type:int,pk"`
	B2bUnitFlag      string          `bun:"b2b_unit_flag,type:char,default:('Y')"`
	TallyFactor      sql.NullFloat64 `bun:"tally_factor,type:decimal(19,9),nullzero"`
	WwmsFlag         sql.NullString  `bun:"wwms_flag,type:char,default:('N')"`
	ProdOrderFactor  sql.NullInt32   `bun:"prod_order_factor,type:int,nullzero"`
	MinimumOrderQty  sql.NullFloat64 `bun:"minimum_order_qty,type:decimal(19,4),nullzero"`
}

type ItemUomModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns an slice of ItemUom by the given InvMastUid
func (m ItemUomModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*ItemUom, error) {
	var itemUoms []*ItemUom
	err := m.bun.NewSelect().Model(&itemUoms).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return itemUoms, nil
}

// Delete deletes the ItemUom from the database.
func (m ItemUomModel) Delete(ctx context.Context, itemUom *ItemUom) error {
	_, err := m.bun.NewDelete().Model(itemUom).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
