package go_vyze

import (
	"errors"
	"reflect"
	"strings"
)

// Types and constants

// FieldType represents the source of a value for a field.
type FieldType string

const (
	FieldTypeID      FieldType = "id"
	FieldTypeName    FieldType = "name"
	FieldTypeSize    FieldType = "size"
	FieldTypeCreated FieldType = "created"
	FieldTypeUser    FieldType = "user"
	FieldTypeData    FieldType = "data"
)

// FormatType specifies the way of mapping binary data to or from JSON values.
type FormatType string

const (
	FormatTypeRaw     FormatType = "raw"
	FormatTypeHex     FormatType = "hex"
	FormatTypeBase64  FormatType = "base64"
	FormatTypeString  FormatType = "string"
	FormatTypeInteger FormatType = "integer"
	FormatTypeFloat   FormatType = "float"
	FormatTypeBoolean FormatType = "boolean"
)

// ReturnType specifies how data from multiple relations should be returned.
type ReturnType string

const (
	ReturnTypeObjects        ReturnType = "objects"
	ReturnTypeRelations      ReturnType = "relations"
	ReturnTypePairs          ReturnType = "pairs"
	ReturnTypeMap            ReturnType = "map"
	ReturnTypeKeyedObjects   ReturnType = "keyed_objects"
	ReturnTypeKeyedRelations ReturnType = "keyed_relations"
	ReturnTypeKeyedPairs     ReturnType = "keyed_pairs"
)

// MappingType specifies how multiple relations should be mapped onto or from JSON values.
type MappingType string

const (
	MappingTypePrimitive MappingType = "primitive"
	MappingTypeList      MappingType = "list"
)

// ResourceType specifies the type of source of one or multiple instances.
type ResourceType string

const (
	ResourceTypeSpecials ResourceType = "specials"
	ResourceTypeInstance ResourceType = "instance"
)

// Write resource

// WriteResource obtains a nested structure composed of maps, slices and primitive values from a Go variable.
func WriteResource(obj any) (any, error) {
	objType := reflect.TypeOf(obj)
	valType := reflect.ValueOf(obj)
	for {
		objKind := objType.Kind()
		if objKind == reflect.Pointer {
			objType = objType.Elem()
			valType = valType.Elem()
			obj = valType.Interface()
			continue
		}
		break
	}
	switch objType.Kind() {
	case reflect.Struct:
		return writeResourceStruct(obj)
	case reflect.Slice:
		return writeResourceSlice(obj)
	case reflect.String:
		return valType.Interface().(string), nil
	case reflect.Bool:
		return valType.Interface().(bool), nil
	case reflect.Float32:
		return valType.Interface().(float32), nil
	case reflect.Float64:
		return valType.Interface().(float64), nil
	case reflect.Int:
		return float64(valType.Interface().(int)), nil
	case reflect.Int8:
		return float64(valType.Interface().(int8)), nil
	case reflect.Int16:
		return float64(valType.Interface().(int16)), nil
	case reflect.Int32:
		return float64(valType.Interface().(int32)), nil
	case reflect.Int64:
		return float64(valType.Interface().(int64)), nil
	case reflect.Uint:
		return float64(valType.Interface().(uint)), nil
	case reflect.Uint8:
		return float64(valType.Interface().(uint8)), nil
	case reflect.Uint16:
		return float64(valType.Interface().(uint16)), nil
	case reflect.Uint32:
		return float64(valType.Interface().(uint32)), nil
	case reflect.Uint64:
		return float64(valType.Interface().(uint64)), nil
	case reflect.Array:
		id := valType.Interface().(ID)
		return id.Hex(), nil
	default:
		return valType.Interface(), nil
	}
}

// MustWriteResource behaves like WriteResource but panics in case of an error.
func MustWriteResource(obj any) any {
	val, err := WriteResource(obj)
	if err != nil {
		panic(err)
	}
	return val
}

func writeResourceStruct(obj any) (map[string]any, error) {
	val := map[string]any{}
	objType := reflect.TypeOf(obj)
	valType := reflect.ValueOf(obj)
	objKind := valType.Kind()
	if objKind != reflect.Struct {
		return nil, errors.New("require struct")
	}
	for i := 0; i < valType.NumField(); i++ {
		objFieldValue := valType.Field(i)
		field, err := fieldFromTag(objType.Field(i))
		if err != nil {
			continue
		}
		val[field.Name], _ = WriteResource(objFieldValue.Interface())
	}
	return val, nil
}

