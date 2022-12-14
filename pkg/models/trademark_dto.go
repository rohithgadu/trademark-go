package models

import (
	"github.com/jinzhu/gorm"
	"go-demo/pkg/config"
)

//	type FinalStruct struct {
//		Root Root `json:"root"`
//	}
var db *gorm.DB

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Root{})
}

func (record Root) CreateRecord() {
	db.NewRecord(record)
	db.Create(&record)
}

func GetRecordById(Id int64) (*Root, *gorm.DB) {
	var getRecord Root
	db := db.Where("ID=?", Id).Find(&getRecord)
	return &getRecord, db
}

type Root struct {
	gorm.Model
	Cacom              string             `json:"cacom"`
	Catmk              string             `json:"catmk"`
	COM                string             `json:"com"`
	IPOVersion         string             `json:"ipo_version"`
	NS3                string             `json:"ns_3"`
	NS4                string             `json:"ns_4"`
	RequestExamination RequestExamination `json:"req	uest_examination"`
	RequestSearch      RequestExamination `json:"request_search"`
	St96_Version       string             `json:"st_96_version"`
	Tmk                string             `json:"tmk"`
	TrademarkBag       RequestExamination `json:"trademark_bag"`
	XMLName            XMLName            `json:"xml_name"`
	Xs                 string             `json:"xs"`
	Xsi                string             `json:"xsi"`
}

type ViennaClassificationElement struct {
	CreatedAt            string             `json:"CreatedAt"`
	DeletedAt            DeletedAt          `json:"DeletedAt"`
	ID                   string             `json:"ID"`
	UpdatedAt            string             `json:"UpdatedAt"`
	Type                 string             `json:"type"`
	ViennaCategory       string             `json:"vienna_category"`
	ViennaDescriptionBag RequestExamination `json:"vienna_description_bag"`
	ViennaDivision       string             `json:"vienna_division"`
	ViennaSection        string             `json:"vienna_section"`
}

type ViennaClassification struct {
	Element []ViennaClassificationElement `json:"element"`
}

type UseRightElement struct {
	CreatedAt         string             `json:"CreatedAt"`
	DeletedAt         DeletedAt          `json:"DeletedAt"`
	ID                string             `json:"ID"`
	UpdatedAt         string             `json:"UpdatedAt"`
	UseRightIndicator string             `json:"use_right_indicator"`
	UseRightText      RequestExamination `json:"use_right_text"`
}

type UseRight struct {
	Element []UseRightElement `json:"element"`
}

type TrademarkClass struct {
	TrademarkClassCode           string             `json:"trademark_class_code"`
	TrademarkClassDescriptionBag RequestExamination `json:"trademark_class_description_bag"`
}

type Legislation struct {
	CreatedAt                 string             `json:"CreatedAt"`
	DeletedAt                 DeletedAt          `json:"DeletedAt"`
	ID                        string             `json:"ID"`
	UpdatedAt                 string             `json:"UpdatedAt"`
	LegislationCode           string             `json:"legislation_code"`
	LegislationDescriptionBag RequestExamination `json:"legislation_description_bag"`
}

type NationalTrademarkInformation struct {
	CreatedAt                                string             `json:"CreatedAt"`
	DeletedAt                                DeletedAt          `json:"DeletedAt"`
	ID                                       string             `json:"ID"`
	UpdatedAt                                string             `json:"UpdatedAt"`
	AllowedDate                              string             `json:"allowed_date"`
	ClaimBag                                 RequestExamination `json:"claim_bag"`
	InterestedPartyBag                       RequestExamination `json:"interested_party_bag"`
	Legislation                              Legislation        `json:"legislation"`
	MarkCurrentStatusInternalDescriptionText string             `json:"mark_current_status_internal_description_text"`
	MarkFeatureCategoryBag                   RequestExamination `json:"mark_feature_category_bag"`
	RegisterCategory                         string             `json:"register_category"`
	RenewalDate                              string             `json:"renewal_date"`
	TrademarkClass                           TrademarkClass     `json:"trademark_class"`
	Type                                     string             `json:"type"`
}

