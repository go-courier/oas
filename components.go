package oas

import (
	"encoding/json"
	"fmt"
)

type Components struct {
	ComponentsObject
	SpecExtensions
}

func (i Components) MarshalJSON() ([]byte, error) {
	return flattenMarshalJSON(i.ComponentsObject, i.SpecExtensions)
}

func (i *Components) UnmarshalJSON(data []byte) error {
	return flattenUnmarshalJSON(data, &i.ComponentsObject, &i.SpecExtensions)
}

type ComponentsObject struct {
	Schemas       map[string]*Schema      `json:"schemas,omitempty"`
	Responses     map[string]*Response    `json:"responses,omitempty"`
	Parameters    map[string]*Parameter   `json:"parameters,omitempty"`
	WithExamples
	RequestBodies map[string]*RequestBody `json:"requestBodies,omitempty"`
	WithHeaders
	WithSecuritySchemes
	WithLinks
	WithCallbacks
}

func (object *ComponentsObject) AddSchema(id string, s *Schema) {
	if s == nil {
		return
	}
	if object.Schemas == nil {
		object.Schemas = make(map[string]*Schema)
	}
	object.Schemas[id] = s
}

func (object *ComponentsObject) AddResponse(id string, r *Response) {
	if r == nil {
		return
	}
	if object.Responses == nil {
		object.Responses = make(map[string]*Response)
	}
	object.Responses[id] = r
}

func (object *ComponentsObject) AddParameter(id string, p *Parameter) {
	if p == nil {
		return
	}
	if object.Parameters == nil {
		object.Parameters = make(map[string]*Parameter)
	}
	object.Parameters[id] = p
}

func (object *ComponentsObject) AddRequestBody(id string, e *RequestBody) {
	if e == nil {
		return
	}
	if object.RequestBodies == nil {
		object.RequestBodies = make(map[string]*RequestBody)
	}
	object.RequestBodies[id] = e
}

func (object *ComponentsObject) RefString(group string, id string) string {
	return fmt.Sprintf("#/components/%s/%s", group, id)
}

func (object *ComponentsObject) RefSchema(id string) *Schema {
	if object.Schemas == nil || object.Schemas[id] == nil {
		return nil
	}
	return RefSchema(object.RefString("schemas", id))
}

func (object *ComponentsObject) RefResponse(id string) *Response {
	if object.Responses == nil || object.Responses[id] == nil {
		return nil
	}
	s := &Response{}
	s.Ref = object.RefString("responses", id)
	return s
}

func (object *ComponentsObject) RefParameter(id string) *Parameter {
	if object.Parameters == nil || object.Parameters[id] == nil {
		return nil
	}
	s := &Parameter{}
	s.Ref = object.RefString("parameters", id)
	return s
}

func (object *ComponentsObject) RefExample(id string) *Example {
	if object.Examples == nil || object.Examples[id] == nil {
		return nil
	}
	s := &Example{}
	s.Ref = object.RefString("examples", id)
	return s
}

func (object *ComponentsObject) RefRequestBody(id string) *RequestBody {
	if object.RequestBodies == nil || object.RequestBodies[id] == nil {
		return nil
	}
	s := &RequestBody{}
	s.Ref = object.RefString("requestBodies", id)
	return s
}

func (object *ComponentsObject) RefHeader(id string) *Header {
	if object.Headers == nil || object.Headers[id] == nil {
		return nil
	}
	s := &Header{}
	s.Ref = object.RefString("headers", id)
	return s
}

func (object *ComponentsObject) RefLink(id string) *Link {
	if object.Links == nil || object.Headers[id] == nil {
		return nil
	}
	s := &Link{}
	s.Ref = object.RefString("links", id)
	return s
}

func (object *ComponentsObject) RefCallback(id string) *Callback {
	if object.Callbacks == nil || object.Callbacks[id] == nil {
		return nil
	}
	s := &Callback{}
	s.Ref = object.RefString("callbacks", id)
	return s
}

func (object *ComponentsObject) RequireSecurity(id string, scopes ...string) SecurityRequirement {
	if object.SecuritySchemes == nil || object.SecuritySchemes[id] == nil {
		return nil
	}
	ss := object.SecuritySchemes[id]
	if ss.Type == SecurityTypeOAuth2 {
		return SecurityRequirement{
			id: scopes,
		}
	}
	return SecurityRequirement{
		id: []string{},
	}
}

type Reference struct {
	Ref string `json:"$ref,omitempty"`
}

func (ref Reference) MarshalJSONRefFirst(values ...interface{}) ([]byte, error) {
	if ref.Ref != "" {
		return json.Marshal(ref)
	}
	return flattenMarshalJSON(values...)
}

func (ref *Reference) UnmarshalJSONRefFirst(data []byte, values ...interface{}) error {
	if err := json.Unmarshal(data, &ref); err != nil {
		return err
	}
	if ref.Ref != "" {
		return nil
	}
	return flattenUnmarshalJSON(data, values...)
}