func writeResourceSlice(obj any) ([]any, error) {
	val := []any{}
	valType := reflect.ValueOf(obj)
	objKind := valType.Kind()
	if objKind != reflect.Slice {
		return nil, errors.New("require slice")
	}
	for i := 0; i < valType.Len(); i++ {
		objFieldValue := valType.Index(i)
		v, _ := WriteResource(objFieldValue.Interface())
		val = append(val, v)
	}
	return val, nil
}

// Read resource

// ReadResource turns a nested structure composed of maps, slices and primitive values into a Go variable.
func ReadResource(val any, obj any) error {
	objType := reflect.TypeOf(obj)
	valType := reflect.ValueOf(obj)
	return readResourceVal(val, objType, valType)
}

// MustReadResource behaves similar to ReadResource but creates a new instance from the passed type and panics in case
// of an error.
func MustReadResource[T any](val any, obj T) T {
	objType := reflect.TypeOf(obj)
	objVal := reflect.New(objType)
	if err := ReadResource(val, objVal.Interface()); err != nil {
		panic(err)
	}
	return objVal.Elem().Interface().(T)
}

func readResourceVal(val any, objType reflect.Type, valType reflect.Value) error {
	if val == nil {
		return nil
	}
	for {
		objKind := objType.Kind()
		if objKind == reflect.Pointer {
			objType = objType.Elem()
			valType = valType.Elem()
			continue
		}
		break
	}
	switch objType.Kind() {
	case reflect.Struct:
		_ = readResourceStruct(val.(map[string]any), valType, objType)
	case reflect.Slice:
		_ = readResourceSlice(val.([]any), valType, objType)
	case reflect.String:
		valType.SetString(val.(string))
	case reflect.Bool:
		valType.SetBool(val.(bool))
	case reflect.Float32:
		valType.Set(reflect.ValueOf(float32(val.(float64))))
	case reflect.Float64:
		valType.Set(reflect.ValueOf(val.(float64)))
	case reflect.Int:
		valType.Set(reflect.ValueOf(int(val.(float64))))
	case reflect.Int8:
		valType.Set(reflect.ValueOf(int8(val.(float64))))
	case reflect.Int16:
		valType.Set(reflect.ValueOf(int16(val.(float64))))
	case reflect.Int32:
		valType.Set(reflect.ValueOf(int32(val.(float64))))
	case reflect.Int64:
		valType.Set(reflect.ValueOf(int64(val.(float64))))
	case reflect.Uint:
		valType.Set(reflect.ValueOf(uint(val.(float64))))
	case reflect.Uint8:
		valType.Set(reflect.ValueOf(uint8(val.(float64))))
	case reflect.Uint16:
		valType.Set(reflect.ValueOf(uint16(val.(float64))))
	case reflect.Uint32:
		valType.Set(reflect.ValueOf(uint32(val.(float64))))
	case reflect.Uint64:
		valType.Set(reflect.ValueOf(uint64(val.(float64))))
	case reflect.Array:
		id, _ := ParseID(val.(string))
		valType.Set(reflect.ValueOf(id))
	default:
		valType.Set(reflect.ValueOf(val))
	}
	return nil
}

func readResourceStruct(val map[string]any, valType reflect.Value, objType reflect.Type) error {
	objKind := valType.Kind()
	if objKind != reflect.Struct {
		return errors.New("require struct")
	}
	for i := 0; i < valType.NumField(); i++ {
		objFieldValue := valType.Field(i)
		objFieldType := objType.Field(i).Type
		if !objFieldValue.CanSet() {
			continue
		}
		field, err := fieldFromTag(objType.Field(i))
		if err != nil {
			continue
		}
		ptr := false
		if objFieldValue.Kind() == reflect.Pointer {
			objFieldType = objFieldType.Elem()
			ptr = true
		}
		obj := reflect.New(objFieldType)
		_ = readResourceVal(val[field.Name], objFieldType, obj.Elem())
		if !ptr {
			obj = obj.Elem()
		}
		objFieldValue.Set(obj)
	}
	return nil
}

