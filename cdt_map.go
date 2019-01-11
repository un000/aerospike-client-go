// Copyright 2013-2019 Aerospike, Inc.
//
// Portions may be licensed to Aerospike, Inc. under one or more contributor
// license agreements WHICH ARE COMPATIBLE WITH THE APACHE LICENSE, VERSION 2.0.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package aerospike

const (
	_CDT_MAP_SET_TYPE                       = 64
	_CDT_MAP_ADD                            = 65
	_CDT_MAP_ADD_ITEMS                      = 66
	_CDT_MAP_PUT                            = 67
	_CDT_MAP_PUT_ITEMS                      = 68
	_CDT_MAP_REPLACE                        = 69
	_CDT_MAP_REPLACE_ITEMS                  = 70
	_CDT_MAP_INCREMENT                      = 73
	_CDT_MAP_DECREMENT                      = 74
	_CDT_MAP_CLEAR                          = 75
	_CDT_MAP_REMOVE_BY_KEY                  = 76
	_CDT_MAP_REMOVE_BY_INDEX                = 77
	_CDT_MAP_REMOVE_BY_RANK                 = 79
	_CDT_MAP_REMOVE_KEY_LIST                = 81
	_CDT_MAP_REMOVE_BY_VALUE                = 82
	_CDT_MAP_REMOVE_VALUE_LIST              = 83
	_CDT_MAP_REMOVE_BY_KEY_INTERVAL         = 84
	_CDT_MAP_REMOVE_BY_INDEX_RANGE          = 85
	_CDT_MAP_REMOVE_BY_VALUE_INTERVAL       = 86
	_CDT_MAP_REMOVE_BY_RANK_RANGE           = 87
	_CDT_MAP_REMOVE_BY_KEY_REL_INDEX_RANGE  = 88
	_CDT_MAP_REMOVE_BY_VALUE_REL_RANK_RANGE = 89
	_CDT_MAP_SIZE                           = 96
	_CDT_MAP_GET_BY_KEY                     = 97
	_CDT_MAP_GET_BY_INDEX                   = 98
	_CDT_MAP_GET_BY_RANK                    = 100
	_CDT_MAP_GET_BY_VALUE                   = 102
	_CDT_MAP_GET_BY_KEY_INTERVAL            = 103
	_CDT_MAP_GET_BY_INDEX_RANGE             = 104
	_CDT_MAP_GET_BY_VALUE_INTERVAL          = 105
	_CDT_MAP_GET_BY_RANK_RANGE              = 106
	_CDT_MAP_GET_BY_KEY_LIST                = 107
	_CDT_MAP_GET_BY_VALUE_LIST              = 108
	_CDT_MAP_GET_BY_KEY_REL_INDEX_RANGE     = 109
	_CDT_MAP_GET_BY_VALUE_REL_RANK_RANGE    = 110
)

type mapOrderType int

// Map storage order.
var MapOrder = struct {
	// Map is not ordered. This is the default.
	UNORDERED mapOrderType // 0

	// Order map by key.
	KEY_ORDERED mapOrderType // 1

	// Order map by key, then value.
	KEY_VALUE_ORDERED mapOrderType // 3
}{0, 1, 3}

type mapReturnType int

// Map return type. Type of data to return when selecting or removing items from the map.
var MapReturnType = struct {
	// Do not return a result.
	NONE mapReturnType

	// Return key index order.
	//
	// 0 = first key
	// N = Nth key
	// -1 = last key
	INDEX mapReturnType

	// Return reverse key order.
	//
	// 0 = last key
	// -1 = first key
	REVERSE_INDEX mapReturnType

	// Return value order.
	//
	// 0 = smallest value
	// N = Nth smallest value
	// -1 = largest value
	RANK mapReturnType

	// Return reserve value order.
	//
	// 0 = largest value
	// N = Nth largest value
	// -1 = smallest value
	REVERSE_RANK mapReturnType

	// Return count of items selected.
	COUNT mapReturnType

	// Return key for single key read and key list for range read.
	KEY mapReturnType

	// Return value for single key read and value list for range read.
	VALUE mapReturnType

	// Return key/value items. The possible return types are:
	//
	// map[interface{}]interface{} : Returned for unordered maps
	// []MapPair : Returned for range results where range order needs to be preserved.
	KEY_VALUE mapReturnType

	// Invert meaning of map command and return values.  For example:
	// MapRemoveByKeyRange(binName, keyBegin, keyEnd, MapReturnType.KEY | MapReturnType.INVERTED)
	// With the INVERTED flag enabled, the keys outside of the specified key range will be removed and returned.
	INVERTED mapReturnType
}{
	0, 1, 2, 3, 4, 5, 6, 7, 8, 0x10000,
}