type NationalRepresentative struct {
	CreatedAt   string    `json:"CreatedAt"`
	DeletedAt   DeletedAt `json:"DeletedAt"`
	ID          string    `json:"ID"`
	UpdatedAt   string    `json:"UpdatedAt"`
	CommentText string    `json:"comment_text"`
	Contact     Contact   `json:"contact"`
}

type MarkRepresentation struct {
	CreatedAt           string             `json:"CreatedAt"`
	DeletedAt           DeletedAt          `json:"DeletedAt"`
	ID                  string             `json:"ID"`
	UpdatedAt           string             `json:"UpdatedAt"`
	MarkFeatureCategory string             `json:"mark_feature_category"`
	MarkReproduction    RequestExamination `json:"mark_reproduction"`
}

type Trademark struct {
	CreatedAt                     string                       `json:"CreatedAt"`
	DeletedAt                     DeletedAt                    `json:"DeletedAt"`
	ID                            string                       `json:"ID"`
	UpdatedAt                     string                       `json:"UpdatedAt"`
	ApplicantBag                  RequestExamination           `json:"applicant_bag"`
	ApplicationDate               string                       `json:"application_date"`
	ApplicationLanguageCode       LanguageCode                 `json:"application_language_code"`
	ApplicationNumber             ApplicationNumber            `json:"application_number"`
	ExpiryDate                    string                       `json:"expiry_date"`
	FilingPlace                   string                       `json:"filing_place"`
	GoodsServicesBag              RequestExamination           `json:"goods_services_bag"`
	MarkCategory                  string                       `json:"mark_category"`
	MarkCurrentStatusCode         string                       `json:"mark_current_status_code"`
	MarkCurrentStatusDate         string                       `json:"mark_current_status_date"`
	MarkEventBag                  RequestExamination           `json:"mark_event_bag"`
	MarkFeatureDescription        string                       `json:"mark_feature_description"`
	MarkRepresentation            MarkRepresentation           `json:"mark_representation"`
	NationalRepresentative        NationalRepresentative       `json:"national_representative"`
	NationalTrademarkInformation  NationalTrademarkInformation `json:"national_trademark_information"`
	NonUseCancelledIndicator      string                       `json:"non_use_cancelled_indicator"`
	OperationCategory             string                       `json:"operation_category"`
	OppositionPeriodEndDate       string                       `json:"opposition_period_end_date"`
	OppositionPeriodStartDate     string                       `json:"opposition_period_start_date"`
	PriorityBag                   RequestExamination           `json:"priority_bag"`
	PublicationBag                RequestExamination           `json:"publication_bag"`
	ReceivingOfficeCode           string                       `json:"receiving_office_code"`
	ReceivingOfficeDate           string                       `json:"receiving_office_date"`
	RegistrationDate              string                       `json:"registration_date"`
	RegistrationNumber            string                       `json:"registration_number"`
	RegistrationOfficeCode        string                       `json:"registration_office_code"`
	TradeDistinctivenessIndicator string                       `json:"trade_distinctiveness_indicator"`
	UseRight                      UseRight                     `json:"use_right"`
}

type PriorityElement struct {
	CreatedAt                     string             `json:"CreatedAt"`
	DeletedAt                     DeletedAt          `json:"DeletedAt"`
	ID                            string             `json:"ID"`
	UpdatedAt                     string             `json:"UpdatedAt"`
	ApplicationNumber             RequestExamination `json:"application_number"`
	PriorityApplicationFilingDate string             `json:"priority_application_filing_date"`
	PriorityCountryCode           string             `json:"priority_country_code"`
	PriorityPartialGoodsServices  RequestExamination `json:"priority_partial_goods_services"`
}

type Priority struct {
	Element []PriorityElement `json:"element"`
}

type PostalStructuredAddress struct {
	CreatedAt            string                 `json:"CreatedAt"`
	DeletedAt            DeletedAt              `json:"DeletedAt"`
	ID                   string                 `json:"ID"`
	UpdatedAt            string                 `json:"UpdatedAt"`
	AddressLineText      LegislationDescription `json:"address_line_text"`
	CountryCode          string                 `json:"country_code"`
	LanguageCode         LanguageCode           `json:"language_code"`
	GeographicRegionName *RequestExamination    `json:"geographic_region_name,omitempty"`
	PostalCode           *string                `json:"postal_code,omitempty"`
}

