package component

import (
	"context"
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type AnyType struct {
	InnerXML string `xml:",innerxml"`
}

type AnyURI string

type NCName string

type ComponentType string

const (
	ComponentTypeNone ComponentType = "None"

	ComponentTypeOtherEquipment ComponentType = "OtherEquipment"

	ComponentTypeNetwork ComponentType = "Network"

	ComponentTypeCable ComponentType = "Cable"

	ComponentTypeConnectivityHardware ComponentType = "ConnectivityHardware"

	ComponentTypePort ComponentType = "Port"

	ComponentTypePower ComponentType = "Power"

	ComponentTypeNonItFixedAsset ComponentType = "NonItFixedAsset"

	ComponentTypePeripheral ComponentType = "Peripheral"

	ComponentTypeSpace ComponentType = "Space"

	ComponentTypeSite ComponentType = "Site"

	ComponentTypeBuilding ComponentType = "Building"

	ComponentTypeFloor ComponentType = "Floor"

	ComponentTypeDataCenter ComponentType = "DataCenter"

	ComponentTypeArea ComponentType = "Area"

	ComponentTypeRack ComponentType = "Rack"

	ComponentTypeStoreRoom ComponentType = "StoreRoom"

	ComponentTypeConduit ComponentType = "Conduit"

	ComponentTypePerson ComponentType = "Person"

	ComponentTypePit ComponentType = "Pit"

	ComponentTypeStorage ComponentType = "Storage"

	ComponentTypeServer ComponentType = "Server"

	ComponentTypeMPO ComponentType = "MPO"

	ComponentTypeWorkflowGeneral ComponentType = "WorkflowGeneral"

	ComponentTypeWorkflowInstall ComponentType = "WorkflowInstall"

	ComponentTypeWorkflowDecommission ComponentType = "WorkflowDecommission"

	ComponentTypeWorkflowModify ComponentType = "WorkflowModify"

	ComponentTypeWorkflowRename ComponentType = "WorkflowRename"

	ComponentTypeWorkflowMove ComponentType = "WorkflowMove"

	ComponentTypeWorkflowIncident ComponentType = "WorkflowIncident"

	ComponentTypeWorkflowOther ComponentType = "WorkflowOther"

	ComponentTypeWorkflowAdmin ComponentType = "WorkflowAdmin"
)

type ZeroRackUnit string

const (
	ZeroRackUnitNone ZeroRackUnit = "None"

	ZeroRackUnitLeft ZeroRackUnit = "Left"

	ZeroRackUnitRight ZeroRackUnit = "Right"

	ZeroRackUnitTop ZeroRackUnit = "Top"

	ZeroRackUnitBottom ZeroRackUnit = "Bottom"
)

type MountingType string

const (
	MountingTypeNone MountingType = "None"

	MountingTypeRackFrontAndBack MountingType = "RackFrontAndBack"

	MountingTypeRackFront MountingType = "RackFront"

	MountingTypeRackBack MountingType = "RackBack"

	MountingTypeFloorStanding MountingType = "FloorStanding"

	MountingTypeWallMounted MountingType = "WallMounted"

	MountingTypeDesktop MountingType = "Desktop"

	MountingTypeBlade MountingType = "Blade"

	MountingTypeRackFrontAndBackReversed MountingType = "RackFrontAndBackReversed"
)

type ConnectionType string

const (
	ConnectionTypeNone ConnectionType = "None"

	ConnectionTypeNetworkCopper ConnectionType = "NetworkCopper"

	ConnectionTypeNetworkFibreMM ConnectionType = "NetworkFibreMM"

	ConnectionTypeNetworkFibreSM ConnectionType = "NetworkFibreSM"

	ConnectionTypeSAN ConnectionType = "SAN"

	ConnectionTypePower ConnectionType = "Power"

	ConnectionTypeOther ConnectionType = "Other"

	ConnectionTypeNetworkCopperPOE ConnectionType = "NetworkCopperPOE"

	ConnectionTypeVoice ConnectionType = "Voice"

	ConnectionTypeVoiceAnalog ConnectionType = "VoiceAnalog"
)

type PowerActualDerivation string

const (
	PowerActualDerivationInstantaneous PowerActualDerivation = "Instantaneous"

	PowerActualDerivationPeak PowerActualDerivation = "Peak"

	PowerActualDerivationAverage PowerActualDerivation = "Average"
)

type NumberingOrigin string

const (
	NumberingOriginFromBottom NumberingOrigin = "FromBottom"

	NumberingOriginFromTop NumberingOrigin = "FromTop"
)

type RoleAccess string

const (
	RoleAccessNone RoleAccess = "None"

	RoleAccessView RoleAccess = "View"

	RoleAccessModify RoleAccess = "Modify"

	RoleAccessFull RoleAccess = "Full"

	RoleAccessAdmin RoleAccess = "Admin"
)

type AttributeDataType string

const (
	AttributeDataTypeNone AttributeDataType = "None"

	AttributeDataTypeAlpha AttributeDataType = "Alpha"

	AttributeDataTypeNumeric AttributeDataType = "Numeric"

	AttributeDataTypeAlphaNumeric AttributeDataType = "AlphaNumeric"

	AttributeDataTypeSelection AttributeDataType = "Selection"

	AttributeDataTypeDate AttributeDataType = "Date"

	AttributeDataTypeLookup AttributeDataType = "Lookup"

	AttributeDataTypeUrl AttributeDataType = "Url"
)

type CreateComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponent"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CreateComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentResponse"`

	CreateComponentResult *Component `xml:"CreateComponentResult,omitempty" json:"CreateComponentResult,omitempty"`
}

type CreateComponentWithDescriptionAndBarcode struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithDescriptionAndBarcode"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Description string `xml:"description,omitempty" json:"description,omitempty"`

	Barcode string `xml:"barcode,omitempty" json:"barcode,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CreateComponentWithDescriptionAndBarcodeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithDescriptionAndBarcodeResponse"`

	CreateComponentWithDescriptionAndBarcodeResult *Component `xml:"CreateComponentWithDescriptionAndBarcodeResult,omitempty" json:"CreateComponentWithDescriptionAndBarcodeResult,omitempty"`
}

type CreateComponentWithSubComponents struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSubComponents"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CreateComponentWithSubComponentsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSubComponentsResponse"`

	CreateComponentWithSubComponentsResult *Component `xml:"CreateComponentWithSubComponentsResult,omitempty" json:"CreateComponentWithSubComponentsResult,omitempty"`
}

type CreateComponentWithSpecifiedSubcomponents struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSpecifiedSubcomponents"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Description string `xml:"description,omitempty" json:"description,omitempty"`

	Barcode string `xml:"barcode,omitempty" json:"barcode,omitempty"`

	SubcomponentGroups *ArrayOfSubcomponentGroup `xml:"subcomponentGroups,omitempty" json:"subcomponentGroups,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CreateComponentWithSpecifiedSubcomponentsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSpecifiedSubcomponentsResponse"`

	CreateComponentWithSpecifiedSubcomponentsResult *Component `xml:"CreateComponentWithSpecifiedSubcomponentsResult,omitempty" json:"CreateComponentWithSpecifiedSubcomponentsResult,omitempty"`
}

type CreateComponentWithSubComponentsWithDescriptionAndBarcode struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSubComponentsWithDescriptionAndBarcode"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Description string `xml:"description,omitempty" json:"description,omitempty"`

	Barcode string `xml:"barcode,omitempty" json:"barcode,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse"`

	CreateComponentWithSubComponentsWithDescriptionAndBarcodeResult *Component `xml:"CreateComponentWithSubComponentsWithDescriptionAndBarcodeResult,omitempty" json:"CreateComponentWithSubComponentsWithDescriptionAndBarcodeResult,omitempty"`
}

type DeleteComponentByID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi DeleteComponentByID"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type DeleteComponentByIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi DeleteComponentByIDResponse"`
}

type DeleteComponentByComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi DeleteComponentByComponent"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type DeleteComponentByComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi DeleteComponentByComponentResponse"`
}

type GetChildComponents struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponents"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetChildComponentsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponentsResponse"`

	GetChildComponentsResult *ArrayOfComponent `xml:"GetChildComponentsResult,omitempty" json:"GetChildComponentsResult,omitempty"`
}

type GetChildComponentsByNameType struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponentsByNameType"`

	ComponentName string `xml:"componentName,omitempty" json:"componentName,omitempty"`

	ComponentType *ComponentType `xml:"componentType,omitempty" json:"componentType,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetChildComponentsByNameTypeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponentsByNameTypeResponse"`

	GetChildComponentsByNameTypeResult *ArrayOfComponent `xml:"GetChildComponentsByNameTypeResult,omitempty" json:"GetChildComponentsByNameTypeResult,omitempty"`
}

type GetChildComponentsByIDType struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponentsByIDType"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	ComponentType *ComponentType `xml:"componentType,omitempty" json:"componentType,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetChildComponentsByIDTypeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildComponentsByIDTypeResponse"`

	GetChildComponentsByIDTypeResult *ArrayOfComponent `xml:"GetChildComponentsByIDTypeResult,omitempty" json:"GetChildComponentsByIDTypeResult,omitempty"`
}

type GetChildRacks struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildRacks"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`
}

type GetChildRacksResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChildRacksResponse"`

	GetChildRacksResult *ArrayOfComponent `xml:"GetChildRacksResult,omitempty" json:"GetChildRacksResult,omitempty"`
}

type FilterComponentsByAccess struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FilterComponentsByAccess"`

	Ids *ArrayOfInt `xml:"ids,omitempty" json:"ids,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`

	Access *RoleAccess `xml:"access,omitempty" json:"access,omitempty"`
}

type FilterComponentsByAccessResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FilterComponentsByAccessResponse"`

	FilterComponentsByAccessResult *ArrayOfInt `xml:"FilterComponentsByAccessResult,omitempty" json:"FilterComponentsByAccessResult,omitempty"`
}

type FindSpacePath struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FindSpacePath"`

	StartingSpace int32 `xml:"startingSpace,omitempty" json:"startingSpace,omitempty"`

	TargetSpace int32 `xml:"targetSpace,omitempty" json:"targetSpace,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type FindSpacePathResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FindSpacePathResponse"`

	FindSpacePathResult *ArrayOfInt `xml:"FindSpacePathResult,omitempty" json:"FindSpacePathResult,omitempty"`
}

type GetAllSpaces struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllSpaces"`
}

type GetAllSpacesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllSpacesResponse"`

	GetAllSpacesResult *Component `xml:"GetAllSpacesResult,omitempty" json:"GetAllSpacesResult,omitempty"`
}

type GetAllZones struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllZones"`
}

type GetAllZonesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllZonesResponse"`

	GetAllZonesResult *ArrayOfZone `xml:"GetAllZonesResult,omitempty" json:"GetAllZonesResult,omitempty"`
}

