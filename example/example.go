package main

import (
	"github.com/dant89/golang-cablesolve-api-wrapper/clients/component"
	"github.com/dant89/golang-cablesolve-api-wrapper/clients/template"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
	"log"
	"net/http"
	"strconv"
)

const CableSolveHost = ""     // Specify host URL
const CableSolveUsername = "" // Username to act as - this has to exist to avoid 500 SOAP errors

func main() {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// Get components by an array of template IDs
	// runs a subroutine per ID to speed up response of requests
	router.GET("/components", func(c *gin.Context) {
		// Slice of template IDs
		templateIds := []int32{}
		result := getComponentsByTemplateIds(templateIds)
		c.JSON(http.StatusOK, result)
	})

	router.GET("/components/:id", func(c *gin.Context) {
		componentId := c.Param("id")
		componentId32, _ := strconv.ParseInt(componentId, 0, 32)
		result := getComponentById(int32(componentId32))
		c.JSON(http.StatusOK, result)
	})

	router.GET("/components/:id/children", func(c *gin.Context) {
		componentId := c.Param("id")
		componentId32, _ := strconv.ParseInt(componentId, 0, 32)
		result := getComponentChildrenById(int32(componentId32))
		c.JSON(http.StatusOK, result)
	})

	router.GET("/components/:id/parent", func(c *gin.Context) {
		componentId := c.Param("id")
		componentId32, _ := strconv.ParseInt(componentId, 0, 32)
		result := getComponentParentById(int32(componentId32))
		c.JSON(http.StatusOK, result)
	})

	router.GET("/templates", func(c *gin.Context) {
		result := getTemplates()
		c.JSON(http.StatusOK, result)
	})

	router.GET("/templates/:id", func(c *gin.Context) {
		templateId := c.Param("id")
		templateId32, _ := strconv.ParseInt(templateId, 0, 32)
		result := getTemplate(int32(templateId32))
		c.JSON(http.StatusOK, result)
	})

	router.Run(":3004")
}

// Add further component clients as required
func getComponentClient() component.ComponentServicesSoap {
	client := soap.NewClient("https://" + CableSolveHost + "/cswebapi/ComponentServices.asmx?WSDL")
	return component.NewComponentServicesSoap(client)
}

func getTemplateClient() template.TemplateServicesSoap {
	client := soap.NewClient("https://" + CableSolveHost + "/cswebapi/TemplateServices.asmx?WSDL")
	return template.NewTemplateServicesSoap(client)
}

func getComponentById(componentId int32) *component.Component {
	service := getComponentClient()

	log.Println("Loading single component.")
	res, err := service.GetCompleteComponentByID(&component.GetCompleteComponentByID{
		Id:       componentId,
		Username: CableSolveUsername,
	})
	if err != nil {
		log.Fatalf("Failed to load single component: %v", err)
		return nil
	}

	return res.GetCompleteComponentByIDResult.Component
}

func getComponentChildrenById(componentId int32) []*component.Component {
	service := getComponentClient()

	log.Println("Loading single component children.")
	res, err := service.GetChildComponents(&component.GetChildComponents{
		ComponentId: componentId,
		Username:    CableSolveUsername,
	})
	if err != nil {
		log.Fatalf("Failed to load single component children: %v", err)
		return nil
	}

	return res.GetChildComponentsResult.Component
}

func getComponentParentById(componentId int32) *component.Component {
	service := getComponentClient()

	log.Println("Loading single component parent.")
	res, err := service.GetParentRack(&component.GetParentRack{
		ChildID: componentId,
	})
	if err != nil {
		log.Fatalf("Failed to load single component parent: %v", err)
		return nil
	}

	return res.GetParentRackResult
}

func getComponentsByTemplateIds(templateIds []int32) []*component.Component {

	var components []*component.Component
	ch := make(chan []*component.Component, 20)

	i := 0
	for _, templateId := range templateIds {
		go getComponentsByTemplateId(ch, templateId)
		i++
	}

	for l := 1; l <= i; l++ {
		components = append(components, <-ch...)
	}

	return components
}

func getComponentsByTemplateId(c chan []*component.Component, templateId int32) {
	service := getComponentClient()

	log.Printf("Loading subsection of components with templateId: %v", templateId)
	res, err := service.GetComponentsByTemplateId(&component.GetComponentsByTemplateId{
		TemplateId: templateId,
	})
	if err != nil {
		log.Fatalf("Failed to load subsection of components with templateId: %v %v", templateId, err)
	}

	c <- res.GetComponentsByTemplateIdResult.Component
}

func getTemplates() []*template.Template {
	service := getTemplateClient()

	log.Println("Loading templates.")
	res, err := service.GetAllTemplates(&template.GetAllTemplates{})
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
		return nil
	}

	return res.GetAllTemplatesResult.Template
}

func getTemplate(templateId int32) *template.Template {
	service := getTemplateClient()

	log.Println("Loading single template.")
	res, err := service.GetTemplateByID(&template.GetTemplateByID{
		Id: templateId,
	})
	if err != nil {
		log.Fatalf("Failed to load single template: %v", err)
		return nil
	}

	return res.GetTemplateByIDResult
}
