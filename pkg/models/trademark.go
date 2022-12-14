package models

import (
	"encoding/xml"
	"github.com/jinzhu/gorm"
	"go-demo/pkg/config"
)

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

type TrademarkApplication struct {
	XMLName xml.Name `xml:"TrademarkApplication" json:"xml_name"`
	//Text          string   `xml:",chardata" json:"text,omitempty"`
	Catmk         string `xml:"catmk,attr" json:"catmk,omitempty"`
	Ns4           string `xml:"ns4,attr" json:"ns_4,omitempty"`
	Ns3           string `xml:"ns3,attr" json:"ns_3,omitempty"`
	Com           string `xml:"com,attr" json:"com,omitempty"`
	Cacom         string `xml:"cacom,attr" json:"cacom,omitempty"`
	Xs            string `xml:"xs,attr" json:"xs,omitempty"`
	Tmk           string `xml:"tmk,attr" json:"tmk,omitempty"`
	Xsi           string `xml:"xsi,attr" json:"xsi,omitempty"`
	St96Version   string `xml:"st96Version,attr" json:"st_96_version,omitempty"`
	IpoVersion    string `xml:"ipoVersion,attr" json:"ipo_version,omitempty"`
	RequestSearch struct {
		gorm.Model
		//Text                  string `xml:",chardata" json:"text,omitempty"`
		RequestSearchCategory string `xml:"RequestSearchCategory" json:"request_search_category,omitempty"`
	} `xml:"RequestSearch" json:"request_search"`
	RequestExamination struct {
		gorm.Model
		//Text                       string `xml:",chardata" json:"text,omitempty"`
		RequestExaminationCategory string `xml:"RequestExaminationCategory" json:"request_examination_category,omitempty"`
	} `xml:"RequestExamination" json:"request_examination"`
	TrademarkBag struct {
		gorm.Model
		//Text      string `xml:",chardata" json:"text,omitempty"`
		Trademark struct {
			gorm.Model
			//Text                   string `xml:",chardata" json:"text,omitempty"`
			OperationCategory      string `xml:"operationCategory,attr" json:"operation_category,omitempty"`
			RegistrationOfficeCode string `xml:"RegistrationOfficeCode" json:"registration_office_code,omitempty"`
			ReceivingOfficeCode    string `xml:"ReceivingOfficeCode" json:"receiving_office_code,omitempty"`
			ReceivingOfficeDate    string `xml:"ReceivingOfficeDate" json:"receiving_office_date,omitempty"`
			ApplicationNumber      struct {
				gorm.Model
				//Text                  string `xml:",chardata" json:"text,omitempty"`
				IPOfficeCode          string `xml:"IPOfficeCode" json:"ip_office_code,omitempty"`
				ST13ApplicationNumber string `xml:"ST13ApplicationNumber" json:"st_13_application_number,omitempty"`
			} `xml:"ApplicationNumber" json:"application_number"`
			RegistrationNumber      string `xml:"RegistrationNumber" json:"registration_number,omitempty"`
			ApplicationDate         string `xml:"ApplicationDate" json:"application_date,omitempty"`
			RegistrationDate        string `xml:"RegistrationDate" json:"registration_date,omitempty"`
			FilingPlace             string `xml:"FilingPlace" json:"filing_place,omitempty"`
			ApplicationLanguageCode string `xml:"ApplicationLanguageCode" json:"application_language_code,omitempty"`
			ExpiryDate              string `xml:"ExpiryDate" json:"expiry_date,omitempty"`
			MarkCurrentStatusCode   string `xml:"MarkCurrentStatusCode" json:"mark_current_status_code,omitempty"`
			MarkCurrentStatusDate   string `xml:"MarkCurrentStatusDate" json:"mark_current_status_date,omitempty"`
			MarkCategory            string `xml:"MarkCategory" json:"mark_category,omitempty"`
			MarkRepresentation      struct {
				gorm.Model
				//Text                string `xml:",chardata" json:"text,omitempty"`
				MarkFeatureCategory string `xml:"MarkFeatureCategory" json:"mark_feature_category,omitempty"`
				MarkReproduction    struct {
					gorm.Model
					//Text         string `xml:",chardata" json:"text,omitempty"`
					MarkImageBag struct {
						gorm.Model
						//Text      string `xml:",chardata" json:"text,omitempty"`
						MarkImage struct {
							gorm.Model
							//Text                    string `xml:",chardata" json:"text,omitempty"`
							FileName                string `xml:"FileName" json:"file_name,omitempty"`
							ImageFormatCategory     string `xml:"ImageFormatCategory" json:"image_format_category,omitempty"`
							MarkImageClassification struct {
								gorm.Model
								//Text                               string `xml:",chardata" json:"text,omitempty"`
								FigurativeElementClassificationBag struct {
									gorm.Model
									//Text                    string `xml:",chardata" json:"text,omitempty"`
									ViennaClassificationBag struct {
										gorm.Model
										//Text                 string `xml:",chardata" json:"text,omitempty"`
										ViennaClassification []struct {
											gorm.Model
											//Text                 string `xml:",chardata" json:"text,omitempty"`
											Type                 string `xml:"type,attr" json:"type,omitempty"`
											ViennaCategory       string `xml:"ViennaCategory" json:"vienna_category,omitempty"`
											ViennaDivision       string `xml:"ViennaDivision" json:"vienna_division,omitempty"`
											ViennaSection        string `xml:"ViennaSection" json:"vienna_section,omitempty"`
											ViennaDescriptionBag struct {
												gorm.Model
												//Text              string `xml:",chardata" json:"text,omitempty"`
												ViennaDescription []struct {
													gorm.Model
													//Text         string `xml:",chardata" json:"text,omitempty"`
													LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
												} `xml:"ViennaDescription" json:"vienna_description,omitempty"`
											} `xml:"ViennaDescriptionBag" json:"vienna_description_bag"`
										} `xml:"ViennaClassification" json:"vienna_classification,omitempty"`
									} `xml:"ViennaClassificationBag" json:"vienna_classification_bag"`
								} `xml:"FigurativeElementClassificationBag" json:"figurative_element_classification_bag"`
							} `xml:"MarkImageClassification" json:"mark_image_classification"`
						} `xml:"MarkImage" json:"mark_image"`
					} `xml:"MarkImageBag" json:"mark_image_bag"`
				} `xml:"MarkReproduction" json:"mark_reproduction"`
			} `xml:"MarkRepresentation" json:"mark_representation"`
			NonUseCancelledIndicator      string `xml:"NonUseCancelledIndicator" json:"non_use_cancelled_indicator,omitempty"`
			TradeDistinctivenessIndicator string `xml:"TradeDistinctivenessIndicator" json:"trade_distinctiveness_indicator,omitempty"`
			UseRight                      []struct {
				gorm.Model
				//Text              string `xml:",chardata" json:"text,omitempty"`
				UseRightIndicator string `xml:"UseRightIndicator" json:"use_right_indicator,omitempty"`
				UseRightText      struct {
					gorm.Model
					//Text         string `xml:",chardata" json:"text,omitempty"`
					LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
				} `xml:"UseRightText" json:"use_right_text"`
			} `xml:"UseRight" json:"use_right,omitempty"`
			//CommentText               string `xml:"CommentText" json:"comment_text,omitempty"`
			OppositionPeriodStartDate string `xml:"OppositionPeriodStartDate" json:"opposition_period_start_date,omitempty"`
			OppositionPeriodEndDate   string `xml:"OppositionPeriodEndDate" json:"opposition_period_end_date,omitempty"`
			GoodsServicesBag          struct {
				gorm.Model
				//Text          string `xml:",chardata" json:"text,omitempty"`
				GoodsServices struct {
					gorm.Model
					//Text                           string `xml:",chardata" json:"text,omitempty"`
					GoodsServicesClassificationBag struct {
						gorm.Model
						//Text                        string `xml:",chardata" json:"text,omitempty"`
						GoodsServicesClassification []struct {
							gorm.Model
							//Text                   string `xml:",chardata" json:"text,omitempty"`
							ClassificationKindCode string `xml:"ClassificationKindCode" json:"classification_kind_code,omitempty"`
							CommentText            string `xml:"CommentText" json:"comment_text,omitempty"`
							ClassNumber            string `xml:"ClassNumber" json:"class_number,omitempty"`
							ClassTitleText         string `xml:"ClassTitleText" json:"class_title_text,omitempty"`
						} `xml:"GoodsServicesClassification" json:"goods_services_classification,omitempty"`
					} `xml:"GoodsServicesClassificationBag" json:"goods_services_classification_bag"`
					NationalGoodsServices struct {
						gorm.Model
						//Text                       string `xml:",chardata" json:"text,omitempty"`
						NationalClassTotalQuantity string `xml:"NationalClassTotalQuantity" json:"national_class_total_quantity,omitempty"`
					} `xml:"NationalGoodsServices" json:"national_goods_services"`
					NationalFilingBasis struct {
						gorm.Model
						//Text         string `xml:",chardata" json:"text,omitempty"`
						CurrentBasis struct {
							gorm.Model
							//Text                              string `xml:",chardata" json:"text,omitempty"`
							BasisForeignApplicationIndicator  string `xml:"BasisForeignApplicationIndicator" json:"basis_foreign_application_indicator,omitempty"`
							BasisForeignRegistrationIndicator string `xml:"BasisForeignRegistrationIndicator" json:"basis_foreign_registration_indicator,omitempty"`
							BasisUseIndicator                 string `xml:"BasisUseIndicator" json:"basis_use_indicator,omitempty"`
							BasisIntentToUseIndicator         string `xml:"BasisIntentToUseIndicator" json:"basis_intent_to_use_indicator,omitempty"`
							NoBasisIndicator                  string `xml:"NoBasisIndicator" json:"no_basis_indicator,omitempty"`
						} `xml:"CurrentBasis" json:"current_basis"`
					} `xml:"NationalFilingBasis" json:"national_filing_basis"`
					ClassificationKindCode string `xml:"ClassificationKindCode" json:"classification_kind_code,omitempty"`
					ClassDescriptionBag    struct {
						gorm.Model
						//Text             string `xml:",chardata" json:"text,omitempty"`
						ClassDescription []struct {
							gorm.Model
							//Text                         string `xml:",chardata" json:"text,omitempty"`
							GoodsServicesDescriptionText struct {
								gorm.Model
								//Text           string `xml:",chardata" json:"text,omitempty"`
								SequenceNumber string `xml:"sequenceNumber,attr" json:"sequence_number,omitempty"`
								LanguageCode   string `xml:"languageCode,attr" json:"language_code,omitempty"`
							} `xml:"GoodsServicesDescriptionText" json:"goods_services_description_text"`
						} `xml:"ClassDescription" json:"class_description,omitempty"`
					} `xml:"ClassDescriptionBag" json:"class_description_bag"`
				} `xml:"GoodsServices" json:"goods_services"`
			} `xml:"GoodsServicesBag" json:"goods_services_bag"`
			PriorityBag struct {
				gorm.Model
				//Text     string `xml:",chardata" json:"text,omitempty"`
				Priority []struct {
					gorm.Model
					//Text                string `xml:",chardata" json:"text,omitempty"`
					PriorityCountryCode string `xml:"PriorityCountryCode" json:"priority_country_code,omitempty"`
					ApplicationNumber   struct {
						gorm.Model
						//Text                  string `xml:",chardata" json:"text,omitempty"`
						ApplicationNumberText string `xml:"ApplicationNumberText" json:"application_number_text,omitempty"`
					} `xml:"ApplicationNumber" json:"application_number"`
					PriorityApplicationFilingDate string `xml:"PriorityApplicationFilingDate" json:"priority_application_filing_date,omitempty"`
					PriorityPartialGoodsServices  struct {
						gorm.Model
						//Text                string `xml:",chardata" json:"text,omitempty"`
						ClassDescriptionBag struct {
							gorm.Model
							//Text             string `xml:",chardata" json:"text,omitempty"`
							ClassDescription struct {
								gorm.Model
								//Text                         string `xml:",chardata" json:"text,omitempty"`
								GoodsServicesDescriptionText struct {
									gorm.Model
									//Text           string `xml:",chardata" json:"text,omitempty"`
									SequenceNumber string `xml:"sequenceNumber,attr" json:"sequence_number,omitempty"`
									LanguageCode   string `xml:"languageCode,attr" json:"language_code,omitempty"`
								} `xml:"GoodsServicesDescriptionText" json:"goods_services_description_text"`
							} `xml:"ClassDescription" json:"class_description"`
						} `xml:"ClassDescriptionBag" json:"class_description_bag"`
					} `xml:"PriorityPartialGoodsServices" json:"priority_partial_goods_services"`
				} `xml:"Priority" json:"priority,omitempty"`
			} `xml:"PriorityBag" json:"priority_bag"`
			PublicationBag struct {
				gorm.Model  //Text        string `xml:",chardata" json:"text,omitempty"`
				Publication struct {
					gorm.Model
					//Text                  string `xml:",chardata" json:"text,omitempty"`
					PublicationIdentifier string `xml:"PublicationIdentifier" json:"publication_identifier,omitempty"`
					NationalPublication   struct {
						gorm.Model
						//Text                      string `xml:",chardata" json:"text,omitempty"`
						PublicationStatusCategory string `xml:"PublicationStatusCategory" json:"publication_status_category,omitempty"`
						PublicationActionDate     string `xml:"PublicationActionDate" json:"publication_action_date,omitempty"`
					} `xml:"NationalPublication" json:"national_publication"`
				} `xml:"Publication" json:"publication"`
			} `xml:"PublicationBag" json:"publication_bag"`
			ApplicantBag struct {
				gorm.Model
				//Text      string `xml:",chardata" json:"text,omitempty"`
				Applicant struct {
					gorm.Model
					//Text            string `xml:",chardata" json:"text,omitempty"`
					LegalEntityName string `xml:"LegalEntityName" json:"legal_entity_name,omitempty"`
					Contact         struct {
						gorm.Model
						//Text         string `xml:",chardata" json:"text,omitempty"`
						LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
						Name         struct {
							gorm.Model
							//Text       string `xml:",chardata" json:"text,omitempty"`
							EntityName struct {
								gorm.Model
								//Text         string `xml:",chardata" json:"text,omitempty"`
								LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
							} `xml:"EntityName" json:"entity_name"`
						} `xml:"Name" json:"name"`
						PostalAddressBag struct {
							gorm.Model
							//Text          string `xml:",chardata" json:"text,omitempty"`
							PostalAddress struct {
								gorm.Model
								//Text                    string `xml:",chardata" json:"text,omitempty"`
								PostalStructuredAddress struct {
									gorm.Model
									//Text            string `xml:",chardata" json:"text,omitempty"`
									LanguageCode    string `xml:"languageCode,attr" json:"language_code,omitempty"`
									AddressLineText []struct {
										gorm.Model
										//Text           string `xml:",chardata" json:"text,omitempty"`
										SequenceNumber string `xml:"sequenceNumber,attr" json:"sequence_number,omitempty"`
									} `xml:"AddressLineText" json:"address_line_text,omitempty"`
									CountryCode string `xml:"CountryCode" json:"country_code,omitempty"`
								} `xml:"PostalStructuredAddress" json:"postal_structured_address"`
							} `xml:"PostalAddress" json:"postal_address"`
						} `xml:"PostalAddressBag" json:"postal_address_bag"`
					} `xml:"Contact" json:"contact"`
					NationalLegalEntityCode string `xml:"NationalLegalEntityCode" json:"national_legal_entity_code,omitempty"`
				} `xml:"Applicant" json:"applicant"`
			} `xml:"ApplicantBag" json:"applicant_bag"`
			NationalRepresentative struct {
				gorm.Model
				//Text        string `xml:",chardata" json:"text,omitempty"`
				CommentText string `xml:"CommentText" json:"comment_text,omitempty"`
				Contact     struct {
					gorm.Model
					//Text         string `xml:",chardata" json:"text,omitempty"`
					LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
					Name         struct {
						gorm.Model
						//Text       string `xml:",chardata" json:"text,omitempty"`
						EntityName struct {
							gorm.Model
							//Text         string `xml:",chardata" json:"text,omitempty"`
							LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
						} `xml:"EntityName" json:"entity_name"`
					} `xml:"Name" json:"name"`
					PostalAddressBag struct {
						gorm.Model
						//Text          string `xml:",chardata" json:"text,omitempty"`
						PostalAddress struct {
							gorm.Model
							//Text                    string `xml:",chardata" json:"text,omitempty"`
							PostalStructuredAddress struct {
								gorm.Model
								//Text            string `xml:",chardata" json:"text,omitempty"`
								LanguageCode    string `xml:"languageCode,attr" json:"language_code,omitempty"`
								AddressLineText []struct {
									gorm.Model
									//Text           string `xml:",chardata" json:"text,omitempty"`
									SequenceNumber string `xml:"sequenceNumber,attr" json:"sequence_number,omitempty"`
								} `xml:"AddressLineText" json:"address_line_text,omitempty"`
								GeographicRegionName struct {
									gorm.Model
									//Text                     string `xml:",chardata" json:"text,omitempty"`
									GeographicRegionCategory string `xml:"geographicRegionCategory,attr" json:"geographic_region_category,omitempty"`
								} `xml:"GeographicRegionName" json:"geographic_region_name"`
								CountryCode string `xml:"CountryCode" json:"country_code,omitempty"`
								PostalCode  string `xml:"PostalCode" json:"postal_code,omitempty"`
							} `xml:"PostalStructuredAddress" json:"postal_structured_address"`
						} `xml:"PostalAddress" json:"postal_address"`
					} `xml:"PostalAddressBag" json:"postal_address_bag"`
				} `xml:"Contact" json:"contact"`
			} `xml:"NationalRepresentative" json:"national_representative"`
			MarkEventBag struct {
				gorm.Model
				//Text      string `xml:",chardata" json:"text,omitempty"`
				MarkEvent []struct {
					gorm.Model
					//Text              string `xml:",chardata" json:"text,omitempty"`
					MarkEventCategory string `xml:"MarkEventCategory" json:"mark_event_category,omitempty"`
					NationalMarkEvent struct {
						gorm.Model
						//Text                                     string `xml:",chardata" json:"text,omitempty"`
						Type                                     string `xml:"type,attr" json:"type,omitempty"`
						MarkEventCode                            string `xml:"MarkEventCode" json:"mark_event_code,omitempty"`
						MarkEventDescriptionText                 string `xml:"MarkEventDescriptionText" json:"mark_event_description_text,omitempty"`
						MarkEventOtherLanguageDescriptionTextBag struct {
							//Text                     string `xml:",chardata" json:"text,omitempty"`
							MarkEventDescriptionText struct {
								gorm.Model
								//Text         string `xml:",chardata" json:"text,omitempty"`
								LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
							} `xml:"MarkEventDescriptionText" json:"mark_event_description_text"`
						} `xml:"MarkEventOtherLanguageDescriptionTextBag" json:"mark_event_other_language_description_text_bag"`
						MarkEventAdditionalText string `xml:"MarkEventAdditionalText" json:"mark_event_additional_text,omitempty"`
					} `xml:"NationalMarkEvent" json:"national_mark_event"`
					MarkEventDate         string `xml:"MarkEventDate" json:"mark_event_date,omitempty"`
					MarkEventResponseDate string `xml:"MarkEventResponseDate" json:"mark_event_response_date,omitempty"`
				} `xml:"MarkEvent" json:"mark_event,omitempty"`
			} `xml:"MarkEventBag" json:"mark_event_bag"`
			MarkFeatureDescription       string `xml:"MarkFeatureDescription" json:"mark_feature_description,omitempty"`
			NationalTrademarkInformation struct {
				gorm.Model
				//Text                                     string `xml:",chardata" json:"text,omitempty"`
				Type                                     string `xml:"type,attr" json:"type,omitempty"`
				RegisterCategory                         string `xml:"RegisterCategory" json:"register_category,omitempty"`
				MarkCurrentStatusInternalDescriptionText string `xml:"MarkCurrentStatusInternalDescriptionText" json:"mark_current_status_internal_description_text,omitempty"`
				RenewalDate                              string `xml:"RenewalDate" json:"renewal_date,omitempty"`
				AllowedDate                              string `xml:"AllowedDate" json:"allowed_date,omitempty"`
				TrademarkClass                           struct {
					//Text                         string `xml:",chardata" json:"text,omitempty"`
					TrademarkClassCode           string `xml:"TrademarkClassCode" json:"trademark_class_code,omitempty"`
					TrademarkClassDescriptionBag struct {
						gorm.Model
						//Text                      string `xml:",chardata" json:"text,omitempty"`
						TrademarkClassDescription []struct {
							gorm.Model
							//Text         string `xml:",chardata" json:"text,omitempty"`
							LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
						} `xml:"TrademarkClassDescription" json:"trademark_class_description,omitempty"`
					} `xml:"TrademarkClassDescriptionBag" json:"trademark_class_description_bag"`
				} `xml:"TrademarkClass" json:"trademark_class"`
				Legislation struct {
					gorm.Model
					//Text                      string `xml:",chardata" json:"text,omitempty"`
					LegislationCode           string `xml:"LegislationCode" json:"legislation_code,omitempty"`
					LegislationDescriptionBag struct {
						gorm.Model
						//Text                   string `xml:",chardata" json:"text,omitempty"`
						LegislationDescription []struct {
							gorm.Model
							//Text         string `xml:",chardata" json:"text,omitempty"`
							LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
						} `xml:"LegislationDescription" json:"legislation_description,omitempty"`
					} `xml:"LegislationDescriptionBag" json:"legislation_description_bag"`
				} `xml:"Legislation" json:"legislation"`
				InterestedPartyBag struct {
					gorm.Model
					//Text            string `xml:",chardata" json:"text,omitempty"`
					InterestedParty []struct {
						gorm.Model
						//Text                    string `xml:",chardata" json:"text,omitempty"`
						InterestedPartyCategory string `xml:"InterestedPartyCategory" json:"interested_party_category,omitempty"`
						Contact                 struct {
							gorm.Model
							//Text         string `xml:",chardata" json:"text,omitempty"`
							LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
							Name         struct {
								gorm.Model
								//Text       string `xml:",chardata" json:"text,omitempty"`
								EntityName struct {
									gorm.Model
									//Text         string `xml:",chardata" json:"text,omitempty"`
									LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
								} `xml:"EntityName" json:"entity_name"`
							} `xml:"Name" json:"name"`
							PostalAddressBag struct {
								gorm.Model
								//Text          string `xml:",chardata" json:"text,omitempty"`
								PostalAddress struct {
									gorm.Model
									//Text                    string `xml:",chardata" json:"text,omitempty"`
									PostalStructuredAddress struct {
										gorm.Model
										//Text            string `xml:",chardata" json:"text,omitempty"`
										LanguageCode    string `xml:"languageCode,attr" json:"language_code,omitempty"`
										AddressLineText []struct {
											gorm.Model
											//Text           string `xml:",chardata" json:"text,omitempty"`
											SequenceNumber string `xml:"sequenceNumber,attr" json:"sequence_number,omitempty"`
										} `xml:"AddressLineText" json:"address_line_text,omitempty"`
										CountryCode string `xml:"CountryCode" json:"country_code,omitempty"`
									} `xml:"PostalStructuredAddress" json:"postal_structured_address"`
								} `xml:"PostalAddress" json:"postal_address"`
							} `xml:"PostalAddressBag" json:"postal_address_bag"`
						} `xml:"Contact" json:"contact"`
					} `xml:"InterestedParty" json:"interested_party,omitempty"`
				} `xml:"InterestedPartyBag" json:"interested_party_bag"`
				ClaimBag struct {
					gorm.Model
					//Text  string `xml:",chardata" json:"text,omitempty"`
					Claim []struct {
						gorm.Model
						//Text                 string `xml:",chardata" json:"text,omitempty"`
						ClaimCategoryType    string `xml:"ClaimCategoryType" json:"claim_category_type,omitempty"`
						ClaimTypeDescription struct {
							gorm.Model
							//Text         string `xml:",chardata" json:"text,omitempty"`
							LanguageCode string `xml:"languageCode,attr" json:"language_code,omitempty"`
						} `xml:"ClaimTypeDescription" json:"claim_type_description"`
						ClaimNumber string `xml:"ClaimNumber" json:"claim_number,omitempty"`
						ClaimDate   struct {
							gorm.Model
							//Text                string `xml:",chardata" json:"text,omitempty"`
							StructuredClaimDate string `xml:"StructuredClaimDate" json:"structured_claim_date,omitempty"`
						} `xml:"ClaimDate" json:"claim_date"`
						ClaimCountryCode            string `xml:"ClaimCountryCode" json:"claim_country_code,omitempty"`
						ClaimForeignRegistrationNbr string `xml:"ClaimForeignRegistrationNbr" json:"claim_foreign_registration_nbr,omitempty"`
						ClaimText                   string `xml:"ClaimText" json:"claim_text,omitempty"`
						GoodsServicesReferenceBag   struct {
							gorm.Model
							//Text                   string `xml:",chardata" json:"text,omitempty"`
							GoodsServicesReference []struct {
								gorm.Model
								//Text                             string `xml:",chardata" json:"text,omitempty"`
								GoodsServicesReferenceIdentifier string `xml:"GoodsServicesReferenceIdentifier" json:"goods_services_reference_identifier,omitempty"`
							} `xml:"GoodsServicesReference" json:"goods_services_reference,omitempty"`
						} `xml:"GoodsServicesReferenceBag" json:"goods_services_reference_bag"`
					} `xml:"Claim" json:"claim,omitempty"`
				} `xml:"ClaimBag" json:"claim_bag"`
				MarkFeatureCategoryBag struct {
					gorm.Model
					//Text                string `xml:",chardata" json:"text,omitempty"`
					MarkFeatureCategory string `xml:"MarkFeatureCategory" json:"mark_feature_category,omitempty"`
				} `xml:"MarkFeatureCategoryBag" json:"mark_feature_category_bag"`
			} `xml:"NationalTrademarkInformation" json:"national_trademark_information"`
		} `xml:"Trademark" json:"trademark"`
	} `xml:"TrademarkBag" json:"trademark_bag"`
}