type GetChangeHistory struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChangeHistory"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetChangeHistoryResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChangeHistoryResponse"`

	GetChangeHistoryResult *ArrayOfChangeHistory `xml:"GetChangeHistoryResult,omitempty" json:"GetChangeHistoryResult,omitempty"`
}

type GetComponentAttributeValueByAttributeName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeValueByAttributeName"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetComponentAttributeValueByAttributeNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeValueByAttributeNameResponse"`

	GetComponentAttributeValueByAttributeNameResult string `xml:"GetComponentAttributeValueByAttributeNameResult,omitempty" json:"GetComponentAttributeValueByAttributeNameResult,omitempty"`
}

type GetComponentAttributeValueByComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeValueByComponent"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetComponentAttributeValueByComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeValueByComponentResponse"`

	GetComponentAttributeValueByComponentResult string `xml:"GetComponentAttributeValueByComponentResult,omitempty" json:"GetComponentAttributeValueByComponentResult,omitempty"`
}

type GetParentRack struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetParentRack"`

	ChildID int32 `xml:"childID,omitempty" json:"childID,omitempty"`
}

type GetParentRackResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetParentRackResponse"`

	GetParentRackResult *Component `xml:"GetParentRackResult,omitempty" json:"GetParentRackResult,omitempty"`
}

type SaveComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi SaveComponent"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type SaveComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi SaveComponentResponse"`

	SaveComponentResult bool `xml:"SaveComponentResult,omitempty" json:"SaveComponentResult,omitempty"`
}

type TryGetComponentAttributeValueByAttributeName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryGetComponentAttributeValueByAttributeName"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type TryGetComponentAttributeValueByAttributeNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryGetComponentAttributeValueByAttributeNameResponse"`

	TryGetComponentAttributeValueByAttributeNameResult bool `xml:"TryGetComponentAttributeValueByAttributeNameResult,omitempty" json:"TryGetComponentAttributeValueByAttributeNameResult,omitempty"`

	Result string `xml:"result,omitempty" json:"result,omitempty"`
}

type TryGetComponentAttributeValueByComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryGetComponentAttributeValueByComponent"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type TryGetComponentAttributeValueByComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryGetComponentAttributeValueByComponentResponse"`

	TryGetComponentAttributeValueByComponentResult bool `xml:"TryGetComponentAttributeValueByComponentResult,omitempty" json:"TryGetComponentAttributeValueByComponentResult,omitempty"`

	Result string `xml:"result,omitempty" json:"result,omitempty"`
}

type TryUpdateComponentAttributeValueByComponentID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryUpdateComponentAttributeValueByComponentID"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type TryUpdateComponentAttributeValueByComponentIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryUpdateComponentAttributeValueByComponentIDResponse"`

	TryUpdateComponentAttributeValueByComponentIDResult bool `xml:"TryUpdateComponentAttributeValueByComponentIDResult,omitempty" json:"TryUpdateComponentAttributeValueByComponentIDResult,omitempty"`
}

type TryUpdateComponentAttributeValue struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryUpdateComponentAttributeValue"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type TryUpdateComponentAttributeValueResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi TryUpdateComponentAttributeValueResponse"`

	TryUpdateComponentAttributeValueResult bool `xml:"TryUpdateComponentAttributeValueResult,omitempty" json:"TryUpdateComponentAttributeValueResult,omitempty"`
}

type UpdateComponentAttributeValueByComponentID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi UpdateComponentAttributeValueByComponentID"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type UpdateComponentAttributeValueByComponentIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi UpdateComponentAttributeValueByComponentIDResponse"`
}

type UpdateComponentAttributeValueByComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi UpdateComponentAttributeValueByComponent"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type UpdateComponentAttributeValueByComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi UpdateComponentAttributeValueByComponentResponse"`
}

type MoveToSpace struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi MoveToSpace"`

	Component *Component `xml:"component,omitempty" json:"component,omitempty"`

	Destination *Component `xml:"destination,omitempty" json:"destination,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type MoveToSpaceResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi MoveToSpaceResponse"`
}

type GetComponentAttributeByComponentID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeByComponentID"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	AttributeId int32 `xml:"attributeId,omitempty" json:"attributeId,omitempty"`
}

type GetComponentAttributeByComponentIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributeByComponentIDResponse"`

	GetComponentAttributeByComponentIDResult *ComponentAttribute `xml:"GetComponentAttributeByComponentIDResult,omitempty" json:"GetComponentAttributeByComponentIDResult,omitempty"`
}

type GetComponentAttributesByComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributesByComponent"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`

	Access *RoleAccess `xml:"access,omitempty" json:"access,omitempty"`
}

type GetComponentAttributesByComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentAttributesByComponentResponse"`

	GetComponentAttributesByComponentResult *ArrayOfComponentAttribute `xml:"GetComponentAttributesByComponentResult,omitempty" json:"GetComponentAttributesByComponentResult,omitempty"`
}

type GetComponentByID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentByID"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`
}

type GetComponentByIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentByIDResponse"`

	GetComponentByIDResult *Component `xml:"GetComponentByIDResult,omitempty" json:"GetComponentByIDResult,omitempty"`
}

type GetComponentIDsByTemplate struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentIDsByTemplate"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`
}

type GetComponentIDsByTemplateResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentIDsByTemplateResponse"`

	GetComponentIDsByTemplateResult *ArrayOfInt `xml:"GetComponentIDsByTemplateResult,omitempty" json:"GetComponentIDsByTemplateResult,omitempty"`
}

type GetComponentsByAttributeName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByAttributeName"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`
}

type GetComponentsByAttributeNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByAttributeNameResponse"`

	GetComponentsByAttributeNameResult *ArrayOfComponent `xml:"GetComponentsByAttributeNameResult,omitempty" json:"GetComponentsByAttributeNameResult,omitempty"`
}

type GetComponentsByAttributeNameAndValue struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByAttributeNameAndValue"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`
}

type GetComponentsByAttributeNameAndValueResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByAttributeNameAndValueResponse"`

	GetComponentsByAttributeNameAndValueResult *ArrayOfComponent `xml:"GetComponentsByAttributeNameAndValueResult,omitempty" json:"GetComponentsByAttributeNameAndValueResult,omitempty"`
}

type GetComponentsByBarcode struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByBarcode"`

	Barcode string `xml:"barcode,omitempty" json:"barcode,omitempty"`
}

type GetComponentsByBarcodeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByBarcodeResponse"`

	GetComponentsByBarcodeResult *ArrayOfComponent `xml:"GetComponentsByBarcodeResult,omitempty" json:"GetComponentsByBarcodeResult,omitempty"`
}

type GetComponentsByLevel struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByLevel"`

	RooComponent int32 `xml:"rooComponent,omitempty" json:"rooComponent,omitempty"`

	Level int32 `xml:"level,omitempty" json:"level,omitempty"`

	IncludeCable bool `xml:"includeCable,omitempty" json:"includeCable,omitempty"`

	IncludePorts bool `xml:"includePorts,omitempty" json:"includePorts,omitempty"`
}

type GetComponentsByLevelResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByLevelResponse"`

	GetComponentsByLevelResult *ArrayOfComponent `xml:"GetComponentsByLevelResult,omitempty" json:"GetComponentsByLevelResult,omitempty"`
}

type GetComponentsByName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByName"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

type GetComponentsByNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByNameResponse"`

	GetComponentsByNameResult *ArrayOfComponent `xml:"GetComponentsByNameResult,omitempty" json:"GetComponentsByNameResult,omitempty"`
}

type GetComponentsInStoreRoomBySite struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsInStoreRoomBySite"`

	SiteId int32 `xml:"siteId,omitempty" json:"siteId,omitempty"`
}

type GetComponentsInStoreRoomBySiteResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsInStoreRoomBySiteResponse"`

	GetComponentsInStoreRoomBySiteResult *ArrayOfComponent `xml:"GetComponentsInStoreRoomBySiteResult,omitempty" json:"GetComponentsInStoreRoomBySiteResult,omitempty"`
}

type GetComponentsByTemplateId struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByTemplateId"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`
}

type GetComponentsByTemplateIdResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentsByTemplateIdResponse"`

	GetComponentsByTemplateIdResult *ArrayOfComponent `xml:"GetComponentsByTemplateIdResult,omitempty" json:"GetComponentsByTemplateIdResult,omitempty"`
}

type FilterComponentsByParentId struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FilterComponentsByParentId"`

	ParentId int32 `xml:"parentId,omitempty" json:"parentId,omitempty"`

	Components *ArrayOfComponent `xml:"components,omitempty" json:"components,omitempty"`
}

type FilterComponentsByParentIdResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi FilterComponentsByParentIdResponse"`

	FilterComponentsByParentIdResult *ArrayOfComponent `xml:"FilterComponentsByParentIdResult,omitempty" json:"FilterComponentsByParentIdResult,omitempty"`
}

type GetCompleteComponentsByTemplateIdAttributeNameAndValue struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByTemplateIdAttributeNameAndValue"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse"`

	GetCompleteComponentsByTemplateIdAttributeNameAndValueResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByTemplateIdAttributeNameAndValueResult,omitempty" json:"GetCompleteComponentsByTemplateIdAttributeNameAndValueResult,omitempty"`
}

type GetCompleteComponentsByAttributeNameAndValue struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByAttributeNameAndValue"`

	AttributeName string `xml:"attributeName,omitempty" json:"attributeName,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByAttributeNameAndValueResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByAttributeNameAndValueResponse"`

	GetCompleteComponentsByAttributeNameAndValueResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByAttributeNameAndValueResult,omitempty" json:"GetCompleteComponentsByAttributeNameAndValueResult,omitempty"`
}

type GetCompleteComponentsByTemplate struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByTemplate"`

	TemplateId int32 `xml:"templateId,omitempty" json:"templateId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByTemplateResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByTemplateResponse"`

	GetCompleteComponentsByTemplateResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByTemplateResult,omitempty" json:"GetCompleteComponentsByTemplateResult,omitempty"`
}

type GetCompleteComponentByID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentByID"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentByIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentByIDResponse"`

	GetCompleteComponentByIDResult *CompleteComponent `xml:"GetCompleteComponentByIDResult,omitempty" json:"GetCompleteComponentByIDResult,omitempty"`
}

type GetCompleteComponentsByIDs struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByIDs"`

	Ids *ArrayOfInt `xml:"ids,omitempty" json:"ids,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByIDsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByIDsResponse"`

	GetCompleteComponentsByIDsResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByIDsResult,omitempty" json:"GetCompleteComponentsByIDsResult,omitempty"`
}

type GetCompleteComponentsByName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByName"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByNameResponse"`

	GetCompleteComponentsByNameResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByNameResult,omitempty" json:"GetCompleteComponentsByNameResult,omitempty"`
}

type GetCompleteComponentsByBarcode struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByBarcode"`

	Barcode string `xml:"barcode,omitempty" json:"barcode,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentsByBarcodeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentsByBarcodeResponse"`

	GetCompleteComponentsByBarcodeResult *ArrayOfCompleteComponent `xml:"GetCompleteComponentsByBarcodeResult,omitempty" json:"GetCompleteComponentsByBarcodeResult,omitempty"`
}

