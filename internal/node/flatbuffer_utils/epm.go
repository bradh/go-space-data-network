package flatbuffer_utils

import (
	"context"
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	EPM "github.com/DigitalArsenal/space-data-network/internal/spacedatastandards/EPM"
	"github.com/emersion/go-vcard"
	flatbuffers "github.com/google/flatbuffers/go"
)

var EPMFID string = "$EPM"

func CreateEPM(
	distinguishedName string,
	legalName string,
	familyName string,
	givenName string,
	additionalName string,
	honorificPrefix string,
	honorificSuffix string,
	jobTitle string,
	occupation string,
	alternateNames []string,
	email string,
	telephone string,
	country string,
	region string,
	locality string,
	postalCode string,
	street string,
	poBox string,
) []byte {
	builder := flatbuffers.NewBuilder(0)

	// Create string offsets for all fields that are of string type
	dnOffset := builder.CreateString(distinguishedName)
	legalNameOffset := builder.CreateString(legalName)
	familyNameOffset := builder.CreateString(familyName)
	givenNameOffset := builder.CreateString(givenName)
	additionalNameOffset := builder.CreateString(additionalName)
	honorificPrefixOffset := builder.CreateString(honorificPrefix)
	honorificSuffixOffset := builder.CreateString(honorificSuffix)
	jobTitleOffset := builder.CreateString(jobTitle)
	occupationOffset := builder.CreateString(occupation)
	emailOffset := builder.CreateString(email)
	telephoneOffset := builder.CreateString(telephone)

	// Create string offsets for address fields
	countryOffset := builder.CreateString(country)
	regionOffset := builder.CreateString(region)
	localityOffset := builder.CreateString(locality)
	postalCodeOffset := builder.CreateString(postalCode)
	streetOffset := builder.CreateString(street)
	poBoxOffset := builder.CreateString(poBox)

	// Create vectors for alternate names
	alternateNamesOffsets := make([]flatbuffers.UOffsetT, len(alternateNames))
	for i, name := range alternateNames {
		alternateNamesOffsets[i] = builder.CreateString(name)
	}
	EPM.EPMStartALTERNATE_NAMESVector(builder, len(alternateNamesOffsets))
	for i := len(alternateNamesOffsets) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(alternateNamesOffsets[i])
	}
	alternateNamesVec := builder.EndVector(len(alternateNamesOffsets))

	// Start the Address table
	EPM.AddressStart(builder)
	EPM.AddressAddCOUNTRY(builder, countryOffset)
	EPM.AddressAddREGION(builder, regionOffset)
	EPM.AddressAddLOCALITY(builder, localityOffset)
	EPM.AddressAddPOSTAL_CODE(builder, postalCodeOffset)
	EPM.AddressAddSTREET(builder, streetOffset)
	EPM.AddressAddPOST_OFFICE_BOX_NUMBER(builder, poBoxOffset)
	addressOffset := EPM.AddressEnd(builder)

	// Start the EPM table
	EPM.EPMStart(builder)
	EPM.EPMAddDN(builder, dnOffset)
	EPM.EPMAddLEGAL_NAME(builder, legalNameOffset)
	EPM.EPMAddFAMILY_NAME(builder, familyNameOffset)
	EPM.EPMAddGIVEN_NAME(builder, givenNameOffset)
	EPM.EPMAddADDITIONAL_NAME(builder, additionalNameOffset)
	EPM.EPMAddHONORIFIC_PREFIX(builder, honorificPrefixOffset)
	EPM.EPMAddHONORIFIC_SUFFIX(builder, honorificSuffixOffset)
	EPM.EPMAddJOB_TITLE(builder, jobTitleOffset)
	EPM.EPMAddOCCUPATION(builder, occupationOffset)
	EPM.EPMAddALTERNATE_NAMES(builder, alternateNamesVec)
	EPM.EPMAddEMAIL(builder, emailOffset)
	EPM.EPMAddTELEPHONE(builder, telephoneOffset)
	EPM.EPMAddADDRESS(builder, addressOffset)

	// Finish the EPM table
	epm := EPM.EPMEnd(builder)
	builder.Finish(epm)

	// Return the byte slice containing the EPM object
	return builder.FinishedBytes()
}

// createStringVector is a helper function that creates a vector of strings in the FlatBuffers builder.
func createStringVector(builder *flatbuffers.Builder, items []string) flatbuffers.UOffsetT {
	offsets := make([]flatbuffers.UOffsetT, len(items))
	for i, item := range items {
		offsets[i] = builder.CreateString(item)
	}

	builder.StartVector(4, len(items), 4) // Specify the size of each element to be 4 bytes (UOffsetT size)
	for i := len(offsets) - 1; i >= 0; i-- {
		builder.PrependUOffsetT(offsets[i])
	}

	return builder.EndVector(len(items))
}

func SerializeEPM(builder *flatbuffers.Builder, epm flatbuffers.UOffsetT) []byte {
	builder.FinishSizePrefixedWithFileIdentifier(epm, []byte(EPMFID))
	return builder.FinishedBytes()
}

