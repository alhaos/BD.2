package hl7

import (
	"fmt"
	"strings"
)

// PV1 represents the HL7 v2.4 Patient Visit (PV1) segment structure
type PV1 struct {
	SetID                   string   // PV1-1: Set ID - PV1
	PatientClass            string   // PV1-2: Patient Class
	AssignedPatientLocation string   // PV1-3: Assigned Patient Location
	AdmissionType           string   // PV1-4: Admission Type
	PreadmitNumber          string   // PV1-5: Preadmit Number
	PriorPatientLocation    []string // PV1-6: Prior Patient Location
	AttendingDoctor         string   // PV1-7: Attending Doctor
	ReferringDoctor         string   // PV1-8: Referring Doctor
	ConsultingDoctor        string   // PV1-9: Consulting Doctor
	HospitalService         string   // PV1-10: Hospital Service
	TemporaryLocation       string   // PV1-11: Temporary Location
	PreadmitTestIndicator   string   // PV1-12: Preadmit Test Indicator
	ReAdmissionIndicator    string   // PV1-13: Re-admission Indicator
	AdmitSource             string   // PV1-14: Admit Source
	AmbulatoryStatus        string   // PV1-15: Ambulatory Status
	VIPIndicator            string   // PV1-16: VIP Indicator
	AdmittingDoctor         string   // PV1-17: Admitting Doctor
	PatientType             string   // PV1-18: Patient Type
	VisitNumber             string   // PV1-19: Visit Number
	FinancialClass          string   // PV1-20: Financial Class
}

func NewPV1(line string) (*PV1, error) {
	// Fields are separated by '|'
	fields := strings.Split(line, "|")
	if len(fields) < 53 {
		return nil, fmt.Errorf("Invalid PV1 line, less than 53 fields: %s", line)
	}

	pv1 := &PV1{
		SetID:                   fields[1],
		PatientClass:            fields[2],
		AssignedPatientLocation: fields[3],
		AdmissionType:           fields[4],
		PreadmitNumber:          fields[5],
		PriorPatientLocation:    SplitComponents(fields[6]),
		AttendingDoctor:         fields[7],
		ReferringDoctor:         fields[8],
		ConsultingDoctor:        fields[9],
		HospitalService:         fields[10],
		TemporaryLocation:       fields[11],
		PreadmitTestIndicator:   fields[12],
		ReAdmissionIndicator:    fields[13],
		AdmitSource:             fields[14],
		AmbulatoryStatus:        fields[15],
		VIPIndicator:            fields[16],
		AdmittingDoctor:         fields[17],
		PatientType:             fields[18],
		VisitNumber:             fields[19],
		FinancialClass:          fields[20],
	}

	return pv1, nil
}

func (pv1 *PV1) ToString() string {
	// Collect all fields into a slice
	fields := []string{
		"PV1", // Segment Identifier
		pv1.SetID,
		pv1.PatientClass,
		pv1.AssignedPatientLocation,
		pv1.AdmissionType,
		pv1.PreadmitNumber,
		strings.Join(pv1.PriorPatientLocation, "^"),
		pv1.AttendingDoctor,
		pv1.ReferringDoctor,
		pv1.ConsultingDoctor,
		pv1.HospitalService,
		pv1.TemporaryLocation,
		pv1.PreadmitTestIndicator,
		pv1.ReAdmissionIndicator,
		pv1.AdmitSource,
		pv1.AmbulatoryStatus,
		pv1.VIPIndicator,
		pv1.AdmittingDoctor,
		pv1.PatientType,
		pv1.VisitNumber,
		pv1.FinancialClass,
	}

	// Join the fields with '|'
	return strings.Join(fields, "|")
}