type MarkImage struct {
	CreatedAt               string             `json:"CreatedAt"`
	DeletedAt               DeletedAt          `json:"DeletedAt"`
	ID                      string             `json:"ID"`
	UpdatedAt               string             `json:"UpdatedAt"`
	FileName                string             `json:"file_name"`
	ImageFormatCategory     string             `json:"image_format_category"`
	MarkImageClassification RequestExamination `json:"mark_image_classification"`
}

type MarkEventOtherLanguageDescriptionTextBag struct {
	MarkEventDescriptionText RequestExamination `json:"mark_event_description_text"`
}

type NationalMarkEvent struct {
	CreatedAt                                string                                   `json:"CreatedAt"`
	DeletedAt                                DeletedAt                                `json:"DeletedAt"`
	ID                                       string                                   `json:"ID"`
	UpdatedAt                                string                                   `json:"UpdatedAt"`
	MarkEventCode                            string                                   `json:"mark_event_code"`
	MarkEventDescriptionText                 string                                   `json:"mark_event_description_text"`
	MarkEventOtherLanguageDescriptionTextBag MarkEventOtherLanguageDescriptionTextBag `json:"mark_event_other_language_description_text_bag"`
	Type                                     Type                                     `json:"type"`
	MarkEventAdditionalText                  *string                                  `json:"mark_event_additional_text,omitempty"`
}

type MarkEventElement struct {
	CreatedAt             string            `json:"CreatedAt"`
	DeletedAt             DeletedAt         `json:"DeletedAt"`
	ID                    string            `json:"ID"`
	UpdatedAt             string            `json:"UpdatedAt"`
	MarkEventCategory     string            `json:"mark_event_category"`
	MarkEventDate         string            `json:"mark_event_date"`
	NationalMarkEvent     NationalMarkEvent `json:"national_mark_event"`
	MarkEventResponseDate *string           `json:"mark_event_response_date,omitempty"`
}

type MarkEvent struct {
	Element []MarkEventElement `json:"element"`
}

type LegislationDescription struct {
	Element []RequestExamination `json:"element"`
}

type InterestedPartyElement struct {
	CreatedAt               string    `json:"CreatedAt"`
	DeletedAt               DeletedAt `json:"DeletedAt"`
	ID                      string    `json:"ID"`
	UpdatedAt               string    `json:"UpdatedAt"`
	Contact                 Contact   `json:"contact"`
	InterestedPartyCategory string    `json:"interested_party_category"`
}

type InterestedParty struct {
	Element []InterestedPartyElement `json:"element"`
}

type GoodsServicesReference struct {
	Element *ElementUnion `json:"element"`
}

type GoodsServices struct {
	CreatedAt                      string             `json:"CreatedAt"`
	DeletedAt                      DeletedAt          `json:"DeletedAt"`
	ID                             string             `json:"ID"`
	UpdatedAt                      string             `json:"UpdatedAt"`
	ClassDescriptionBag            RequestExamination `json:"class_description_bag"`
	ClassificationKindCode         string             `json:"classification_kind_code"`
	GoodsServicesClassificationBag RequestExamination `json:"goods_services_classification_bag"`
	NationalFilingBasis            RequestExamination `json:"national_filing_basis"`
	NationalGoodsServices          RequestExamination `json:"national_goods_services"`
}

type ClassDescription struct {
	Element                      []RequestExamination `json:"element,omitempty"`
	CreatedAt                    *string              `json:"CreatedAt,omitempty"`
	DeletedAt                    *DeletedAt           `json:"DeletedAt,omitempty"`
	ID                           *string              `json:"ID,omitempty"`
	UpdatedAt                    *string              `json:"UpdatedAt,omitempty"`
	GoodsServicesDescriptionText *RequestExamination  `json:"goods_services_description_text,omitempty"`
}

