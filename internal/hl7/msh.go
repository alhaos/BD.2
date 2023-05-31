package hl7

import (
	"fmt"
	"strings"
)

type MSH struct {
	FieldSeparator                string   // Поле 1: Разделитель полей (Field Separator)
	EncodingCharacters            string   // Поле 2: Символы кодирования (Encoding Characters)
	SendingApplication            string   // Поле 3: Отправляющее приложение (Sending Application)
	SendingFacility               string   // Поле 4: Отправляющая организация (Sending Facility)
	ReceivingApplication          string   // Поле 5: Получающее приложение (Receiving Application)
	ReceivingFacility             string   // Поле 6: Получающая организация (Receiving Facility)
	DateTimeOfMessage             string   // Поле 7: Дата/время сообщения (Date/Time Of Message)
	Security                      string   // Поле 8: Безопасность (Security)
	MessageType                   []string // Поле 9: Тип сообщения (Message Type)
	MessageControlID              string   // Поле 10: Идентификатор управления сообщением (Message Control ID)
	ProcessingID                  string   // Поле 11: ID обработки (Processing ID)
	VersionID                     string   // Поле 12: ID версии (Version ID)
	SequenceNumber                string   // Поле 13: Порядковый номер (Sequence Number)
	ContinuationPointer           string   // Поле 14: Указатель продолжения (Continuation Pointer)
	AcceptAcknowledgement         string   // Поле 15: Принятие подтверждения (Accept Acknowledgement)
	ApplicationAcknowledgmentType string   // Поле 16: Тип подтверждения приложения (Application Acknowledgment Type)
}

func NewMSH(input string) (*MSH, error) {
	fields := strings.Split(input, "|")
	if len(fields) < 16 {
		return nil, fmt.Errorf("input does not contain enough fields")
	}

	return &MSH{
		FieldSeparator:                fields[0],
		EncodingCharacters:            fields[1],
		SendingApplication:            fields[2],
		SendingFacility:               fields[3],
		ReceivingApplication:          fields[4],
		ReceivingFacility:             fields[5],
		DateTimeOfMessage:             fields[6],
		Security:                      fields[7],
		MessageType:                   SplitComponents(fields[8]),
		MessageControlID:              fields[9],
		ProcessingID:                  fields[10],
		VersionID:                     fields[11],
		SequenceNumber:                fields[12],
		ContinuationPointer:           fields[13],
		AcceptAcknowledgement:         fields[14],
		ApplicationAcknowledgmentType: fields[15],
	}, nil
}

func (m *MSH) String() string {

	messageType := strings.Join(m.MessageType, "^")

	return fmt.Sprintf("%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s|%s",
		m.FieldSeparator,
		m.EncodingCharacters,
		m.SendingApplication,
		m.SendingFacility,
		m.ReceivingApplication,
		m.ReceivingFacility,
		m.DateTimeOfMessage,
		m.Security,
		messageType,
		m.MessageControlID,
		m.ProcessingID,
		m.VersionID,
		m.SequenceNumber,
		m.ContinuationPointer,
		m.AcceptAcknowledgement,
		m.ApplicationAcknowledgmentType,
	)
}
