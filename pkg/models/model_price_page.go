package models

import (
	"context"
	"database/sql"
	"time"

	"github.com/uptrace/bun"
)

type PricePage struct {
	bun.BaseModel              `bun:"table:price_page"`
	PricePageUid               int32           `bun:"price_page_uid,type:int,pk"`
	PricePageTypeCd            int16           `bun:"price_page_type_cd,type:smallint"`
	MfgClassId                 sql.NullString  `bun:"mfg_class_id,type:varchar(255),nullzero"`
	SupplierId                 sql.NullFloat64 `bun:"supplier_id,type:decimal(19,0),nullzero"`
	ProductGroupId             sql.NullString  `bun:"product_group_id,type:varchar(8),nullzero"`
	DiscountGroupId            sql.NullString  `bun:"discount_group_id,type:varchar(8),nullzero"`
	MajorGroupId               sql.NullString  `bun:"major_group_id,type:varchar(8),nullzero"`
	Description                sql.NullString  `bun:"description,type:varchar(255),nullzero"`
	PricingMethodCd            int16           `bun:"pricing_method_cd,type:smallint"`
	SourcePriceCd              sql.NullInt16   `bun:"source_price_cd,type:smallint,nullzero"`
	Price                      sql.NullFloat64 `bun:"price,type:decimal(19,9),nullzero"`
	EffectiveDate              time.Time       `bun:"effective_date,type:datetime"`
	ExpirationDate             time.Time       `bun:"expiration_date,type:datetime"`
	ContractNumber             sql.NullString  `bun:"contract_number,type:varchar(40),nullzero"`
	CalculationMethodCd        int16           `bun:"calculation_method_cd,type:smallint"`
	CalculationValue1          float64         `bun:"calculation_value1,type:decimal(19,6)"`
	CalculationValue2          float64         `bun:"calculation_value2,type:decimal(19,6)"`
	CalculationValue3          float64         `bun:"calculation_value3,type:decimal(19,6)"`
	CalculationValue4          float64         `bun:"calculation_value4,type:decimal(19,6)"`
	CalculationValue5          float64         `bun:"calculation_value5,type:decimal(19,6)"`
	CalculationValue6          float64         `bun:"calculation_value6,type:decimal(19,6)"`
	CalculationValue7          float64         `bun:"calculation_value7,type:decimal(19,6)"`
	CalculationValue8          float64         `bun:"calculation_value8,type:decimal(19,6)"`
	CalculationValue9          float64         `bun:"calculation_value9,type:decimal(19,6)"`
	CalculationValue10         float64         `bun:"calculation_value10,type:decimal(19,6)"`
	CalculationValue11         float64         `bun:"calculation_value11,type:decimal(19,6)"`
	CalculationValue12         float64         `bun:"calculation_value12,type:decimal(19,6)"`
	CalculationValue13         float64         `bun:"calculation_value13,type:decimal(19,6)"`
	CalculationValue14         float64         `bun:"calculation_value14,type:decimal(19,6)"`
	CalculationValue15         float64         `bun:"calculation_value15,type:decimal(19,6)"`
	Break1                     float64         `bun:"break1,type:decimal(19,9)"`
	Break2                     float64         `bun:"break2,type:decimal(19,9)"`
	Break3                     float64         `bun:"break3,type:decimal(19,9)"`
	Break4                     float64         `bun:"break4,type:decimal(19,9)"`
	Break5                     float64         `bun:"break5,type:decimal(19,9)"`
	Break6                     float64         `bun:"break6,type:decimal(19,9)"`
	Break7                     float64         `bun:"break7,type:decimal(19,9)"`
	Break8                     float64         `bun:"break8,type:decimal(19,9)"`
	Break9                     float64         `bun:"break9,type:decimal(19,9)"`
	Break10                    float64         `bun:"break10,type:decimal(19,9)"`
	Break11                    float64         `bun:"break11,type:decimal(19,9)"`
	Break12                    float64         `bun:"break12,type:decimal(19,9)"`
	Break13                    float64         `bun:"break13,type:decimal(19,9)"`
	Break14                    float64         `bun:"break14,type:decimal(19,9)"`
	OtherCost1                 sql.NullFloat64 `bun:"other_cost1,type:decimal(19,9),nullzero"`
	OtherCost2                 sql.NullFloat64 `bun:"other_cost2,type:decimal(19,9),nullzero"`
	OtherCost3                 sql.NullFloat64 `bun:"other_cost3,type:decimal(19,9),nullzero"`
	OtherCost4                 sql.NullFloat64 `bun:"other_cost4,type:decimal(19,9),nullzero"`
	OtherCost5                 sql.NullFloat64 `bun:"other_cost5,type:decimal(19,9),nullzero"`
	OtherCost6                 sql.NullFloat64 `bun:"other_cost6,type:decimal(19,9),nullzero"`
	OtherCost7                 sql.NullFloat64 `bun:"other_cost7,type:decimal(19,9),nullzero"`
	OtherCost8                 sql.NullFloat64 `bun:"other_cost8,type:decimal(19,9),nullzero"`
	OtherCost9                 sql.NullFloat64 `bun:"other_cost9,type:decimal(19,9),nullzero"`
	OtherCost10                sql.NullFloat64 `bun:"other_cost10,type:decimal(19,9),nullzero"`
	OtherCost11                sql.NullFloat64 `bun:"other_cost11,type:decimal(19,9),nullzero"`
	OtherCost12                sql.NullFloat64 `bun:"other_cost12,type:decimal(19,9),nullzero"`
	OtherCost13                sql.NullFloat64 `bun:"other_cost13,type:decimal(19,9),nullzero"`
	OtherCost14                sql.NullFloat64 `bun:"other_cost14,type:decimal(19,9),nullzero"`
	OtherCost15                sql.NullFloat64 `bun:"other_cost15,type:decimal(19,9),nullzero"`
	Uom1                       sql.NullString  `bun:"uom1,type:varchar(8),nullzero"`
	Uom2                       sql.NullString  `bun:"uom2,type:varchar(8),nullzero"`
	Uom3                       sql.NullString  `bun:"uom3,type:varchar(8),nullzero"`
	Uom4                       sql.NullString  `bun:"uom4,type:varchar(8),nullzero"`
	Uom5                       sql.NullString  `bun:"uom5,type:varchar(8),nullzero"`
	Uom6                       sql.NullString  `bun:"uom6,type:varchar(8),nullzero"`
	Uom7                       sql.NullString  `bun:"uom7,type:varchar(8),nullzero"`
	Uom8                       sql.NullString  `bun:"uom8,type:varchar(8),nullzero"`
	Uom9                       sql.NullString  `bun:"uom9,type:varchar(8),nullzero"`
	Uom10                      sql.NullString  `bun:"uom10,type:varchar(8),nullzero"`
	Uom11                      sql.NullString  `bun:"uom11,type:varchar(8),nullzero"`
	Uom12                      sql.NullString  `bun:"uom12,type:varchar(8),nullzero"`
	Uom13                      sql.NullString  `bun:"uom13,type:varchar(8),nullzero"`
	Uom14                      sql.NullString  `bun:"uom14,type:varchar(8),nullzero"`
	TotalingMethodCd           int16           `bun:"totaling_method_cd,type:smallint"`
	TotalingBasisCd            int16           `bun:"totaling_basis_cd,type:smallint"`
	RowStatusFlag              int16           `bun:"row_status_flag,type:smallint"`
	DateLastModified           time.Time       `bun:"date_last_modified,type:datetime"`
	DateCreated                time.Time       `bun:"date_created,type:datetime"`
	LastMaintainedBy           string          `bun:"last_maintained_by,type:varchar(30)"`
	OtherCostTypeCd            sql.NullInt16   `bun:"other_cost_type_cd,type:smallint,nullzero"`
	OtherCostValue             sql.NullFloat64 `bun:"other_cost_value,type:decimal(19,9),nullzero"`
	OtherCostSourceCd          sql.NullInt16   `bun:"other_cost_source_cd,type:smallint,nullzero"`
	CostCalculationMethodCd    sql.NullInt16   `bun:"cost_calculation_method_cd,type:smallint,nullzero"`
	CostCalculationValue       sql.NullFloat64 `bun:"cost_calculation_value,type:decimal(19,4),nullzero"`
	CommissionCostTypeCd       sql.NullInt16   `bun:"commission_cost_type_cd,type:smallint,nullzero"`
	CommissionCostValue        sql.NullFloat64 `bun:"commission_cost_value,type:decimal(19,9),nullzero"`
	CommissionCostSourceCd     sql.NullInt16   `bun:"commission_cost_source_cd,type:smallint,nullzero"`
	CommissionCostCalcMethodCd sql.NullInt16   `bun:"commission_cost_calc_method_cd,type:smallint,nullzero"`
	CommissionCostCalcValue    sql.NullFloat64 `bun:"commission_cost_calc_value,type:decimal(19,4),nullzero"`
	PricePageId                sql.NullString  `bun:"price_page_id,type:varchar(20),nullzero"`
	CustomerId                 sql.NullFloat64 `bun:"customer_id,type:decimal(19,0),nullzero"`
	CustomerPartNo             sql.NullString  `bun:"customer_part_no,type:varchar(40),nullzero"`
	CompanyId                  sql.NullString  `bun:"company_id,type:varchar(8),nullzero"`
	InvMastUid                 sql.NullInt32   `bun:"inv_mast_uid,type:int,nullzero"`
	CreatedBy                  sql.NullString  `bun:"created_by,type:varchar(255),default:(suser_sname())"`
	CommCostCalcValue1         sql.NullFloat64 `bun:"comm_cost_calc_value1,type:decimal(19,9),default:(0)"`
	CommCostCalcValue2         sql.NullFloat64 `bun:"comm_cost_calc_value2,type:decimal(19,9),default:(0)"`
	CommCostCalcValue3         sql.NullFloat64 `bun:"comm_cost_calc_value3,type:decimal(19,9),default:(0)"`
	CommCostCalcValue4         sql.NullFloat64 `bun:"comm_cost_calc_value4,type:decimal(19,9),default:(0)"`
	CommCostCalcValue5         sql.NullFloat64 `bun:"comm_cost_calc_value5,type:decimal(19,9),default:(0)"`
	CommCostCalcValue6         sql.NullFloat64 `bun:"comm_cost_calc_value6,type:decimal(19,9),default:(0)"`
	CommCostCalcValue7         sql.NullFloat64 `bun:"comm_cost_calc_value7,type:decimal(19,9),default:(0)"`
	CommCostCalcValue8         sql.NullFloat64 `bun:"comm_cost_calc_value8,type:decimal(19,9),default:(0)"`
	CommCostCalcValue9         sql.NullFloat64 `bun:"comm_cost_calc_value9,type:decimal(19,9),default:(0)"`
	CommCostCalcValue10        sql.NullFloat64 `bun:"comm_cost_calc_value10,type:decimal(19,9),default:(0)"`
	CommCostCalcValue11        sql.NullFloat64 `bun:"comm_cost_calc_value11,type:decimal(19,9),default:(0)"`
	CommCostCalcValue12        sql.NullFloat64 `bun:"comm_cost_calc_value12,type:decimal(19,9),default:(0)"`
	CommCostCalcValue13        sql.NullFloat64 `bun:"comm_cost_calc_value13,type:decimal(19,9),default:(0)"`
	CommCostCalcValue14        sql.NullFloat64 `bun:"comm_cost_calc_value14,type:decimal(19,9),default:(0)"`
	CommCostCalcValue15        sql.NullFloat64 `bun:"comm_cost_calc_value15,type:decimal(19,9),default:(0)"`
	CalculatorType             string          `bun:"calculator_type,type:char,default:('B')"`
	CurrencyId                 sql.NullFloat64 `bun:"currency_id,type:decimal(19,0),nullzero"`
	ValuesCurrencyId           sql.NullFloat64 `bun:"values_currency_id,type:decimal(19,0),nullzero"`
	ApplyPpToMroCd             int16           `bun:"apply_pp_to_mro_cd,type:smallint,default:((1104))"`
	DateLastSentOn832          sql.NullTime    `bun:"date_last_sent_on_832,type:datetime,nullzero"`
	DatePageDeleted            sql.NullTime    `bun:"date_page_deleted,type:datetime,nullzero"`
	PriceFamilyUid             sql.NullInt32   `bun:"price_family_uid,type:int,nullzero"`
	PeerPricePageUid           sql.NullInt32   `bun:"peer_price_page_uid,type:int,nullzero"`
	ContractLineUid            sql.NullInt32   `bun:"contract_line_uid,type:int,nullzero"`
	OnContractFlag             string          `bun:"on_contract_flag,type:char,default:('N')"`
	StrategicPriceAppliesToCd  int32           `bun:"strategic_price_applies_to_cd,type:int,default:((2636))"`
	RoundToNextDollar          sql.NullInt16   `bun:"round_to_next_dollar,type:tinyint,nullzero"`
	ApplyFreightFactor         string          `bun:"apply_freight_factor,type:char,default:('N')"`
	FreightFactorSourceCd      sql.NullInt32   `bun:"freight_factor_source_cd,type:int,nullzero"`
	NoChargeFlag               sql.NullString  `bun:"no_charge_flag,type:char,nullzero"`
	NonStockItemsOnlyFlag      string          `bun:"non_stock_items_only_flag,type:char,default:('N')"`
	ApplyPpToSopCd             int16           `bun:"apply_pp_to_sop_cd,type:smallint,default:((1103))"`
	PriceOverride              string          `bun:"price_override,type:char,default:('N')"`
	RebateMargin               sql.NullFloat64 `bun:"rebate_margin,type:decimal(19,9),nullzero"`
	UpperMarginVariance        sql.NullFloat64 `bun:"upper_margin_variance,type:decimal(19,9),nullzero"`
	LowerMarginVariance        sql.NullFloat64 `bun:"lower_margin_variance,type:decimal(19,9),nullzero"`
	ExcludeOrderLevelDiscFlag  sql.NullString  `bun:"exclude_order_level_disc_flag,type:char,nullzero"`
	RolledItemPricingTypeCd    sql.NullInt32   `bun:"rolled_item_pricing_type_cd,type:int,nullzero"`
}

type PricePageModel struct {
	bun bun.IDB
}

// GetByInvMastUid returns a slice of PricePage by the given InvMastUid
func (m PricePageModel) GetByInvMastUid(ctx context.Context, invMastUid int32) ([]*PricePage, error) {
	var pricePages []*PricePage
	err := m.bun.NewSelect().Model(&pricePages).Where("inv_mast_uid = ?", invMastUid).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return pricePages, nil
}

// Delete removes the record from the database.
func (m PricePageModel) Delete(ctx context.Context, pricePage *PricePage) error {
	_, err := m.bun.NewDelete().Model(pricePage).WherePK().Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
