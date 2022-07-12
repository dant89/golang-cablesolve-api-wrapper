package template

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

type ZeroRackUnit string

const (
	ZeroRackUnitNone ZeroRackUnit = "None"

	ZeroRackUnitLeft ZeroRackUnit = "Left"

	ZeroRackUnitRight ZeroRackUnit = "Right"

	ZeroRackUnitTop ZeroRackUnit = "Top"

	ZeroRackUnitBottom ZeroRackUnit = "Bottom"
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

type GetAllTemplates struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllTemplates"`
}

type GetAllTemplatesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllTemplatesResponse"`

	GetAllTemplatesResult *ArrayOfTemplate `xml:"GetAllTemplatesResult,omitempty" json:"GetAllTemplatesResult,omitempty"`
}

type GetAllTemplatesTemplate struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllTemplatesTemplate"`
}

type GetAllTemplatesTemplateResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetAllTemplatesTemplateResponse"`

	GetAllTemplatesTemplateResult *Template `xml:"GetAllTemplatesTemplateResult,omitempty" json:"GetAllTemplatesTemplateResult,omitempty"`
}

type GetTemplateAttributes struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateAttributes"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`
}

type GetTemplateAttributesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateAttributesResponse"`

	GetTemplateAttributesResult *ArrayOfTemplateAttribute `xml:"GetTemplateAttributesResult,omitempty" json:"GetTemplateAttributesResult,omitempty"`
}

type GetTemplateByID struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateByID"`

	Id int32 `xml:"id,omitempty" json:"id,omitempty"`
}

type GetTemplateByIDResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateByIDResponse"`

	GetTemplateByIDResult *Template `xml:"GetTemplateByIDResult,omitempty" json:"GetTemplateByIDResult,omitempty"`
}

type GetTemplatesByLevel struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByLevel"`

	Level int32 `xml:"level,omitempty" json:"level,omitempty"`
}

type GetTemplatesByLevelResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByLevelResponse"`

	GetTemplatesByLevelResult *ArrayOfTemplate `xml:"GetTemplatesByLevelResult,omitempty" json:"GetTemplatesByLevelResult,omitempty"`
}

type GetTemplatesByName struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByName"`

	Name string `xml:"name,omitempty" json:"name,omitempty"`
}

type GetTemplatesByNameResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByNameResponse"`

	GetTemplatesByNameResult *ArrayOfTemplate `xml:"GetTemplatesByNameResult,omitempty" json:"GetTemplatesByNameResult,omitempty"`
}

type GetTemplatesByType struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByType"`

	Type_ *ComponentType `xml:"type,omitempty" json:"type,omitempty"`
}

type GetTemplatesByTypeResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByTypeResponse"`

	GetTemplatesByTypeResult *ArrayOfTemplate `xml:"GetTemplatesByTypeResult,omitempty" json:"GetTemplatesByTypeResult,omitempty"`
}

type GetTemplateChildren struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateChildren"`

	Parent int32 `xml:"parent,omitempty" json:"parent,omitempty"`
}

type GetTemplateChildrenResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateChildrenResponse"`

	GetTemplateChildrenResult *ArrayOfTemplate `xml:"GetTemplateChildrenResult,omitempty" json:"GetTemplateChildrenResult,omitempty"`
}

type GetTemplateByAttributeValue struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateByAttributeValue"`

	AttributeID int32 `xml:"attributeID,omitempty" json:"attributeID,omitempty"`

	Value string `xml:"value,omitempty" json:"value,omitempty"`
}

type GetTemplateByAttributeValueResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplateByAttributeValueResponse"`

	GetTemplateByAttributeValueResult *ArrayOfTemplate `xml:"GetTemplateByAttributeValueResult,omitempty" json:"GetTemplateByAttributeValueResult,omitempty"`
}

type GetWorkflowTemplates struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetWorkflowTemplates"`
}

type GetWorkflowTemplatesResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetWorkflowTemplatesResponse"`

	GetWorkflowTemplatesResult *ArrayOfTemplate `xml:"GetWorkflowTemplatesResult,omitempty" json:"GetWorkflowTemplatesResult,omitempty"`
}

type GetTemplatesByIds struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByIds"`

	Ids *ArrayOfInt `xml:"ids,omitempty" json:"ids,omitempty"`
}

