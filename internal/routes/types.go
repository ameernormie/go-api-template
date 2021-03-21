package routes

import (
	"database/sql"
)

// Company : Company
type Company struct {
	KEDCD              string          `json:"kedCode"`
	NAME               string          `json:"name"`
	CORP_NO            string          `json:"corporateNumber"`
	BZNO               string          `json:"businessNumber"`
	ENP_SZE            string          `json:"businessSize"`
	IPO_CD_NM          string          `json:"IPOStatus"`
	ESTB_DT            string          `json:"foundedDate"`
	TEL_NO             string          `json:"tel"`
	ADDR               string          `json:"address"`
	PD_NM              string          `json:"mainProducts"`
	BASE_DT_1          string          `json:"baseDate1"`
	BASE_DT_2          string          `json:"baseDate2"`
	BASE_DT_3          string          `json:"baseDate3"`
	SALES_1            sql.NullInt64   `json:"sales1"`
	SALES_2            sql.NullInt64   `json:"sales2"`
	SALES_3            sql.NullInt64   `json:"sales3"`
	SALES_LABEL_1      string          `json:"salesLabel1"`
	SALES_LABEL_2      string          `json:"salesLabel2"`
	SALES_LABEL_3      string          `json:"salesLabel3"`
	OPR_INCOME_1       sql.NullInt64   `json:"operatingIncome1"`
	OPR_INCOME_2       sql.NullInt64   `json:"operatingIncome2"`
	OPR_INCOME_3       sql.NullInt64   `json:"operatingIncome3"`
	OPR_INCOME_LABEL_1 string          `json:"operatingIncomeLabel1"`
	OPR_INCOME_LABEL_2 string          `json:"operatingIncomeLabel2"`
	OPR_INCOME_LABEL_3 string          `json:"operatingIncomeLabel3"`
	EW                 string          `json:"warningGrade"`
	CR_GRD             string          `json:"creditGrade"`
	LAT                sql.NullFloat64 `json:"lat"`
	LNG                sql.NullFloat64 `json:"lng"`
	EM_CN              sql.NullInt64   `json:"numEmployee"`
}

// Supplier : Supplier
type Supplier struct {
	ID            string          `json:"userID"`
	BASE_YY       string          `json:"baseYear"`
	TXPL_BZNO     string          `json:"businessNumber"`
	TXPL_KEDCD    string          `json:"kedCode"`
	ENP_NM        string          `json:"name"`
	TXPL_TCN      sql.NullInt64   `json:"transactionCount"`
	TXPL_TSPPR    sql.NullInt64   `json:"transactionAmount"`
	TXPL_T_TAXAM  sql.NullInt64   `json:"transactionTax"`
	REG_CCD       string          `json:"registerType"`
	KSIC10_BZC_CD string          `json:"industryCode"`
	BZC_CD_LCD    string          `json:"industryDescription"`
	TOT_AM        sql.NullInt64   `json:"totalAmount"`
	ADDR          string          `json:"address"`
	LAT           sql.NullFloat64 `json:"lat"`
	LOT           sql.NullFloat64 `json:"lng"`
	RATIO         sql.NullFloat64 `json:"ratio"`
	EM_CN         sql.NullInt64   `json:"numEmployee"`
	IPO_CD_NM     string          `json:"IPOStatus"`
	ESTB_DT       string          `json:"foundedDate"`
	EW            string          `json:"warningGrade"`
}

// AlternativeSupplier: AlternativeSupplier
type AlternativeSupplier struct {
	NO         sql.NullInt64 `json:"seq"`
	KEDCD      string        `json:"kedCode"`
	ENP_NM     string        `json:"name"`
	IPO_CD_NM  string        `json:"IPOStatus"`
	ENP_SCD_NM string        `json:"status"`
	ADDR       string        `json:"address"`
	PD_NM      string        `json:"mainProducts"`
	SALES_SZE  string        `json:"salesSize"`
	INDUSTRY   string        `json:"industry"`
	EM_CN      sql.NullInt64 `json:"numEmployee"`
	TEL_NO     string        `json:"tel"`
}