type GetCompleteComponentSpaceByAttributeIdentifier struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentSpaceByAttributeIdentifier"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	AttributeId int32 `xml:"attributeId,omitempty" json:"attributeId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetCompleteComponentSpaceByAttributeIdentifierResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetCompleteComponentSpaceByAttributeIdentifierResponse"`

	GetCompleteComponentSpaceByAttributeIdentifierResult *CompleteComponent `xml:"GetCompleteComponentSpaceByAttributeIdentifierResult,omitempty" json:"GetCompleteComponentSpaceByAttributeIdentifierResult,omitempty"`
}

type GetComponentSpaceByAttributeIdentifier struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentSpaceByAttributeIdentifier"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	AttributeId int32 `xml:"attributeId,omitempty" json:"attributeId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetComponentSpaceByAttributeIdentifierResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentSpaceByAttributeIdentifierResponse"`

	GetComponentSpaceByAttributeIdentifierResult *Component `xml:"GetComponentSpaceByAttributeIdentifierResult,omitempty" json:"GetComponentSpaceByAttributeIdentifierResult,omitempty"`
}

type GetComponentChildren struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentChildren"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`

	IncludePorts bool `xml:"includePorts,omitempty" json:"includePorts,omitempty"`

	IncludeCables bool `xml:"includeCables,omitempty" json:"includeCables,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetComponentChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentChildrenResponse"`

	GetComponentChildrenResult *ArrayOfComponent `xml:"GetComponentChildrenResult,omitempty" json:"GetComponentChildrenResult,omitempty"`
}

type GetComponentPorts struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentPorts"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`
}

type GetComponentPortsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentPortsResponse"`

	GetComponentPortsResult *ArrayOfComponent `xml:"GetComponentPortsResult,omitempty" json:"GetComponentPortsResult,omitempty"`
}

type GetZone struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetZone"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`
}

type GetZoneResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetZoneResponse"`

	GetZoneResult *Zone `xml:"GetZoneResult,omitempty" json:"GetZoneResult,omitempty"`
}

type SaveComponentAttributes struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi SaveComponentAttributes"`

	ComponentAttributes *ArrayOfComponentAttribute `xml:"componentAttributes,omitempty" json:"componentAttributes,omitempty"`

	IsSaveSystemFields bool `xml:"isSaveSystemFields,omitempty" json:"isSaveSystemFields,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type SaveComponentAttributesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi SaveComponentAttributesResponse"`

	SaveComponentAttributesResult *ArrayOfComponentAttributeSaveResult `xml:"SaveComponentAttributesResult,omitempty" json:"SaveComponentAttributesResult,omitempty"`
}

type RenameComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi RenameComponent"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	NewName string `xml:"newName,omitempty" json:"newName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type RenameComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi RenameComponentResponse"`
}

type CopyComponent struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CopyComponent"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	SpaceId int32 `xml:"spaceId,omitempty" json:"spaceId,omitempty"`

	NewName string `xml:"newName,omitempty" json:"newName,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type CopyComponentResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi CopyComponentResponse"`

	CopyComponentResult int32 `xml:"CopyComponentResult,omitempty" json:"CopyComponentResult,omitempty"`
}

type IsSubcomponentChangeRequired struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi IsSubcomponentChangeRequired"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	NewTemplateId int32 `xml:"newTemplateId,omitempty" json:"newTemplateId,omitempty"`
}

type IsSubcomponentChangeRequiredResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi IsSubcomponentChangeRequiredResponse"`

	IsSubcomponentChangeRequiredResult bool `xml:"IsSubcomponentChangeRequiredResult,omitempty" json:"IsSubcomponentChangeRequiredResult,omitempty"`
}

type GetComponentRoleAccess struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentRoleAccess"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`

	Password string `xml:"password,omitempty" json:"password,omitempty"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`
}

type GetComponentRoleAccessResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetComponentRoleAccessResponse"`

	GetComponentRoleAccessResult *RoleAccess `xml:"GetComponentRoleAccessResult,omitempty" json:"GetComponentRoleAccessResult,omitempty"`
}

type GetPortHostingComponentsByID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetPortHostingComponentsByID"`

	PortHostId int32 `xml:"portHostId,omitempty" json:"portHostId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type GetPortHostingComponentsByIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetPortHostingComponentsByIDResponse"`

	GetPortHostingComponentsByIDResult *ArrayOfCompleteComponent `xml:"GetPortHostingComponentsByIDResult,omitempty" json:"GetPortHostingComponentsByIDResult,omitempty"`
}

type GetChassisInDataCenterByBladeTemplate struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChassisInDataCenterByBladeTemplate"`

	DataCenterId int32 `xml:"dataCenterId,omitempty" json:"dataCenterId,omitempty"`

	BladeTemplateId int32 `xml:"bladeTemplateId,omitempty" json:"bladeTemplateId,omitempty"`
}

type GetChassisInDataCenterByBladeTemplateResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetChassisInDataCenterByBladeTemplateResponse"`

	GetChassisInDataCenterByBladeTemplateResult *ArrayOfComponent `xml:"GetChassisInDataCenterByBladeTemplateResult,omitempty" json:"GetChassisInDataCenterByBladeTemplateResult,omitempty"`
}

type GetAvailableSlotNumbersForChassis struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAvailableSlotNumbersForChassis"`

	ChassisComponentId int32 `xml:"chassisComponentId,omitempty" json:"chassisComponentId,omitempty"`

	BladeTemplateId int32 `xml:"bladeTemplateId,omitempty" json:"bladeTemplateId,omitempty"`
}

type GetAvailableSlotNumbersForChassisResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAvailableSlotNumbersForChassisResponse"`

	GetAvailableSlotNumbersForChassisResult *ArrayOfInt `xml:"GetAvailableSlotNumbersForChassisResult,omitempty" json:"GetAvailableSlotNumbersForChassisResult,omitempty"`
}

type EnsureSubcomponentDefaultsExist struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi EnsureSubcomponentDefaultsExist"`

	ComponentId int32 `xml:"componentId,omitempty" json:"componentId,omitempty"`

	Username string `xml:"username,omitempty" json:"username,omitempty"`
}

type EnsureSubcomponentDefaultsExistResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi EnsureSubcomponentDefaultsExistResponse"`
}

type Component struct {
	*Dto

	AppCat string `xml:"AppCat,omitempty" json:"AppCat,omitempty"`

	AuditorID string `xml:"AuditorID,omitempty" json:"AuditorID,omitempty"`

	Barcode string `xml:"Barcode,omitempty" json:"Barcode,omitempty"`

	ChannelID int32 `xml:"ChannelID,omitempty" json:"ChannelID,omitempty"`

	Type *ComponentType `xml:"Type,omitempty" json:"Type,omitempty"`

	Connectivity byte `xml:"Connectivity,omitempty" json:"Connectivity,omitempty"`

	DeratingFactor float64 `xml:"DeratingFactor,omitempty" json:"DeratingFactor,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	Detail byte `xml:"Detail,omitempty" json:"Detail,omitempty"`

	HasChild int32 `xml:"HasChild,omitempty" json:"HasChild,omitempty"`

	Height float64 `xml:"Height,omitempty" json:"Height,omitempty"`

	IconID int32 `xml:"IconID,omitempty" json:"IconID,omitempty"`

	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	IsAudit bool `xml:"IsAudit,omitempty" json:"IsAudit,omitempty"`

	LastAuditDate soap.XSDDateTime `xml:"LastAuditDate,omitempty" json:"LastAuditDate,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Pairs int32 `xml:"Pairs,omitempty" json:"Pairs,omitempty"`

	SeqNum int32 `xml:"SeqNum,omitempty" json:"SeqNum,omitempty"`

	SpaceID int32 `xml:"SpaceID,omitempty" json:"SpaceID,omitempty"`

	SpaceIconID int32 `xml:"SpaceIconID,omitempty" json:"SpaceIconID,omitempty"`

	SpaceName string `xml:"SpaceName,omitempty" json:"SpaceName,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty" json:"TemplateID,omitempty"`

	TemplateName string `xml:"TemplateName,omitempty" json:"TemplateName,omitempty"`

	Length float64 `xml:"Length,omitempty" json:"Length,omitempty"`

	Status string `xml:"Status,omitempty" json:"Status,omitempty"`

	XPos int32 `xml:"XPos,omitempty" json:"XPos,omitempty"`

	YPos int32 `xml:"YPos,omitempty" json:"YPos,omitempty"`

	ZPos int32 `xml:"ZPos,omitempty" json:"ZPos,omitempty"`

	Width float64 `xml:"Width,omitempty" json:"Width,omitempty"`

	SiteID int32 `xml:"SiteID,omitempty" json:"SiteID,omitempty"`

	BuildingID int32 `xml:"BuildingID,omitempty" json:"BuildingID,omitempty"`

	FloorID int32 `xml:"FloorID,omitempty" json:"FloorID,omitempty"`

	DataCenterID int32 `xml:"DataCenterID,omitempty" json:"DataCenterID,omitempty"`

	AreaID int32 `xml:"AreaID,omitempty" json:"AreaID,omitempty"`

	StoreRoomID int32 `xml:"StoreRoomID,omitempty" json:"StoreRoomID,omitempty"`

	RackID int32 `xml:"RackID,omitempty" json:"RackID,omitempty"`

	ServiceID int32 `xml:"ServiceID,omitempty" json:"ServiceID,omitempty"`

	UnitHeight float64 `xml:"UnitHeight,omitempty" json:"UnitHeight,omitempty"`

	EquipmentWidth int32 `xml:"EquipmentWidth,omitempty" json:"EquipmentWidth,omitempty"`

	UnitLocation float64 `xml:"UnitLocation,omitempty" json:"UnitLocation,omitempty"`

	HorizontalOffset int32 `xml:"HorizontalOffset,omitempty" json:"HorizontalOffset,omitempty"`

	ZeroRackUnit *ZeroRackUnit `xml:"ZeroRackUnit,omitempty" json:"ZeroRackUnit,omitempty"`

	MountingType *MountingType `xml:"MountingType,omitempty" json:"MountingType,omitempty"`

	SlotNumber int32 `xml:"SlotNumber,omitempty" json:"SlotNumber,omitempty"`

	SlotSize int32 `xml:"SlotSize,omitempty" json:"SlotSize,omitempty"`

	GridReference string `xml:"GridReference,omitempty" json:"GridReference,omitempty"`

	ConnectionType *ConnectionType `xml:"ConnectionType,omitempty" json:"ConnectionType,omitempty"`

	PowerPlate float64 `xml:"PowerPlate,omitempty" json:"PowerPlate,omitempty"`

	PowerActual float64 `xml:"PowerActual,omitempty" json:"PowerActual,omitempty"`

	PowerActualDerivation *PowerActualDerivation `xml:"PowerActualDerivation,omitempty" json:"PowerActualDerivation,omitempty"`

	HeatDissipationPlate float64 `xml:"HeatDissipationPlate,omitempty" json:"HeatDissipationPlate,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty"`

	IpAddress string `xml:"IpAddress,omitempty" json:"IpAddress,omitempty"`

	ZoneID int32 `xml:"ZoneID,omitempty" json:"ZoneID,omitempty"`

	RackUnitHeight float64 `xml:"RackUnitHeight,omitempty" json:"RackUnitHeight,omitempty"`

	RackWidth int32 `xml:"RackWidth,omitempty" json:"RackWidth,omitempty"`

	RackOrientation int32 `xml:"RackOrientation,omitempty" json:"RackOrientation,omitempty"`

	NumberingOrigin *NumberingOrigin `xml:"NumberingOrigin,omitempty" json:"NumberingOrigin,omitempty"`

	OverlappingAllowed bool `xml:"OverlappingAllowed,omitempty" json:"OverlappingAllowed,omitempty"`

	PatchingGroup int32 `xml:"PatchingGroup,omitempty" json:"PatchingGroup,omitempty"`

	CoolingMax float64 `xml:"CoolingMax,omitempty" json:"CoolingMax,omitempty"`

	WeightMax float64 `xml:"WeightMax,omitempty" json:"WeightMax,omitempty"`

	PowerMax float64 `xml:"PowerMax,omitempty" json:"PowerMax,omitempty"`
}