// Unique key map write type.
type mapWriteMode struct {
	itemCommand  int
	itemsCommand int
}

// MapWriteMode should only be used for server versions < 4.3.
// MapWriteFlags are recommended for server versions >= 4.3.
var MapWriteMode = struct {
	// If the key already exists, the item will be overwritten.
	// If the key does not exist, a new item will be created.
	UPDATE *mapWriteMode

	// If the key already exists, the item will be overwritten.
	// If the key does not exist, the write will fail.
	UPDATE_ONLY *mapWriteMode

	// If the key already exists, the write will fail.
	// If the key does not exist, a new item will be created.
	CREATE_ONLY *mapWriteMode
}{
	&mapWriteMode{_CDT_MAP_PUT, _CDT_MAP_PUT_ITEMS},
	&mapWriteMode{_CDT_MAP_REPLACE, _CDT_MAP_REPLACE_ITEMS},
	&mapWriteMode{_CDT_MAP_ADD, _CDT_MAP_ADD_ITEMS},
}

/**
 * Map write bit flags.
 * Requires server versions >= 4.3.
 */
const (
	// MapWriteFlagsDefault is the Default. Allow create or update.
	MapWriteFlagsDefault = 0

	// MapWriteFlagsCreateOnly means: If the key already exists, the item will be denied.
	// If the key does not exist, a new item will be created.
	MapWriteFlagsCreateOnly = 1

	// MapWriteFlagsUpdateOnly means: If the key already exists, the item will be overwritten.
	// If the key does not exist, the item will be denied.
	MapWriteFlagsUpdateOnly = 2

	// MapWriteFlagsNoFail means: Do not raise error if a map item is denied due to write flag constraints.
	MapWriteFlagsNoFail = 4

	// MapWriteFlagsNoFail means: Allow other valid map items to be committed if a map item is denied due to
	// write flag constraints.
	MapWriteFlagsPartial = 8
)

// MapPolicy directives when creating a map and writing map items.
type MapPolicy struct {
	attributes   mapOrderType
	flags        int
	itemCommand  int
	itemsCommand int
}

// NewMapPolicy creates a MapPolicy with WriteMode. Use with servers before v4.3
func NewMapPolicy(order mapOrderType, writeMode *mapWriteMode) *MapPolicy {
	return &MapPolicy{
		attributes:   order,
		flags:        MapWriteFlagsDefault,
		itemCommand:  writeMode.itemCommand,
		itemsCommand: writeMode.itemsCommand,
	}
}

// NewMapPolicyWithFlags creates a MapPolicy with WriteFlags. Use with servers v4.3+
func NewMapPolicyWithFlags(order mapOrderType, flags int) *MapPolicy {
	return &MapPolicy{
		attributes:   order,
		flags:        flags,
		itemCommand:  MapWriteMode.UPDATE.itemCommand,
		itemsCommand: MapWriteMode.UPDATE.itemsCommand,
	}
}

// DefaultMapPolicy returns the default map policy
func DefaultMapPolicy() *MapPolicy {
	return NewMapPolicy(MapOrder.UNORDERED, MapWriteMode.UPDATE)
}

func newMapSetPolicyEncoder(op *Operation, packer BufferEx) (int, error) {
	return packCDTParamsAsArray(packer, _CDT_MAP_SET_TYPE, op.binValue.(IntegerValue))
}