// AddtionalSupplier: AddtionalSupplier
type AddtionalSupplier struct {
	BZNO      string        `json:"businessNumber"`
	ENP_NM    string        `json:"name"`
	INDUSTRY  string        `json:"industry"`
	ESTB_DT   string        `json:"foundedDate"`
	EM_CN     sql.NullInt64 `json:"numEmployee"`
	ENP_SZE   string        `json:"businessSize"`
	IPO_CD_NM string        `json:"IPOStatus"`
	KEDCD     string        `json:"kedCode"`
}

// CompanySearchResult: CompanySearchResult
type CompanySearchResult struct {
	NO       sql.NullInt64 `json:"seq"`
	BZNO     string        `json:"businessNumber"`
	KEDCD    string        `json:"kedCode"`
	ENP_NM   string        `json:"name"`
	REPER_NM string        `json:"CEO"`
	INDUSTRY string        `json:"industry"`
}

// CityTown: CityTown
type CityTown struct {
	BZC_SGRP_CD string          `json:"industryCode"`
	SIDO_NM     string          `json:"stateName"`
	SIGUNGU_NM  string          `json:"cityName"`
	SIGUNGU_CD  string          `json:"cityCode"`
	LCDONG_NM   string          `json:"townName"`
	LCDONG_CD   string          `json:"townCode"`
	EMPLMT_PO   sql.NullFloat64 `json:"employmentScore"`
	STAB_PO     sql.NullFloat64 `json:"stabilityScore"`
	GRW_PO      sql.NullFloat64 `json:"growthScore"`
	TECH_PO     sql.NullFloat64 `json:"techScore"`
	TX_PO       sql.NullFloat64 `json:"tranactionScore"`
	T_PO        sql.NullFloat64 `json:"totalScore"`
	COMMENT1    string          `json:"comment1"`
	COMMENT2    string          `json:"comment2"`
}

// Grid: Grid
type Grid struct {
	GRID            string `json:"GRID"`
	EM_PO           string `json:"EM_PO"`
	ENTCP_PO        string `json:"ENTCP_PO"`
	RTIR_PO         string `json:"RTIR_PO"`
	YWAGE_AVG_PO    string `json:"YWAGE_AVG_PO"`
	EM_INCR_RT_PO   string `json:"EM_INCR_RT_PO"`
	CR_GRD_PO       string `json:"CR_GRD_PO"`
	EW_GRD_PO       string `json:"EW_GRD_PO"`
	BZHIS_PO        string `json:"BZHIS_PO"`
	LQDT_CN_PO      string `json:"LQDT_CN_PO"`
	ENP_CN_PO       string `json:"ENP_CN_PO"`
	SAM_PO          string `json:"SAM_PO"`
	BZPF_PO         string `json:"BZPF_PO"`
	NE_PO           string `json:"NE_PO"`
	GRS_RT_PO       string `json:"GRS_RT_PO"`
	BZPF_INCR_RT_PO string `json:"BZPF_INCR_RT_PO"`
	PTNT_PO         string `json:"PTNT_PO"`
	CERT_PO         string `json:"CERT_PO"`
	RSRCHI_PO       string `json:"RSRCHI_PO"`
	UM_PO           string `json:"UM_PO"`
	RND_XPN_PO      string `json:"RND_XPN_PO"`
	MATPRTS_EQIP_PO string `json:"MATPRTS_EQIP_PO"`
	SALEPL_CN_PO    string `json:"SALEPL_CN_PO"`
	PUCHPL_CN_PO    string `json:"PUCHPL_CN_PO"`
	SALEPL_EW_PO    string `json:"SALEPL_EW_PO"`
	PUCHPL_EW_PO    string `json:"PUCHPL_EW_PO"`
	EMPLMT_PO       string `json:"EMPLMT_PO"`
	STAB_PO         string `json:"STAB_PO"`
	GRW_PO          string `json:"GRW_PO"`
	TECH_PO         string `json:"TECH_PO"`
	TX_PO           string `json:"TX_PO"`
	T_PO            string `json:"T_PO"`
	COMMENT1        string `json:"comment1"`
	COMMENT2        string `json:"comment2"`
}
