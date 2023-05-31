package hl7

import (
	"fmt"
	"strings"
)

type PID struct {
	SetIDPID              string   // PID-1
	PatientID             []string // PID-2
	PatientIdentifierList []string // PID-3
	AlternatePatientID    []string // PID-4
	PatientName           []string // PID-5
	MotherMaidenName      []string // PID-6
	DateTimeOfBirth       string   // PID-7
	Sex                   string   // PID-8
	PatientAlias          []string // PID-9
	Race                  []string // PID-10
	PatientAddress        []string // PID-11
	CountryCode           string   // PID-12
	PhoneNumberHome       []string // PID-13
	PhoneNumberBusiness   []string // PID-14
	PrimaryLanguage       string   // PID-15
	MaritalStatus         string   // PID-16
	Religion              string   // PID-17
	PatientAccountNumber  string   // PID-18
	SSNNumberPatient      string   // PID-19
	DriversLicenseNumber  string   // PID-20
}

func NewPID(line string) (*PID, error) {
	fields := strings.Split(line, "|")
	if len(fields) < 20 {
		return nil, fmt.Errorf("not enough fields in line: %d", len(fields))
	}

	pid := &PID{
		SetIDPID:              fields[0],
		PatientID:             SplitComponents(fields[1]),
		PatientIdentifierList: SplitComponents(fields[2]),
		AlternatePatientID:    SplitComponents(fields[3]),
		PatientName:           SplitComponents(fields[4]),
		MotherMaidenName:      SplitComponents(fields[5]),
		DateTimeOfBirth:       fields[6],
		Sex:                   fields[7],
		PatientAlias:          SplitComponents(fields[8]),
		Race:                  SplitComponents(fields[9]),
		PatientAddress:        SplitComponents(fields[10]),
		CountryCode:           fields[11],
		PhoneNumberHome:       SplitComponents(fields[12]),
		PhoneNumberBusiness:   SplitComponents(fields[13]),
		PrimaryLanguage:       fields[14],
		MaritalStatus:         fields[15],
		Religion:              fields[16],
		PatientAccountNumber:  fields[17],
		SSNNumberPatient:      fields[18],
		DriversLicenseNumber:  fields[19],
	}

	return pid, nil
}

// String method for PID segment.
func (pid *PID) String() string {

	result := make([]string, 0)

	result = append(result, pid.SetIDPID)
	result = append(result, strings.Join(pid.PatientID, "^"))
	result = append(result, strings.Join(pid.PatientIdentifierList, "^"))
	result = append(result, strings.Join(pid.AlternatePatientID, "^"))
	result = append(result, strings.Join(pid.PatientName, "^"))
	result = append(result, strings.Join(pid.MotherMaidenName, "^"))
	result = append(result, pid.DateTimeOfBirth)
	result = append(result, pid.Sex)
	result = append(result, strings.Join(pid.PatientAlias, "^"))
	result = append(result, strings.Join(pid.Race, "^"))
	result = append(result, strings.Join(pid.PatientAddress, "^"))
	result = append(result, pid.CountryCode)
	result = append(result, strings.Join(pid.PhoneNumberHome, "^"))
	result = append(result, strings.Join(pid.PhoneNumberBusiness, "^"))
	result = append(result, pid.PrimaryLanguage)
	result = append(result, pid.MaritalStatus)
	result = append(result, pid.Religion)
	result = append(result, pid.PatientAccountNumber)
	result = append(result, pid.SSNNumberPatient)
	result = append(result, pid.DriversLicenseNumber)

	return strings.Join(result, "|")
}
