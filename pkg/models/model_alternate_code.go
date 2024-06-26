package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type AlternateCode struct {
	bun.BaseModel      `bun:"table:alternate_code"`
	AlternateCode      string         `bun:"alternate_code,type:varchar(40),pk"`
	DeleteFlag         string         `bun:"delete_flag,type:char"`
	DateCreated        time.Time      `bun:"date_created,type:datetime"`
	DateLastModified   time.Time      `bun:"date_last_modified,type:datetime"`
	LastMaintainedBy   string         `bun:"last_maintained_by,type:varchar(30),default:(user_name())"`
	InvMastUid         int32          `bun:"inv_mast_uid,type:int,pk"`
	AlternateCodeDesc  sql.NullString `bun:"alternate_code_desc,type:varchar(40),nullzero"`
	AlternateCodeUid   int32          `bun:"alternate_code_uid,type:int"`
	CreatedBy          string         `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	SourceTypeCd       int32          `bun:"source_type_cd,type:int,default:((2051))"`
	DefaultUom         sql.NullString `bun:"default_uom,type:varchar(8),nullzero"`
	ExcludeFromB2bFlag sql.NullString `bun:"exclude_from_b2b_flag,type:char,nullzero"`
}

type AlternateCodeModel struct {
	bun bun.IDB
}

func (m *AlternateCodeModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*AlternateCode, error) {
	var alternateCodes []*AlternateCode
	err := m.bun.NewSelect().Model(&alternateCodes).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return alternateCodes, err
}

// Delete deletes the alternate code
func (m *AlternateCodeModel) Delete(ctx context.Context, alternateCode *AlternateCode) error {
	_, err := m.bun.NewDelete().Model(alternateCode).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