func newMapSetPolicy(binName string, attributes mapOrderType) *Operation {
	return &Operation{
		opType:   MAP_MODIFY,
		binName:  binName,
		binValue: IntegerValue(attributes),
		encoder:  newMapSetPolicyEncoder,
	}
}

func newMapCreatePutEncoder(op *Operation, packer BufferEx) (int, error) {
	return packCDTIfcParamsAsArray(packer, int16(*op.opSubType), op.binValue.(ListValue))
}

/////////////////////////

// Unique key map bin operations. Create map operations used by the client operate command.
// The default unique key map is unordered.
//
// All maps maintain an index and a rank.  The index is the item offset from the start of the map,
// for both unordered and ordered maps.  The rank is the sorted index of the value component.
// Map supports negative indexing for index and rank.
//
// Index examples:
//
// Index 0: First item in map.
// Index 4: Fifth item in map.
// Index -1: Last item in map.
// Index -3: Third to last item in map.
// Index 1 Count 2: Second and third items in map.
// Index -3 Count 3: Last three items in map.
// Index -5 Count 4: Range between fifth to last item to second to last item inclusive.
//
// Rank examples:
//
// Rank 0: Item with lowest value rank in map.
// Rank 4: Fifth lowest ranked item in map.
// Rank -1: Item with highest ranked value in map.
// Rank -3: Item with third highest ranked value in map.
// Rank 1 Count 2: Second and third lowest ranked items in map.
// Rank -3 Count 3: Top three ranked items in map.

// MapSetPolicyOp creates set map policy operation.
// Server sets map policy attributes.  Server returns null.
//
// The required map policy attributes can be changed after the map is created.
func MapSetPolicyOp(policy *MapPolicy, binName string) *Operation {
	return newMapSetPolicy(binName, policy.attributes)
}

// MapPutOp creates map put operation.
// Server writes key/value item to map bin and returns map size.
//
// The required map policy dictates the type of map to create when it does not exist.
// The map policy also specifies the mode used when writing items to the map.
func MapPutOp(policy *MapPolicy, binName string, key interface{}, value interface{}) *Operation {
	if policy.flags != 0 {
		ops := _CDT_MAP_PUT

		// Replace doesn't allow map attributes because it does not create on non-existing key.
		return &Operation{
			opType:    MAP_MODIFY,
			opSubType: &ops,
			binName:   binName,
			binValue:  ListValue([]interface{}{key, value, IntegerValue(policy.attributes), IntegerValue(policy.flags)}),
			encoder:   newMapCreatePutEncoder,
		}
	} else {
		if policy.itemCommand == _CDT_MAP_REPLACE {
			// Replace doesn't allow map attributes because it does not create on non-existing key.
			return &Operation{
				opType:    MAP_MODIFY,
				opSubType: &policy.itemCommand,
				binName:   binName,
				binValue:  ListValue([]interface{}{key, value}),
				encoder:   newMapCreatePutEncoder,
			}
		}

		return &Operation{
			opType:    MAP_MODIFY,
			opSubType: &policy.itemCommand,
			binName:   binName,
			binValue:  ListValue([]interface{}{key, value, IntegerValue(policy.attributes)}),
			encoder:   newMapCreatePutEncoder,
		}
	}
}

// MapPutItemsOp creates map put items operation
// Server writes each map item to map bin and returns map size.
//
// The required map policy dictates the type of map to create when it does not exist.
// The map policy also specifies the mode used when writing items to the map.
func MapPutItemsOp(policy *MapPolicy, binName string, amap map[interface{}]interface{}) *Operation {
	if policy.flags != 0 {
		ops := _CDT_MAP_PUT_ITEMS

		// Replace doesn't allow map attributes because it does not create on non-existing key.
		return &Operation{
			opType:    MAP_MODIFY,
			opSubType: &ops,
			binName:   binName,
			binValue:  ListValue([]interface{}{amap, IntegerValue(policy.attributes), IntegerValue(policy.flags)}),
			encoder:   newCDTCreateOperationEncoder,
		}
	} else {
		if policy.itemsCommand == int(_CDT_MAP_REPLACE_ITEMS) {
			// Replace doesn't allow map attributes because it does not create on non-existing key.
			return &Operation{
				opType:    MAP_MODIFY,
				opSubType: &policy.itemsCommand,
				binName:   binName,
				binValue:  ListValue([]interface{}{MapValue(amap)}),
				encoder:   newCDTCreateOperationEncoder,
			}
		}

		return &Operation{
			opType:    MAP_MODIFY,
			opSubType: &policy.itemsCommand,
			binName:   binName,
			binValue:  ListValue([]interface{}{MapValue(amap), IntegerValue(policy.attributes)}),
			encoder:   newCDTCreateOperationEncoder,
		}
	}
}