func readResourceSlice(val []any, valType reflect.Value, objType reflect.Type) error {
	objKind := objType.Kind()
	if objKind != reflect.Slice {
		return errors.New("require slice")
	}
	objSlice := reflect.MakeSlice(objType, len(val), len(val))
	ptr := false
	objType = objType.Elem()
	if objType.Kind() == reflect.Pointer {
		objType = objType.Elem()
		ptr = true
	}
	for i := 0; i < len(val); i++ {
		objFieldValue := objSlice.Index(i)
		obj := reflect.New(objType)
		_ = readResourceVal(val[i], objType, obj.Elem())
		if !ptr {
			obj = obj.Elem()
		}
		objFieldValue.Set(obj)
	}
	valType.Set(objSlice)
	return nil
}

// Extract from struct

// ExtractSchema extracts a ResourceSchema instance from a go structure.
func ExtractSchema(obj any, resolver Resolver, base string) (ResourceSchema, error) {
	objType := reflect.TypeOf(obj)
	objKind := objType.Kind()
	for {
		if objKind == reflect.Pointer || objKind == reflect.Slice {
			objType = objType.Elem()
			objKind = objType.Kind()
			continue
		}
		if objKind != reflect.Struct {
			return ResourceSchema{}, errors.New("require struct")
		}
		break
	}
	rs := ResourceSchema{}
	for i := 0; i < objType.NumField(); i++ {
		objFieldType := objType.Field(i)
		field, err := fieldFromTag(objFieldType)
		if err != nil {
			continue
		}
		rs.Fields = append(rs.Fields, field)
	}
	if resolver != nil {
		if err := rs.Resolve(resolver, base); err != nil {
			return ResourceSchema{}, err
		}
	}
	return rs, nil
}

// MustExtractSchema behaves like ExtractSchema but panics in case of an error.
func MustExtractSchema(obj any, univ *Universe, base string) *ResourceSchema {
	sch, err := ExtractSchema(obj, univ, base)
	if err != nil {
		return nil
	}
	return &sch
}

// Resource schema

// ResourceMapping represents the type of mapping used to map object or relation field values onto a JSON object.
type ResourceMapping struct {
	ValueType   MappingType
	ValueParams map[string]any
}

func NewPrimitiveMapping() ResourceMapping {
	return ResourceMapping{
		ValueType:   MappingTypePrimitive,
		ValueParams: map[string]any{},
	}
}

func NewListMapping() ResourceMapping {
	return ResourceMapping{
		ValueType:   MappingTypeList,
		ValueParams: map[string]any{},
	}
}

func (rm *ResourceMapping) Set(key string, value any) *ResourceMapping {
	rm.ValueParams[key] = value
	return rm
}

// ResourceField represents a key in the JSON object to be rendered to or from a ResourceSchema.
type ResourceField struct {
	// Name is the name of the field as it appears in the
	Name string `json:"name"`

	// Relation is the ID or the identifier of the relation of this field
	Relation *string `json:"relation,omitempty"`

	// Mapping is the type of mapping for this field
	Mapping MappingType `json:"mapping"`

	// MappingParams are the parameters of the mapping specified in Mapping
	MappingParams map[string]any `json:"mappingParams,omitempty"`

	// Field is the type of this field
	Field FieldType `json:"field"`

	// Format is the format used for the values carried by this field
	Format FormatType `json:"format"`
}

func (f ResourceField) GetMappingParam(key string, def string) any {
	val, ok := f.MappingParams[key]
	if !ok {
		return def
	}
	return val
}

func (f *ResourceField) SetMappingParam(key string, val any) {
	if f.MappingParams == nil {
		f.MappingParams = map[string]any{}
	}
	f.MappingParams[key] = val
}

type ResourceFilters []ResourceFilter

type ResourceFilter struct {
	Name     string       `json:"name"`
	Value    interface{}  `json:"value"`
	Operator OperatorType `json:"operator"`
}

type ResourceOrder struct {
	Name      string `json:"name"`
	Ascending bool   `json:"ascending"`
}

type ResourceOrders []ResourceOrder

