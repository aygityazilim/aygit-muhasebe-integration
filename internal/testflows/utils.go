package testflows

import (
	"strings"
)

// replaceUUIDinXML - Basit UBL UUID değiştirme yardımcı fonksiyonu (Testler için)
func replaceUUIDinXML(xmlStr, newUUID string) string {
	// Eski UUID'yi <cbc:UUID> etiketinin içeriğinden bularak bulup, tüm geçişlerini (AdditionalDocumentReference vb. dahil) değiştiririz.
	startIndex := strings.Index(xmlStr, "<cbc:UUID>")
	endIndex := strings.Index(xmlStr, "</cbc:UUID>")

	if startIndex != -1 && endIndex != -1 {
		// Extract just the old UUID string without the tags
		oldUUID := xmlStr[startIndex+len("<cbc:UUID>") : endIndex]
		// Replace all occurrences of this exact old UUID with the new one
		return strings.ReplaceAll(xmlStr, oldUUID, newUUID)
	}
	return xmlStr
}