type Dto struct {
}

type ArrayOfSubcomponentGroup struct {
	SubcomponentGroup []*SubcomponentGroup `xml:"SubcomponentGroup,omitempty" json:"SubcomponentGroup,omitempty"`
}

type SubcomponentGroup struct {
	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Min int32 `xml:"Min,omitempty" json:"Min,omitempty"`

	Max int32 `xml:"Max,omitempty" json:"Max,omitempty"`

	Count int32 `xml:"Count,omitempty" json:"Count,omitempty"`

	Wizard bool `xml:"Wizard,omitempty" json:"Wizard,omitempty"`

	Inherit bool `xml:"Inherit,omitempty" json:"Inherit,omitempty"`

	Restart bool `xml:"Restart,omitempty" json:"Restart,omitempty"`

	Separator string `xml:"Separator,omitempty" json:"Separator,omitempty"`

	TemplateID int32 `xml:"TemplateID,omitempty" json:"TemplateID,omitempty"`

	Subcomponents *ArrayOfSubcomponent `xml:"Subcomponents,omitempty" json:"Subcomponents,omitempty"`
}

type ArrayOfSubcomponent struct {
	Subcomponent []*Subcomponent `xml:"Subcomponent,omitempty" json:"Subcomponent,omitempty"`
}

type Subcomponent struct {
	Template *Template `xml:"Template,omitempty" json:"Template,omitempty"`

	Default int32 `xml:"Default,omitempty" json:"Default,omitempty"`

	Count int32 `xml:"Count,omitempty" json:"Count,omitempty"`

	OriginalCount int32 `xml:"OriginalCount,omitempty" json:"OriginalCount,omitempty"`

	SubcomponentGroupId int32 `xml:"SubcomponentGroupId,omitempty" json:"SubcomponentGroupId,omitempty"`
}