// ResourceSchema represents a mapping between JSON objects and VYZE objects.
type ResourceSchema struct {
	Fields []ResourceField `json:"fields"`
}

// NewResourceSchema creates a new ResourceSchema instance.
func NewResourceSchema() *ResourceSchema {
	return &ResourceSchema{
		Fields: []ResourceField{},
	}
}

// GetField returns a field of this schema identified by its name.
func (r ResourceSchema) GetField(name string) *ResourceField {
	for _, f := range r.Fields {
		if f.Name == name {
			return &f
		}
	}
	return nil
}

// AddObjectField attaches a new object field to the resource schema.
func (r *ResourceSchema) AddObjectField(name string, fieldType FieldType, formatType FormatType) *ResourceSchema {
	r.Fields = append(r.Fields, ResourceField{
		Name:          name,
		Mapping:       "primitive",
		MappingParams: nil,
		Field:         fieldType,
		Format:        formatType,
	})
	return r
}

// AddRelationField attaches a new relation field to the resource schema.
func (r *ResourceSchema) AddRelationField(name string, relation string, fieldType FieldType, formatType FormatType, mapping ResourceMapping) *ResourceSchema {
	r.Fields = append(r.Fields, ResourceField{
		Name:          name,
		Mapping:       mapping.ValueType,
		MappingParams: mapping.ValueParams,
		Relation:      &relation,
		Field:         fieldType,
		Format:        formatType,
	})
	return r
}

func (r *ResourceSchema) Resolve(resolver Resolver, base string) error {
	for i, f := range r.Fields {
		if f.Relation == nil {
			continue
		}
		oldRelStr := *f.Relation
		reverse := false
		if strings.HasPrefix(oldRelStr, "-") {
			oldRelStr = oldRelStr[1:]
			reverse = true
		}
		relID := resolver.Resolve(oldRelStr, base)
		if relID.IsNull() {
			return errors.New(oldRelStr + " not found")
		}
		newRelStr := relID.Hex()
		if reverse {
			newRelStr = "-" + newRelStr
		}
		f.Relation = &newRelStr
		r.Fields[i] = f
	}
	return nil
}

// Resource

// ResourceInstance represents a resource object for single instances.
type ResourceInstance struct {
	ObjectID ID             `json:"objectId"`
	Schema   ResourceSchema `json:"schema"`
}

func (ri ResourceInstance) toGetRequest() GetResourceRequest {
	return GetResourceRequest{
		ObjectID: ri.ObjectID,
		Object:   ResourceTypeInstance,
		Schema:   ri.Schema,
	}
}

func (ri ResourceInstance) toPutRequest(value any, layer ID) PutResourceRequest {
	return PutResourceRequest{
		ObjectID: ri.ObjectID,
		Object:   ResourceTypeInstance,
		Schema:   ri.Schema,
		Value:    value,
		Layer:    layer,
	}
}

// ResourceSpecials represents a resource object for lists of instances.
type ResourceSpecials struct {
	ObjectID ID              `json:"objectId"`
	Schema   ResourceSchema  `json:"schema"`
	Filter   ResourceFilters `json:"filter,omitempty"`
	Order    ResourceOrders  `json:"order,omitempty"`
	Offset   *int            `json:"offset,omitempty"`
	Limit    *int            `json:"limit,omitempty"`
}

func (rs ResourceSpecials) toGetRequest() GetResourceRequest {
	return GetResourceRequest{
		ObjectID: rs.ObjectID,
		Object:   ResourceTypeSpecials,
		Schema:   rs.Schema,
		Filter:   rs.Filter,
		Order:    rs.Order,
		Offset:   rs.Offset,
		Limit:    rs.Limit,
	}
}

func (rs ResourceSpecials) toPutRequest(values any, layer ID) PutResourceRequest {
	return PutResourceRequest{
		ObjectID: rs.ObjectID,
		Object:   ResourceTypeSpecials,
		Schema:   rs.Schema,
		Value:    values,
		Layer:    layer,
	}
}

// NewResourceInstance creates a new resource object for single instances.
func NewResourceInstance(id ID, schema *ResourceSchema) *ResourceInstance {
	return &ResourceInstance{
		ObjectID: id,
		Schema:   *schema,
	}
}

