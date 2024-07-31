package set

import "github.com/CrestFallenTurtle/go-evil/utility/structure/json"

// Sets the js that wil be used
func Javascript(js_content string, data_object *json.Json_t) {

	data_object.Set_js(js_content)

}