type ClaimElement struct {
	CreatedAt                   string             `json:"CreatedAt"`
	DeletedAt                   DeletedAt          `json:"DeletedAt"`
	ID                          string             `json:"ID"`
	UpdatedAt                   string             `json:"UpdatedAt"`
	ClaimCategoryType           string             `json:"claim_category_type"`
	ClaimCountryCode            *string            `json:"claim_country_code,omitempty"`
	ClaimDate                   RequestExamination `json:"claim_date"`
	ClaimForeignRegistrationNbr *string            `json:"claim_foreign_registration_nbr,omitempty"`
	ClaimNumber                 string             `json:"claim_number"`
	ClaimText                   string             `json:"claim_text"`
	ClaimTypeDescription        RequestExamination `json:"claim_type_description"`
	GoodsServicesReferenceBag   RequestExamination `json:"goods_services_reference_bag"`
}

type Claim struct {
	Element []ClaimElement `json:"element"`
}

type Contact struct {
	CreatedAt        string             `json:"CreatedAt"`
	DeletedAt        DeletedAt          `json:"DeletedAt"`
	ID               string             `json:"ID"`
	UpdatedAt        string             `json:"UpdatedAt"`
	LanguageCode     LanguageCode       `json:"language_code"`
	Name             RequestExamination `json:"name"`
	PostalAddressBag RequestExamination `json:"postal_address_bag"`
}

type Applicant struct {
	CreatedAt               string    `json:"CreatedAt"`
	DeletedAt               DeletedAt `json:"DeletedAt"`
	ID                      string    `json:"ID"`
	UpdatedAt               string    `json:"UpdatedAt"`
	Contact                 Contact   `json:"contact"`
	LegalEntityName         string    `json:"legal_entity_name"`
	NationalLegalEntityCode string    `json:"national_legal_entity_code"`
}

type RequestExamination struct {
	CreatedAt                          string                       `json:"CreatedAt"`
	DeletedAt                          DeletedAt                    `json:"DeletedAt"`
	ID                                 string                       `json:"ID"`
	UpdatedAt                          string                       `json:"UpdatedAt"`
	RequestExaminationCategory         *string                      `json:"request_examination_category,omitempty"`
	RequestSearchCategory              *string                      `json:"request_search_category,omitempty"`
	Trademark                          *Trademark                   `json:"trademark,omitempty"`
	Applicant                          *Applicant                   `json:"applicant,omitempty"`
	EntityName                         *RequestExamination          `json:"entity_name,omitempty"`
	LanguageCode                       *LanguageCode                `json:"language_code,omitempty"`
	PostalAddress                      *RequestExamination          `json:"postal_address,omitempty"`
	PostalStructuredAddress            *PostalStructuredAddress     `json:"postal_structured_address,omitempty"`
	SequenceNumber                     *string                      `json:"sequence_number,omitempty"`
	GoodsServices                      *GoodsServices               `json:"goods_services,omitempty"`
	ClassDescription                   *ClassDescription            `json:"class_description,omitempty"`
	GoodsServicesDescriptionText       *RequestExamination          `json:"goods_services_description_text,omitempty"`
	GoodsServicesClassification        *GoodsServicesClassification `json:"goods_services_classification,omitempty"`
	CurrentBasis                       *CurrentBasis                `json:"current_basis,omitempty"`
	NationalClassTotalQuantity         *string                      `json:"national_class_total_quantity,omitempty"`
	MarkEvent                          *MarkEvent                   `json:"mark_event,omitempty"`
	MarkImageBag                       *RequestExamination          `json:"mark_image_bag,omitempty"`
	MarkImage                          *MarkImage                   `json:"mark_image,omitempty"`
	FigurativeElementClassificationBag *RequestExamination          `json:"figurative_element_classification_bag,omitempty"`
	ViennaClassificationBag            *RequestExamination          `json:"vienna_classification_bag,omitempty"`
	ViennaClassification               *ViennaClassification        `json:"vienna_classification,omitempty"`
	ViennaDescription                  *LegislationDescription      `json:"vienna_description,omitempty"`
	GeographicRegionCategory           *string                      `json:"geographic_region_category,omitempty"`
	Claim                              *Claim                       `json:"claim,omitempty"`
	StructuredClaimDate                *string                      `json:"structured_claim_date,omitempty"`
	GoodsServicesReference             *GoodsServicesReference      `json:"goods_services_reference,omitempty"`
	GoodsServicesReferenceIdentifier   *string                      `json:"goods_services_reference_identifier,omitempty"`
	InterestedParty                    *InterestedParty             `json:"interested_party,omitempty"`
	LegislationDescription             *LegislationDescription      `json:"legislation_description,omitempty"`
	MarkFeatureCategory                *string                      `json:"mark_feature_category,omitempty"`
	TrademarkClassDescription          *LegislationDescription      `json:"trademark_class_description,omitempty"`
	Priority                           *Priority                    `json:"priority,omitempty"`
	ApplicationNumberText              *string                      `json:"application_number_text,omitempty"`
	ClassDescriptionBag                *RequestExamination          `json:"class_description_bag,omitempty"`
	Publication                        *Publication                 `json:"publication,omitempty"`
}

