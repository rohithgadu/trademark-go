package models

import "gorm.io/datatypes"

type TrademarkTable struct {
	ID                int            `json:"id"`
	ApplicationNumber int            `json:"application_number"`
	TrademarkName     string         `json:"trademark_name"`
	Country           string         `json:"country"`
	Attributes        datatypes.JSON `json:"attributes"`
}

type TrademarkApplication struct {
	TrademarkBag struct {
		Trademark struct {
			RegistrationOfficeCode string `json:"registration_office_code,omitempty" xml:"RegistrationOfficeCode,omitempty"`
			ReceivingOfficeCode    string `json:"receiving_office_code,omitempty" xml:"ReceivingOfficeCode,omitempty"`
			ReceivingOfficeDate    string `json:"receiving_office_date,omitempty" xml:"ReceivingOfficeDate,omitempty"`
			ApplicationNumber      struct {
				IPOfficeCode          string `json:"ip_office_code,omitempty" xml:"IPOfficeCode,omitempty"`
				ST13ApplicationNumber int64  `json:"st_13_application_number,omitempty" xml:"ST13ApplicationNumber,omitempty"`
			} `json:"application_number,omitempty" xml:"ApplicationNumber,omitempty"`
			ApplicationDate         string `json:"application_date,omitempty" xml:"ApplicationDate,omitempty"`
			RegistrationNumber      string `json:"registration_number,omitempty" xml:"RegistrationNumber,omitempty"`
			RegistrationDate        string `json:"registration_date,omitempty" xml:"RegistrationDate,omitempty"`
			FilingPlace             string `json:"filing_place,omitempty" xml:"FilingPlace,omitempty"`
			ApplicationLanguageCode string `json:"application_language_code,omitempty" xml:"ApplicationLanguageCode,omitempty"`
			MarkCurrentStatusCode   string `json:"mark_current_status_code,omitempty" xml:"MarkCurrentStatusCode,omitempty"`
			MarkCurrentStatusDate   string `json:"mark_current_status_date,omitempty" xml:"MarkCurrentStatusDate,omitempty"`
			MarkCategory            string `json:"mark_category,omitempty" xml:"MarkCategory,omitempty"`
			MarkRepresentation      struct {
				MarkFeatureCategory string `json:"mark_feature_category,omitempty" xml:"MarkFeatureCategory,omitempty"`
				MarkReproduction    struct {
					WordMarkSpecification struct {
						MarkVerbalElementText            string `json:"mark_verbal_element_text,omitempty" xml:"MarkVerbalElementText,omitempty"`
						MarkSignificantVerbalElementText string `json:"mark_significant_verbal_element_text,omitempty" xml:"MarkSignificantVerbalElementText,omitempty"`
						MarkStandardCharacterIndicator   bool   `json:"mark_standard_character_indicator,omitempty" xml:"MarkStandardCharacterIndicator,omitempty"`
					} `json:"word_mark_specification,omitempty" xml:"WordMarkSpecification,omitempty"`
				} `json:"mark_reproduction,omitempty" xml:"MarkReproduction,omitempty"`
			} `json:"mark_representation,omitempty" xml:"MarkRepresentation,omitempty"`
			NonUseCancelledIndicator      bool `json:"non_use_cancelled_indicator,omitempty" xml:"NonUseCancelledIndicator,omitempty"`
			TradeDistinctivenessIndicator bool `json:"trade_distinctiveness_indicator,omitempty" xml:"TradeDistinctivenessIndicator,omitempty"`
			UseRight                      struct {
				UseRightIndicator bool `json:"use_right_indicator,omitempty" xml:"UseRightIndicator,omitempty"`
			} `json:"use_right,omitempty" xml:"UseRight,omitempty"`
			CommentText      string `json:"comment_text,omitempty" xml:"CommentText,omitempty"`
			GoodsServicesBag struct {
				GoodsServices struct {
					GoodsServicesClassificationBag struct {
						GoodsServicesClassification []struct {
							ClassificationKindCode string `json:"classification_kind_code,omitempty" xml:"ClassificationKindCode,omitempty"`
							CommentText            string `json:"comment_text,omitempty" xml:"CommentText,omitempty"`
							ClassNumber            int    `json:"class_number,omitempty" xml:"ClassNumber,omitempty"`
							ClassTitleText         string `json:"class_title_text,omitempty" xml:"ClassTitleText,omitempty"`
						} `json:"goods_services_classification,omitempty" xml:"GoodsServicesClassification,omitempty"`
					} `json:"goods_services_classification_bag,omitempty" xml:"GoodsServicesClassificationBag,omitempty"`
					NationalGoodsServices struct {
						NationalClassTotalQuantity int `json:"national_class_total_quantity,omitempty" xml:"NationalClassTotalQuantity,omitempty"`
					} `json:"national_goods_services,omitempty" xml:"NationalGoodsServices,omitempty"`
					NationalFilingBasis struct {
						CurrentBasis struct {
							BasisForeignApplicationIndicator  bool `json:"basis_foreign_application_indicator,omitempty" xml:"BasisForeignApplicationIndicator,omitempty"`
							BasisForeignRegistrationIndicator bool `json:"basis_foreign_registration_indicator,omitempty" xml:"BasisForeignRegistrationIndicator,omitempty"`
							BasisUseIndicator                 bool `json:"basis_use_indicator,omitempty" xml:"BasisUseIndicator,omitempty"`
							BasisIntentToUseIndicator         bool `json:"basis_intent_to_use_indicator,omitempty" xml:"BasisIntentToUseIndicator,omitempty"`
							NoBasisIndicator                  bool `json:"no_basis_indicator,omitempty" xml:"NoBasisIndicator,omitempty"`
						} `json:"current_basis,omitempty" xml:"CurrentBasis,omitempty"`
					} `json:"national_filing_basis,omitempty" xml:"NationalFilingBasis,omitempty"`
					ClassificationKindCode string `json:"classification_kind_code,omitempty" xml:"ClassificationKindCode,omitempty"`
					ClassDescriptionBag    struct {
						ClassDescription []struct {
							GoodsServicesDescriptionText string `json:"goods_services_description_text,omitempty" xml:"GoodsServicesDescriptionText,omitempty"`
						} `json:"class_description,omitempty" xml:"ClassDescription,omitempty"`
					} `json:"class_description_bag,omitempty" xml:"ClassDescriptionBag,omitempty"`
				} `json:"goods_services,omitempty" xml:"GoodsServices,omitempty"`
			} `json:"goods_services_bag,omitempty" xml:"GoodsServicesBag,omitempty"`
			ApplicantBag struct {
				Applicant struct {
					LegalEntityName string `json:"legal_entity_name,omitempty" xml:"LegalEntityName,omitempty"`
					Contact         struct {
						Name struct {
							EntityName string `json:"entity_name,omitempty" xml:"EntityName,omitempty"`
						} `json:"name,omitempty" xml:"Name,omitempty"`
						PostalAddressBag struct {
							PostalAddress struct {
								PostalStructuredAddress struct {
									AddressLineText      []string `json:"address_line_text,omitempty" xml:"AddressLineText,omitempty"`
									GeographicRegionName string   `json:"geographic_region_name,omitempty" xml:"GeographicRegionName,omitempty"`
									CountryCode          string   `json:"country_code,omitempty" xml:"CountryCode,omitempty"`
									PostalCode           string   `json:"postal_code,omitempty" xml:"PostalCode,omitempty"`
								} `json:"postal_structured_address,omitempty" xml:"PostalStructuredAddress,omitempty"`
							} `json:"postal_address,omitempty" xml:"PostalAddress,omitempty"`
						} `json:"postal_address_bag,omitempty" xml:"PostalAddressBag,omitempty"`
					} `json:"contact,omitempty" xml:"Contact,omitempty"`
					NationalLegalEntityCode string `json:"national_legal_entity_code,omitempty" xml:"NationalLegalEntityCode,omitempty"`
				} `json:"applicant,omitempty" xml:"Applicant,omitempty"`
			} `json:"applicant_bag,omitempty" xml:"ApplicantBag,omitempty"`
			MarkEventBag struct {
				MarkEvent []struct {
					MarkEventCategory string `json:"mark_event_category,omitempty" xml:"MarkEventCategory,omitempty"`
					NationalMarkEvent struct {
						MarkEventCode                            int    `json:"mark_event_code,omitempty" xml:"MarkEventCode,omitempty"`
						MarkEventDescriptionText                 string `json:"mark_event_description_text,omitempty" xml:"MarkEventDescriptionText,omitempty"`
						MarkEventOtherLanguageDescriptionTextBag struct {
							MarkEventDescriptionText string `json:"mark_event_description_text,omitempty" xml:"MarkEventDescriptionText,omitempty"`
						} `json:"mark_event_other_language_description_text_bag,omitempty" xml:"MarkEventOtherLanguageDescriptionTextBag,omitempty"`
						MarkEventAdditionalText string `json:"mark_event_additional_text,omitempty,omitempty" xml:"MarkEventAdditionalText,omitempty"`
					} `json:"national_mark_event,omitempty" xml:"NationalMarkEvent,omitempty"`
					MarkEventDate         string `json:"mark_event_date,omitempty" xml:"MarkEventDate,omitempty"`
					MarkEventResponseDate string `json:"mark_event_response_date,omitempty,omitempty" xml:"MarkEventResponseDate,omitempty"`
				} `json:"mark_event,omitempty" xml:"MarkEvent,omitempty"`
			} `json:"mark_event_bag,omitempty" xml:"MarkEventBag,omitempty"`
			NationalTrademarkInformation struct {
				RegisterCategory                         string `json:"register_category,omitempty" xml:"RegisterCategory,omitempty"`
				ApplicationAbandonedDate                 string `json:"application_abandoned_date,omitempty" xml:"ApplicationAbandonedDate,omitempty"`
				RenewalDate                              string `json:"renewal_date,omitempty" xml:"RenewalDate,omitempty"`
				AllowedDate                              string `json:"allowed_date,omitempty" xml:"AllowedDate,omitempty"`
				MarkCurrentStatusInternalDescriptionText string `json:"mark_current_status_internal_description_text,omitempty" xml:"MarkCurrentStatusInternalDescriptionText,omitempty"`
				TrademarkClass                           struct {
					TrademarkClassCode           int `json:"trademark_class_code,omitempty" xml:"TrademarkClassCode,omitempty"`
					TrademarkClassDescriptionBag struct {
						TrademarkClassDescription []string `json:"trademark_class_description,omitempty" xml:"TrademarkClassDescription,omitempty"`
					} `json:"trademark_class_description_bag,omitempty" xml:"TrademarkClassDescriptionBag,omitempty"`
				} `json:"trademark_class,omitempty" xml:"TrademarkClass,omitempty"`
				Legislation struct {
					LegislationCode           int `json:"legislation_code,omitempty" xml:"LegislationCode,omitempty"`
					LegislationDescriptionBag struct {
						LegislationDescription []string `json:"legislation_description,omitempty" xml:"LegislationDescription,omitempty"`
					} `json:"legislation_description_bag,omitempty" xml:"LegislationDescriptionBag,omitempty"`
				} `json:"legislation,omitempty" xml:"Legislation,omitempty"`
				InterestedPartyBag struct {
					InterestedParty struct {
						InterestedPartyCategory string `json:"interested_party_category,omitempty" xml:"InterestedPartyCategory,omitempty"`
						Contact                 struct {
							Name struct {
								EntityName string `json:"entity_name,omitempty" xml:"EntityName,omitempty"`
							} `json:"name,omitempty" xml:"Name,omitempty"`
							PostalAddressBag struct {
								PostalAddress struct {
									PostalStructuredAddress struct {
										AddressLineText      []string `json:"address_line_text,omitempty" xml:"AddressLineText,omitempty"`
										GeographicRegionName string   `json:"geographic_region_name,omitempty" xml:"GeographicRegionName,omitempty"`
										CountryCode          string   `json:"country_code,omitempty" xml:"CountryCode,omitempty"`
										PostalCode           string   `json:"postal_code,omitempty" xml:"PostalCode,omitempty"`
									} `json:"postal_structured_address,omitempty" xml:"PostalStructuredAddress,omitempty"`
								} `json:"postal_address,omitempty" xml:"PostalAddress,omitempty"`
							} `json:"postal_address_bag,omitempty" xml:"PostalAddressBag,omitempty"`
						} `json:"contact,omitempty" xml:"Contact,omitempty"`
					} `json:"interested_party,omitempty" xml:"InterestedParty,omitempty"`
				} `json:"interested_party_bag,omitempty" xml:"InterestedPartyBag,omitempty"`
				ClaimBag struct {
					Claim struct {
						ClaimCategoryType    int    `json:"claim_category_type,omitempty" xml:"ClaimCategoryType,omitempty"`
						ClaimTypeDescription string `json:"claim_type_description,omitempty" xml:"ClaimTypeDescription,omitempty"`
						ClaimNumber          int    `json:"claim_number,omitempty" xml:"ClaimNumber,omitempty"`
						ClaimCode            int    `json:"claim_code,omitempty" xml:"ClaimCode,omitempty"`
						ClaimDate            struct {
							UnstructuredClaimDate struct {
								ClaimYear  int `json:"claim_year,omitempty" xml:"ClaimYear,omitempty"`
								ClaimMonth int `json:"claim_month,omitempty" xml:"ClaimMonth,omitempty"`
							} `json:"unstructured_claim_date,omitempty" xml:"UnstructuredClaimDate,omitempty"`
						} `json:"claim_date,omitempty" xml:"ClaimDate,omitempty"`
						ClaimText                 string `json:"claim_text,omitempty" xml:"ClaimText,omitempty"`
						GoodsServicesReferenceBag struct {
							GoodsServicesReference []struct {
								GoodsServicesReferenceIdentifier string `json:"goods_services_reference_identifier,omitempty" xml:"GoodsServicesReferenceIdentifier,omitempty"`
							} `json:"goods_services_reference,omitempty" xml:"GoodsServicesReference,omitempty"`
						} `json:"goods_services_reference_bag,omitempty" xml:"GoodsServicesReferenceBag,omitempty"`
					} `json:"claim,omitempty" xml:"Claim,omitempty"`
				} `json:"claim_bag,omitempty" xml:"ClaimBag,omitempty"`
				IndexHeadingBag struct {
					IndexHeading []struct {
						IndexHeadingText string `json:"index_heading_text,omitempty" xml:"IndexHeadingText,omitempty"`
					} `json:"index_heading,omitempty" xml:"IndexHeading,omitempty"`
				} `json:"index_heading_bag,omitempty" xml:"IndexHeadingBag,omitempty"`
				MarkFeatureCategoryBag struct {
					MarkFeatureCategory string `json:"mark_feature_category,omitempty" xml:"MarkFeatureCategory,omitempty"`
				} `json:"mark_feature_category_bag,omitempty" xml:"MarkFeatureCategoryBag,omitempty"`
			} `json:"national_trademark_information,omitempty" xml:"NationalTrademarkInformation,omitempty"`
		} `json:"trademark,omitempty" xml:"Trademark,omitempty"`
	} `json:"trademark_bag,omitempty" xml:"TrademarkBag,omitempty"`
}
