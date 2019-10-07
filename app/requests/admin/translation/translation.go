package translation

import (
	"github.com/thedevsaddam/govalidator"
	"golang-starter/app/models"
	"golang-starter/helpers"
	"net/http"
)

/**
* validate store category request
 */
func StoreUpdate(r *http.Request , request *models.Translation) *govalidator.Validator {
	lang := helpers.GetCurrentLangFromHttp(r)
	/// Validation rules
	rules := govalidator.MapData{
		"value":    []string{"required", "min:2", "max:255"},
		"lang":  []string{"required", "min:2", "max:10"},
		"page":  []string{"required", "min:2", "max:30"},
		"slug":  []string{"required", "min:2", "max:50"},
	}

	messages := govalidator.MapData{
		"value": []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "255")},
		"lang":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "10")},
		"page":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "30")},
		"slug":  []string{helpers.Required(lang), helpers.Min(lang, "2"), helpers.Max(lang, "50")},
	}

	opts := govalidator.Options{
		Request: r,     // request object
		Rules:   rules, // rules map
		Data:    request,
		Messages:        messages, // custom message map (Optional)
		RequiredDefault: true, // all the field to be pass the rules
	}
	return govalidator.New(opts)
}