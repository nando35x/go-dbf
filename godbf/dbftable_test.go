package godbf

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

const testEncoding = "UTF-8"

// For reference: https://www.dbase.com/Knowledgebase/INT/db7_file_fmt.htm

func TestDbfTable_New(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeZero())
}

func TestDbfTable_AddBooleanField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testBool"
	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("L"))
}

func TestDbfTable_AddBooleanField_TooLongGetsTruncated(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "FieldName!"
	suppliedFieldName := expectedFieldName + "shouldBeTruncated"

	tableUnderTest.AddBooleanField(suppliedFieldName)

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
}

func TestDbfTable_AddBooleanField_SecondAttemptFails(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "FieldName!"

	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	secondAdditionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(secondAdditionError).ToNot(BeNil())

	t.Log(secondAdditionError)
}

func TestDbfTable_AddBooleanField_ErrorAfterDataEntryStart(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "goodField"

	additionError := tableUnderTest.AddBooleanField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	tableUnderTest.AddNewRecord()

	postDataEntryField := "badField"

	secondAdditionError := tableUnderTest.AddBooleanField(postDataEntryField)
	g.Expect(secondAdditionError).ToNot(BeNil())

	t.Log(secondAdditionError)
}

func TestDbfTable_AddDateField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testDate"
	additionError := tableUnderTest.AddDateField(expectedFieldName)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("D"))
}

func TestDbfTable_AddTextField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testText"
	expectedFieldLength := uint8(20)
	additionError := tableUnderTest.AddTextField(expectedFieldName, expectedFieldLength)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("C"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
}

func TestDbfTable_AddNumberField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testNumber"
	expectedFieldLength := uint8(20)
	expectedFDecimalPlaces := uint8(2)
	additionError := tableUnderTest.AddNumberField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("N"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
	g.Expect(addedField.fieldDecimalPlaces).To(Equal(expectedFDecimalPlaces))
}

func TestDbfTable_AddFloatField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)
	expectedFieldName := "testFloat"
	expectedFieldLength := uint8(20)
	expectedFDecimalPlaces := uint8(2)
	additionError := tableUnderTest.AddFloatField(expectedFieldName, expectedFieldLength, expectedFDecimalPlaces)
	g.Expect(additionError).To(BeNil())

	g.Expect(tableUnderTest.NumberOfRecords()).To(BeZero())
	g.Expect(len(tableUnderTest.Fields())).To(BeNumerically("==", 1))

	addedField := tableUnderTest.Fields()[0]
	g.Expect(addedField.fieldName).To(Equal(expectedFieldName))
	g.Expect(addedField.fieldType).To(Equal("F"))
	g.Expect(addedField.fieldLength).To(Equal(expectedFieldLength))
	g.Expect(addedField.fieldDecimalPlaces).To(Equal(expectedFDecimalPlaces))
}

func TestDbfTable_FieldNames(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	expectedFieldNames := []string{"first", "second"}

	for _, name := range expectedFieldNames {
		additionError := tableUnderTest.AddBooleanField(name)
		g.Expect(additionError).To(BeNil())
	}

	fieldNamesUnderTest := tableUnderTest.FieldNames()
	g.Expect(fieldNamesUnderTest).To(Equal(expectedFieldNames))
}

func TestDbfTable_DecimalPlacesInField_ValidField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	numberFieldName := "numField"
	expectedNumberDecimalPlaces := uint8(0)
	tableUnderTest.AddNumberField(numberFieldName, 5, expectedNumberDecimalPlaces)
	actualNumberDecimalPlaces, numberError := tableUnderTest.DecimalPlacesInField(numberFieldName)

	g.Expect(numberError).To(BeNil())
	g.Expect(actualNumberDecimalPlaces).To(BeNumerically("==", expectedNumberDecimalPlaces))

	floatFieldName := "floatField"
	expectedFloatDecimalPlaces := uint8(2)
	tableUnderTest.AddFloatField(floatFieldName, 10, expectedFloatDecimalPlaces)
	actualFloatDecimalPlaces, floatError := tableUnderTest.DecimalPlacesInField(floatFieldName)

	g.Expect(floatError).To(BeNil())
	g.Expect(actualFloatDecimalPlaces).To(BeNumerically("==", expectedFloatDecimalPlaces))
}

func TestDbfTable_DecimalPlacesInField_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	_, numberError := tableUnderTest.DecimalPlacesInField("missingField")

	g.Expect(numberError).ToNot(BeNil())
	t.Log(numberError)
}

func TestDbfTable_DecimalPlacesInField_InvalidField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	textFieldName := "textField"
	tableUnderTest.AddTextField(textFieldName, 5)
	_, numberError := tableUnderTest.DecimalPlacesInField(textFieldName)

	g.Expect(numberError).ToNot(BeNil())
	t.Log(numberError)
}

func TestDbfTable_GetRowAsSlice_InitiallyEmptyStrings(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	booldFieldName := "boolField"
	tableUnderTest.AddBooleanField(booldFieldName)

	textFieldName := "textField"
	tableUnderTest.AddBooleanField(textFieldName)

	dateFieldName := "dateField"
	tableUnderTest.AddBooleanField(dateFieldName)

	numFieldName := "numField"
	tableUnderTest.AddBooleanField(numFieldName)

	floatFieldName := "floatField"
	tableUnderTest.AddBooleanField(floatFieldName)

	recordIndex := tableUnderTest.AddNewRecord()

	fieldValues := tableUnderTest.GetRowAsSlice(recordIndex)

	for _, value := range fieldValues {
		g.Expect(value).To(Equal(""))
	}
}

