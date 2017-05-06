
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/core/crypto/primitives"
)



 
// CBDoc is a high level smart contract that POCs together business artifact based smart contracts
type CBDoc struct {

}

// UserDetails is for storing User Details

type Document struct{	
	DocumentId string `json:"documentId"`
	Source string `json:"source"`
	Destination string `json:"destination"`
	Status string `json:"status"`	
}

type Counter struct {
	count int
}

func (self Counter) currentValue() int {
	return self.count
}
func (self *Counter) increment() {
	self.count++
}

// Init initializes the smart contracts
func (t *CBDoc) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	// Check if table already exists
	_, err := stub.GetTable("Document")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = stub.CreateTable("Document", []*shim.ColumnDefinition{
		&shim.ColumnDefinition{Name: "documentId", Type: shim.ColumnDefinition_STRING, Key: true},
		&shim.ColumnDefinition{Name: "source", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "destination", Type: shim.ColumnDefinition_STRING, Key: false},
		&shim.ColumnDefinition{Name: "status", Type: shim.ColumnDefinition_STRING, Key: false},
	})
	if err != nil {
		return nil, errors.New("Failed creating Document table.")
	}
	return nil, nil
}

// generate booking number for shipping item
func (t *CBDoc) createDocument(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 4 {
			return nil, fmt.Errorf("Incorrect number of arguments. Expecting 4. Got: %d.", len(args))
		}
		
		documentId:=args[0]
		source:=args[1]
		destination:=args[2]
		status:=args[3]
		
		// Insert a row
		ok, err := stub.InsertRow("Document", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: documentId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: destination}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},			
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}	

//get all booking details for specified document status
func (t *CBDoc) viewDocumentsByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting document status to query")
	}

	status := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("Document", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*Document{}	
	
	for row := range rows {		
		newApp:= new(Document)
		newApp.DocumentId = row.Columns[0].GetString_()
		newApp.Source = row.Columns[1].GetString_()
		newApp.Destination = row.Columns[2].GetString_()
		newApp.Status = row.Columns[3].GetString_()
		
		if newApp.Status == status{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

 //get all booking details for specified document source
func (t *CBDoc) viewDocumentsBySource(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting document source to query")
	}

	source := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("Document", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*Document{}	
	
	for row := range rows {		
		newApp:= new(Document)
		newApp.DocumentId = row.Columns[0].GetString_()
		newApp.Source = row.Columns[1].GetString_()
		newApp.Destination = row.Columns[2].GetString_()
		newApp.Status = row.Columns[3].GetString_()
		
		if newApp.Source == source{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get all booking details for specified document destination
func (t *CBDoc) viewDocumentsByDestination(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting document destination to query")
	}

	destination := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("Document", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*Document{}	
	
	for row := range rows {		
		newApp:= new(Document)
		newApp.DocumentId = row.Columns[0].GetString_()
		newApp.Source = row.Columns[1].GetString_()
		newApp.Destination = row.Columns[2].GetString_()
		newApp.Status = row.Columns[3].GetString_()
		
		if newApp.Destination == destination{
		res2E=append(res2E,newApp)		
		}				
	}
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

//get all booking details for specified document destination
func (t *CBDoc) countDocumentsByStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting document destination to query")
	}

	destination := args[0]
	
	var columns []shim.Column

	rows, err := stub.GetRows("Document", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
			
	res2E:= []*Document{}	
	counter := Counter{0}
		for row := range rows {		
		newApp:= new(Document)
		newApp.DocumentId = row.Columns[0].GetString_()
		newApp.Source = row.Columns[1].GetString_()
		newApp.Destination = row.Columns[2].GetString_()
		newApp.Status = row.Columns[3].GetString_()
		
		if newApp.Destination == destination{
		res2E=append(res2E,newApp)
		counter.increment()
		}				
	}
	
     mapB, _ := json.Marshal(counter.currentValue())
     fmt.Println(string(mapB))
	
     return mapB, nil

}

//update document status by document id
func (t *CBDoc) updateDocumentStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2.")
	}
	documentId := args[0]
	newStatus :=  args[1]
	// Get the row pertaining to this asnNumber
		var columns1 []shim.Column
		col1 := shim.Column{Value: &shim.Column_String_{String_: documentId}}
		columns1 = append(columns1, col1)

		row, err := stub.GetRow("Document", columns1)
		if err != nil {
			return nil, fmt.Errorf("Error: Failed retrieving document with document id %s. Error %s", documentId, err.Error())
		}

		// GetRows returns empty message if key does not exist
		if len(row.Columns) == 0 {
			return nil, nil
		}
		// Delete the row pertaining to this applicationId
		err = stub.DeleteRow(
			"Document",
			columns1,
		)
		if err != nil {
			return nil, errors.New("Failed deleting row.")
		}
		documentId = row.Columns[0].GetString_()
		source := row.Columns[1].GetString_()
		destination := row.Columns[2].GetString_()
		status := newStatus
		
		// Inserting document details
		ok, err := stub.InsertRow("Document", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: documentId}},
				&shim.Column{Value: &shim.Column_String_{String_: source}},
				&shim.Column{Value: &shim.Column_String_{String_: destination}},
				&shim.Column{Value: &shim.Column_String_{String_: status}},
			}})

		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
	return nil, nil	
}
// Invoke invokes the chaincode
func (t *CBDoc) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "createDocument" {
		t := CBDoc{}
		return t.createDocument(stub, args)	
	} 
	
	if function == "updateDocumentStatus" {
		t := CBDoc{}
		return t.updateDocumentStatus(stub, args)	
	}

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *CBDoc) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "viewDocumentsByStatus" {
		t := CBDoc{}
		return t.viewDocumentsByStatus(stub, args)		
	} else if function == "viewDocumentsBySource" { 
		t := CBDoc{}
		return t.viewDocumentsBySource(stub, args)
	}else if function == "viewDocumentsByDestination" { 
		t := CBDoc{}
		return t.viewDocumentsByDestination(stub, args)
	}	
	return nil, errors.New("Invalid query function name.")
}

func main() {
	primitives.SetSecurityLevel("SHA3", 256)
	err := shim.Start(new(CBDoc))
	if err != nil {
		fmt.Printf("Error starting CBDoc: %s", err)
	}
} 