// MapIncrementOp creates map increment operation.
// Server increments values by incr for all items identified by key and returns final result.
// Valid only for numbers.
//
// The required map policy dictates the type of map to create when it does not exist.
// The map policy also specifies the mode used when writing items to the map.
func MapIncrementOp(policy *MapPolicy, binName string, key interface{}, incr interface{}) *Operation {
	return newCDTCreateOperationValues2(_CDT_MAP_INCREMENT, policy.attributes, binName, key, incr)
}

// MapDecrementOp creates map decrement operation.
// Server decrements values by decr for all items identified by key and returns final result.
// Valid only for numbers.
//
// The required map policy dictates the type of map to create when it does not exist.
// The map policy also specifies the mode used when writing items to the map.
func MapDecrementOp(policy *MapPolicy, binName string, key interface{}, decr interface{}) *Operation {
	return newCDTCreateOperationValues2(_CDT_MAP_DECREMENT, policy.attributes, binName, key, decr)
}

// MapClearOp creates map clear operation.
// Server removes all items in map.  Server returns null.
func MapClearOp(binName string) *Operation {
	return newCDTCreateOperationValues0(_CDT_MAP_CLEAR, MAP_MODIFY, binName)
}

// MapRemoveByKeyOp creates map remove operation.
// Server removes map item identified by key and returns removed data specified by returnType.
func MapRemoveByKeyOp(binName string, key interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_BY_KEY, MAP_MODIFY, binName, key, returnType)
}

// MapRemoveByKeyListOp creates map remove operation.
// Server removes map items identified by keys and returns removed data specified by returnType.
func MapRemoveByKeyListOp(binName string, keys []interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_KEY_LIST, MAP_MODIFY, binName, keys, returnType)
}

// MapRemoveByKeyRangeOp creates map remove operation.
// Server removes map items identified by key range (keyBegin inclusive, keyEnd exclusive).
// If keyBegin is null, the range is less than keyEnd.
// If keyEnd is null, the range is greater than equal to keyBegin.
//
// Server returns removed data specified by returnType.
func MapRemoveByKeyRangeOp(binName string, keyBegin interface{}, keyEnd interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateRangeOperation(_CDT_MAP_REMOVE_BY_KEY_INTERVAL, MAP_MODIFY, binName, keyBegin, keyEnd, returnType)
}

// MapRemoveByValueOp creates map remove operation.
// Server removes map items identified by value and returns removed data specified by returnType.
func MapRemoveByValueOp(binName string, value interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_BY_VALUE, MAP_MODIFY, binName, value, returnType)
}

// MapRemoveByValueListOp creates map remove operation.
// Server removes map items identified by values and returns removed data specified by returnType.
func MapRemoveByValueListOp(binName string, values []interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValuesN(_CDT_MAP_REMOVE_VALUE_LIST, MAP_MODIFY, binName, values, returnType)
}

// MapRemoveByValueRangeOp creates map remove operation.
// Server removes map items identified by value range (valueBegin inclusive, valueEnd exclusive).
// If valueBegin is null, the range is less than valueEnd.
// If valueEnd is null, the range is greater than equal to valueBegin.
//
// Server returns removed data specified by returnType.
func MapRemoveByValueRangeOp(binName string, valueBegin interface{}, valueEnd interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateRangeOperation(_CDT_MAP_REMOVE_BY_VALUE_INTERVAL, MAP_MODIFY, binName, valueBegin, valueEnd, returnType)
}

