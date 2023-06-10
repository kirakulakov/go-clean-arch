// Пакет entity определяет основные сущности для бизнес-логики (сервисов), отображения базы данных и
// объектов HTTP-ответа, если это уместно. Каждая группа логики имеет собственный файл.

package entity

type Translation struct {
	Source      string `json:"source" example:"auto"`
	Destination string `json:"destination" example:"en"`
	Original    string `json:"original" example:"текст для перевода"`
	Translation string `json:"translation" example:"text for translation"`
}