func DeserializeEPM(ctx context.Context, stream io.Reader) (*EPM.EPM, error) {
	totalSizeBuf := make([]byte, 4)
	if _, err := io.ReadFull(stream, totalSizeBuf); err != nil {
		return nil, fmt.Errorf("failed to read total size prefix: %v", err)
	}
	totalSize := binary.LittleEndian.Uint32(totalSizeBuf)
	data := make([]byte, 0, totalSize)
	for uint32(len(data)) < totalSize {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			chunkSize := totalSize - uint32(len(data))
			chunk := make([]byte, chunkSize)
			n, err := io.ReadFull(stream, chunk)
			if err != nil {
				return nil, fmt.Errorf("failed to read EPM data: %v", err)
			}
			data = append(data, chunk[:n]...)
		}
	}

	fileID := string(data[4:8])
	if fileID != EPMFID {
		return nil, fmt.Errorf("unexpected file identifier: got %s, want %s", fileID, EPMFID)
	} else {
		fmt.Println(fileID)
	}

	epm := EPM.GetRootAsEPM(data, 0)
	return epm, nil
}

func ConvertTovCard(binaryEPM []byte) string {
	epm := EPM.GetRootAsEPM(binaryEPM, 0)

	card := vcard.Card{}
	versionField := &vcard.Field{Value: "4.0"}
	card.Set("VERSION", versionField)

	if dn := epm.DN(); dn != nil {
		card.Add("FN", &vcard.Field{Value: string(dn)})
	}

	if legalName := epm.LEGAL_NAME(); legalName != nil {
		card.Add("ORG", &vcard.Field{Value: string(legalName)})
	}

	if email := epm.EMAIL(); email != nil {
		card.Add("EMAIL", &vcard.Field{Value: string(email)})
	}

	if telephone := epm.TELEPHONE(); telephone != nil {
		card.Add("TEL", &vcard.Field{Value: string(telephone)})
	}

	address := new(EPM.Address)
	epm.ADDRESS(address) // This populates the 'address' object with data

	// Initialize a slice to hold address components
	addrComponents := []string{}

	// Helper function to safely add address components if they exist
	addIfNotNil := func(b []byte) {
		if b != nil {
			addrComponents = append(addrComponents, string(b))
		}
	}

	// Safely add address components using the helper function
	addIfNotNil(address.STREET())
	addIfNotNil(address.POST_OFFICE_BOX_NUMBER())
	addIfNotNil(address.LOCALITY())
	addIfNotNil(address.REGION())
	addIfNotNil(address.POSTAL_CODE())
	addIfNotNil(address.COUNTRY())

	// Only add the ADR field to the card if there are non-empty address components
	if len(addrComponents) > 0 {
		card.Add("ADR", &vcard.Field{Value: strings.Join(addrComponents, ";")})
	}

	familyNameFB := epm.FAMILY_NAME()
	givenNameFB := epm.GIVEN_NAME()
	familyName := ""
	givenName := ""

	if familyNameFB != nil {
		familyName = string(familyNameFB)
	}

	if givenNameFB != nil {
		givenName = string(givenNameFB)
	}

	if familyName != "" || givenName != "" {
		additionalNameFB := epm.ADDITIONAL_NAME()
		honorificPrefixFB := epm.HONORIFIC_PREFIX()
		honorificSuffixFB := epm.HONORIFIC_SUFFIX()

		additionalName := ""
		honorificPrefix := ""
		honorificSuffix := ""

		if additionalNameFB != nil {
			additionalName = string(additionalNameFB)
		}

		if honorificPrefixFB != nil {
			honorificPrefix = string(honorificPrefixFB)
		}

		if honorificSuffixFB != nil {
			honorificSuffix = string(honorificSuffixFB)
		}

		n := []string{familyName, givenName, additionalName, honorificPrefix, honorificSuffix}
		card.Add("N", &vcard.Field{Value: strings.Join(n, ";")})
	}
	if jobTitle := epm.JOB_TITLE(); jobTitle != nil {
		card.Add("TITLE", &vcard.Field{Value: string(jobTitle)})
	}

	if occupation := epm.OCCUPATION(); occupation != nil {
		card.Add("ROLE", &vcard.Field{Value: string(occupation)})
	}

	alternateNamesLen := epm.ALTERNATE_NAMESLength() // Get the number of alternate names

	for i := 0; i < alternateNamesLen; i++ {
		nameBytes := epm.ALTERNATE_NAMES(i) // Get the alternate name as a byte slice at index i
		if nameBytes != nil {
			name := string(nameBytes) // Convert the byte slice to a string
			card.Add("X-ALTERNATE-NAME", &vcard.Field{Value: name})
		}
	}

	// Implement the logic to add KEYS and MULTIFORMAT_ADDRESS to the card as per your requirements

	// Convert the vCard to a string
	var b strings.Builder
	enc := vcard.NewEncoder(&b)
	if err := enc.Encode(card); err != nil {
		panic(err)
	}

	return b.String()
}