type DeletedAt struct {
	Null string `json:"_null"`
}

type ApplicationNumber struct {
	CreatedAt              string    `json:"CreatedAt"`
	DeletedAt              DeletedAt `json:"DeletedAt"`
	ID                     string    `json:"ID"`
	UpdatedAt              string    `json:"UpdatedAt"`
	IPOfficeCode           string    `json:"ip_office_code"`
	St13_ApplicationNumber string    `json:"st_13_application_number"`
}

type CurrentBasis struct {
	CreatedAt                         string    `json:"CreatedAt"`
	DeletedAt                         DeletedAt `json:"DeletedAt"`
	ID                                string    `json:"ID"`
	UpdatedAt                         string    `json:"UpdatedAt"`
	BasisForeignApplicationIndicator  string    `json:"basis_foreign_application_indicator"`
	BasisForeignRegistrationIndicator string    `json:"basis_foreign_registration_indicator"`
	BasisIntentToUseIndicator         string    `json:"basis_intent_to_use_indicator"`
	BasisUseIndicator                 string    `json:"basis_use_indicator"`
	NoBasisIndicator                  string    `json:"no_basis_indicator"`
}

type GoodsServicesClassification struct {
	Element []GoodsServicesClassificationElement `json:"element"`
}

type GoodsServicesClassificationElement struct {
	CreatedAt              string    `json:"CreatedAt"`
	DeletedAt              DeletedAt `json:"DeletedAt"`
	ID                     string    `json:"ID"`
	UpdatedAt              string    `json:"UpdatedAt"`
	ClassNumber            string    `json:"class_number"`
	ClassTitleText         string    `json:"class_title_text"`
	ClassificationKindCode string    `json:"classification_kind_code"`
	CommentText            string    `json:"comment_text"`
}

type Publication struct {
	CreatedAt             string              `json:"CreatedAt"`
	DeletedAt             DeletedAt           `json:"DeletedAt"`
	ID                    string              `json:"ID"`
	UpdatedAt             string              `json:"UpdatedAt"`
	NationalPublication   NationalPublication `json:"national_publication"`
	PublicationIdentifier string              `json:"publication_identifier"`
}

type NationalPublication struct {
	CreatedAt                 string    `json:"CreatedAt"`
	DeletedAt                 DeletedAt `json:"DeletedAt"`
	ID                        string    `json:"ID"`
	UpdatedAt                 string    `json:"UpdatedAt"`
	PublicationActionDate     string    `json:"publication_action_date"`
	PublicationStatusCategory string    `json:"publication_status_category"`
}

type XMLName struct {
	Local string `json:"Local"`
	Space string `json:"Space"`
}

type LanguageCode string

const (
	En LanguageCode = "en"
	Fr LanguageCode = "fr"
)

type Type string

const (
	CatmkNationalMarkEventType Type = "catmk:NationalMarkEventType"
)

type ElementUnion struct {
	RequestExamination      *RequestExamination
	RequestExaminationArray []RequestExamination
}