// MapRemoveByValueRelativeRankRangeOp creates a map remove by value relative to rank range operation.
// Server removes map items nearest to value and greater by relative rank.
// Server returns removed data specified by returnType.
//
// Examples for map [{4=2},{9=10},{5=15},{0=17}]:
//
// (value,rank) = [removed items]
// (11,1) = [{0=17}]
// (11,-1) = [{9=10},{5=15},{0=17}]
func MapRemoveByValueRelativeRankRangeOp(binName string, value interface{}, rank int, returnType mapReturnType) *Operation {
	return newCDTCreateRangeOperation(_CDT_MAP_REMOVE_BY_VALUE_REL_RANK_RANGE, MAP_MODIFY, binName, value, rank, returnType)
}

// MapRemoveByValueRelativeRankRangeCountOp creates a map remove by value relative to rank range operation.
// Server removes map items nearest to value and greater by relative rank with a count limit.
// Server returns removed data specified by returnType (See {@link MapReturnType}).
//
// Examples for map [{4=2},{9=10},{5=15},{0=17}]:
//
// (value,rank,count) = [removed items]
// (11,1,1) = [{0=17}]
// (11,-1,1) = [{9=10}]
func MapRemoveByValueRelativeRankRangeCountOp(binName string, value interface{}, rank, count int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndexCount(_CDT_MAP_REMOVE_BY_VALUE_REL_RANK_RANGE, MAP_MODIFY, binName, NewValue(value), rank, count, returnType)
}

// MapRemoveByIndexOp creates map remove operation.
// Server removes map item identified by index and returns removed data specified by returnType.
func MapRemoveByIndexOp(binName string, index int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_BY_INDEX, MAP_MODIFY, binName, index, returnType)
}

// MapRemoveByIndexRangeOp creates map remove operation.
// Server removes map items starting at specified index to the end of map and returns removed
// data specified by returnType.
func MapRemoveByIndexRangeOp(binName string, index int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_BY_INDEX_RANGE, MAP_MODIFY, binName, index, returnType)
}

// MapRemoveByIndexRangeCountOp creates map remove operation.
// Server removes "count" map items starting at specified index and returns removed data specified by returnType.
func MapRemoveByIndexRangeCountOp(binName string, index int, count int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationIndexCount(_CDT_MAP_REMOVE_BY_INDEX_RANGE, MAP_MODIFY, binName, index, count, returnType)
}

// MapRemoveByRankOp creates map remove operation.
// Server removes map item identified by rank and returns removed data specified by returnType.
func MapRemoveByRankOp(binName string, rank int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_REMOVE_BY_RANK, MAP_MODIFY, binName, rank, returnType)
}

// MapRemoveByRankRangeOp creates map remove operation.
// Server removes map items starting at specified rank to the last ranked item and returns removed
// data specified by returnType.
func MapRemoveByRankRangeOp(binName string, rank int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationIndex(_CDT_MAP_REMOVE_BY_RANK_RANGE, MAP_MODIFY, binName, rank, returnType)
}

// MapRemoveByRankRangeCountOp creates map remove operation.
// Server removes "count" map items starting at specified rank and returns removed data specified by returnType.
func MapRemoveByRankRangeCountOp(binName string, rank int, count int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationIndexCount(_CDT_MAP_REMOVE_BY_RANK_RANGE, MAP_MODIFY, binName, rank, count, returnType)
}

// MapRemoveByKeyRelativeIndexRangeOp creates a map remove by key relative to index range operation.
// Server removes map items nearest to key and greater by index.
// Server returns removed data specified by returnType.
//
// Examples for map [{0=17},{4=2},{5=15},{9=10}]:
//
// (value,index) = [removed items]
// (5,0) = [{5=15},{9=10}]
// (5,1) = [{9=10}]
// (5,-1) = [{4=2},{5=15},{9=10}]
// (3,2) = [{9=10}]
// (3,-2) = [{0=17},{4=2},{5=15},{9=10}]
func MapRemoveByKeyRelativeIndexRangeOp(binName string, key interface{}, index int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndex(_CDT_MAP_REMOVE_BY_KEY_REL_INDEX_RANGE, MAP_MODIFY, binName, NewValue(key), index, returnType)
}