type GetTemplatesByIdsResponse struct {
	XMLName xml.Name `xml:"http://www.cormant.com/cswebapi GetTemplatesByIdsResponse"`

	GetTemplatesByIdsResult *ArrayOfTemplate `xml:"GetTemplatesByIdsResult,omitempty" json:"GetTemplatesByIdsResult,omitempty"`
}

type ArrayOfTemplate struct {
	Template []*Template `xml:"Template,omitempty" json:"Template,omitempty"`
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

type ArrayOfTemplateAttribute struct {
	TemplateAttribute []*TemplateAttribute `xml:"TemplateAttribute,omitempty" json:"TemplateAttribute,omitempty"`
}

type TemplateAttribute struct {
	Attribute int32 `xml:"Attribute,omitempty" json:"Attribute,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty" json:"DefaultValue,omitempty"`

	Template int32 `xml:"Template,omitempty" json:"Template,omitempty"`
}

type ArrayOfInt struct {
	int []int32 `xml:"int,omitempty" json:"int,omitempty"`
}

type TemplateServicesSoap interface {

	/* Gets all templates. */
	GetAllTemplates(request *GetAllTemplates) (*GetAllTemplatesResponse, error)

	GetAllTemplatesContext(ctx context.Context, request *GetAllTemplates) (*GetAllTemplatesResponse, error)

	/* Gets the "All Templates" template. */
	GetAllTemplatesTemplate(request *GetAllTemplatesTemplate) (*GetAllTemplatesTemplateResponse, error)

	GetAllTemplatesTemplateContext(ctx context.Context, request *GetAllTemplatesTemplate) (*GetAllTemplatesTemplateResponse, error)

	/* Gets all the template attributes of the specified template. */
	GetTemplateAttributes(request *GetTemplateAttributes) (*GetTemplateAttributesResponse, error)

	GetTemplateAttributesContext(ctx context.Context, request *GetTemplateAttributes) (*GetTemplateAttributesResponse, error)

	/* Gets a template based on the specified ID. */
	GetTemplateByID(request *GetTemplateByID) (*GetTemplateByIDResponse, error)

	GetTemplateByIDContext(ctx context.Context, request *GetTemplateByID) (*GetTemplateByIDResponse, error)

	/* Get all templates based on the specified level. */
	GetTemplatesByLevel(request *GetTemplatesByLevel) (*GetTemplatesByLevelResponse, error)

	GetTemplatesByLevelContext(ctx context.Context, request *GetTemplatesByLevel) (*GetTemplatesByLevelResponse, error)

	/* Gets all templates based on the specified name. */
	GetTemplatesByName(request *GetTemplatesByName) (*GetTemplatesByNameResponse, error)

	GetTemplatesByNameContext(ctx context.Context, request *GetTemplatesByName) (*GetTemplatesByNameResponse, error)

	/* Gets all templates with the specified type. This is intended to load Site, DataCenter and Rack templates. */
	GetTemplatesByType(request *GetTemplatesByType) (*GetTemplatesByTypeResponse, error)

	GetTemplatesByTypeContext(ctx context.Context, request *GetTemplatesByType) (*GetTemplatesByTypeResponse, error)

	/* Gets all the children of the specified template. */
	GetTemplateChildren(request *GetTemplateChildren) (*GetTemplateChildrenResponse, error)

	GetTemplateChildrenContext(ctx context.Context, request *GetTemplateChildren) (*GetTemplateChildrenResponse, error)

	/* Gets templates containing the specified attribute with the specified default value. */
	GetTemplateByAttributeValue(request *GetTemplateByAttributeValue) (*GetTemplateByAttributeValueResponse, error)

	GetTemplateByAttributeValueContext(ctx context.Context, request *GetTemplateByAttributeValue) (*GetTemplateByAttributeValueResponse, error)

	/* Gets workflow templates configured for the Default site in Reference Data. The Template.Flag can be cast to the to the Task.TaskType that identifies whether the Task is Install, Decommission, Modify etc.
	There can be multiple templates corresponding to a single Task Type. This allows different defaults for the custom attributes for different types. */
	GetWorkflowTemplates(request *GetWorkflowTemplates) (*GetWorkflowTemplatesResponse, error)

	GetWorkflowTemplatesContext(ctx context.Context, request *GetWorkflowTemplates) (*GetWorkflowTemplatesResponse, error)

	/* Gets templates by an array of IDs. Used to improve performance if several templates are required. */
	GetTemplatesByIds(request *GetTemplatesByIds) (*GetTemplatesByIdsResponse, error)

	GetTemplatesByIdsContext(ctx context.Context, request *GetTemplatesByIds) (*GetTemplatesByIdsResponse, error)
}

type templateServicesSoap struct {
	client *soap.Client
}

func NewTemplateServicesSoap(client *soap.Client) TemplateServicesSoap {
	return &templateServicesSoap{
		client: client,
	}
}

func (service *templateServicesSoap) GetAllTemplatesContext(ctx context.Context, request *GetAllTemplates) (*GetAllTemplatesResponse, error) {
	response := new(GetAllTemplatesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetAllTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetAllTemplates(request *GetAllTemplates) (*GetAllTemplatesResponse, error) {
	return service.GetAllTemplatesContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetAllTemplatesTemplateContext(ctx context.Context, request *GetAllTemplatesTemplate) (*GetAllTemplatesTemplateResponse, error) {
	response := new(GetAllTemplatesTemplateResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetAllTemplatesTemplate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetAllTemplatesTemplate(request *GetAllTemplatesTemplate) (*GetAllTemplatesTemplateResponse, error) {
	return service.GetAllTemplatesTemplateContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplateAttributesContext(ctx context.Context, request *GetTemplateAttributes) (*GetTemplateAttributesResponse, error) {
	response := new(GetTemplateAttributesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplateAttributes", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplateAttributes(request *GetTemplateAttributes) (*GetTemplateAttributesResponse, error) {
	return service.GetTemplateAttributesContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplateByIDContext(ctx context.Context, request *GetTemplateByID) (*GetTemplateByIDResponse, error) {
	response := new(GetTemplateByIDResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplateByID", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplateByID(request *GetTemplateByID) (*GetTemplateByIDResponse, error) {
	return service.GetTemplateByIDContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplatesByLevelContext(ctx context.Context, request *GetTemplatesByLevel) (*GetTemplatesByLevelResponse, error) {
	response := new(GetTemplatesByLevelResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplatesByLevel", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplatesByLevel(request *GetTemplatesByLevel) (*GetTemplatesByLevelResponse, error) {
	return service.GetTemplatesByLevelContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplatesByNameContext(ctx context.Context, request *GetTemplatesByName) (*GetTemplatesByNameResponse, error) {
	response := new(GetTemplatesByNameResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplatesByName", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplatesByName(request *GetTemplatesByName) (*GetTemplatesByNameResponse, error) {
	return service.GetTemplatesByNameContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplatesByTypeContext(ctx context.Context, request *GetTemplatesByType) (*GetTemplatesByTypeResponse, error) {
	response := new(GetTemplatesByTypeResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplatesByType", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplatesByType(request *GetTemplatesByType) (*GetTemplatesByTypeResponse, error) {
	return service.GetTemplatesByTypeContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplateChildrenContext(ctx context.Context, request *GetTemplateChildren) (*GetTemplateChildrenResponse, error) {
	response := new(GetTemplateChildrenResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplateChildren", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplateChildren(request *GetTemplateChildren) (*GetTemplateChildrenResponse, error) {
	return service.GetTemplateChildrenContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplateByAttributeValueContext(ctx context.Context, request *GetTemplateByAttributeValue) (*GetTemplateByAttributeValueResponse, error) {
	response := new(GetTemplateByAttributeValueResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplateByAttributeValue", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplateByAttributeValue(request *GetTemplateByAttributeValue) (*GetTemplateByAttributeValueResponse, error) {
	return service.GetTemplateByAttributeValueContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetWorkflowTemplatesContext(ctx context.Context, request *GetWorkflowTemplates) (*GetWorkflowTemplatesResponse, error) {
	response := new(GetWorkflowTemplatesResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetWorkflowTemplates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetWorkflowTemplates(request *GetWorkflowTemplates) (*GetWorkflowTemplatesResponse, error) {
	return service.GetWorkflowTemplatesContext(
		context.Background(),
		request,
	)
}

func (service *templateServicesSoap) GetTemplatesByIdsContext(ctx context.Context, request *GetTemplatesByIds) (*GetTemplatesByIdsResponse, error) {
	response := new(GetTemplatesByIdsResponse)
	err := service.client.CallContext(ctx, "http://www.cormant.com/cswebapi/GetTemplatesByIds", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *templateServicesSoap) GetTemplatesByIds(request *GetTemplatesByIds) (*GetTemplatesByIdsResponse, error) {
	return service.GetTemplatesByIdsContext(
		context.Background(),
		request,
	)
}