// NewResourceSpecials creates a new resource object for lists of instances.
func NewResourceSpecials(id ID, schema *ResourceSchema) *ResourceSpecials {
	rl := &ResourceSpecials{
		ObjectID: id,
		Schema:   *schema,
	}
	return rl
}

// AddFilter attaches a filter to the resource object.
func (rs *ResourceSpecials) AddFilter(name string, value any, operatorType OperatorType) *ResourceSpecials {
	if rs.Schema.GetField(name) != nil {
		rs.Filter = append(rs.Filter, ResourceFilter{
			Name:     name,
			Value:    value,
			Operator: operatorType,
		})
	}
	return rs
}

// AddOrder attaches an order to the resource object.
func (rs *ResourceSpecials) AddOrder(name string, ascending bool) *ResourceSpecials {
	if rs.Schema.GetField(name) != nil {
		rs.Order = append(rs.Order, ResourceOrder{
			Name:      name,
			Ascending: ascending,
		})
	}
	return rs
}

// SetOffset sets the offset for the resource object.
func (rs *ResourceSpecials) SetOffset(offset int) *ResourceSpecials {
	rs.Offset = &offset
	return rs
}

// SetLimit sets the limit for the resource object ensuring no more than limit instances are returned.
func (rs *ResourceSpecials) SetLimit(limit int) *ResourceSpecials {
	rs.Limit = &limit
	return rs
}

func fieldFromTag(field reflect.StructField) (ResourceField, error) {
	vyzeTag, ok := field.Tag.Lookup("vyze")
	if !ok {
		return ResourceField{}, errors.New("missing tag")
	}
	rf := ResourceField{}
	jsonTag, ok := field.Tag.Lookup("json")
	if !ok {
		return ResourceField{}, errors.New("require json tag")
	}
	jsonOpts := strings.Split(jsonTag, ",")
	if jsonOpts[0] == "" {
		return ResourceField{}, errors.New("require json tag")
	}
	rf.Name = jsonOpts[0]
	vyzeOpts := strings.Split(vyzeTag, ",")
	if len(vyzeOpts) > 0 {
		rf.Field = FieldType(vyzeOpts[0])
	}
	if len(vyzeOpts) > 1 {
		rf.Relation = &vyzeOpts[1]
	}
	fieldType := field.Type
	fieldKind := fieldType.Kind()
	if fieldKind == reflect.Pointer {
		fieldType = fieldType.Elem()
		fieldKind = fieldType.Kind()
	}
	list := false
	if fieldType.Kind() == reflect.Slice {
		fieldType = fieldType.Elem()
		fieldKind = fieldType.Kind()
		list = true
	}
	if !list {
		rf.Mapping = "primitive"
		rf.MappingParams = map[string]any{"change": "set"}
	} else {
		rf.Mapping = "list"
		rf.MappingParams = map[string]any{"change": "set"}
	}
	if len(vyzeOpts) > 2 {
		mappingSplit := strings.Split(vyzeOpts[2], "(")
		if mappingSplit[0] != "" {
			rf.Mapping = MappingType(mappingSplit[0])
		}
		if len(mappingSplit) == 2 {
			mappingParamsStr := mappingSplit[1]
			if mappingParamsStr[len(mappingParamsStr)-1] != ')' {
				return ResourceField{}, errors.New("expected ')' at end of mapping")
			}
			mappingParamsStr = mappingParamsStr[:len(mappingParamsStr)-1]
			mappingParams := strings.Split(mappingParamsStr, ",")
			for _, mappingParam := range mappingParams {
				keyValueSplit := strings.Split(mappingParam, ":")
				rf.SetMappingParam(keyValueSplit[0], keyValueSplit[1])
			}
		}
	}
	if fieldType.Kind() == reflect.Array && fieldType.Elem().Kind() == reflect.Uint8 {
		rf.Format = "hex"
	}
	if rf.Format == "" {
		switch fieldKind {
		case reflect.String:
			rf.Format = FormatTypeString
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			rf.Format = FormatTypeInteger
		case reflect.Float64, reflect.Float32:
			rf.Format = FormatTypeFloat
		case reflect.Bool:
			rf.Format = FormatTypeBoolean
		}
	}
	return rf, nil
}