// Create map remove by key relative to index range operation.
// Server removes map items nearest to key and greater by index with a count limit.
// Server returns removed data specified by returnType.
//
// Examples for map [{0=17},{4=2},{5=15},{9=10}]:
//
// (value,index,count) = [removed items]
// (5,0,1) = [{5=15}]
// (5,1,2) = [{9=10}]
// (5,-1,1) = [{4=2}]
// (3,2,1) = [{9=10}]
// (3,-2,2) = [{0=17}]
func MapRemoveByKeyRelativeIndexRangeCountOp(binName string, key interface{}, index, count int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndexCount(_CDT_MAP_REMOVE_BY_KEY_REL_INDEX_RANGE, MAP_MODIFY, binName, NewValue(key), index, count, returnType)
}

// MapSizeOp creates map size operation.
// Server returns size of map.
func MapSizeOp(binName string) *Operation {
	return newCDTCreateOperationValues0(_CDT_MAP_SIZE, MAP_READ, binName)
}

// MapGetByKeyOp creates map get by key operation.
// Server selects map item identified by key and returns selected data specified by returnType.
func MapGetByKeyOp(binName string, key interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_KEY, MAP_READ, binName, key, returnType)
}

// MapGetByKeyRangeOp creates map get by key range operation.
// Server selects map items identified by key range (keyBegin inclusive, keyEnd exclusive).
// If keyBegin is null, the range is less than keyEnd.
// If keyEnd is null, the range is greater than equal to keyBegin.
//
// Server returns selected data specified by returnType.
func MapGetByKeyRangeOp(binName string, keyBegin interface{}, keyEnd interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateRangeOperation(_CDT_MAP_GET_BY_KEY_INTERVAL, MAP_READ, binName, keyBegin, keyEnd, returnType)
}

// MapGetByKeyRelativeIndexRangeOp creates a map get by key relative to index range operation.
// Server selects map items nearest to key and greater by index.
// Server returns selected data specified by returnType.
//
// Examples for ordered map [{0=17},{4=2},{5=15},{9=10}]:
//
// (value,index) = [selected items]
// (5,0) = [{5=15},{9=10}]
// (5,1) = [{9=10}]
// (5,-1) = [{4=2},{5=15},{9=10}]
// (3,2) = [{9=10}]
// (3,-2) = [{0=17},{4=2},{5=15},{9=10}]
func MapGetByKeyRelativeIndexRangeOp(binName string, key interface{}, index int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndex(_CDT_MAP_GET_BY_KEY_REL_INDEX_RANGE, MAP_READ, binName, NewValue(key), index, returnType)
}

// MapGetByKeyRelativeIndexRangeCountOp creates a map get by key relative to index range operation.
// Server selects map items nearest to key and greater by index with a count limit.
// Server returns selected data specified by returnType (See {@link MapReturnType}).
// <p>
// Examples for ordered map [{0=17},{4=2},{5=15},{9=10}]:
// <ul>
// <li>(value,index,count) = [selected items]</li>
// <li>(5,0,1) = [{5=15}]</li>
// <li>(5,1,2) = [{9=10}]</li>
// <li>(5,-1,1) = [{4=2}]</li>
// <li>(3,2,1) = [{9=10}]</li>
// <li>(3,-2,2) = [{0=17}]</li>
func MapGetByKeyRelativeIndexRangeCountOp(binName string, key interface{}, index, count int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndexCount(_CDT_MAP_GET_BY_KEY_REL_INDEX_RANGE, MAP_READ, binName, NewValue(key), index, count, returnType)
}

// MapGetByKeyListOp creates a map get by key list operation.
// Server selects map items identified by keys and returns selected data specified by returnType.
func MapGetByKeyListOp(binName string, keys []interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_KEY_LIST, MAP_READ, binName, keys, returnType)
}