func TestDbfTable_GetRowAsSlice(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	boolFieldName := "boolField"
	expectedBoolFieldValue := "T"
	tableUnderTest.AddBooleanField(boolFieldName)

	textFieldName := "textField"
	expectedTextFieldValue := "some text"
	tableUnderTest.AddTextField(textFieldName, 10)

	dateFieldName := "dateField"
	expectedDateFieldValue := "20181201"
	tableUnderTest.AddDateField(dateFieldName)

	numFieldName := "numField"
	expectedNumFieldValue := "640"
	tableUnderTest.AddNumberField(numFieldName, 3, 0)

	floatFieldName := "floatField"
	expectedFloatFieldValue := "640.42"
	tableUnderTest.AddFloatField(floatFieldName, 6, 2)

	recordIndex := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, boolFieldName, expectedBoolFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, textFieldName, expectedTextFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, dateFieldName, expectedDateFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, numFieldName, expectedNumFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	fieldValues := tableUnderTest.GetRowAsSlice(recordIndex)

	g.Expect(fieldValues[0]).To(Equal(expectedBoolFieldValue))
	g.Expect(fieldValues[1]).To(Equal(expectedTextFieldValue))
	g.Expect(fieldValues[2]).To(Equal(expectedDateFieldValue))
	g.Expect(fieldValues[3]).To(Equal(expectedNumFieldValue))
	g.Expect(fieldValues[4]).To(Equal(expectedFloatFieldValue))
}

func TestDbfTable_FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	boolFieldName := "boolField"
	expectedBoolFieldValue := "T"
	tableUnderTest.AddBooleanField(boolFieldName)

	textFieldName := "textField"
	expectedTextFieldValue := "some text"
	tableUnderTest.AddTextField(textFieldName, 10)

	dateFieldName := "dateField"
	expectedDateFieldValue := "20181201"
	tableUnderTest.AddDateField(dateFieldName)

	numFieldName := "numField"
	expectedNumFieldValue := "640"
	tableUnderTest.AddNumberField(numFieldName, 3, 0)

	floatFieldName := "floatField"
	expectedFloatFieldValue := "640.42"
	tableUnderTest.AddFloatField(floatFieldName, 6, 2)

	recordIndex := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, boolFieldName, expectedBoolFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, textFieldName, expectedTextFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, dateFieldName, expectedDateFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, numFieldName, expectedNumFieldValue)
	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	g.Expect(tableUnderTest.FieldValueByName(recordIndex, boolFieldName)).To(Equal(expectedBoolFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, textFieldName)).To(Equal(expectedTextFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, dateFieldName)).To(Equal(expectedDateFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, numFieldName)).To(Equal(expectedNumFieldValue))
	g.Expect(tableUnderTest.FieldValueByName(recordIndex, floatFieldName)).To(Equal(expectedFloatFieldValue))
}

func TestDbfTable_FieldValueByName_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	recordIndex := tableUnderTest.AddNewRecord()

	_, error := tableUnderTest.FieldValueByName(recordIndex, "missingField")

	g.Expect(error).ToNot(BeNil())
	t.Log(error)
}

func TestDbfTable_SetFieldValueByName_NonExistentField(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	recordIndex := tableUnderTest.AddNewRecord()

	error := tableUnderTest.SetFieldValueByName(recordIndex, "missingField", "someText")

	g.Expect(error).ToNot(BeNil())
	t.Log(error)
}

func TestDbfTable_Int64FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	intFieldName := "intField"
	expectedIntValue := 640
	expectedIntFieldValue := fmt.Sprintf("%d", expectedIntValue)
	tableUnderTest.AddNumberField(intFieldName, 6, 2)

	recordIndex := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, intFieldName, expectedIntFieldValue)

	actualIntFieldValue, error := tableUnderTest.Int64FieldValueByName(recordIndex, intFieldName)

	g.Expect(error).To(BeNil())
	g.Expect(actualIntFieldValue).To(BeNumerically("==", expectedIntValue))
}

func TestDbfTable_Float64FieldValueByName(t *testing.T) {
	g := NewGomegaWithT(t)

	tableUnderTest := New(testEncoding)

	floatFieldName := "floatField"
	expectedFloatValue := 640.42
	expectedFloatFieldValue := fmt.Sprintf("%.2f", expectedFloatValue)
	tableUnderTest.AddFloatField(floatFieldName, 10, 2)

	recordIndex := tableUnderTest.AddNewRecord()

	tableUnderTest.SetFieldValueByName(recordIndex, floatFieldName, expectedFloatFieldValue)

	actualFloatFieldValue, error := tableUnderTest.Float64FieldValueByName(recordIndex, floatFieldName)

	g.Expect(error).To(BeNil())
	g.Expect(actualFloatFieldValue).To(BeNumerically("==", expectedFloatValue))
}