type Template struct {
	*TemplateBase

	AppCat string `xml:"AppCat,omitempty" json:"AppCat,omitempty"`

	Audit bool `xml:"Audit,omitempty" json:"Audit,omitempty"`

	BarcodeLink bool `xml:"BarcodeLink,omitempty" json:"BarcodeLink,omitempty"`

	Connectivity byte `xml:"Connectivity,omitempty" json:"Connectivity,omitempty"`

	DeratingFactor float64 `xml:"DeratingFactor,omitempty" json:"DeratingFactor,omitempty"`

	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	Detail byte `xml:"Detail,omitempty" json:"Detail,omitempty"`

	Type *ComponentType `xml:"Type,omitempty" json:"Type,omitempty"`

	Icon int32 `xml:"Icon,omitempty" json:"Icon,omitempty"`

	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	Length float64 `xml:"Length,omitempty" json:"Length,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Pairs int32 `xml:"Pairs,omitempty" json:"Pairs,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty" json:"ParentID,omitempty"`

	PrimaryImg int32 `xml:"PrimaryImg,omitempty" json:"PrimaryImg,omitempty"`

	SecondaryImg int32 `xml:"SecondaryImg,omitempty" json:"SecondaryImg,omitempty"`

	Status string `xml:"Status,omitempty" json:"Status,omitempty"`

	UnitHeight float64 `xml:"UnitHeight,omitempty" json:"UnitHeight,omitempty"`

	EquipmentWidth int32 `xml:"EquipmentWidth,omitempty" json:"EquipmentWidth,omitempty"`

	MountingType *MountingType `xml:"MountingType,omitempty" json:"MountingType,omitempty"`

	SlotSize int32 `xml:"SlotSize,omitempty" json:"SlotSize,omitempty"`

	ZeroRackUnit *ZeroRackUnit `xml:"ZeroRackUnit,omitempty" json:"ZeroRackUnit,omitempty"`

	ConnectionType *ConnectionType `xml:"ConnectionType,omitempty" json:"ConnectionType,omitempty"`

	PowerPlate float64 `xml:"PowerPlate,omitempty" json:"PowerPlate,omitempty"`

	PowerActualDerivation *PowerActualDerivation `xml:"PowerActualDerivation,omitempty" json:"PowerActualDerivation,omitempty"`

	HeatDissipationPlate float64 `xml:"HeatDissipationPlate,omitempty" json:"HeatDissipationPlate,omitempty"`

	Weight float64 `xml:"Weight,omitempty" json:"Weight,omitempty"`

	RackUnitHeight float64 `xml:"RackUnitHeight,omitempty" json:"RackUnitHeight,omitempty"`

	RackWidth int32 `xml:"RackWidth,omitempty" json:"RackWidth,omitempty"`

	RackOrientation int32 `xml:"RackOrientation,omitempty" json:"RackOrientation,omitempty"`

	NumberingOrigin *NumberingOrigin `xml:"NumberingOrigin,omitempty" json:"NumberingOrigin,omitempty"`

	OverlappingAllowed bool `xml:"OverlappingAllowed,omitempty" json:"OverlappingAllowed,omitempty"`

	CoolingMax float64 `xml:"CoolingMax,omitempty" json:"CoolingMax,omitempty"`

	WeightMax float64 `xml:"WeightMax,omitempty" json:"WeightMax,omitempty"`

	PowerMax float64 `xml:"PowerMax,omitempty" json:"PowerMax,omitempty"`
}

type TemplateBase struct {
}

type ArrayOfComponent struct {
	Component []*Component `xml:"Component,omitempty" json:"Component,omitempty"`
}

type ArrayOfInt struct {
	int []int32 `xml:"int,omitempty" json:"int,omitempty"`
}

type ArrayOfZone struct {
	Zone []*Zone `xml:"Zone,omitempty" json:"Zone,omitempty"`
}

type Zone struct {
	Description string `xml:"Description,omitempty" json:"Description,omitempty"`

	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	Space int32 `xml:"Space,omitempty" json:"Space,omitempty"`
}

type ArrayOfChangeHistory struct {
	ChangeHistory []*ChangeHistory `xml:"ChangeHistory,omitempty" json:"ChangeHistory,omitempty"`
}

type ChangeHistory struct {
	ChangeTime soap.XSDDateTime `xml:"ChangeTime,omitempty" json:"ChangeTime,omitempty"`

	Engineer string `xml:"Engineer,omitempty" json:"Engineer,omitempty"`

	From int32 `xml:"From,omitempty" json:"From,omitempty"`

	ID int32 `xml:"ID,omitempty" json:"ID,omitempty"`

	Log string `xml:"Log,omitempty" json:"Log,omitempty"`

	Name string `xml:"Name,omitempty" json:"Name,omitempty"`

	Summary string `xml:"Summary,omitempty" json:"Summary,omitempty"`

	To *int32 `xml:"To,omitempty" json:"To,omitempty"`
}

type ComponentAttribute struct {
	Value string `xml:"Value,omitempty" json:"Value,omitempty"`

	Attribute int32 `xml:"Attribute,omitempty" json:"Attribute,omitempty"`

	AttributeGroup int32 `xml:"AttributeGroup,omitempty" json:"AttributeGroup,omitempty"`

	AttributeGroupName string `xml:"AttributeGroupName,omitempty" json:"AttributeGroupName,omitempty"`

	AttributeName string `xml:"AttributeName,omitempty" json:"AttributeName,omitempty"`

	Component int32 `xml:"Component,omitempty" json:"Component,omitempty"`

	DataType *AttributeDataType `xml:"DataType,omitempty" json:"DataType,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty" json:"DefaultValue,omitempty"`
}

type ArrayOfComponentAttribute struct {
	ComponentAttribute []*ComponentAttribute `xml:"ComponentAttribute,omitempty" json:"ComponentAttribute,omitempty"`
}

type ArrayOfCompleteComponent struct {
	CompleteComponent []*CompleteComponent `xml:"CompleteComponent,omitempty" json:"CompleteComponent,omitempty"`
}

type CompleteComponent struct {
	*Component

	SpaceFlag byte `xml:"SpaceFlag,omitempty" json:"SpaceFlag,omitempty"`

	FullPath string `xml:"FullPath,omitempty" json:"FullPath,omitempty"`

	ComponentAttributes *ArrayOfComponentAttribute `xml:"ComponentAttributes,omitempty" json:"ComponentAttributes,omitempty"`
}

type ArrayOfComponentAttributeSaveResult struct {
	ComponentAttributeSaveResult []*ComponentAttributeSaveResult `xml:"ComponentAttributeSaveResult,omitempty" json:"ComponentAttributeSaveResult,omitempty"`
}

type ComponentAttributeSaveResult struct {
	Attribute int32 `xml:"Attribute,omitempty" json:"Attribute,omitempty"`

	Component int32 `xml:"Component,omitempty" json:"Component,omitempty"`

	Error string `xml:"Error,omitempty" json:"Error,omitempty"`

	Value string `xml:"Value,omitempty" json:"Value,omitempty"`
}

type ComponentServicesSoap interface {

	/* Creates a new component. Exception is thrown if the user does not have full zone permissions for the space. */
	CreateComponent(request *CreateComponent) (*CreateComponentResponse, error)

	CreateComponentContext(ctx context.Context, request *CreateComponent) (*CreateComponentResponse, error)

	/* Creates a new component with description and barcode. Exception is thrown if the user does not have full zone permissions for the space */
	CreateComponentWithDescriptionAndBarcode(request *CreateComponentWithDescriptionAndBarcode) (*CreateComponentWithDescriptionAndBarcodeResponse, error)

	CreateComponentWithDescriptionAndBarcodeContext(ctx context.Context, request *CreateComponentWithDescriptionAndBarcode) (*CreateComponentWithDescriptionAndBarcodeResponse, error)

	/* Creates a new component and its default subcomponents.  Exception is thrown if the user does not have full zone permissions for the space */
	CreateComponentWithSubComponents(request *CreateComponentWithSubComponents) (*CreateComponentWithSubComponentsResponse, error)

	CreateComponentWithSubComponentsContext(ctx context.Context, request *CreateComponentWithSubComponents) (*CreateComponentWithSubComponentsResponse, error)

	/* Creates a new component and specified subcomponents. */
	CreateComponentWithSpecifiedSubcomponents(request *CreateComponentWithSpecifiedSubcomponents) (*CreateComponentWithSpecifiedSubcomponentsResponse, error)

	CreateComponentWithSpecifiedSubcomponentsContext(ctx context.Context, request *CreateComponentWithSpecifiedSubcomponents) (*CreateComponentWithSpecifiedSubcomponentsResponse, error)

	/* Creates a new component and its subcomponents.  Exception is thrown if the user does not have full zone permissions for the space */
	CreateComponentWithSubComponentsWithDescriptionAndBarcode(request *CreateComponentWithSubComponentsWithDescriptionAndBarcode) (*CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse, error)

	CreateComponentWithSubComponentsWithDescriptionAndBarcodeContext(ctx context.Context, request *CreateComponentWithSubComponentsWithDescriptionAndBarcode) (*CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse, error)

	/* Deletes a component.  Exception is thrown if the user does not have full zone permissions for the space. */
	DeleteComponentByID(request *DeleteComponentByID) (*DeleteComponentByIDResponse, error)

	DeleteComponentByIDContext(ctx context.Context, request *DeleteComponentByID) (*DeleteComponentByIDResponse, error)

	/* Deletes a component. Exception is thrown if the user does not have full zone permissions for the space. */
	DeleteComponentByComponent(request *DeleteComponentByComponent) (*DeleteComponentByComponentResponse, error)

	DeleteComponentByComponentContext(ctx context.Context, request *DeleteComponentByComponent) (*DeleteComponentByComponentResponse, error)

	/* Gets a list containing all the children of the specified component. The list only includes the child compononents that the specified user has permissions to view. All cables will be excluded from the returned list. */
	GetChildComponents(request *GetChildComponents) (*GetChildComponentsResponse, error)

	GetChildComponentsContext(ctx context.Context, request *GetChildComponents) (*GetChildComponentsResponse, error)

	/* Gets a list containing all the children of the specified component based on what type of component. The list only includes the child compononents that the specified user has permissions to view. All cables will be excluded from the returned list. */
	GetChildComponentsByNameType(request *GetChildComponentsByNameType) (*GetChildComponentsByNameTypeResponse, error)

	GetChildComponentsByNameTypeContext(ctx context.Context, request *GetChildComponentsByNameType) (*GetChildComponentsByNameTypeResponse, error)

	/* Gets a list containing all the children of the specified component based on what type of component. The list only includes the child compononents that the specified user has permissions to view. All cables will be excluded from the returned list. */
	GetChildComponentsByIDType(request *GetChildComponentsByIDType) (*GetChildComponentsByIDTypeResponse, error)

	GetChildComponentsByIDTypeContext(ctx context.Context, request *GetChildComponentsByIDType) (*GetChildComponentsByIDTypeResponse, error)

	/* Gets the specified component's child racks at all levels. No checking for user permissions. */
	GetChildRacks(request *GetChildRacks) (*GetChildRacksResponse, error)

	GetChildRacksContext(ctx context.Context, request *GetChildRacks) (*GetChildRacksResponse, error)

	/* Filters the specified components. */
	FilterComponentsByAccess(request *FilterComponentsByAccess) (*FilterComponentsByAccessResponse, error)

	FilterComponentsByAccessContext(ctx context.Context, request *FilterComponentsByAccess) (*FilterComponentsByAccessResponse, error)

	/* Finds the path between two components. */
	FindSpacePath(request *FindSpacePath) (*FindSpacePathResponse, error)

	FindSpacePathContext(ctx context.Context, request *FindSpacePath) (*FindSpacePathResponse, error)

	/* Gets the "All Spaces" component. */
	GetAllSpaces(request *GetAllSpaces) (*GetAllSpacesResponse, error)

	GetAllSpacesContext(ctx context.Context, request *GetAllSpaces) (*GetAllSpacesResponse, error)

	/* Gets all zones. */
	GetAllZones(request *GetAllZones) (*GetAllZonesResponse, error)

	GetAllZonesContext(ctx context.Context, request *GetAllZones) (*GetAllZonesResponse, error)

	/* Gets the change history of the specified component. */
	GetChangeHistory(request *GetChangeHistory) (*GetChangeHistoryResponse, error)

	GetChangeHistoryContext(ctx context.Context, request *GetChangeHistory) (*GetChangeHistoryResponse, error)

	/* Gets the value of the specified attribute of the default component. Null is returned if username has no zone or attribute group read permissions.  */
	GetComponentAttributeValueByAttributeName(request *GetComponentAttributeValueByAttributeName) (*GetComponentAttributeValueByAttributeNameResponse, error)

	GetComponentAttributeValueByAttributeNameContext(ctx context.Context, request *GetComponentAttributeValueByAttributeName) (*GetComponentAttributeValueByAttributeNameResponse, error)

	/* Gets a value of a the specified component's attribute.  Null is returned if username has no zone or attribute group read permissions.  */
	GetComponentAttributeValueByComponent(request *GetComponentAttributeValueByComponent) (*GetComponentAttributeValueByComponentResponse, error)

	GetComponentAttributeValueByComponentContext(ctx context.Context, request *GetComponentAttributeValueByComponent) (*GetComponentAttributeValueByComponentResponse, error)

	/* Gets any parent racks for the specified child component. */
	GetParentRack(request *GetParentRack) (*GetParentRackResponse, error)

	GetParentRackContext(ctx context.Context, request *GetParentRack) (*GetParentRackResponse, error)

	/* Saves the changes to the specified component. */
	SaveComponent(request *SaveComponent) (*SaveComponentResponse, error)

	SaveComponentContext(ctx context.Context, request *SaveComponent) (*SaveComponentResponse, error)

	/* Gets the value of the default component's attribute. False returned if username has no zone or attribute group read permissions.  */
	TryGetComponentAttributeValueByAttributeName(request *TryGetComponentAttributeValueByAttributeName) (*TryGetComponentAttributeValueByAttributeNameResponse, error)

	TryGetComponentAttributeValueByAttributeNameContext(ctx context.Context, request *TryGetComponentAttributeValueByAttributeName) (*TryGetComponentAttributeValueByAttributeNameResponse, error)

	/* Gets the value of the specified component's attribute.  False returned if username has no zone or attribute group read permissions. */
	TryGetComponentAttributeValueByComponent(request *TryGetComponentAttributeValueByComponent) (*TryGetComponentAttributeValueByComponentResponse, error)

	TryGetComponentAttributeValueByComponentContext(ctx context.Context, request *TryGetComponentAttributeValueByComponent) (*TryGetComponentAttributeValueByComponentResponse, error)

	/* Updates a component's attribute value. False returned if username has no zone or attribute group modify permissions.  */
	TryUpdateComponentAttributeValueByComponentID(request *TryUpdateComponentAttributeValueByComponentID) (*TryUpdateComponentAttributeValueByComponentIDResponse, error)

	TryUpdateComponentAttributeValueByComponentIDContext(ctx context.Context, request *TryUpdateComponentAttributeValueByComponentID) (*TryUpdateComponentAttributeValueByComponentIDResponse, error)

	/* Updates a component's attribute value. False returned if username has no zone or attribute group modify permissions. */
	TryUpdateComponentAttributeValue(request *TryUpdateComponentAttributeValue) (*TryUpdateComponentAttributeValueResponse, error)

	TryUpdateComponentAttributeValueContext(ctx context.Context, request *TryUpdateComponentAttributeValue) (*TryUpdateComponentAttributeValueResponse, error)

	/* Updates a component's attribute value. An exception is thrown if username has no zone or attribute group modify permissions.  */
	UpdateComponentAttributeValueByComponentID(request *UpdateComponentAttributeValueByComponentID) (*UpdateComponentAttributeValueByComponentIDResponse, error)

	UpdateComponentAttributeValueByComponentIDContext(ctx context.Context, request *UpdateComponentAttributeValueByComponentID) (*UpdateComponentAttributeValueByComponentIDResponse, error)

	/* Updates a component's attribute value. An exception is thrown if username has no zone or attribute group modify permissions. */
	UpdateComponentAttributeValueByComponent(request *UpdateComponentAttributeValueByComponent) (*UpdateComponentAttributeValueByComponentResponse, error)

	UpdateComponentAttributeValueByComponentContext(ctx context.Context, request *UpdateComponentAttributeValueByComponent) (*UpdateComponentAttributeValueByComponentResponse, error)

	/* Moves a component to a space. Exception thrown in no zone full permissions for origin or destination space or if component constraints are violated. Logs change history.  */
	MoveToSpace(request *MoveToSpace) (*MoveToSpaceResponse, error)

	MoveToSpaceContext(ctx context.Context, request *MoveToSpace) (*MoveToSpaceResponse, error)

	/* Gets a component's attribute. */
	GetComponentAttributeByComponentID(request *GetComponentAttributeByComponentID) (*GetComponentAttributeByComponentIDResponse, error)

	GetComponentAttributeByComponentIDContext(ctx context.Context, request *GetComponentAttributeByComponentID) (*GetComponentAttributeByComponentIDResponse, error)

	/* Gets a component's attribute. */
	GetComponentAttributesByComponent(request *GetComponentAttributesByComponent) (*GetComponentAttributesByComponentResponse, error)

	GetComponentAttributesByComponentContext(ctx context.Context, request *GetComponentAttributesByComponent) (*GetComponentAttributesByComponentResponse, error)

	/* Gets a component for the specified ID. */
	GetComponentByID(request *GetComponentByID) (*GetComponentByIDResponse, error)

	GetComponentByIDContext(ctx context.Context, request *GetComponentByID) (*GetComponentByIDResponse, error)

	/* Gets the list of component ids for a specified templateID. */
	GetComponentIDsByTemplate(request *GetComponentIDsByTemplate) (*GetComponentIDsByTemplateResponse, error)

	GetComponentIDsByTemplateContext(ctx context.Context, request *GetComponentIDsByTemplate) (*GetComponentIDsByTemplateResponse, error)

	/* Gets a list of Component having the specified attribute. */
	GetComponentsByAttributeName(request *GetComponentsByAttributeName) (*GetComponentsByAttributeNameResponse, error)

	GetComponentsByAttributeNameContext(ctx context.Context, request *GetComponentsByAttributeName) (*GetComponentsByAttributeNameResponse, error)

	/* Gets a list of Component having the specified attribute. */
	GetComponentsByAttributeNameAndValue(request *GetComponentsByAttributeNameAndValue) (*GetComponentsByAttributeNameAndValueResponse, error)

	GetComponentsByAttributeNameAndValueContext(ctx context.Context, request *GetComponentsByAttributeNameAndValue) (*GetComponentsByAttributeNameAndValueResponse, error)

	/* Gets a component with the matching barcode. */
	GetComponentsByBarcode(request *GetComponentsByBarcode) (*GetComponentsByBarcodeResponse, error)

	GetComponentsByBarcodeContext(ctx context.Context, request *GetComponentsByBarcode) (*GetComponentsByBarcodeResponse, error)

	/* Gets components by level. */
	GetComponentsByLevel(request *GetComponentsByLevel) (*GetComponentsByLevelResponse, error)

	GetComponentsByLevelContext(ctx context.Context, request *GetComponentsByLevel) (*GetComponentsByLevelResponse, error)

	/* Gets all components matching the specified name. */
	GetComponentsByName(request *GetComponentsByName) (*GetComponentsByNameResponse, error)

	GetComponentsByNameContext(ctx context.Context, request *GetComponentsByName) (*GetComponentsByNameResponse, error)

	/* Gets equipment type components from any store room in the site. */
	GetComponentsInStoreRoomBySite(request *GetComponentsInStoreRoomBySite) (*GetComponentsInStoreRoomBySiteResponse, error)

	GetComponentsInStoreRoomBySiteContext(ctx context.Context, request *GetComponentsInStoreRoomBySite) (*GetComponentsInStoreRoomBySiteResponse, error)

	/* Gets all components matching the specified template id. */
	GetComponentsByTemplateId(request *GetComponentsByTemplateId) (*GetComponentsByTemplateIdResponse, error)

	GetComponentsByTemplateIdContext(ctx context.Context, request *GetComponentsByTemplateId) (*GetComponentsByTemplateIdResponse, error)

	/* Filters a list of components by parentID. Checks all parents, not just direct parent. */
	FilterComponentsByParentId(request *FilterComponentsByParentId) (*FilterComponentsByParentIdResponse, error)

	FilterComponentsByParentIdContext(ctx context.Context, request *FilterComponentsByParentId) (*FilterComponentsByParentIdResponse, error)

	/* Gets all complete components matching the specified template id, attribute name and attribute value. */
	GetCompleteComponentsByTemplateIdAttributeNameAndValue(request *GetCompleteComponentsByTemplateIdAttributeNameAndValue) (*GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse, error)

	GetCompleteComponentsByTemplateIdAttributeNameAndValueContext(ctx context.Context, request *GetCompleteComponentsByTemplateIdAttributeNameAndValue) (*GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse, error)

	/* Gets all complete components matching the specified attribute name and attribute value. */
	GetCompleteComponentsByAttributeNameAndValue(request *GetCompleteComponentsByAttributeNameAndValue) (*GetCompleteComponentsByAttributeNameAndValueResponse, error)

	GetCompleteComponentsByAttributeNameAndValueContext(ctx context.Context, request *GetCompleteComponentsByAttributeNameAndValue) (*GetCompleteComponentsByAttributeNameAndValueResponse, error)

	/* Gets all complete components matching the specified template id. */
	GetCompleteComponentsByTemplate(request *GetCompleteComponentsByTemplate) (*GetCompleteComponentsByTemplateResponse, error)

	GetCompleteComponentsByTemplateContext(ctx context.Context, request *GetCompleteComponentsByTemplate) (*GetCompleteComponentsByTemplateResponse, error)

	/* Gets complete component for the specified ID. */
	GetCompleteComponentByID(request *GetCompleteComponentByID) (*GetCompleteComponentByIDResponse, error)

	GetCompleteComponentByIDContext(ctx context.Context, request *GetCompleteComponentByID) (*GetCompleteComponentByIDResponse, error)

	/* Gets complete component list for the specified id list. Sorted according to the component id order. */
	GetCompleteComponentsByIDs(request *GetCompleteComponentsByIDs) (*GetCompleteComponentsByIDsResponse, error)

	GetCompleteComponentsByIDsContext(ctx context.Context, request *GetCompleteComponentsByIDs) (*GetCompleteComponentsByIDsResponse, error)

	/* Gets all complete components matching the specified name. */
	GetCompleteComponentsByName(request *GetCompleteComponentsByName) (*GetCompleteComponentsByNameResponse, error)

	GetCompleteComponentsByNameContext(ctx context.Context, request *GetCompleteComponentsByName) (*GetCompleteComponentsByNameResponse, error)

	/* Gets all complete components matching the specified barcode. */
	GetCompleteComponentsByBarcode(request *GetCompleteComponentsByBarcode) (*GetCompleteComponentsByBarcodeResponse, error)

	GetCompleteComponentsByBarcodeContext(ctx context.Context, request *GetCompleteComponentsByBarcode) (*GetCompleteComponentsByBarcodeResponse, error)

	/* Gets all complete components matching the specified component id and attribute id. */
	GetCompleteComponentSpaceByAttributeIdentifier(request *GetCompleteComponentSpaceByAttributeIdentifier) (*GetCompleteComponentSpaceByAttributeIdentifierResponse, error)

	GetCompleteComponentSpaceByAttributeIdentifierContext(ctx context.Context, request *GetCompleteComponentSpaceByAttributeIdentifier) (*GetCompleteComponentSpaceByAttributeIdentifierResponse, error)

	/* Gets all components matching the specified component id and attribute id. */
	GetComponentSpaceByAttributeIdentifier(request *GetComponentSpaceByAttributeIdentifier) (*GetComponentSpaceByAttributeIdentifierResponse, error)

	GetComponentSpaceByAttributeIdentifierContext(ctx context.Context, request *GetComponentSpaceByAttributeIdentifier) (*GetComponentSpaceByAttributeIdentifierResponse, error)

	/* Gets the specified component's children. */
	GetComponentChildren(request *GetComponentChildren) (*GetComponentChildrenResponse, error)

	GetComponentChildrenContext(ctx context.Context, request *GetComponentChildren) (*GetComponentChildrenResponse, error)

	/* Gets the specified component's ports. */
	GetComponentPorts(request *GetComponentPorts) (*GetComponentPortsResponse, error)

	GetComponentPortsContext(ctx context.Context, request *GetComponentPorts) (*GetComponentPortsResponse, error)

	/* Gets a zone for the specified ID. */
	GetZone(request *GetZone) (*GetZoneResponse, error)

	GetZoneContext(ctx context.Context, request *GetZone) (*GetZoneResponse, error)

	/* Saves the specified component attributes. */
	SaveComponentAttributes(request *SaveComponentAttributes) (*SaveComponentAttributesResponse, error)

	SaveComponentAttributesContext(ctx context.Context, request *SaveComponentAttributes) (*SaveComponentAttributesResponse, error)

	/* Renames the component to have the new name. Checks permissions, renames subcomponents, logs history, supports components with + in the name. Throws exception if rename problems or permission problems */
	RenameComponent(request *RenameComponent) (*RenameComponentResponse, error)

	RenameComponentContext(ctx context.Context, request *RenameComponent) (*RenameComponentResponse, error)

	/* Create copy of a component (and its children) into a space. */
	CopyComponent(request *CopyComponent) (*CopyComponentResponse, error)

	CopyComponentContext(ctx context.Context, request *CopyComponent) (*CopyComponentResponse, error)

	/* Check whether a component's subcomponents need to be modified if the template changes. */
	IsSubcomponentChangeRequired(request *IsSubcomponentChangeRequired) (*IsSubcomponentChangeRequiredResponse, error)

	IsSubcomponentChangeRequiredContext(ctx context.Context, request *IsSubcomponentChangeRequired) (*IsSubcomponentChangeRequiredResponse, error)

	/* Gets the highest role access permitted to the component for the requesting user. */
	GetComponentRoleAccess(request *GetComponentRoleAccess) (*GetComponentRoleAccessResponse, error)

	GetComponentRoleAccessContext(ctx context.Context, request *GetComponentRoleAccess) (*GetComponentRoleAccessResponse, error)

	/* Gets an array of complete components:  Hosting server, hosting rack item, and hosting rack. */
	GetPortHostingComponentsByID(request *GetPortHostingComponentsByID) (*GetPortHostingComponentsByIDResponse, error)

	GetPortHostingComponentsByIDContext(ctx context.Context, request *GetPortHostingComponentsByID) (*GetPortHostingComponentsByIDResponse, error)

	/* Get a list of chassis within a data center by blade template. */
	GetChassisInDataCenterByBladeTemplate(request *GetChassisInDataCenterByBladeTemplate) (*GetChassisInDataCenterByBladeTemplateResponse, error)

	GetChassisInDataCenterByBladeTemplateContext(ctx context.Context, request *GetChassisInDataCenterByBladeTemplate) (*GetChassisInDataCenterByBladeTemplateResponse, error)

	/* Get a list of available slot numbers in a chassis (starting at 1), taking into account already occupied slots and the size of the blade to be added. */
	GetAvailableSlotNumbersForChassis(request *GetAvailableSlotNumbersForChassis) (*GetAvailableSlotNumbersForChassisResponse, error)

	GetAvailableSlotNumbersForChassisContext(ctx context.Context, request *GetAvailableSlotNumbersForChassis) (*GetAvailableSlotNumbersForChassisResponse, error)

	/* Ensures that all default subcomponents of the component exist. If some subcomponents already exist with
	   the default names, these will be left alone. Will not delete any subcomponents.
	   This should not be used unless you know that the component's subcomponents are missing, and you
	   want to make sure at least the default counts exist. */
	EnsureSubcomponentDefaultsExist(request *EnsureSubcomponentDefaultsExist) (*EnsureSubcomponentDefaultsExistResponse, error)

	EnsureSubcomponentDefaultsExistContext(ctx context.Context, request *EnsureSubcomponentDefaultsExist) (*EnsureSubcomponentDefaultsExistResponse, error)
}

type componentServicesSoap struct {
	client *soap.Client
}

func NewComponentServicesSoap(client *soap.Client) ComponentServicesSoap {
	return &componentServicesSoap{
		client: client,
	}
}

func (service *componentServicesSoap) CreateComponentContext(ctx context.Context, request *CreateComponent) (*CreateComponentResponse, error) {
	response := new(CreateComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CreateComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CreateComponent(request *CreateComponent) (*CreateComponentResponse, error) {
	return service.CreateComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) CreateComponentWithDescriptionAndBarcodeContext(ctx context.Context, request *CreateComponentWithDescriptionAndBarcode) (*CreateComponentWithDescriptionAndBarcodeResponse, error) {
	response := new(CreateComponentWithDescriptionAndBarcodeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CreateComponentWithDescriptionAndBarcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CreateComponentWithDescriptionAndBarcode(request *CreateComponentWithDescriptionAndBarcode) (*CreateComponentWithDescriptionAndBarcodeResponse, error) {
	return service.CreateComponentWithDescriptionAndBarcodeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) CreateComponentWithSubComponentsContext(ctx context.Context, request *CreateComponentWithSubComponents) (*CreateComponentWithSubComponentsResponse, error) {
	response := new(CreateComponentWithSubComponentsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CreateComponentWithSubComponents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CreateComponentWithSubComponents(request *CreateComponentWithSubComponents) (*CreateComponentWithSubComponentsResponse, error) {
	return service.CreateComponentWithSubComponentsContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) CreateComponentWithSpecifiedSubcomponentsContext(ctx context.Context, request *CreateComponentWithSpecifiedSubcomponents) (*CreateComponentWithSpecifiedSubcomponentsResponse, error) {
	response := new(CreateComponentWithSpecifiedSubcomponentsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CreateComponentWithSpecifiedSubcomponents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CreateComponentWithSpecifiedSubcomponents(request *CreateComponentWithSpecifiedSubcomponents) (*CreateComponentWithSpecifiedSubcomponentsResponse, error) {
	return service.CreateComponentWithSpecifiedSubcomponentsContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) CreateComponentWithSubComponentsWithDescriptionAndBarcodeContext(ctx context.Context, request *CreateComponentWithSubComponentsWithDescriptionAndBarcode) (*CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse, error) {
	response := new(CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CreateComponentWithSubComponentsWithDescriptionAndBarcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CreateComponentWithSubComponentsWithDescriptionAndBarcode(request *CreateComponentWithSubComponentsWithDescriptionAndBarcode) (*CreateComponentWithSubComponentsWithDescriptionAndBarcodeResponse, error) {
	return service.CreateComponentWithSubComponentsWithDescriptionAndBarcodeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) DeleteComponentByIDContext(ctx context.Context, request *DeleteComponentByID) (*DeleteComponentByIDResponse, error) {
	response := new(DeleteComponentByIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/DeleteComponentByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) DeleteComponentByID(request *DeleteComponentByID) (*DeleteComponentByIDResponse, error) {
	return service.DeleteComponentByIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) DeleteComponentByComponentContext(ctx context.Context, request *DeleteComponentByComponent) (*DeleteComponentByComponentResponse, error) {
	response := new(DeleteComponentByComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/DeleteComponentByComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) DeleteComponentByComponent(request *DeleteComponentByComponent) (*DeleteComponentByComponentResponse, error) {
	return service.DeleteComponentByComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChildComponentsContext(ctx context.Context, request *GetChildComponents) (*GetChildComponentsResponse, error) {
	response := new(GetChildComponentsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChildComponents", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChildComponents(request *GetChildComponents) (*GetChildComponentsResponse, error) {
	return service.GetChildComponentsContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChildComponentsByNameTypeContext(ctx context.Context, request *GetChildComponentsByNameType) (*GetChildComponentsByNameTypeResponse, error) {
	response := new(GetChildComponentsByNameTypeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChildComponentsByNameType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChildComponentsByNameType(request *GetChildComponentsByNameType) (*GetChildComponentsByNameTypeResponse, error) {
	return service.GetChildComponentsByNameTypeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChildComponentsByIDTypeContext(ctx context.Context, request *GetChildComponentsByIDType) (*GetChildComponentsByIDTypeResponse, error) {
	response := new(GetChildComponentsByIDTypeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChildComponentsByIDType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChildComponentsByIDType(request *GetChildComponentsByIDType) (*GetChildComponentsByIDTypeResponse, error) {
	return service.GetChildComponentsByIDTypeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChildRacksContext(ctx context.Context, request *GetChildRacks) (*GetChildRacksResponse, error) {
	response := new(GetChildRacksResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChildRacks", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChildRacks(request *GetChildRacks) (*GetChildRacksResponse, error) {
	return service.GetChildRacksContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) FilterComponentsByAccessContext(ctx context.Context, request *FilterComponentsByAccess) (*FilterComponentsByAccessResponse, error) {
	response := new(FilterComponentsByAccessResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/FilterComponentsByAccess", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) FilterComponentsByAccess(request *FilterComponentsByAccess) (*FilterComponentsByAccessResponse, error) {
	return service.FilterComponentsByAccessContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) FindSpacePathContext(ctx context.Context, request *FindSpacePath) (*FindSpacePathResponse, error) {
	response := new(FindSpacePathResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/FindSpacePath", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) FindSpacePath(request *FindSpacePath) (*FindSpacePathResponse, error) {
	return service.FindSpacePathContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetAllSpacesContext(ctx context.Context, request *GetAllSpaces) (*GetAllSpacesResponse, error) {
	response := new(GetAllSpacesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetAllSpaces", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetAllSpaces(request *GetAllSpaces) (*GetAllSpacesResponse, error) {
	return service.GetAllSpacesContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetAllZonesContext(ctx context.Context, request *GetAllZones) (*GetAllZonesResponse, error) {
	response := new(GetAllZonesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetAllZones", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetAllZones(request *GetAllZones) (*GetAllZonesResponse, error) {
	return service.GetAllZonesContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChangeHistoryContext(ctx context.Context, request *GetChangeHistory) (*GetChangeHistoryResponse, error) {
	response := new(GetChangeHistoryResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChangeHistory", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChangeHistory(request *GetChangeHistory) (*GetChangeHistoryResponse, error) {
	return service.GetChangeHistoryContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentAttributeValueByAttributeNameContext(ctx context.Context, request *GetComponentAttributeValueByAttributeName) (*GetComponentAttributeValueByAttributeNameResponse, error) {
	response := new(GetComponentAttributeValueByAttributeNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentAttributeValueByAttributeName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentAttributeValueByAttributeName(request *GetComponentAttributeValueByAttributeName) (*GetComponentAttributeValueByAttributeNameResponse, error) {
	return service.GetComponentAttributeValueByAttributeNameContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentAttributeValueByComponentContext(ctx context.Context, request *GetComponentAttributeValueByComponent) (*GetComponentAttributeValueByComponentResponse, error) {
	response := new(GetComponentAttributeValueByComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentAttributeValueByComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentAttributeValueByComponent(request *GetComponentAttributeValueByComponent) (*GetComponentAttributeValueByComponentResponse, error) {
	return service.GetComponentAttributeValueByComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetParentRackContext(ctx context.Context, request *GetParentRack) (*GetParentRackResponse, error) {
	response := new(GetParentRackResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetParentRack", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetParentRack(request *GetParentRack) (*GetParentRackResponse, error) {
	return service.GetParentRackContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) SaveComponentContext(ctx context.Context, request *SaveComponent) (*SaveComponentResponse, error) {
	response := new(SaveComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/SaveComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) SaveComponent(request *SaveComponent) (*SaveComponentResponse, error) {
	return service.SaveComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) TryGetComponentAttributeValueByAttributeNameContext(ctx context.Context, request *TryGetComponentAttributeValueByAttributeName) (*TryGetComponentAttributeValueByAttributeNameResponse, error) {
	response := new(TryGetComponentAttributeValueByAttributeNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/TryGetComponentAttributeValueByAttributeName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) TryGetComponentAttributeValueByAttributeName(request *TryGetComponentAttributeValueByAttributeName) (*TryGetComponentAttributeValueByAttributeNameResponse, error) {
	return service.TryGetComponentAttributeValueByAttributeNameContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) TryGetComponentAttributeValueByComponentContext(ctx context.Context, request *TryGetComponentAttributeValueByComponent) (*TryGetComponentAttributeValueByComponentResponse, error) {
	response := new(TryGetComponentAttributeValueByComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/TryGetComponentAttributeValueByComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) TryGetComponentAttributeValueByComponent(request *TryGetComponentAttributeValueByComponent) (*TryGetComponentAttributeValueByComponentResponse, error) {
	return service.TryGetComponentAttributeValueByComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) TryUpdateComponentAttributeValueByComponentIDContext(ctx context.Context, request *TryUpdateComponentAttributeValueByComponentID) (*TryUpdateComponentAttributeValueByComponentIDResponse, error) {
	response := new(TryUpdateComponentAttributeValueByComponentIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/TryUpdateComponentAttributeValueByComponentID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) TryUpdateComponentAttributeValueByComponentID(request *TryUpdateComponentAttributeValueByComponentID) (*TryUpdateComponentAttributeValueByComponentIDResponse, error) {
	return service.TryUpdateComponentAttributeValueByComponentIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) TryUpdateComponentAttributeValueContext(ctx context.Context, request *TryUpdateComponentAttributeValue) (*TryUpdateComponentAttributeValueResponse, error) {
	response := new(TryUpdateComponentAttributeValueResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/TryUpdateComponentAttributeValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) TryUpdateComponentAttributeValue(request *TryUpdateComponentAttributeValue) (*TryUpdateComponentAttributeValueResponse, error) {
	return service.TryUpdateComponentAttributeValueContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) UpdateComponentAttributeValueByComponentIDContext(ctx context.Context, request *UpdateComponentAttributeValueByComponentID) (*UpdateComponentAttributeValueByComponentIDResponse, error) {
	response := new(UpdateComponentAttributeValueByComponentIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/UpdateComponentAttributeValueByComponentID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) UpdateComponentAttributeValueByComponentID(request *UpdateComponentAttributeValueByComponentID) (*UpdateComponentAttributeValueByComponentIDResponse, error) {
	return service.UpdateComponentAttributeValueByComponentIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) UpdateComponentAttributeValueByComponentContext(ctx context.Context, request *UpdateComponentAttributeValueByComponent) (*UpdateComponentAttributeValueByComponentResponse, error) {
	response := new(UpdateComponentAttributeValueByComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/UpdateComponentAttributeValueByComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) UpdateComponentAttributeValueByComponent(request *UpdateComponentAttributeValueByComponent) (*UpdateComponentAttributeValueByComponentResponse, error) {
	return service.UpdateComponentAttributeValueByComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) MoveToSpaceContext(ctx context.Context, request *MoveToSpace) (*MoveToSpaceResponse, error) {
	response := new(MoveToSpaceResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/MoveToSpace", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) MoveToSpace(request *MoveToSpace) (*MoveToSpaceResponse, error) {
	return service.MoveToSpaceContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentAttributeByComponentIDContext(ctx context.Context, request *GetComponentAttributeByComponentID) (*GetComponentAttributeByComponentIDResponse, error) {
	response := new(GetComponentAttributeByComponentIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentAttributeByComponentID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentAttributeByComponentID(request *GetComponentAttributeByComponentID) (*GetComponentAttributeByComponentIDResponse, error) {
	return service.GetComponentAttributeByComponentIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentAttributesByComponentContext(ctx context.Context, request *GetComponentAttributesByComponent) (*GetComponentAttributesByComponentResponse, error) {
	response := new(GetComponentAttributesByComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentAttributesByComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentAttributesByComponent(request *GetComponentAttributesByComponent) (*GetComponentAttributesByComponentResponse, error) {
	return service.GetComponentAttributesByComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentByIDContext(ctx context.Context, request *GetComponentByID) (*GetComponentByIDResponse, error) {
	response := new(GetComponentByIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentByID(request *GetComponentByID) (*GetComponentByIDResponse, error) {
	return service.GetComponentByIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentIDsByTemplateContext(ctx context.Context, request *GetComponentIDsByTemplate) (*GetComponentIDsByTemplateResponse, error) {
	response := new(GetComponentIDsByTemplateResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentIDsByTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentIDsByTemplate(request *GetComponentIDsByTemplate) (*GetComponentIDsByTemplateResponse, error) {
	return service.GetComponentIDsByTemplateContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByAttributeNameContext(ctx context.Context, request *GetComponentsByAttributeName) (*GetComponentsByAttributeNameResponse, error) {
	response := new(GetComponentsByAttributeNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByAttributeName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByAttributeName(request *GetComponentsByAttributeName) (*GetComponentsByAttributeNameResponse, error) {
	return service.GetComponentsByAttributeNameContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByAttributeNameAndValueContext(ctx context.Context, request *GetComponentsByAttributeNameAndValue) (*GetComponentsByAttributeNameAndValueResponse, error) {
	response := new(GetComponentsByAttributeNameAndValueResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByAttributeNameAndValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByAttributeNameAndValue(request *GetComponentsByAttributeNameAndValue) (*GetComponentsByAttributeNameAndValueResponse, error) {
	return service.GetComponentsByAttributeNameAndValueContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByBarcodeContext(ctx context.Context, request *GetComponentsByBarcode) (*GetComponentsByBarcodeResponse, error) {
	response := new(GetComponentsByBarcodeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByBarcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByBarcode(request *GetComponentsByBarcode) (*GetComponentsByBarcodeResponse, error) {
	return service.GetComponentsByBarcodeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByLevelContext(ctx context.Context, request *GetComponentsByLevel) (*GetComponentsByLevelResponse, error) {
	response := new(GetComponentsByLevelResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByLevel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByLevel(request *GetComponentsByLevel) (*GetComponentsByLevelResponse, error) {
	return service.GetComponentsByLevelContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByNameContext(ctx context.Context, request *GetComponentsByName) (*GetComponentsByNameResponse, error) {
	response := new(GetComponentsByNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByName(request *GetComponentsByName) (*GetComponentsByNameResponse, error) {
	return service.GetComponentsByNameContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsInStoreRoomBySiteContext(ctx context.Context, request *GetComponentsInStoreRoomBySite) (*GetComponentsInStoreRoomBySiteResponse, error) {
	response := new(GetComponentsInStoreRoomBySiteResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsInStoreRoomBySite", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsInStoreRoomBySite(request *GetComponentsInStoreRoomBySite) (*GetComponentsInStoreRoomBySiteResponse, error) {
	return service.GetComponentsInStoreRoomBySiteContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentsByTemplateIdContext(ctx context.Context, request *GetComponentsByTemplateId) (*GetComponentsByTemplateIdResponse, error) {
	response := new(GetComponentsByTemplateIdResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentsByTemplateId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentsByTemplateId(request *GetComponentsByTemplateId) (*GetComponentsByTemplateIdResponse, error) {
	return service.GetComponentsByTemplateIdContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) FilterComponentsByParentIdContext(ctx context.Context, request *FilterComponentsByParentId) (*FilterComponentsByParentIdResponse, error) {
	response := new(FilterComponentsByParentIdResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/FilterComponentsByParentId", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) FilterComponentsByParentId(request *FilterComponentsByParentId) (*FilterComponentsByParentIdResponse, error) {
	return service.FilterComponentsByParentIdContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByTemplateIdAttributeNameAndValueContext(ctx context.Context, request *GetCompleteComponentsByTemplateIdAttributeNameAndValue) (*GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse, error) {
	response := new(GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByTemplateIdAttributeNameAndValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByTemplateIdAttributeNameAndValue(request *GetCompleteComponentsByTemplateIdAttributeNameAndValue) (*GetCompleteComponentsByTemplateIdAttributeNameAndValueResponse, error) {
	return service.GetCompleteComponentsByTemplateIdAttributeNameAndValueContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByAttributeNameAndValueContext(ctx context.Context, request *GetCompleteComponentsByAttributeNameAndValue) (*GetCompleteComponentsByAttributeNameAndValueResponse, error) {
	response := new(GetCompleteComponentsByAttributeNameAndValueResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByAttributeNameAndValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByAttributeNameAndValue(request *GetCompleteComponentsByAttributeNameAndValue) (*GetCompleteComponentsByAttributeNameAndValueResponse, error) {
	return service.GetCompleteComponentsByAttributeNameAndValueContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByTemplateContext(ctx context.Context, request *GetCompleteComponentsByTemplate) (*GetCompleteComponentsByTemplateResponse, error) {
	response := new(GetCompleteComponentsByTemplateResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByTemplate(request *GetCompleteComponentsByTemplate) (*GetCompleteComponentsByTemplateResponse, error) {
	return service.GetCompleteComponentsByTemplateContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentByIDContext(ctx context.Context, request *GetCompleteComponentByID) (*GetCompleteComponentByIDResponse, error) {
	response := new(GetCompleteComponentByIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentByID(request *GetCompleteComponentByID) (*GetCompleteComponentByIDResponse, error) {
	return service.GetCompleteComponentByIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByIDsContext(ctx context.Context, request *GetCompleteComponentsByIDs) (*GetCompleteComponentsByIDsResponse, error) {
	response := new(GetCompleteComponentsByIDsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByIDs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByIDs(request *GetCompleteComponentsByIDs) (*GetCompleteComponentsByIDsResponse, error) {
	return service.GetCompleteComponentsByIDsContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByNameContext(ctx context.Context, request *GetCompleteComponentsByName) (*GetCompleteComponentsByNameResponse, error) {
	response := new(GetCompleteComponentsByNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByName(request *GetCompleteComponentsByName) (*GetCompleteComponentsByNameResponse, error) {
	return service.GetCompleteComponentsByNameContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentsByBarcodeContext(ctx context.Context, request *GetCompleteComponentsByBarcode) (*GetCompleteComponentsByBarcodeResponse, error) {
	response := new(GetCompleteComponentsByBarcodeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentsByBarcode", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentsByBarcode(request *GetCompleteComponentsByBarcode) (*GetCompleteComponentsByBarcodeResponse, error) {
	return service.GetCompleteComponentsByBarcodeContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetCompleteComponentSpaceByAttributeIdentifierContext(ctx context.Context, request *GetCompleteComponentSpaceByAttributeIdentifier) (*GetCompleteComponentSpaceByAttributeIdentifierResponse, error) {
	response := new(GetCompleteComponentSpaceByAttributeIdentifierResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetCompleteComponentSpaceByAttributeIdentifier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetCompleteComponentSpaceByAttributeIdentifier(request *GetCompleteComponentSpaceByAttributeIdentifier) (*GetCompleteComponentSpaceByAttributeIdentifierResponse, error) {
	return service.GetCompleteComponentSpaceByAttributeIdentifierContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentSpaceByAttributeIdentifierContext(ctx context.Context, request *GetComponentSpaceByAttributeIdentifier) (*GetComponentSpaceByAttributeIdentifierResponse, error) {
	response := new(GetComponentSpaceByAttributeIdentifierResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentSpaceByAttributeIdentifier", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentSpaceByAttributeIdentifier(request *GetComponentSpaceByAttributeIdentifier) (*GetComponentSpaceByAttributeIdentifierResponse, error) {
	return service.GetComponentSpaceByAttributeIdentifierContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentChildrenContext(ctx context.Context, request *GetComponentChildren) (*GetComponentChildrenResponse, error) {
	response := new(GetComponentChildrenResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentChildren", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentChildren(request *GetComponentChildren) (*GetComponentChildrenResponse, error) {
	return service.GetComponentChildrenContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentPortsContext(ctx context.Context, request *GetComponentPorts) (*GetComponentPortsResponse, error) {
	response := new(GetComponentPortsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentPorts", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentPorts(request *GetComponentPorts) (*GetComponentPortsResponse, error) {
	return service.GetComponentPortsContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetZoneContext(ctx context.Context, request *GetZone) (*GetZoneResponse, error) {
	response := new(GetZoneResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetZone", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetZone(request *GetZone) (*GetZoneResponse, error) {
	return service.GetZoneContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) SaveComponentAttributesContext(ctx context.Context, request *SaveComponentAttributes) (*SaveComponentAttributesResponse, error) {
	response := new(SaveComponentAttributesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/SaveComponentAttributes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) SaveComponentAttributes(request *SaveComponentAttributes) (*SaveComponentAttributesResponse, error) {
	return service.SaveComponentAttributesContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) RenameComponentContext(ctx context.Context, request *RenameComponent) (*RenameComponentResponse, error) {
	response := new(RenameComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/RenameComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) RenameComponent(request *RenameComponent) (*RenameComponentResponse, error) {
	return service.RenameComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) CopyComponentContext(ctx context.Context, request *CopyComponent) (*CopyComponentResponse, error) {
	response := new(CopyComponentResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/CopyComponent", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) CopyComponent(request *CopyComponent) (*CopyComponentResponse, error) {
	return service.CopyComponentContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) IsSubcomponentChangeRequiredContext(ctx context.Context, request *IsSubcomponentChangeRequired) (*IsSubcomponentChangeRequiredResponse, error) {
	response := new(IsSubcomponentChangeRequiredResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/IsSubcomponentChangeRequired", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) IsSubcomponentChangeRequired(request *IsSubcomponentChangeRequired) (*IsSubcomponentChangeRequiredResponse, error) {
	return service.IsSubcomponentChangeRequiredContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetComponentRoleAccessContext(ctx context.Context, request *GetComponentRoleAccess) (*GetComponentRoleAccessResponse, error) {
	response := new(GetComponentRoleAccessResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetComponentRoleAccess", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetComponentRoleAccess(request *GetComponentRoleAccess) (*GetComponentRoleAccessResponse, error) {
	return service.GetComponentRoleAccessContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetPortHostingComponentsByIDContext(ctx context.Context, request *GetPortHostingComponentsByID) (*GetPortHostingComponentsByIDResponse, error) {
	response := new(GetPortHostingComponentsByIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetPortHostingComponentsByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetPortHostingComponentsByID(request *GetPortHostingComponentsByID) (*GetPortHostingComponentsByIDResponse, error) {
	return service.GetPortHostingComponentsByIDContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetChassisInDataCenterByBladeTemplateContext(ctx context.Context, request *GetChassisInDataCenterByBladeTemplate) (*GetChassisInDataCenterByBladeTemplateResponse, error) {
	response := new(GetChassisInDataCenterByBladeTemplateResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetChassisInDataCenterByBladeTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetChassisInDataCenterByBladeTemplate(request *GetChassisInDataCenterByBladeTemplate) (*GetChassisInDataCenterByBladeTemplateResponse, error) {
	return service.GetChassisInDataCenterByBladeTemplateContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) GetAvailableSlotNumbersForChassisContext(ctx context.Context, request *GetAvailableSlotNumbersForChassis) (*GetAvailableSlotNumbersForChassisResponse, error) {
	response := new(GetAvailableSlotNumbersForChassisResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetAvailableSlotNumbersForChassis", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) GetAvailableSlotNumbersForChassis(request *GetAvailableSlotNumbersForChassis) (*GetAvailableSlotNumbersForChassisResponse, error) {
	return service.GetAvailableSlotNumbersForChassisContext(
		context.Background(),
		request,
	)
}

func (service *componentServicesSoap) EnsureSubcomponentDefaultsExistContext(ctx context.Context, request *EnsureSubcomponentDefaultsExist) (*EnsureSubcomponentDefaultsExistResponse, error) {
	response := new(EnsureSubcomponentDefaultsExistResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/EnsureSubcomponentDefaultsExist", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *componentServicesSoap) EnsureSubcomponentDefaultsExist(request *EnsureSubcomponentDefaultsExist) (*EnsureSubcomponentDefaultsExistResponse, error) {
	return service.EnsureSubcomponentDefaultsExistContext(
		context.Background(),
		request,
	)
}