// MapGetByValueOp creates map get by value operation.
// Server selects map items identified by value and returns selected data specified by returnType.
func MapGetByValueOp(binName string, value interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_VALUE, MAP_READ, binName, value, returnType)
}

// MapGetByValueRangeOp creates map get by value range operation.
// Server selects map items identified by value range (valueBegin inclusive, valueEnd exclusive)
// If valueBegin is null, the range is less than valueEnd.
// If valueEnd is null, the range is greater than equal to valueBegin.
//
// Server returns selected data specified by returnType.
func MapGetByValueRangeOp(binName string, valueBegin interface{}, valueEnd interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateRangeOperation(_CDT_MAP_GET_BY_VALUE_INTERVAL, MAP_READ, binName, valueBegin, valueEnd, returnType)
}

// MapGetByValueRelativeRankRangeOp creates a map get by value relative to rank range operation.
// Server selects map items nearest to value and greater by relative rank.
// Server returns selected data specified by returnType.
//
// Examples for map [{4=2},{9=10},{5=15},{0=17}]:
//
// (value,rank) = [selected items]
// (11,1) = [{0=17}]
// (11,-1) = [{9=10},{5=15},{0=17}]
func MapGetByValueRelativeRankRangeOp(binName string, value interface{}, rank int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndex(_CDT_MAP_GET_BY_VALUE_REL_RANK_RANGE, MAP_READ, binName, NewValue(value), rank, returnType)
}

// MapGetByValueRelativeRankRangeCountOp creates a map get by value relative to rank range operation.
// Server selects map items nearest to value and greater by relative rank with a count limit.
// Server returns selected data specified by returnType.
//
// Examples for map [{4=2},{9=10},{5=15},{0=17}]:
//
// (value,rank,count) = [selected items]
// (11,1,1) = [{0=17}]
// (11,-1,1) = [{9=10}]
func MapGetByValueRelativeRankRangeCountOp(binName string, value interface{}, rank, count int, returnType mapReturnType) *Operation {
	return newCDTMapCreateOperationRelativeIndexCount(_CDT_MAP_GET_BY_VALUE_REL_RANK_RANGE, MAP_READ, binName, NewValue(value), rank, count, returnType)
}

func MapGetByValueListOp(binName string, values []interface{}, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_VALUE_LIST, MAP_READ, binName, values, returnType)
}

// MapGetByIndexOp creates map get by index operation.
// Server selects map item identified by index and returns selected data specified by returnType.
func MapGetByIndexOp(binName string, index int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_INDEX, MAP_READ, binName, index, returnType)
}

// MapGetByIndexRangeOp creates map get by index range operation.
// Server selects map items starting at specified index to the end of map and returns selected
// data specified by returnType.
func MapGetByIndexRangeOp(binName string, index int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_INDEX_RANGE, MAP_READ, binName, index, returnType)
}

// MapGetByIndexRangeCountOp creates map get by index range operation.
// Server selects "count" map items starting at specified index and returns selected data specified by returnType.
func MapGetByIndexRangeCountOp(binName string, index int, count int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationIndexCount(_CDT_MAP_GET_BY_INDEX_RANGE, MAP_READ, binName, index, count, returnType)
}

// MapGetByRankOp creates map get by rank operation.
// Server selects map item identified by rank and returns selected data specified by returnType.
func MapGetByRankOp(binName string, rank int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_RANK, MAP_READ, binName, rank, returnType)
}

// MapGetByRankRangeOp creates map get by rank range operation.
// Server selects map items starting at specified rank to the last ranked item and returns selected
// data specified by returnType.
func MapGetByRankRangeOp(binName string, rank int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationValue1(_CDT_MAP_GET_BY_RANK_RANGE, MAP_READ, binName, rank, returnType)
}

// MapGetByRankRangeCountOp creates map get by rank range operation.
// Server selects "count" map items starting at specified rank and returns selected data specified by returnType.
func MapGetByRankRangeCountOp(binName string, rank int, count int, returnType mapReturnType) *Operation {
	return newCDTCreateOperationIndexCount(_CDT_MAP_GET_BY_RANK_RANGE, MAP_READ, binName, rank, count, returnType)
}
